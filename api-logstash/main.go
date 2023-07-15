package main

import (
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Membuat instance Logrus
	logger := logrus.New()

	// Membuat hook Logstash
	logstashHost := os.Getenv("LOGSTASH_HOST")
	logstashPort := os.Getenv("LOGSTASH_PORT")
	appName := os.Getenv("APP_NAME")

	hook, err := logrustash.NewLogstashHook("tcp", logstashHost+":"+logstashPort, appName)
	if err != nil {
		logrus.Fatal(err)
	}
	logger.Hooks.Add(hook)

	// Menggunakan Logrus sebagai logger default Gin
	gin.DefaultWriter = logger.Writer()
	gin.DefaultErrorWriter = logger.Writer()

	// Membuat instance Gin
	r := gin.Default()

	// Menambahkan middleware untuk logging request
	r.Use(func(c *gin.Context) {
		// Melacak waktu awal eksekusi
		startTime := time.Now()

		// Melacak User-Agent
		userAgent := c.Request.UserAgent()

		// Melacak IP pengguna
		ip := c.ClientIP()

		// Melacak path URL
		path := c.Request.URL.Path

		// Membuat log entry dengan informasi yang dilacak
		entry := logger.WithFields(logrus.Fields{
			"UserAgent":   userAgent,
			"IP":          ip,
			"Path":        path,
			"ElapsedTime": time.Since(startTime),
		})

		// Mengirimkan log entry ke Logstash
		entry.Info("Request received")

		// Melanjutkan penanganan request
		c.Next()
	})

	// Menambahkan route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Menjalankan server
	r.Run(":8080")
}
