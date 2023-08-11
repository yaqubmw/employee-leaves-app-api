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
	"gorm.io/gorm"
)

type Server struct {
	db             *gorm.DB
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

	// All controllers here
	controller.NewLeaveTypeController(s.useCaseManager.LeaveTypeUseCase(), s.engine)
	controller.NewPositionController(s.useCaseManager.PositionUseCase(), s.engine)
	controller.NewStatusLeaveController(s.engine, s.useCaseManager.StatusLeaveUseCase())
	controller.NewQuotaLeaveController(s.engine, s.useCaseManager.QuotaLeaveUseCase())
	controller.NewRoleController(s.engine, s.useCaseManager.RoleUseCase())
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
