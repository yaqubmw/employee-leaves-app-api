package delivery

import (
	"employeeleave/config"
	"employeeleave/delivery/controller/api"
	"employeeleave/repository"
	"employeeleave/usecase"
	"employeeleave/utils/exceptions"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	emplUC usecase.EmplUseCase
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
	api.NewEmplController(s.emplUC, s.engine)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	emplRepo := repository.NewEmplRepository(db)
	emplUseCase := usecase.NewEmplUseCase(emplRepo)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		emplUC: emplUseCase,
		engine: engine,
		host:   host,
	}
}
