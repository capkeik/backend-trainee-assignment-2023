package main

import (
	"context"
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/config"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/controller"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/pg"
	pg2 "github.com/capkeik/backend-trainee-assignment-2023/internal/repository/pg"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/repository/static"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Starting Segmentation Service")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	log.Println("Reading config")
	cfg := config.Get()
	_ = cfg

	db, err := InitDB()
	if err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	saveDir := fmt.Sprintf("%v/../../data/csv-records/", currentDir)
	userRepo := pg2.NewUserRepo(db)
	segmentRepo := pg2.NewSegmentRepo(db)
	recordsRepo := pg2.NewRecordRepo(db)
	csvRepo := static.NewCSV(saveDir)

	recordService := service.NewRecordService(db)
	userService := service.NewUserService(&userRepo, &recordService)

	recordController := controller.NewRecords(&recordsRepo, &csvRepo)
	userController := controller.NewUsers(ctx, userService)
	segmentController := controller.NewSegments(ctx, &segmentRepo)

	// Init echo
	e := echo.New()
	e.Use(middleware.Static("static"))

	// /user routes
	userRoutes := e.Group("/user")
	userRoutes.GET("/:id", userController.Get)
	userRoutes.POST("", userController.Create)
	userRoutes.PATCH("", userController.UpdateSegments)

	// /segment routes
	segmentRoutes := e.Group("/segment")
	segmentRoutes.POST("", segmentController.Create)
	segmentRoutes.DELETE("", segmentController.Delete)

	// /records routes

	recordsRoutes := e.Group("/records")
	recordsRoutes.GET("", recordController.Get)
	recordsRoutes.GET("/download/csv/:filename", recordController.Download)

	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}

func InitDB() (*gorm.DB, error) {
	log.Println("Opening DB connection")
	db, err := pg.Connect()
	if err != nil {
		return nil, fmt.Errorf("%s, %w", "Error initializing database:", err)
	}
	//sqlDb, err := db.DB()
	////defer func(sqlDb *sql.DB) {
	////	log.Println("Closing DB connection")
	////	err = sqlDb.Close()
	////	if err != nil {
	////		err = fmt.Errorf("%s, %w", "Error closing database:", err)
	////	}
	////}(sqlDb)

	return db, err
}
