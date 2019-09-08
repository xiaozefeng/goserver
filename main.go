package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/xiaozefeng/goserver/config"
	"github.com/xiaozefeng/goserver/model"
	"github.com/xiaozefeng/goserver/router"
	"github.com/xiaozefeng/goserver/router/middleware"

	"github.com/lexkong/log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "server config file path")
)

func main() {
	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	model.DB.Init()

	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine
	g := gin.New()

	// load routers
	router.Load(g,
		middleware.Logging(),
		middleware.RequestId(),
	)

	// ping server
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully")
	}()

	// start server
	port := viper.GetString("port")
	log.Infof("Start to listening the incoming requests on http address: %s", port)
	log.Infof(http.ListenAndServe(port, g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the Server by sending a GET request to `/health`
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second * 1)
	}

	return errors.New("can not connect to the router")
}
