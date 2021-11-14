package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangchenpeng-home/blog-service/global"
	"github.com/wangchenpeng-home/blog-service/pkg/app"
	"github.com/wangchenpeng-home/blog-service/pkg/email"
	"github.com/wangchenpeng-home/blog-service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		PassWord: global.EmailSetting.PassWord,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recovery err: %v"
				global.Logger.WithCallersFrames().Errorf(c, s, err)

				// 增加邮件中间件通知
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息，%v", err),
					)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
