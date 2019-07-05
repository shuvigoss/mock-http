package app

import (
	"io/ioutil"
	"mock-http/ui"
	"mock-http/util"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bingoohuang/now"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	startupTime now.Now
	R           *gin.Engine
	UI          *ui.Context
	MR          *gin.Engine
}

func (a App) Start() {
	a.Route()
	a.Run()
}

func (a App) Run() {
	addr := viper.GetString("addr")
	mockAddr := viper.GetString("mockAddr")

	// restart by self
	server := &http.Server{Addr: addr, Handler: a.R}
	mockServer := &http.Server{Addr: mockAddr, Handler: a.MR}

	updatePidFile()

	go func() {
		logrus.Infof("mock-http ui started to run on addr %s", addr)
		if err := gracehttp.Serve(server); err != nil {
			panic(err)
		}
	}()
	go func() {
		logrus.Infof("mock-http mock started to run on addr %s", mockAddr)
		if err := gracehttp.Serve(mockServer); err != nil {
			panic(err)
		}
	}()

	select {}
}

func CreateApp() *App {
	util.InitFlags()

	app := &App{}
	app.startupTime = now.MakeNow()

	log := util.InitLog()
	app.R = util.InitGin(log)
	app.MR = util.InitGin(log)
	app.UI = ui.CreateContext()
	_ = util.InitDb()

	return app
}

// kill -USR2 {pid} 会执行重启
func updatePidFile() {
	pidFile := "var/pid"
	envPidFile := os.Getenv("PID_FILE")
	if envPidFile != "" {
		pidFile = envPidFile
	}

	bytes, err := ioutil.ReadFile(pidFile)
	if err != nil {
		logrus.Errorf("read pid file error %s", err.Error())
		return
	}

	oldPid, err := strconv.Atoi(strings.TrimSpace(string(bytes)))
	if err != nil {
		logrus.Errorf("trans pid file error %s", err.Error())
		return
	}

	logrus.Infof("old pid is %d, new pid is %d", oldPid, os.Getpid())

	if os.Getpid() != oldPid {
		_ = ioutil.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0644)
	}
}
