package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/Praveenkusuluri08/api/routes"
	dbconfig "github.com/Praveenkusuluri08/dbConfig"
	"github.com/Praveenkusuluri08/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	// 打印环境变量以确保它们已正确加载
	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))

	dbconfig.DBConnect()
}

var wg = sync.WaitGroup{}

var staticFiles embed.FS

func main() {
	r := gin.Default()

	// 添加日志中间件
	fmt.Println("Adding log middleware")
	r.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Printf("Response: %d", c.Writer.Status())
	})

	// 添加中间件以打印本地文件路径
	r.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			requestedPath := c.Request.URL.Path
			localPath := filepath.Join(".", requestedPath)
			absPath, err := filepath.Abs(localPath)
			if err != nil {
				log.Printf("Error getting absolute path: %v", err)
			} else {
				log.Printf("Local file path: %s", absPath)
			}
		}
		c.Next()
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// r.StaticFS("/static/css/", http.FS(staticFiles))
	// r.StaticFS("/static/javascript/", http.FS(staticFiles))
	r.Static("/static/css", "./static/css")
	r.Static("/static/javascript", "./static/javascript")

	dbconfig.DBConnect()
	routes.UserRoutes(r)
	routes.TaskRoutes(r)
	r.Use(gin.Logger())

	// wg.Add(1)
	// go func() {
	err := store.CreateTables()
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Tables created successfully.")
	// 	wg.Done()
	// }()

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	wg.Wait()
}
