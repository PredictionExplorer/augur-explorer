package main
import (
	"os"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)
func main() {

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [input_file] [output_file]\n\t\t"+
			"Converts Random Walk video to Twitter compatible format\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	input_file_name := os.Args[1]
	output_file_name := os.Args[2]

	err := ffmpeg.Input(input_file_name).
		Output(output_file_name, ffmpeg.KwArgs{"s": "640x480"}).
		OverWriteOutput().ErrorToStdOut().Run()
	fmt.Printf("Error: %v\n",err)
}
