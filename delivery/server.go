package delivery

import (
	"employeeleave/config"
	"employeeleave/delivery/controller"
	"employeeleave/manager"
	"employeeleave/utils/exceptions"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine *gin.Engine
	host   string
}

func (s *Server) Run() {
	s.setupControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) setupControllers() {
	// semua controller disini
	controller.NewRoleController(s.engine, s.useCaseManager.RoleUseCase())
	controller.NewHistoryController(s.engine, s.useCaseManager.HistoryUseCase())
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
		engine: engine,
		host:   host,
	}
}
