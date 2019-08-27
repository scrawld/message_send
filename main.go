/* ######################################################################
# File Name: main.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-08-26 18:17:18
####################################################################### */
package main

import (
	"flag"
	"fmt"
	"message_sender/libs/config"
	"message_sender/libs/loops"
	"message_sender/models"
	"message_sender/models/media"
	"message_sender/work_handler"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/cihub/seelog"
)

// pass through when build project, go build -ldflags "main.__version__ 1.2.1" app
var (
	__version__ string
	pwd         = flag.String("d", "", "work directory")
	cfg         = flag.String("c", "conf/app.toml", "config file, relative path")
	wg          sync.WaitGroup
	loopws      *loops.Loops
	message     *work_handler.MessageHandler
	exit        = make(chan int)
)

func init() {
	flag.Parse()

	if *pwd == "" {
		*pwd, _ = os.Getwd()
	}
	os.Setenv("VERSION", __version__)
	os.Setenv("WORKDIR", *pwd)
}

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		fmt.Printf("Signal %d received, App is about to stop...\n", sig)
		go loopws.Stop()
		go message.Stop()
		wg.Wait()
		time.Sleep(time.Second)
		fmt.Printf("App has gone away\n")
		close(exit)
	}()
}

func main() {
	// configuration
	fmt.Printf("Using configuration %s\n", *cfg)
	if strings.HasPrefix(*cfg, "/") == false {
		*cfg = path.Join(os.Getenv("WORKDIR"), *cfg)
	}
	if err := config.SetFileAndLoad(*cfg); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// logger
	logFile := config.Get().Basic.LogFile
	fmt.Printf("Using log configuration %s\n", logFile)
	if strings.HasPrefix(logFile, "/") == false {
		logFile = path.Join(os.Getenv("WORKDIR"), logFile)
	}
	var err error
	var logger seelog.LoggerInterface
	if logger, err = seelog.LoggerFromConfigAsFile(logFile); err != nil {
		fmt.Printf("Log configuration parse error: %s\n", err)
		os.Exit(-1)
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	// GC
	debug.SetGCPercent(config.Get().Gc.Percent)

	// init models
	models.Init()
	media.Start()

	// register stop listener
	registerSignalHandler()

	loopws = loops.New()
	// 任务拆分
	loopws.AddFunc(time.Second, func() { wg.Add(1); defer wg.Done(); work_handler.NewTaskHandler().Run() })

	loopws.Start()

	wg.Add(1)
	message = work_handler.NewMessageHandler()
	message.Start()
	wg.Done()

	fmt.Printf("App start ok, running as %d\n", os.Getpid())
	select {
	case <-exit:
	}
}
