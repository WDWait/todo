package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
	创建data base步骤
	1, 读取配置文件的链接信息
	2, 使用 sql.Open() 链接
	3, 使用 sql.Ping() 测试链接
	4, 赋值全局变量 DB
*/

// DB 全局变量
var DB *sql.DB

// 初始化 db
func initDB() {

	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 从环境变量获取数据库配置
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// dsn: data source name
	// dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// 链接数据库
	open, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// 测试链接
	err = DB.Ping()
	if err != nil {
		log.Fatalf("fatalf error ping database: %v", err)
	}

	// 赋值全局变量 db
	DB = open

	// 再次判断
	if DB != nil {
		log.Fatalf("Fatalf err DB is nil")
	}
	log.Printf("Connected to database %s", dsn)

	// 默认先建一张表 todos
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Fatalf err creating table: %v", createTableSQL)
	}
}
