package main

import (
	"fmt"

	"github.com/Miavega/patrones-gof/src/patrones/creacionales/factorymethod"
)

func main() {
	audio, _ := factorymethod.GetFormat("Audio")
	video, _ := factorymethod.GetFormat("Video")
	_, err := factorymethod.GetFormat("Otger")
	fmt.Println(err)

	printDetails(audio)
	printDetails(video)
}

func printDetails(g factorymethod.IFormat) {
	g.Play()
}
