package user

import (
	"github.com/gin-gonic/gin"
	"test.com/project-user/router"
)

func init() {
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
