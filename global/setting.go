package global

import (
	"github.com/wangchenpeng-home/blog-service/pkg/logger"
	"github.com/wangchenpeng-home/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)
