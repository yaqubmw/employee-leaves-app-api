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
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))

	// semua controller disini
	// controller.NewLeaveTypeController(s.useCaseManager.LeaveTypeUseCase(), s.engine)
	// controller.NewPositionController(s.useCaseManager.PositionUseCase(), s.engine)
	controller.NewUserController(s.engine, s.useCaseManager.UserUseCase())
	controller.NewAuthController(s.engine, s.useCaseManager.AuthUseCase())
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
