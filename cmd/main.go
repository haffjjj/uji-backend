package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/haffjjj/myblog-backend/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/viper"

	_httpDelivery "github.com/haffjjj/uji-backend/delivery/http"
	_courseRepo "github.com/haffjjj/uji-backend/repository/course"
	_courseUsecase "github.com/haffjjj/uji-backend/usecase/course"
)

func init() {
	viper.SetConfigFile("config")
	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		dbHost  = viper.GetString("database.mongodb.host")
		dbPort  = viper.GetString("database.mongodb.port")
		dbUName = viper.GetString("database.mongodb.username")
		dbPass  = viper.GetString("database.mongodb.password")
		port    = viper.GetString("port")
	)

	mgoClient, err := mongo.Connect(context.TODO(), fmt.Sprint("mongodb://", dbUName, ":", dbPass, "@", dbHost, dbPort))
	if err != nil {
		log.Fatal(err)
	}
	err = mgoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer mgoClient.Disconnect(context.TODO())

	// ===========

	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = &utils.Validator{Validator: validator.New()}

	// ===========

	courseRepo := _courseRepo.NewMongoCourseRepository(mgoClient)
	courseUsecase := _courseUsecase.NewCourseUsecase(courseRepo)

	_httpDelivery.NewTagHandler(e, courseUsecase)

	// ===========

	e.Logger.Fatal(e.Start(port))
}