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
	userUC usecase.UserUseCase
	authUC usecase.AuthUseCase
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
	controller.NewUserController(s.engine, s.userUC)
	controller.NewAuthController(s.engine, s.authUC)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	authRepo := userUseCase
	authUseCase := usecase.NewAuthUseCase(authRepo)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		authUC: authUseCase,
		userUC: userUseCase,
		engine: engine,
		host:   host,
	}
}
