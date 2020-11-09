package lenss

import (
	"errors"
	"image"
	"time"

	"github.com/tkmn0/lenss/vpx"
)

type VCodec string
type Image vpx.Image
type Frame struct {
	*image.RGBA
	Timecode   time.Duration
	IsKeyframe bool
}

const (
	CodecVP8 VCodec = "vp8"
	CodecVP9 VCodec = "vp9"
)

var (
	errorInitializeInvalidCodec = errors.New("invalid codec initialized error")
	errorCreateVpxImage         = errors.New("error to create vpx image")
)
