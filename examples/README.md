# encoder / decoder example

## How to run

1. Create ivf file  
   ex) ffmpeg -i `YOUR_VIDEO.mp4 ./video/example.ivf`

2. Decode created ivf file with simple_decoder  
   `cd ./simple_decoder`  
   `go run . -i ../video/example.ivf -o ../video/decoded.raw`

   you can check decoded file with ffmpeg  
   `ex) ffplay -f rawvideo -pix_fmt yuv420p -s 1920x1080 ../video/decoded.raw`

3. Encode deocoded file with simple_encoder  
   `cd ./simple_encoder`  
   `go run . -i ../video/decoded.raw -o ../video/encoded.ivf -w YOUR_VIDEO_WIDTH -h YOUR_VIDEO_HEIGHT -k KEY_FRAME_COUNT -f FRAME_COUNT `  
   `ex) go run . -i ../video/decoded.raw -o ../video/encoded.ivf -w 1920 -h 1080 -k 5 -f 300`
   `ffplay ../video/encoded.ivf`
