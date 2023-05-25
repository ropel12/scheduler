package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ropel12/scheduler/config"
	"github.com/ropel12/scheduler/helper"
	"github.com/ropel12/scheduler/repository"
	"github.com/ropel12/scheduler/service"

	"github.com/robfig/cron/v3"
)

func main() {

	conf, err := config.InitConfiguration()
	helper.PanicIfError(err)
	db, err := config.GetConnection(conf)
	helper.PanicIfError(err)
	schoolrepo := repository.NewRepo()
	schoolserv := service.NewService(db, schoolrepo)
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	defer scheduler.Stop()
	scheduler.AddFunc("*/1 * * * *", schoolserv.UpdateTestResult)
	// scheduler.AddFunc("0 0 1 1 *", func() { SendAutomail("New Year") })
	go scheduler.Start()
	log.Println("[INFO] Starting Service Scheduler")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
