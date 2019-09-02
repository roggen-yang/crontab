package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/roggen-yang/crontab/master"
)

var (
	configFile string
)

func initArgs() {
	flag.StringVar(&configFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = master.InitCofnig(configFile); err != nil {
		goto ERR
	}

	// 任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	// 启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		fmt.Println("heart beat")
		time.Sleep(1 * time.Second)
	}

ERR:
	fmt.Println(err)
}
