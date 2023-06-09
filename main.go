package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/ropel12/scheduler/config"
	"github.com/ropel12/scheduler/helper"
	"github.com/ropel12/scheduler/pkg"
	"github.com/ropel12/scheduler/repository"
	"github.com/ropel12/scheduler/service"
)

func main() {

	conf, err := config.InitConfiguration()
	helper.PanicIfError(err)
	db, err := config.GetConnection(conf)
	helper.PanicIfError(err)
	schoolrepo := repository.NewRepo()
	nsq, err := pkg.NewNSQ(conf)
	helper.PanicIfError(err)
	pusher := pkg.NewPusher(conf)
	schoolserv := service.NewService(db, schoolrepo, nsq, pusher)
	local := time.Now().Location()
	scheduler := cron.New(cron.WithLocation(local))
	defer scheduler.Stop()
	scheduler.AddFunc("@every 10s", schoolserv.UpdateTestResult)
	scheduler.AddFunc("@every 17s", schoolserv.SendMonthlyBilling)
	go scheduler.Start()
	log.Println("[INFO] Starting Service Scheduler")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	nsq.Stop()
	log.Println("[INFO]  Scheduler Service Stopped")

}
