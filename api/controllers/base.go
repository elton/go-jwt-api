package controllers

import (
	"fmt"
	"log"
	"net/http"

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
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected the %s database", DbDriver)
		}
	}

	// database migration
	server.DB.Debug().AutoMigrate(&models.User{})
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(add string) {
	fmt.Printf("Listening to port%s\n", add)
	log.Fatal(http.ListenAndServe(add, server.Router))
}
