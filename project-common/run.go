package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, svrName string, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	//保证下面的优雅启停
	go func() {
		log.Printf("web server running in %s \n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发 kill -2
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project %s...\n", svrName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown, cause by : %v", svrName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout")
	}
	log.Printf("%s stop success...\n", svrName)
}
