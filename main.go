package main

import (
	"os"
	"strings"
	"tts/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/listAllVoice", api.GetAllTTSVoices)
	r.POST("/generalMp3", api.GenerateVoice)

	port := ":19020"
	if len(os.Args) > 1 {
		port = os.Args[1]
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
	}

	r.Run(port)
}
