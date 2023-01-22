package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	docs "github.com/minilabmemo/go-rest-arch/cmd/app-core/docs" //swag init -pd
	internal "github.com/minilabmemo/go-rest-arch/services"
	"github.com/minilabmemo/go-rest-arch/services/apis"
	"github.com/minilabmemo/go-rest-arch/services/config"
	"github.com/minilabmemo/go-rest-arch/services/logger"
	"github.com/minilabmemo/go-rest-arch/services/models"
	"github.com/minilabmemo/go-rest-arch/services/router/delivery/ginrouter"
	"github.com/minilabmemo/go-rest-arch/services/router/repository/mongo"
	"github.com/minilabmemo/go-rest-arch/services/router/usecase"

	"go.uber.org/zap"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		log.Printf("LoadConfig error(%v) ", err)
		os.Exit(1)
	}
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample service server.
func main() {
	start := time.Now()
	errs := make(chan error, 3)
	listenForInterrupt(errs)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	startup(ctx, errs)
	defer stopMain(ctx)
	docs.SwaggerInfo.Title = fmt.Sprintf("Swagger %s API", config.ConfigData.Service.Name)
	zap.S().Infof("Service started in: %v", time.Since(start))
	zap.S().Infof("Version %s", internal.Version)
	zap.S().Infof("%s.Host %s", internal.ClientMongo, config.ConfigData.Clients[internal.ClientMongo].Host)
	zap.S().Infof("%s.Database %s", internal.ClientMongo, config.ConfigData.Clients[internal.ClientMongo].More["database"])
	zap.S().Infof("%s.collection_todo %s", internal.ClientMongo, config.ConfigData.Clients[internal.ClientMongo].More["collection_todo"])
	c := <-errs
	zap.S().Warnf("terminating: %v", c)
}

var client models.MongoRepository

func startup(ctx context.Context, errCh chan error) {
	logger.InitLogger()
	var err error
	client, err = mongo.NewMongoClient(context.TODO(), config.ConfigData.Clients[internal.ClientMongo])
	if err != nil {
		zap.S().Errorf("NewMongoClient: %v", err)
		return
	}

	startGinHttpServer(errCh)

}

func stopMain(ctx context.Context) {
	if err := client.Close(ctx); err != nil {
		zap.S().Errorf("MongoClient Close: %v", err)
	}
}

func listenForInterrupt(errChan chan error) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGALRM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
}

func startGinHttpServer(err chan error) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	loadRoutes(engine)
	apis.StartHttpServer(err, engine)

}

// @BasePath  /service/api/v1
func loadRoutes(engine *gin.Engine) {
	engineGrp := engine.Group("service/api/v1")

	iu := usecase.NewInfoUsecase(*config.ConfigData)
	ginrouter.NewInfoHandler(engineGrp, iu)

	cu := usecase.NewCardUsecase(*config.ConfigData, mongo.NewMongoCardRepository(client.Connection(), "card", "todo"))
	ginrouter.NewTodoHandler(engineGrp, cu)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	//TODO mid

}
