package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iamshubha/go-init/internal/client"
	"github.com/iamshubha/go-init/internal/dao"
	"github.com/iamshubha/go-init/internal/service"
	"golang.org/x/exp/slog"
)

type Server struct {
	router  *gin.Engine
	logger  *slog.Logger
	db      dao.DBImpl
	service service.ServiceImpl
}

// Connect function is use for connect mysql database
// and create database if the database is not already created
func (s *Server) Connect() {

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil) // 👈
	s.logger = slog.New(jsonHandler)
	ginSetMode(os.Getenv("R_MODE"))
	s.router = gin.Default()
	s.db = dao.NewSqlDB()
	c := client.NewClient(s.logger)
	s.service = service.NewService(c, s.db, s.logger)
	s.startCronJob()
	s.InitializeRouter()
	s.router.Run(func() string {
		port := os.Getenv("SERVER_PORT")
		if port == "" {
			return "8080"
		}
		return port
	}())
	// s.logger.Error(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), s.router).Error())
}

func (s *Server) startCronJob() {

	// c := cron.New()
	// Run the cron job every day at 12:00 AM
	// c.AddFunc("0 0 0 * * *", s.service.SetExchangeData)
	// c.Start()
}

// InitializeRouter function is for initialize routs
func (s *Server) InitializeRouter() {
	// v1 := s.router.Group("/api/v1")

}

func ginSetMode(s string) {
	if s == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if s == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)

	}
}
