package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"driverserv/api/controler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r.GET("/metrics", prometheusHandler())

	r.POST("driver/add", controler.Add)
	r.PUT("driver/update", controler.Update)
	r.DELETE("driver/delete/:name", controler.Delete)
	r.GET("driver/get", controler.Get)

	r.POST("driverinfo/add/:name", controler.AddDriver)
	r.GET("driverinfo/get", controler.GetDriver)

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

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
