package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	ch "github.com/intaen/apigocovid/covid/delivery/http"
	cr "github.com/intaen/apigocovid/covid/repository"
	cu "github.com/intaen/apigocovid/covid/usecase"
	"github.com/intaen/apigocovid/pkg/db/postgres"
	"github.com/intaen/apigocovid/pkg/scheduler"

	// mw "github.com/intaen/apigocovid/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/intaen/apigocovid/docs"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

// @title GOCOVID
// @version 1.0
// @description This page is API documentation to get data about Covid-19
// @schemes http
// @host localhost:9030
// @contact.name Developer
// @contact.email intanmarsjaf@outlook.com
func main() {
	// Init connection database
	psql, err := postgres.Init()
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	}

	// Create connection pool
	sql, err := postgres.ConnectionPool(psql)
	if err != nil {
		log.Fatalf("Connection pool: %s", err)
	}
	log.Println(fmt.Sprintf("Postgres connected, Status: %#v", sql.Stats()))
	defer sql.Close()

	// Create router
	r := gin.Default()
	// r.Use(mw.RequestLoggerMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(static.Serve("/", static.LocalFile("./html", true)))
	handler := cors.Default().Handler(r)

	// Handle wrong method
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) { c.JSON(405, gin.H{"message": "Method Not Allowed"}) })
	// Handle no route
	r.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"message": "Page Not Found"}) })

	// Swagger
	url := ginSwagger.URL(viper.GetString("server.host") + ":" + viper.GetString("server.address") + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))                                             // http://localhost:9030/swagger/index.html

	// Initiate Repository, Usecase, Handler
	cRepo := cr.NewCovidRepository(psql)
	cUsecase := cu.NewCovidUsecase(cRepo)
	ch.NewCovidHandler(r, cUsecase)

	// Start Scheduler
	scheduler.StartScheduler(cUsecase)

	// Setting timeout
	nHandler := http.TimeoutHandler(handler, 20*time.Second, "Timeout!")

	// Start server
	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.address"), nHandler))
}
