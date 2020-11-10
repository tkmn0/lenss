package lenss

import (
	"image"

	"github.com/tkmn0/lenss/vpx"
)

func (i *Image) ImageRGBA() *image.RGBA {
	return (*vpx.Image)(i).ImageRGBA()
}

func (i *Image) ImageYCbCr() *image.YCbCr {
	return (*vpx.Image)(i).ImageYCbCr()
}

func (i *Image) FrameSize() int {
	return ((*vpx.Image)(i)).VpxFrameSize()
}

func (i *Image) YuvPlaneBuffer() []byte {
	return ((*vpx.Image)(i)).VpxPlaneBuffer()
}
