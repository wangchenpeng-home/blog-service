package main

import (
	"github.com/wangchenpeng-home/blog-service/global"
	"github.com/wangchenpeng-home/blog-service/internal/model"
	"github.com/wangchenpeng-home/blog-service/internal/routers"
	setting2 "github.com/wangchenpeng-home/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setDBEngine()
	if err != nil {
		log.Fatalf("init.setDBEngine err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	
	log.Printf("global.ServerSetting: %v", global.ServerSetting)
	log.Printf("global.ServerSetting: %v", global.AppSetting)
	log.Printf("global.ServerSetting: %v", global.DatabaseSetting)
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
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}