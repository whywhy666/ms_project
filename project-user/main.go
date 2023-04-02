package main

import (
	"github.com/gin-gonic/gin"
	common "test.com/project-common"
	_ "test.com/project-user/api"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	common.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
