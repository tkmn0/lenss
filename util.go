package lenss

import (
	"image"

	"github.com/tkmn0/lenss/vpx"
)

func (img *Image) ImageRGBA() *image.RGBA {
	return (*vpx.Image)(img).ImageRGBA()
}

func (img *Image) ImageYCbCr() *image.YCbCr {
	return (*vpx.Image)(img).ImageYCbCr()
}

func (i *Image) FrameSize() int {
	return ((*vpx.Image)(i)).VpxFrameSize()
}

func (i *Image) YuvPlaneBuffer() []byte {
	return ((*vpx.Image)(i)).VpxPlaneBuffer()
}
