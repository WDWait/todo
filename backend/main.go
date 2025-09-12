package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/backend/config"
	"todo/backend/controller"
	"todo/backend/repository"
	"todo/backend/router"
	"todo/backend/service"

	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 初始化数据库连接
	config.InitDB()
	defer config.DB.Close()

	// 初始化依赖
	todoRepo := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	// 设置路由
	r := router.SetupRouter(todoController)

	// 获取服务器端口
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
