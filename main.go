package main

import (
	"fmt"
	"github.com/wangchenpeng-home/blog-service/global"
	"github.com/wangchenpeng-home/blog-service/internal/model"
	"github.com/wangchenpeng-home/blog-service/internal/routers"
	"github.com/wangchenpeng-home/blog-service/pkg/logger"
	setting2 "github.com/wangchenpeng-home/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("global.ServerSetting: %v", global.ServerSetting)
	log.Printf("global.ServerSetting: %v", global.AppSetting)
	log.Printf("global.ServerSetting: %v", global.DatabaseSetting)

	// 测试日志组件
	global.Logger.Infof("%s: my-go-programming/%s", "eddycjy", "blog-service")

	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	fmt.Println(global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
