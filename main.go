package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/cactus/config"
	"github.com/leetpy/cactus/model"
	"github.com/leetpy/cactus/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()
	// Routes.
	router.Load(
		// Cores.
		g,

		// // Middlwares.
		// middleware.Logging(),
		// middleware.RequestId(),
	)

	http.ListenAndServe(viper.GetString("addr"), g)
}
