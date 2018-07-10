package vitelog

import (
	"github.com/sirupsen/logrus"
	"github.com/lestrrat-go/file-rotatelogs"
	"time"
	"github.com/micro/go-config"
	"log"
	"github.com/rifflock/lfshook"
	"io/ioutil"
)


var Logger *logrus.Logger

func InitLogger () {
	if Logger != nil {
		return
	}
	env := config.Get("env").String("dev")
	if Logger == nil {
		Logger = logrus.New()
		if env == "production" {
			logDir := config.Get("log", "dirname").String("./logs/")
			infoWriter, err := rotatelogs.New(
				logDir+"info.log.%Y%m%d",
				rotatelogs.WithLinkName(logDir + "info"),
				rotatelogs.WithMaxAge(30*time.Duration(86400)*time.Second),
				rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
			)

			if err != nil {
				log.Fatal("GetLogger failed. Error is " + err.Error())
			}

			errorWriter, err := rotatelogs.New(
				logDir+"error.log.%Y%m%d",
				rotatelogs.WithLinkName(logDir + "error"),
				rotatelogs.WithMaxAge(30*time.Duration(86400)*time.Second),
				rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
			)

			if err != nil {
				log.Fatal("GetLogger failed. Error is " + err.Error())
			}


			warnWriter, err := rotatelogs.New(
				logDir+"error.log.%Y%m%d",
				rotatelogs.WithLinkName(logDir + "warn"),
				rotatelogs.WithMaxAge(30*time.Duration(86400)*time.Second),
				rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
			)

			if err != nil {
				log.Fatal("GetLogger failed. Error is " + err.Error())
			}

			Logger.Hooks.Add(lfshook.NewHook(
				lfshook.WriterMap{
					logrus.InfoLevel:  infoWriter,
					logrus.WarnLevel: warnWriter,
					logrus.ErrorLevel: errorWriter,
				},
				&logrus.TextFormatter{},
			))

			Logger.SetOutput(ioutil.Discard)
		}
	}
}