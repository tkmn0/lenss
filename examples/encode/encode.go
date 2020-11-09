package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unsafe"

	"github.com/tkmn0/lenss"
	ivfwriter "github.com/tkmn0/lenss/examples/util/ivf_writer"
)

func main() {
	infile_path := flag.String("i", "", "input file path (raw video file)")
	outfile_path := flag.String("o", "", "output file path (encoded ivf file)")
	video_width := flag.Int("w", 0, "video width")
	video_height := flag.Int("h", 0, "video heght")
	key_frame_interval := flag.Int("k", 0, "key frame interval")
	frames := flag.Int("f", 0, "frames to encode")
	flag.Parse()
	fmt.Println("input file path:", *infile_path)
	fmt.Println("output file path:", *outfile_path)
	fmt.Println("video width:", *video_width)
	fmt.Println("video height:", *video_height)
	fmt.Println("key frame interval:", *key_frame_interval)
	fmt.Println("frame count:", *frames)

	infile := lenss.OpenFile(*infile_path, "rb")
	defer lenss.CloseFile(infile)

	enc, err := lenss.NewEncoder(lenss.CodecVP8, *video_width, *video_height, 30, 200, *key_frame_interval)
	if err != nil {
		panic(err.Error())
	}

	ivf, err := ivfwriter.New(*outfile_path)
	if err != nil {
		panic(err.Error())
	}
	defer ivf.Close()

	go func() {
		for {
			encoded := <-enc.Output
			ivf.WriteData(encoded)
			fmt.Printf(".")
		}
	}()

	enc.ProcessFromFile(unsafe.Pointer(infile))

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
