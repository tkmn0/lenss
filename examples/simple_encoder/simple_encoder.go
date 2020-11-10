package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

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

	infile, err := os.Open(*infile_path)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()

	enc, err := lenss.NewEncoder(lenss.CodecVP8, *video_width, *video_height, 30, 200, *key_frame_interval)
	if err != nil {
		panic(err.Error())
	}
	enc.Process()

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

	go func() {
		for {
			buf := make([]byte, enc.Image().FrameSize())
			_, err := infile.Read(buf)
			if err != nil {
				break
			}
			enc.Input <- buf
		}
	}()

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
