package main

import (
	"encoding/json"
	"github.com/asticode/go-astisub/blue/model"
	"github.com/asticode/go-astisub/blue/util"
	"log"
	"os"
)

func main() {

	var blueSubtitles []model.Subtitle
	format := "vtt"
	data, _ := os.ReadFile("subtitles.json")
	err := json.Unmarshal(data, &blueSubtitles)
	if err != nil {
		log.Panicln(err)
		return
	}
	asticodeSubtitle := util.BlueModelToAsticodeModel(blueSubtitles, format)
	filename := "/Users/chris/Desktop/subtitles." + format
	if err := asticodeSubtitle.Write(filename); err != nil {
		log.Panicln(err)
		return
	}
}
