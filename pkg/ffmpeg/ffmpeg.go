package ffmpeg

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const converterCommandName = "ffmpeg"

type OutputFormat struct {
	Codec        string `json:"codec"`
	SamplingRate int    `json:"samplingrate"`
	Channel      int    `json:"channel"`
	Bitrate      string `json:"bitrate"`
}

func generateArgs(ofm *OutputFormat, in *os.File, outName string) []string {
	args := make([]string, 0)
	args = append(args, "-i", in.Name())
	if ofm.SamplingRate != 0 {
		args = append(args, "-ar", fmt.Sprint(ofm.SamplingRate))
	}
	if ofm.Channel != 0 {
		args = append(args, "-ac", fmt.Sprint(ofm.Channel))
	}
	if ofm.Bitrate != "" {
		args = append(args, "-b:a", ofm.Bitrate)
	}
	if ofm.Codec != "" {
		outName = convertedFileName(outName, ofm.Codec)
	}
	args = append(args, outName)
	return args
}

func convertedFileName(fileName, extension string) string {
	return strings.Split(fileName, ".")[0] + "." + strings.ToLower(extension)
}

type AudioConverter struct{}

func (ac *AudioConverter) Process(ofm *OutputFormat, in *os.File, outName string) (*os.File, error) {
	args := generateArgs(ofm, in, outName)

	err := processCommand(args)
	if err != nil {
		return nil, fmt.Errorf("process: %s", err.Error())
	}
	out, err := os.Open(convertedFileName(outName, ofm.Codec))
	if err != nil {
		return nil, fmt.Errorf("process: %s", err.Error())
	}
	return out, nil
}

func processCommand(args []string) error {
	procCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(procCtx, converterCommandName, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("processCommand: %s\nout: %s\nstderr: %s", err.Error(), out.String(), stderr.String())
	}
	return nil
}
