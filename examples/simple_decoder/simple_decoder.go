package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tkmn0/lenss"
	ivfreader "github.com/tkmn0/lenss/examples/util/ivf_reader"
)

func main() {
	infile_path := flag.String("i", "", "ivf file to read")
	outfile_path := flag.String("o", "", "output file path")
	flag.Parse()

	infile, _ := os.Open(*infile_path)
	outfile, _ := os.Create(*outfile_path)
	defer infile.Close()
	defer outfile.Close()

	ivf, header, err := ivfreader.NewWith(infile)
	if err != nil {
		fmt.Println("ivf error:", err.Error())
	}

	dec, err := lenss.NewVpxDecoder(lenss.CodecVP8)
	if err != nil {
		fmt.Println("decoder create error:", err.Error())
	}

	dec.Process()
	go func() {
		fmt.Println("writing to raw file...")
		for {
			decoded := <-dec.Output
			buff := lenss.YuvPlaneBuffer(decoded)
			outfile.Write(buff)
			fmt.Printf(".")
		}
	}()

	sleepTime := time.Millisecond * time.Duration((float32(header.TimebaseNumerator)/float32(header.TimebaseDenominator))*1000)
	for {
		frame, _, err := ivf.ParseNextFrame()
		if err != nil {
			if err == io.EOF {
				fmt.Println("file read")
			} else {
				fmt.Println("read error:", err.Error())
			}
			break
		}
		time.Sleep(sleepTime)
		dec.Srouce <- frame
	}

}
