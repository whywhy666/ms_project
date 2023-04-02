package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Route(r *gin.Engine)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
