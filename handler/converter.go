package handler

import (
	"encoding/base64"
	"encoding/json"
	"github.com/cnrywjd11/online-audio-converter/pkg/ffmpeg"
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

func ConvertAudioHandler(c echo.Context) error {
	audio, err := c.FormFile("audio")
	if err != nil {
		return err
	}

	// Source
	src, err := audio.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(audio.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	defer os.Remove(dst.Name())

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	in := dst
	outName := "new_" + in.Name()

	ofm, err := GetOutputFormat(&c)
	if err != nil {
		return err
	}

	ac := &ffmpeg.AudioConverter{}
	out, err := ac.Process(ofm, in, outName)
	if err != nil {
		return err
	}
	defer os.Remove(out.Name())

	return c.File(out.Name())
}

func GetOutputFormat(c *echo.Context) (*ffmpeg.OutputFormat, error) {
	aaf := (*c).Request().Header.Get("Accept-Audio-Format")
	b, err := base64.StdEncoding.DecodeString(aaf)
	if err != nil {
		return nil, err
	}
	var ofm ffmpeg.OutputFormat
	if err := json.Unmarshal(b, &ofm); err != nil {
		return nil, err
	}
	return &ofm, nil
}
