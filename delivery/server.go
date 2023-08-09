package delivery

// import (
// 	"fmt"

// 	"employeeleave/config"
// 	"employeeleave/delivery/controller/api"
// 	"employeeleave/delivery/middleware"
// 	"employeeleave/utils/exceptions"

// 	"employeeleave/manager"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// )

// type Server struct {
// 	useCaseManager manager.UseCaseManager
// 	engine         *gin.Engine
// 	host           string
// 	log            *logrus.Logger
// }

// func (s *Server) Run() {
// 	s.setupControllers()
// 	err := s.engine.Run(s.host)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (s *Server) setupControllers() {
// 	s.engine.Use(middleware.LogRequestMiddleware(s.log))
// 	api.NewEmployeeController(s.engine, s.useCaseManager.EmployeeUseCase())

// }

// func NewServer() *Server {
// 	cfg, err := config.NewConfig()
// 	exceptions.CheckError(err)
// 	infraManager, _ := manager.NewInfraManager(cfg)
// 	repoManager := manager.NewRepoManager(infraManager)
// 	useCaseManager := manager.NewUseCaseManager(repoManager)
// 	engine := gin.Default()
// 	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
// 	return &Server{
// 		useCaseManager: useCaseManager,
// 		engine:         engine,
// 		host:           host,
// 		log:            logrus.New(),
// 	}
// }
