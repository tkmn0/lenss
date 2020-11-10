package lenss

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"image"

	"github.com/tkmn0/lenss/vpx"
)

type File C.FILE

func OpenFile(filePath string, mode string) *File {
	return (*File)(C.fopen(C.CString(filePath), C.CString(mode)))
}

func CloseFile(file *File) {
	C.fclose((*C.FILE)(file))
}

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
