package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arvindkhoisnam/challenges/02/routeHandlers"
	"github.com/gin-gonic/gin"
)



func routes(app *gin.Engine){
	api := app.Group("/api/v1")
	api.Use(LoggerMiddleware())
	{	
		api.GET("/health",routeHandlers.HealthCheck)
		api.GET("/models",routeHandlers.Models)
		api.POST("/completion",routeHandlers.Completion)
		api.POST("/chat",routeHandlers.Chat)
	}
}
func main(){
	app := gin.Default()
	routes(app)

	server := &http.Server{
		Addr: ":3000",
		Handler: app,
	}

	quit := make(chan os.Signal,1)

	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)

	go func ()  {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Fatalf("listen: %s\n", err)
		} 
	}()
	log.Println("Server started on :3000")

	<- quit
	log.Println("Shutting down server...")

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	
	log.Println("Server exited gracefully")

}

func LoggerMiddleware()gin.HandlerFunc{
	return func(c *gin.Context)  {
		start := time.Now()

		c.Set("userID", 123)
		c.Next()
		end := time.Since(start)
		fmt.Println(c.Request.RemoteAddr)
		log.Printf("%s %s %s %v", c.Request.Method, c.Request.RequestURI, c.Request.Proto, end)
	}
}
