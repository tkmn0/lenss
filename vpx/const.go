// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 06 Nov 2020 05:41:36 JST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package vpx

/*
#cgo pkg-config: vpx
#include <vpx/vpx_encoder.h>
#include <vpx/vpx_decoder.h>
#include <vpx/vp8.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// Vp8Fourcc as defined in lenss/<predefine>:24
	Vp8Fourcc = 808996950
	// Vp9Fourcc as defined in lenss/<predefine>:25
	Vp9Fourcc = 809062486
	// TsMaxPeriodicity as defined in vpx/vpx_encoder.h:37
	TsMaxPeriodicity = 16
	// TsMaxLayers as defined in vpx/vpx_encoder.h:40
	TsMaxLayers = 5
	// MaxLayers as defined in vpx/vpx_encoder.h:43
	MaxLayers = 12
	// SsMaxLayers as defined in vpx/vpx_encoder.h:46
	SsMaxLayers = 5
	// SsDefaultLayers as defined in vpx/vpx_encoder.h:49
	SsDefaultLayers = 1
	// EncoderABIVersion as defined in vpx/vpx_encoder.h:59
	EncoderABIVersion = 23
	// CodecCapPsnr as defined in vpx/vpx_encoder.h:71
	CodecCapPsnr = 65536
	// CodecCapOutputPartition as defined in vpx/vpx_encoder.h:78
	CodecCapOutputPartition = 131072
	// CodecUsePsnr as defined in vpx/vpx_encoder.h:87
	CodecUsePsnr = 65536
	// CodecUseOutputPartition as defined in vpx/vpx_encoder.h:89
	CodecUseOutputPartition = 131072
	// CodecUseHighbitdepth as defined in vpx/vpx_encoder.h:90
	CodecUseHighbitdepth = 262144
	// FrameIsKey as defined in vpx/vpx_encoder.h:116
	FrameIsKey = 1
	// FrameIsDroppable as defined in vpx/vpx_encoder.h:119
	FrameIsDroppable = 2
	// FrameIsInvisible as defined in vpx/vpx_encoder.h:121
	FrameIsInvisible = 4
	// FrameIsFragment as defined in vpx/vpx_encoder.h:123
	FrameIsFragment = 8
	// ErrorResilientDefault as defined in vpx/vpx_encoder.h:133
	ErrorResilientDefault = 1
	// ErrorResilientPartitions as defined in vpx/vpx_encoder.h:138
	ErrorResilientPartitions = 2
	// EflagForceKf as defined in vpx/vpx_encoder.h:260
	EflagForceKf = 1
	// DlRealtime as defined in vpx/vpx_encoder.h:830
	DlRealtime = 1
	// DlGoodQuality as defined in vpx/vpx_encoder.h:832
	DlGoodQuality = 1000000
	// DlBestQuality as defined in vpx/vpx_encoder.h:834
	DlBestQuality = 0
	// DeclspecDeprecated as defined in vpx/vpx_codec.h:64

	// CodecABIVersion as defined in vpx/vpx_codec.h:90
	CodecABIVersion = 9
	// CodecCapDecoder as defined in vpx/vpx_codec.h:156
	CodecCapDecoder = 1
	// CodecCapEncoder as defined in vpx/vpx_codec.h:157
	CodecCapEncoder = 2
	// CodecCapHighbitdepth as defined in vpx/vpx_codec.h:161
	CodecCapHighbitdepth = 4
	// ImageABIVersion as defined in vpx/vpx_image.h:30
	ImageABIVersion = 5
	// ImageFormatPlanar as defined in vpx/vpx_image.h:32
	ImageFormatPlanar = 256
	// ImageFormatUvFlip as defined in vpx/vpx_image.h:33
	ImageFormatUvFlip = 512
	// ImageFormatHasAlpha as defined in vpx/vpx_image.h:34
	ImageFormatHasAlpha = 1024
	// ImageFormatHighbitdepth as defined in vpx/vpx_image.h:35
	ImageFormatHighbitdepth = 2048
	// PlanePacked as defined in vpx/vpx_image.h:95
	PlanePacked = 0
	// PlaneY as defined in vpx/vpx_image.h:96
	PlaneY = 0
	// PlaneU as defined in vpx/vpx_image.h:97
	PlaneU = 1
	// PlaneV as defined in vpx/vpx_image.h:98
	PlaneV = 2
	// PlaneAlpha as defined in vpx/vpx_image.h:99
	PlaneAlpha = 3
	// Inline as defined in vpx/vpx_integer.h:23
	Inline = 0
	// DecoderABIVersion as defined in vpx/vpx_decoder.h:43
	DecoderABIVersion = 12
	// CodecCapPutSlice as defined in vpx/vpx_decoder.h:54
	CodecCapPutSlice = 65536
	// CodecCapPutFrame as defined in vpx/vpx_decoder.h:55
	CodecCapPutFrame = 131072
	// CodecCapPostproc as defined in vpx/vpx_decoder.h:56
	CodecCapPostproc = 262144
	// CodecCapErrorConcealment as defined in vpx/vpx_decoder.h:58
	CodecCapErrorConcealment = 524288
	// CodecCapInputFragments as defined in vpx/vpx_decoder.h:60
	CodecCapInputFragments = 1048576
	// CodecCapFrameThreading as defined in vpx/vpx_decoder.h:62
	CodecCapFrameThreading = 2097152
	// CodecCapExternalFrameBuffer as defined in vpx/vpx_decoder.h:64
	CodecCapExternalFrameBuffer = 4194304
	// CodecUsePostproc as defined in vpx/vpx_decoder.h:73
	CodecUsePostproc = 65536
	// CodecUseErrorConcealment as defined in vpx/vpx_decoder.h:75
	CodecUseErrorConcealment = 131072
	// CodecUseInputFragments as defined in vpx/vpx_decoder.h:78
	CodecUseInputFragments = 262144
	// CodecUseFrameThreading as defined in vpx/vpx_decoder.h:80
	CodecUseFrameThreading = 524288
	// MaximumWorkBuffers as defined in vpx/vpx_frame_buffer.h:29
	MaximumWorkBuffers = 8
	// Vp9MaximumRefBuffers as defined in vpx/vpx_frame_buffer.h:33
	Vp9MaximumRefBuffers = 8
)

// CodecCxPktKind as declared in vpx/vpx_encoder.h:146
type CodecCxPktKind int32

// CodecCxPktKind enumeration from vpx/vpx_encoder.h:146
const (
	CodecCxFramePkt   = C.VPX_CODEC_CX_FRAME_PKT
	CodecStatsPkt     = C.VPX_CODEC_STATS_PKT
	CodecFpmbStatsPkt = C.VPX_CODEC_FPMB_STATS_PKT
	CodecPsnrPkt      = C.VPX_CODEC_PSNR_PKT
	CodecCustomPkt    = C.VPX_CODEC_CUSTOM_PKT
)

// EncPass as declared in vpx/vpx_encoder.h:228
type EncPass int32

// EncPass enumeration from vpx/vpx_encoder.h:228
const (
	RcOnePass   EncPass = C.VPX_RC_ONE_PASS
	RcFirstPass EncPass = C.VPX_RC_FIRST_PASS
	RcLastPass  EncPass = C.VPX_RC_LAST_PASS
)

// RcMode as declared in vpx/vpx_encoder.h:231
type RcMode int32

// RcMode enumeration from vpx/vpx_encoder.h:231
const (
	Vbr = C.VPX_VBR
	Cbr = C.VPX_CBR
	Cq  = C.VPX_CQ
	Q   = C.VPX_Q
)

// KfMode as declared in vpx/vpx_encoder.h:246
type KfMode int32

// KfMode enumeration from vpx/vpx_encoder.h:246
const (
	KfFixed    = C.VPX_KF_FIXED
	KfAuto     = C.VPX_KF_AUTO
	KfDisabled = C.VPX_KF_DISABLED
)

// BitDepth as declared in vpx/vpx_codec.h:224
type BitDepth int32

// BitDepth enumeration from vpx/vpx_codec.h:224
const (
	Bits8  BitDepth = C.VPX_BITS_8
	Bits10 BitDepth = C.VPX_BITS_10
	Bits12 BitDepth = C.VPX_BITS_12
)

// ImageFormat as declared in vpx/vpx_image.h:51
type ImageFormat int32

// ImageFormat enumeration from vpx/vpx_image.h:51
const (
	ImageFormatNone   ImageFormat = C.VPX_IMG_FMT_NONE
	ImageFormatYv12   ImageFormat = C.VPX_IMG_FMT_YV12
	ImageFormatI420   ImageFormat = C.VPX_IMG_FMT_I420
	ImageFormatI422   ImageFormat = C.VPX_IMG_FMT_I422
	ImageFormatI444   ImageFormat = C.VPX_IMG_FMT_I444
	ImageFormatI440   ImageFormat = C.VPX_IMG_FMT_I440
	ImageFormatNv12   ImageFormat = C.VPX_IMG_FMT_NV12
	ImageFormatI42016 ImageFormat = C.VPX_IMG_FMT_I42016
	ImageFormatI42216 ImageFormat = C.VPX_IMG_FMT_I42216
	ImageFormatI44416 ImageFormat = C.VPX_IMG_FMT_I44416
	ImageFormatI44016 ImageFormat = C.VPX_IMG_FMT_I44016
)

// ColorSpace as declared in vpx/vpx_image.h:63
type ColorSpace int32

// ColorSpace enumeration from vpx/vpx_image.h:63
const (
	ColorSpaceUnknown  ColorSpace = C.VPX_CS_UNKNOWN
	ColorSpaceBt601    ColorSpace = C.VPX_CS_BT_601
	ColorSpaceBt709    ColorSpace = C.VPX_CS_BT_709
	ColorSpaceSmpte170 ColorSpace = C.VPX_CS_SMPTE_170
	ColorSpaceSmpte240 ColorSpace = C.VPX_CS_SMPTE_240
	ColorSpaceBt2020   ColorSpace = C.VPX_CS_BT_2020
	ColorSpaceReserved ColorSpace = C.VPX_CS_RESERVED
	ColorSpaceSrgb     ColorSpace = C.VPX_CS_SRGB
)

// ColorRange as declared in vpx/vpx_image.h:69
type ColorRange int32

// ColorRange enumeration from vpx/vpx_image.h:69
const (
	CrStudioRange ColorRange = C.VPX_CR_STUDIO_RANGE
	CrFullRange   ColorRange = C.VPX_CR_FULL_RANGE
)

// CodecErr as declared in vpx/vpx_codec.h:145
type CodecErr int32

// CodecErr enumeration from vpx/vpx_codec.h:145
const (
	CodecOk             CodecErr = C.VPX_CODEC_OK
	CodecError          CodecErr = C.VPX_CODEC_ERROR
	CodecMemError       CodecErr = C.VPX_CODEC_MEM_ERROR
	CodecABIMismatch    CodecErr = C.VPX_CODEC_ABI_MISMATCH
	CodecIncapable      CodecErr = C.VPX_CODEC_INCAPABLE
	CodecUnsupBitstream CodecErr = C.VPX_CODEC_UNSUP_BITSTREAM
	CodecUnsupFeature   CodecErr = C.VPX_CODEC_UNSUP_FEATURE
	CodecCorruptFrame   CodecErr = C.VPX_CODEC_CORRUPT_FRAME
	CodecInvalidParam   CodecErr = C.VPX_CODEC_INVALID_PARAM
	CodecListEnd        CodecErr = C.VPX_CODEC_LIST_END
)
