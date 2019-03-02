package httpsvr

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	"kbtg.tech/template/interfaces"
	"kbtg.tech/template/share"
	golog "log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type HttpServer struct {
	Config  *share.Config
	Service interfaces.IPCBService
}

func (server *HttpServer) Run() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//	urlPrefix := ""
	//	r.GET(urlPrefix+"/version", func(c *gin.Context) {
	//		c.JSON(200, map[string]string{
	//			"build":   buildstamp,
	//			"githash": githash,
	//		})
	//	})

	//r.POST("/url1", server.Service.RETRCOMACCTDTLSMW(""))

	srv := &http.Server{
		Addr:    ":" + server.Config.HttpPort,
		Handler: r,
	}

	go func() {
		golog.Println("Starting HTTP", server.Config.HttpPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			golog.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	golog.Println("Shutdown HttpServer ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		golog.Fatal("HttpServer Shutdown:", err)
	}
	defer cancel()

	return r
}
