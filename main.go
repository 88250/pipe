package main

import (
	"io"
	"net/http"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const TABLE_PREFIX = "b3_solo_go_"

type User struct {
	gorm.Model
	Name string `sql:"type:VARCHAR(16) CHARACTER SET utf8 COLLATE utf8_general_ci"`
}

func main() {
	//gin.DisableConsoleColor()

	f, err := os.Create("solo.log")
	if nil != err {
		panic(err)
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetLevel(log.DebugLevel)

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	router := gin.Default()
	//app.Use(favicon.New("./favicon.ico"))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Exiting Solo ...")

		if err := server.Close(); nil != err {
			log.Fatal("Server Close:", err)
		}
	}()

	server.ListenAndServe()

	db.Close()
	log.Println("Solo exited")
}
