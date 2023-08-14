package delivery

import (
	"employeeleave/config"
	"employeeleave/delivery/controller"
	"employeeleave/delivery/middleware"
	"employeeleave/manager"
	"employeeleave/utils/exceptions"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (s *Server) Run() {
	s.setupControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	// semua controller disini
	controller.NewEmplController(s.engine, s.useCaseManager.EmployeeUseCase())
	controller.NewTransactionController(s.engine, s.useCaseManager.TransactionUseCase())

}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		log:            logrus.New(),
	}
}
