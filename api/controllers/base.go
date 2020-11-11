package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elton/go-jwt-api/api/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbHost, DbPort, DbName string) {
	var err error
	if DbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{PrepareStmt: true, DisableForeignKeyConstraintWhenMigrating: true})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected the %s database", DbDriver)
		}
	}

	// database connection pool settings.
	// refer to https://www.alexedwards.net/blog/configuring-sqldb
	sqlDB, _ := server.DB.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(25)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(25)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	// database migration
	server.DB.Debug().AutoMigrate(&models.User{})
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(add string) {
	fmt.Printf("Listening to port %s", add)
	log.Fatal(http.ListenAndServe(add, server.Router))
}
