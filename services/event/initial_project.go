package init_project

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/pkg/database"
	"project-adhyaksa/pkg/httpserver"
	"project-adhyaksa/pkg/logger"
	"project-adhyaksa/pkg/upload"
	v1 "project-adhyaksa/services/event/internal"
	"project-adhyaksa/services/event/internal/repository"
	"project-adhyaksa/services/event/internal/service"
	"project-adhyaksa/services/event/internal/usecase"
	"syscall"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.uber.org/zap"
)

func NewProject() {
	//setup logger
	path := `./services/event/log/error.log`
	l, err := logger.NewZapFileLogger(path)
	if err != nil {
		log.Fatal("logger error ", err)
	}
	zap.ReplaceGlobals(zap.Must(l, err))

	//setup config
	config, err := config.NewConfig("./services/event/")
	if err != nil {
		log.Fatal("config error ", err)
	}

	//setup db
	db := database.ConnectMYSQL("mysql", config)
	config.Db = db

	//setup handler
	handler := gin.New()

	//allow cors
	handler.Use(cors.AllowAll())

	//start server
	httpServer := httpserver.New(handler, httpserver.Port(config.Port))
	log.Println("server is running on port:", config.Port)

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.Cloudinary.CloudName, config.Cloudinary.ApiKey, config.Cloudinary.ApiScret)
	if err != nil {
		log.Fatal("cloudinary error: ", err)
	}

	//dependency injection
	uploadCloudinary := upload.NewCloudinaryUpload(cld, config)
	repository := repository.InitRepository(config)
	service := service.InitService(repository, uploadCloudinary)
	usecase := usecase.InitUseCase(service)
	v1.NewEvent(handler, usecase)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: ", s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}
	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
