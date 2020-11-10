package lenss

import (
	"github.com/tkmn0/lenss/vpx"
)

type VpxDecoder struct {
	ctx    *vpx.CodecCtx
	iface  *vpx.CodecIface
	Srouce chan []byte
	Output chan *Image
}

func NewVpxDecoder(codec VCodec) (*VpxDecoder, error) {
	dec := &VpxDecoder{
		ctx:    vpx.NewCodecCtx(),
		Srouce: make(chan []byte),
		Output: make(chan *Image),
	}
	switch codec {
	case CodecVP8:
		dec.iface = vpx.DecoderIfaceVP8()
	case CodecVP9:
		dec.iface = vpx.DecoderIfaceVP9()
	default:
		return nil, errorInitializeInvalidCodec
	}
	err := vpx.Error(vpx.CodecDecInitVer(dec.ctx, dec.iface, nil, 0, vpx.DecoderABIVersion))
	if err != nil {
		return nil, errorCreateVpxImage
	}
	return dec, nil
}

func (v *VpxDecoder) Process() {
	go func() {
		for {
			src := <-v.Srouce
			dataSize := uint32(len(src))
			if len(src) <= 0 {
				return
			}

			err := vpx.Error(vpx.CodecDecode(v.ctx, string(src), dataSize, nil, vpx.DlRealtime))
			if err != nil {
				return
			}

			var iter *vpx.CodecIter
			img := vpx.CodecGetFrame(v.ctx, &iter)
			for img != nil {
				img.Deref()
				v.Output <- (*Image)(img)
				img = vpx.CodecGetFrame(v.ctx, &iter)
			}
		}
	}()
}
