package api

import (
	"unsafe"

	"github.com/tkmn0/lenss"
	"github.com/tkmn0/spp-go/spp"
)

var encoders []*lenss.VpxEncoder
var decoders []*lenss.VpxDecoder
var encodedQueues map[uintptr]*spp.Queue
var decodedQueues map[uintptr]*spp.Queue
var converter *spp.Converter

func Initialize() {
	encoders = make([]*lenss.VpxEncoder, 0)
	decoders = make([]*lenss.VpxDecoder, 0)
	encodedQueues = make(map[uintptr]*spp.Queue)
	converter = spp.NewConverter()
}

func CreateEncoder(codec string, w int, h int, fps int, bitrate int, keyframeInterval int) uintptr {
	enc, err := lenss.NewEncoder(lenss.VCodec(codec), w, h, bitrate, fps, keyframeInterval)
	if err != nil {
		return 0
	}
	enc.Process()

	encoders = append(encoders, enc)
	encPtr := uintptr(unsafe.Pointer(enc))
	encQueue := spp.NewQueue()
	encodedQueues[encPtr] = encQueue
	handleEncoderCallback(enc, encQueue)
	return encPtr
}

func CreateDecoder(codec string) uintptr {
	dec, err := lenss.NewVpxDecoder(lenss.VCodec(codec))
	if err != nil {
		return 0
	}
	dec.Process()

	decoders = append(decoders, dec)
	decPtr := uintptr(unsafe.Pointer(dec))
	decQueue := spp.NewQueue()
	decodedQueues[decPtr] = decQueue
	handleDecoderCallback(dec, decQueue)
	return decPtr
}

func EncodeFrame(p unsafe.Pointer, yuv []byte) {
	enc := (*lenss.VpxEncoder)(p)
	if enc == nil {
		return
	}
	enc.Input <- yuv
}

func DecodeFrame(p unsafe.Pointer, frame []byte) {
	dec := (*lenss.VpxDecoder)(p)
	if dec == nil {
		return
	}
	dec.Srouce <- frame
}

func ReadEncodedData(p unsafe.Pointer) uintptr {
	q, exits := encodedQueues[uintptr(p)]
	if !exits {
		return emptyEvent()
	}
	ptr := q.Dequeue()
	if ptr == nil {
		return emptyEvent()
	}
	return ptr.(uintptr)
}

func ReadDecodedData(p unsafe.Pointer) uintptr {
	q, exists := decodedQueues[uintptr(p)]
	if !exists {
		return emptyEvent()
	}
	ptr := q.Dequeue()
	if ptr == nil {
		return emptyEvent()
	}
	return ptr.(uintptr)
}

func emptyEvent() uintptr {
	return converter.ConvertBytesToPtr(spp.PayloadTypeEmpty, nil)
}

func handleEncoderCallback(enc *lenss.VpxEncoder, q *spp.Queue) {
	go func() {
		for {
			buff := <-enc.Output
			ptr := converter.ConvertBytesToPtr(spp.PayloadTypeBytes, buff)
			q.Enqueue(ptr)
		}
	}()
}

func handleDecoderCallback(dec *lenss.VpxDecoder, q *spp.Queue) {
	go func() {
		for {
			img := <-dec.Output
			buff := img.YuvPlaneBuffer()
			ptr := converter.ConvertBytesToPtr(spp.PayloadTypeBytes, buff)
			q.Enqueue(ptr)
		}
	}()
}
