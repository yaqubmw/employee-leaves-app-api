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
	statusLeaveUC usecase.StatusLeaveUseCase
	quotaLeaveUC usecase.QuotaLeaveUseCase
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
	controller.NewStatusLeaveController(s.engine, s.statusLeaveUC)
	controller.NewQuotaLeaveController(s.engine, s.quotaLeaveUC)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	statusLeaveRepo := repository.NewStatusLeaveRepository(db)
	statusLeaveUseCase := usecase.NewStatusLeaveUseCase(statusLeaveRepo)
	qoutaRepo := repository.NewQuotaLeaveRepository(db)
	qoutaUseCase := usecase.NewQuotaLeaveUseCase(qoutaRepo)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		statusLeaveUC: statusLeaveUseCase,
		quotaLeaveUC:  qoutaUseCase,
		engine:        engine,
		host:          host,
	}
}
