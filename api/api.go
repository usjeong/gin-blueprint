package api

import (
	"log"

	"quicket.co.kr/qush/api/model"
	"quicket.co.kr/qush/conf"

	"github.com/gin-gonic/gin"
	"github.com/usjeong/gin-blueprint/api"
)

// App app 객체
type App struct {
	Env *viper.Viper
}

// App Http서버 객체
var (
	app *App
	db  *model.DBPool
)

// SetRouter set up router group
func SetRouter(router *gin.Engine) {
	topic := router.Group("topic")
	{
		topic.GET("/ping", api.PingContext)
	}
}

// NewApp Http서버 객체 생성
func NewApp(caseOne *conf.CaseOne) {
	log.Println("init Qush app")
	app = &App{
		Env: caseOne.Env,
	}

	db = &model.DBPool{
		Master: caseOne.DBWriter,
		Slave:  caseOne.DBReader,
	}

}
