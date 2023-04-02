package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	common "test.com/project-common"
	"test.com/project-common/logs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	// 1.获取参数
	mobile := ctx.PostForm("mobile")
	// 2.校验参数
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, rsp.Fail(model.NoLegalMobile, "手机号不合法"))
		return
	}
	// 3.生成验证码 随机4或8位
	code := "123456"
	// 4.调用短信平台 (三方 放入go协程中执行 接口可以快速响应)
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功，发送短信 Info")
		logs.LG.Debug("短信平台调用成功，发送短信 Debug")
		logs.LG.Error("短信平台调用成功，发送短信 Error")
		// redis 假设后续存在mysql or mango 中
		// 5.存储验证码 redis当中 过期时间15分钟
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("验证码登录redis出错, cause by: %v \n", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER:%s : %s", mobile, code)
	}()
	ctx.JSON(http.StatusOK, rsp.Success(code))
}
