package main

import (
	"net/http"
	"os"
	"strconv"
	"student-api/internal/container"
	"student-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := container.New()
	if err != nil {
		panic(err)
	}

	if app.Config.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middleware.GinStructuredLogger(app.Logger))
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.POST("/", app.StudentHandler.CreateStudent)
			students.GET("/", app.StudentHandler.ListStudents)
			students.GET("/:id", app.StudentHandler.GetStudent) // {id} -> :id
			students.DELETE("/:id", app.StudentHandler.DeleteStudent)
		}
	}

	addr := ":" + strconv.Itoa(app.Config.Server.Port)
	app.Logger.Info("Server started with Gin", "addr", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		app.Logger.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}
}
