package lenss

import (
	"github.com/tkmn0/lenss/vpx"
)

type VpxEncoder struct {
	ctx              *vpx.CodecCtx
	iface            *vpx.CodecIface
	img              vpx.Image
	cfg              *vpx.CodecEncCfg
	frameCount       int
	keyframeInterval int
	Input            chan []byte
	Output           chan []byte
}

func NewVpxEncoder(codec VCodec, width int, height int, fps int, bitrate int, keyframe int) (*VpxEncoder, error) {
	enc := &VpxEncoder{
		ctx:              vpx.NewCodecCtx(),
		frameCount:       0,
		keyframeInterval: keyframe,
		Input:            make(chan []byte),
		Output:           make(chan []byte),
	}

	w := uint32(width)
	h := uint32(height)

	switch codec {
	case CodecVP8:
		enc.iface = vpx.EncoderIfaceVP8()
	case CodecVP9:
		enc.iface = vpx.EncoderIfaceVP9()
	}

	if enc.iface == nil {
		return nil, errorInitializeInvalidCodec
	}

	if vpx.ImageAlloc(&enc.img, vpx.ImageFormatI420, w, h, 0) == nil {
		return nil, errorCreateVpxImage
	}

	enc.cfg = &vpx.CodecEncCfg{}
	err := vpx.Error(vpx.CodecEncConfigDefault(enc.iface, enc.cfg, 0))
	if err != nil {
		return nil, err
	}

	enc.cfg.GW = w
	enc.cfg.GH = h
	enc.cfg.GTimebase.Num = 1
	enc.cfg.GTimebase.Den = int32(fps)
	enc.cfg.RcTargetBitrate = uint32(bitrate)
	enc.cfg.GErrorResilient = 1
	enc.cfg.Load()

	err = vpx.Error(vpx.CodecEncInitVer(enc.ctx, enc.iface, enc.cfg, 0, vpx.EncoderABIVersion))
	if err != nil {
		return nil, err
	}

	return enc, nil
}

func (e *VpxEncoder) Image() *Image {
	return (*Image)(&e.img)
}

func (e *VpxEncoder) Process() {
	go func() {
		for {
			yuv := <-e.Input
			(&e.img).VpxImageRead(yuv)

			var flags vpx.EncFrameFlags
			var iter *vpx.CodecIter

			if e.keyframeInterval > 0 && e.frameCount%e.keyframeInterval == 0 {
				flags |= vpx.EflagForceKf
			}

			err := vpx.Error(vpx.CodecEncode(e.ctx, &e.img, vpx.CodecPts(e.frameCount), 1, flags, vpx.DlRealtime))
			if err != nil {
				return
			}
			e.frameCount++

			for {
				pkt := vpx.CodecGetCxData(e.ctx, &iter)
				if pkt != nil && pkt.Kind == vpx.CodecCxFramePkt {
					buffer := vpx.VpxGetBufferFromPkt(pkt)
					if buffer == nil {
						break
					}
					e.Output <- buffer
				} else {
					break
				}
			}
		}
	}()
}
