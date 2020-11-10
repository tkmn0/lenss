package lenss

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"image"
	"unsafe"

	"github.com/tkmn0/lenss/vpx"
)

type File C.FILE

func OpenFile(filePath string, mode string) *File {
	return (*File)(C.fopen(C.CString(filePath), C.CString(mode)))
}

func CloseFile(file *File) {
	C.fclose((*C.FILE)(file))
}

func WriteVpxImage(img *Image, file *File) {
	vpx.VpxImageWrite((*vpx.Image)(img), unsafe.Pointer(file))
}

func (img *Image) ImageRGBA() *image.RGBA {
	return (*vpx.Image)(img).ImageRGBA()
}

func (img *Image) ImageYCbCr() *image.YCbCr {
	return (*vpx.Image)(img).ImageYCbCr()
}

func FrameSize(img *Image) int {
	return vpx.VpxFrameSize((*vpx.Image)(img))
}

func YuvPlaneBuffer(img *Image) []byte {
	return vpx.VpxPlaneBuffer((*vpx.Image)(img))
}
