package api

import (
	"net/http"
	"tts/util"

	"github.com/gin-gonic/gin"
)

func GetAllTTSVoices(c *gin.Context) {
	voices, err := (&util.TTS{}).ListVoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := []map[string]string{}
	for _, voice := range voices {
		result = append(result, map[string]string{
			"name":     voice.Name(),
			"language": voice.Language(),
			"country":  voice.Country(),
			"general":  voice.General(),
			"id":       voice.NameParameter(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func GenerateVoice(c *gin.Context) {
	var requestBody struct {
		Voice string `json:"voiceId"`
		Text  string `json:"text"`
		Speed string `json:"speed"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	voice := requestBody.Voice
	text := requestBody.Text
	speed := requestBody.Speed

	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text is required"})
		return
	}
	if voice == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voice is required"})
		return
	}
	voiceObj := util.VoiceFromRawString(voice, "")
	ret, err := voiceObj.GenerateVoice(text, speed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"output": ret})
}
