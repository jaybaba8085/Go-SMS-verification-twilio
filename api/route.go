package api

import "github.com/gin-gonic/gin"

type Config struct {
	Server *gin.Engine
}

func (config *Config) Routes() {
	config.Server.POST("/otp", config.sendSMS())
	config.Server.POST("/verifyOTP", config.verifySMS())
}