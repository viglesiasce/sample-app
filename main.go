package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var color = "green"
var version = os.Getenv("VERSION")
var redisUrl = os.Getenv("REDIS_URL")
var rdb = redis.NewClient(&redis.Options{
	Addr:     redisUrl,
	Password: "", // no password set
	DB:       0,  // use default DB
})
var rdbCtx = context.Background()

func main() {
	port := ":8080"

	flag.Parse()

	r := gin.Default()
	log.Printf("Backend version: %s\n", version)

	r.LoadHTMLGlob("templates/*")
	r.GET("/", handleIndex)
	r.GET("/version", handleVersion)
	r.GET("/healthz", handleHealthz)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func handleIndex(c *gin.Context) {
	log.Printf("Received request from %s at %s", c.Request.RemoteAddr, c.Request.URL.EscapedPath())
	p := PodMetadata{}
	counter, err := incrCounter(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	err = p.Populate(version, counter, color)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", p)
}

func incrCounter(c *gin.Context) (string, error) {
	err := rdb.Incr(rdbCtx, "counter").Err()
	if err != nil {
		return "", err
	}
	val, err := rdb.Get(rdbCtx, "counter").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func handleVersion(c *gin.Context) {
	c.String(http.StatusOK, "%s", c.Value("version"))
}

func handleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "", "")
}
