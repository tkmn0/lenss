package vpx

/*
#cgo pkg-config: vpx
#include <vpx/vpx_encoder.h>
#include <vpx/vpx_decoder.h>
#include <vpx/vp8.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdio.h>

int
vpx_img_plane_width(const vpx_image_t *img, int plane) {
  if (plane > 0 && img->x_chroma_shift > 0)
    return (img->d_w + 1) >> img->x_chroma_shift;
  else
    return img->d_w;
}

int
vpx_img_plane_height(const vpx_image_t *img, int plane) {
  if (plane > 0 && img->y_chroma_shift > 0)
    return (img->d_h + 1) >> img->y_chroma_shift;
  else
    return img->d_h;
}

void
vpx_img_write(const vpx_image_t *img, FILE *file) {
  int plane;

  for (plane = 0; plane < 3; ++plane) {
    const unsigned char *buf = img->planes[plane];
    const int stride = img->stride[plane];
    const int w = vpx_img_plane_width(img, plane) *
                  ((img->fmt & VPX_IMG_FMT_HIGHBITDEPTH) ? 2 : 1);
    const int h = vpx_img_plane_height(img, plane);
    int y;

    for (y = 0; y < h; ++y) {
      fwrite(buf, 1, w, file);
      buf += stride;
    }
  }
}

int
vpx_img_read(vpx_image_t *img, void *bs) {
  int plane;

  for (plane = 0; plane < 3; ++plane) {
    unsigned char *buf = img->planes[plane];
    const int stride = img->stride[plane];
    const int w = vpx_img_plane_width(img, plane) *
                  ((img->fmt & VPX_IMG_FMT_HIGHBITDEPTH) ? 2 : 1);
    const int h = vpx_img_plane_height(img, plane);
    int y;
    for (y = 0; y < h; ++y) {
      memcpy(buf, bs, w);
      buf += stride;
      bs += w;
    }
  }
  return 1;
}

int
vpx_image_height(vpx_image_t *img, int plane) {
  return vpx_img_plane_height(img, plane);
}

int
vpx_image_stride(vpx_image_t *img, int plane){
  return img->stride[plane];
}

typedef struct Bytes {
  void *bs;
  int size;
} BytesType;

BytesType read_frame_data(vpx_codec_cx_pkt_t *pkt) {
	BytesType bytes = {NULL, 0};
	bytes.bs = pkt->data.frame.buf;
	bytes.size = pkt->data.frame.sz;
  return bytes;
}

BytesType vpx_plane_buffer(vpx_image_t *img, int plane, int h){
  BytesType bytes = {NULL, 0};
  unsigned char *buf = img->planes[plane];
  const int stride = img->stride[plane];
  const int w = vpx_img_plane_width(img, plane) *
                  ((img->fmt & VPX_IMG_FMT_HIGHBITDEPTH) ? 2 : 1);
  buf += stride * h;
  bytes.bs = buf;
  bytes.size = w;
  return bytes;
}
*/
import "C"
import (
	"unsafe"
)

func VpxImageRead(img *Image, yuv []byte) {
	C.vpx_img_read(img.refc09455e3, unsafe.Pointer(&yuv[0]))
}

func (img *Image) VpxPlaneBuffer() []byte {
	buffer := []byte{}
	for i := 0; i < 3; i++ {
		h := int(C.vpx_image_height(img.refc09455e3, C.int(i)))

		for y := 0; y < h; y++ {
			data := C.vpx_plane_buffer(img.refc09455e3, C.int(i), C.int(y))
			if data.bs != nil {
				buffer = append(buffer, C.GoBytes(data.bs, data.size)...)
			}
		}
	}

	return buffer
}

func VpxFrameSize(img *Image) int {
	frame := 0
	for plane := 0; plane < 3; plane++ {
		h := int(C.vpx_image_height(img.refc09455e3, C.int(plane)))
		s := int(C.vpx_image_stride(img.refc09455e3, C.int(plane)))
		frame += h * s
	}
	return frame
}

func VpxGetBufferFromPkt(pkt *CodecCxPkt) []byte {
	data := C.read_frame_data(pkt.refa671fc83)
	if data.bs == nil {
		return nil
	}
	return C.GoBytes(data.bs, data.size)
}

func VpxImageWrite(img *Image, file unsafe.Pointer) {
	C.vpx_img_write(img.refc09455e3, (*C.FILE)(file))
}
