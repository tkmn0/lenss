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

// int
// vpx_img_read(vpx_image_t *img, void *bs) {
//   int plane;
//   for (plane = 0; plane < 3; ++plane) {
//     unsigned char *buf = img->planes[plane];
//     const int stride = img->stride[plane];
//     const int w = vpx_img_plane_width(img, plane) *
//                   ((img->fmt & VPX_IMG_FMT_HIGHBITDEPTH) ? 2 : 1);
//     const int h = vpx_img_plane_height(img, plane);
//     int y;
//     for (y = 0; y < h; ++y) {
//       memcpy(buf, bs, w);
//       buf += stride;
//       bs += w;
//     }
//   }
//   return 1;
// }

int
vpx_img_read_from_file(vpx_image_t *img, FILE *file) {
  int plane;

  for (plane = 0; plane < 3; ++plane) {
    unsigned char *buf = img->planes[plane];
    const int stride = img->stride[plane];
    const int w = vpx_img_plane_width(img, plane) *
                  ((img->fmt & VPX_IMG_FMT_HIGHBITDEPTH) ? 2 : 1);
    const int h = vpx_img_plane_height(img, plane);
    int y;

    for (y = 0; y < h; ++y) {
      if (fread(buf, 1, w, file) != (size_t)w) return 0;
      buf += stride;
    }
  }

  return 1;
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
*/
import "C"
import (
	"unsafe"
)

// func VpxImageRead(img *Image, yuv []byte) {
// 	// cimg, cimgAllocMap := img.PassRef()
// 	// runtime.KeepAlive(cimgAllocMap)
// 	// C.vpx_img_read(cimg, unsafe.Pointer(&yuv[0]))
// }

func VpxImageReadFromFile(img *Image, file unsafe.Pointer) int {
	return int(C.vpx_img_read_from_file(img.refc09455e3, (*C.FILE)(file)))
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
