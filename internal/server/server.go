package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShushiAPI struct {
	Addr string // Адрес запуска сервера
}

func NewShushiAPI(addr string) *ShushiAPI {
	return &ShushiAPI{Addr: addr}
}

func (s *ShushiAPI) StartAPI() error {
	router := gin.Default()

	router.GET("/", baseHandler)
	router.POST("/member", addMemberHandler)

	return router.Run(s.Addr)
}

func baseHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Добро пожаловать на сервер ShushiShop")
}

func addMemberHandler(ctx *gin.Context) {
	ctx.String(http.StatusNotFound, "Это закроытое общество, НАХУЙ!")
}
