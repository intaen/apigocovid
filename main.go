package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	chh "apigocovid/chart/delivery/http"
	chr "apigocovid/chart/repository"
	chu "apigocovid/chart/usecase"
	cvh "apigocovid/covid/delivery/http"
	cvr "apigocovid/covid/repository"
	cvu "apigocovid/covid/usecase"
	"apigocovid/pkg/db/postgres"
	"apigocovid/pkg/db/redis"
	"apigocovid/pkg/scheduler"

	// mw "apigocovid/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "apigocovid/docs"

	"github.com/gin-gonic/contrib/static"
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
	psqlClient, err := postgres.NewSQLClient()
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	}

	// Create connection pool
	sql, err := postgres.ConnectionPool(psqlClient)
	if err != nil {
		log.Fatalf("ConnectionPool: %s", err)
	}
	log.Println(fmt.Sprintf("Postgres connected, Status: %#v", sql.Stats()))
	defer sql.Close()

	// Init connection redis
	rd := redis.NewRedisClient()
	defer rd.Close()
	log.Println("Redis connected", rd)

	// Create client redis
	var ctx *gin.Context
	redisClient := redis.CreateClient(rd, ctx)

	// Create router
	r := gin.Default()
	// r.Use(mw.RequestLoggerMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("./views", true)))
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

	// Initiate Repository
	chRepo := chr.NewChartRepository(redisClient, psqlClient)
	cvRepo := cvr.NewCovidRepository(redisClient, psqlClient)

	// Initiate Usecase
	chUsecase := chu.NewChartUsecase(chRepo)
	cvUsecase := cvu.NewCovidUsecase(cvRepo)

	// Initiate Handler
	chh.NewChartHandler(r, chUsecase)
	cvh.NewCovidHandler(r, chUsecase, cvUsecase)

	// Start Scheduler
	scheduler.StartScheduler(cvUsecase)

	// Setting timeout
	nHandler := http.TimeoutHandler(handler, 20*time.Second, "Timeout!")

	// Start server
	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.address"), nHandler))
}