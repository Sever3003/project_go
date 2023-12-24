package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Gwoop/Driver/api/controler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func API() {
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	//насройка Header cors для валидации и защиты
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "content-type", "accept", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           10 * time.Minute,
	}))

	//endpoint
	r.POST("location/add", controler.Add)
	r.PUT("location/update/:uuid", controler.Update)
	r.DELETE("location/delete/:uuid", controler.Delete)
	r.GET("location/get", controler.Get)

	//конфигурация сервера, время чтения, записи и т.д.
	s := &http.Server{
		Addr:           os.Getenv("IP"),
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
