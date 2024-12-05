package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func generateRandomString() (string, error) {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

type GenerateVoiceInfo struct {
	CMD          string `json:"cmd"`
	Output       string `json:"output"`
	Mp3Hex       string `json:"mp3_hex"`
	Subtitleshex string `json:"subtitles_hex"`
}

func (v *Voice) GenerateVoice(text string, rate string) (*GenerateVoiceInfo, error) {
	// Generate a random filename
	randomStr, err := generateRandomString()
	if err != nil {
		return nil, err
	}
	timestamp := time.Now().UnixMilli()
	filename := randomStr + "_" + strconv.FormatInt(timestamp, 10)

	// Write the text to a file
	err = os.WriteFile(filename+".txt", []byte(text), 0644)
	if nil != err {
		return nil, err
	}

	// Run the command general mp3 and subtitles
	cmd := exec.Command(
		"python3.9", "-m",
		"edge_tts",
		"-f", filename+".txt",
		"--voice", v.NameParameter(),
		"--write-media", filename+".mp3",
		"--write-subtitles", filename+".vtt",
		"--proxy", "http://127.0.0.1:1082",
	)
	if rate != "" {
		cmd.Args = append(cmd.Args, "--rate="+rate)
	}

	fmt.Println(cmd.Args)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Read the generated mp3 file
	mp3Buf, err := os.ReadFile(filename + ".mp3")
	if err != nil {
		return nil, err
	}
	mp3BufHex := fmt.Sprintf("%x", mp3Buf)

	// read the generated vtt file
	subtitlesBuf, err := os.ReadFile(filename + ".vtt")
	if err != nil {
		return nil, err
	}
	subtitlesBufHex := fmt.Sprintf("%x", subtitlesBuf)

	return &GenerateVoiceInfo{
		CMD:          cmd.String(),
		Output:       string(output),
		Mp3Hex:       mp3BufHex,
		Subtitleshex: subtitlesBufHex,
	}, nil
}
