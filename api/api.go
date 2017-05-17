package api

import (
	"log"

	"quicket.co.kr/qush/api/model"
	"quicket.co.kr/qush/conf"

	"github.com/gin-gonic/gin"
	"github.com/usjeong/gin-blueprint/api"
)

// App inject configuration
type App struct {
	Env *viper.Viper
}

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

// NewApp create app
func NewApp(caseOne *conf.CaseOne) {
	log.Println("init app")
	app = &App{
		Env: caseOne.Env,
	}

	db = &model.DBPool{
		Master: caseOne.DBWriter,
		Slave:  caseOne.DBReader,
	}

}
