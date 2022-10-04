package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/minilabmemo/go-rest-arch/internal/api"
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
	c := <-errs
	zap.S().Warnf("terminating: %v", c)
}

func startup(err chan error) {
	logger.InitLogger()
	api.StartHttpServer(err)

}

func stopMain() {

}

func listenForInterrupt(errChan chan error) {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-c)
	}()
}
