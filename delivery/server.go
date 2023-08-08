package delivery

import (
	"employeeleave/config"
	"employeeleave/delivery/controller"
	"employeeleave/repository"
	"employeeleave/usecase"
	"employeeleave/utils/exceptions"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	roleUC usecase.RoleUseCase
	historyUC usecase.HistoryUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	// semua controller disini
	controller.NewUserController(s.engine, s.roleUC)
	controller.NewHistoryController(s.engine, s.historyUC)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	roleRepo := repository.NewRoleRepository(db)
	roleUseCase := usecase.NewRoleUseCase(roleRepo)
	historyRepo := repository.NewHistoryRepository(db)
	historyUseCase := usecase.NewHistoryUseCase(historyRepo)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		roleUC: roleUseCase,
		historyUC: historyUseCase,
		engine: engine,
		host:   host,
	}
}
