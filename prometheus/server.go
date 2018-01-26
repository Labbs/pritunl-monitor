package prometheus

import (
	"net/http"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func Limiter(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1000000)
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
	}()

	c.Next()
}

func Register(engine *gin.Engine) {
	engine.Use(Limiter)
	engine.Use(Recovery)

	engine.GET("/metrics", metricsHandler())
}

func metricsHandler() gin.HandlerFunc {
	prmHandler := prometheus.Handler()

	return func(c *gin.Context) {
		prmHandler.ServeHTTP(c.Writer, c.Request)
	}
}

func Start() (err error) {
	Update()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			err := Update()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	Register(router)

	server := &http.Server{
		Addr:           ":9099",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 4096,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}
