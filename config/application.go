package config

import (
	"os"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/yiziz/gin-template/config/initializers"
	"github.com/yiziz/gin-template/services/env"
)

// AppEnv determines what env the app is running in
var AppEnv string

// SetAppEnv sets the server's envirnoment var
func SetAppEnv() {
	AppEnv = os.Getenv("GIN_ENV")
	if AppEnv == "" {
		AppEnv = "development"
	}
}

func setAppMode() {
	switch AppEnv {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	}
}

// Run spins up the app
func Run() {
	env.SetEnvVars()
	SetAppEnv()
	setAppMode()

	jobrunner.Start()

	r := gin.Default()
	initializers.Initialize(r, AppEnv)
	if AppEnv == "development" {
		r.Run(":8000")
	} else {
		r.Run(":80")
	}
}
