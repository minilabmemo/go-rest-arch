package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal"
	"github.com/minilabmemo/go-rest-arch/internal/apis"
	"github.com/minilabmemo/go-rest-arch/internal/card/delivery/ginrouter"
	"github.com/minilabmemo/go-rest-arch/internal/card/usecase"
	"github.com/minilabmemo/go-rest-arch/internal/config"
	"github.com/minilabmemo/go-rest-arch/internal/logger"
	"go.uber.org/zap"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		log.Printf("LoadConfig error(%v) ", err)
		os.Exit(1)
	}
}

func main() {
	start := time.Now()
	errs := make(chan error, 3)
	listenForInterrupt(errs)
	startup(errs)
	defer stopMain()

	zap.S().Infof("Service started in: %v", time.Since(start))
	zap.S().Infof("Version %s", internal.Version)
	c := <-errs
	zap.S().Warnf("terminating: %v", c)
}

func startup(err chan error) {
	logger.InitLogger()

	startGinHttpServer(err)

}

func stopMain() {

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

func loadRoutes(engine *gin.Engine) {
	engineGrp := engine.Group("service/api/v1")

	iu := usecase.NewInfoUsecase(*config.ConfigData)
	ginrouter.NewInfoHandler(engineGrp, iu)

	//TODO mid

}
