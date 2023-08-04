package handler

import (
	"github.com/Yusup1907/banking-api/src/manager"
	"github.com/Yusup1907/banking-api/src/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	serviceManager manager.ServiceManager
	engine         *gin.Engine
}

func (s *server) Run() {
	store := cookie.NewStore([]byte(utils.KEY))
	s.engine.Use(sessions.Sessions("session", store))

	// Handler
	NewAuthenticationHandler(s.engine, s.serviceManager.GetNasabahService())
	NewNasabahHandler(s.engine, s.serviceManager.GetNasabahService())

	s.engine.Run(":8080")
}

func NewServer() Server {
	repository := manager.NewRepositoryManager()
	if repository == nil {
		// Jika repository belum diinisialisasi dengan benar, return nil atau tangani sesuai kebutuhan Anda
		return nil
	}

	service := manager.NewServiceManager(repository)

	engine := gin.Default()

	return &server{
		serviceManager: service,
		engine:         engine,
	}
}
