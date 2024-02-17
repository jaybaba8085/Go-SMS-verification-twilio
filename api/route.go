package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (config *Config) Routes() {
	config.Router.POST("/otp", config.sendSMS())
	config.Router.POST("/verifyOTP", config.verifySMS())
}