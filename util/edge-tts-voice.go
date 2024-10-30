package util

import (
	"os/exec"
	"strings"
)

type TTS struct{}

type Voice struct {
	NameStr    string
	GeneralStr string
}

func VoiceFromRawString(id, general string) Voice {
	return Voice{
		NameStr:    "Name: " + id,
		GeneralStr: "Gender: " + general,
	}
}

func (v *Voice) NameParameter() string {
	return strings.TrimPrefix(strings.TrimSpace(v.NameStr), "Name: ")
}

func (v *Voice) Name() string {
	parts := strings.Split(v.NameStr, "-")
	if len(parts) < 3 {
		return v.NameStr
	}
	return parts[2]
}

func (v *Voice) Country() string {
	parts := strings.Split(v.NameStr, "-")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

func (v *Voice) Language() string {
	parts := strings.Split(v.NameStr, "-")
	if len(parts) < 1 {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(parts[0], "Name:"))
}

func (v *Voice) General() string {
	return strings.TrimSpace(strings.TrimPrefix(v.GeneralStr, "Gender:"))
}

// run linux cmd `python3.9 -m edge_tts -l` and return the output
func (t *TTS) ListVoices() ([]Voice, error) {
	cmd := exec.Command("python3.9", "-m", "edge_tts", "-l")
	// cmd := exec.Command("edge_tts", "-l")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	line := 0
	rets := []Voice{}
	for line < len(lines) {
		nameLine := lines[line]
		genderLine := lines[line+1]
		rets = append(rets, Voice{
			NameStr:    nameLine,
			GeneralStr: genderLine,
		})
		line += 3
	}

	return rets, nil
}
