package main

import (
	"log"
	"net/http"
	"time"

	cors "gopkg.in/gin-contrib/cors.v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/usjeong/gin-blueprint/api"
	"github.com/usjeong/gin-blueprint/conf"
)

func main() {
	caseOne := conf.NewCaseOne("production")
	env := caseOne.Env
	api.NewApp(caseOne)

	redisHost := env.GetString("Store")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	sessionStore, err := sessions.NewRedisStore(10, "tcp", redisHost, env.GetString("StorePassword"), []byte("secret"))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("session type: redis")

	r.Use(cors.Default())
	r.Use(sessions.Sessions("appsession", sessionStore))
	api.SetRouter(r)

	pprof.Register(r, nil)

	server := &http.Server{
		Addr:           env.GetString("Listen"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   35 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start server\n\tport: %s\n", server.Addr)

	if server.ListenAndServe() != nil {
		log.Fatal(err)
	}
}
