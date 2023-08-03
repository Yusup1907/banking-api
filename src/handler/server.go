package handler

import (
	"github.com/Yusup1907/banking-api/src/manager"
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
	NewNasabahHandler(s.engine, s.serviceManager.GetNasabahService())

	s.engine.Run(":8080")
}

func NewServer() Server {
	filePath := "../data/nasabah.json"
	repository := manager.NewRepositoryManager(filePath)
	if repository == nil {
		// Jika repository belum diinisialisasi dengan benar, return nil atau tangani sesuai kebutuhan Anda
		return nil
	}

	secretKey := "asdkjashd&%adsadbadsaduihj_adsiduasidhasdbasdjliuwhdaskd"

	service := manager.NewServiceManager(repository, secretKey)

	engine := gin.Default()

	return &server{
		serviceManager: service,
		engine:         engine,
	}
}
