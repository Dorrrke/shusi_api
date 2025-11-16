package server

import (
	"github.com/Dorrrke/shusi_api/internal/models"
	"github.com/Dorrrke/shusi_api/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShushiAPI struct {
	Addr string
}

func NewShushiAPI(addr string) *ShushiAPI {
	return &ShushiAPI{Addr: addr}
}

func (s *ShushiAPI) Start() error {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.POST("/member", addMemberHandler)

	return router.Run(s.Addr)
}

func addMemberHandler(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.NewString()

	storage.NewUserStorage().AddUser(user)

	ctx.JSON(200, user)
}
