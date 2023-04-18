package main

import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/slices"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "Click Test")
	var click string
	var clickSlice []time.Time

	for !rl.WindowShouldClose() {
		currentTime := time.Now()
		rl.BeginDrawing()
		if rl.IsMouseButtonPressed(0) {
			clickSlice = append(clickSlice, currentTime)
		}
		if rl.IsMouseButtonPressed(1) {
			clickSlice = append(clickSlice, currentTime)
		}
		now := time.Now()
		if len(clickSlice) > 0 {
			for i := 0; i <= len(clickSlice); i++ {
				if now.Sub(clickSlice[0]) > 1000000000 {
					clickSlice = slices.Delete(clickSlice, 0, 1)
				}
			}
		}
		if len(clickSlice) == 1 {
			click = " click "
		} else {
			click = " clicks "
		}
		rl.ClearBackground(rl.RayWhite)
		fontSize := int32(rl.GetRenderWidth() / 40)
		textWidth := rl.MeasureText(strconv.Itoa(len(clickSlice))+click+"per second", fontSize)
		textPositionX := int32(rl.GetRenderWidth()/2 - int(textWidth)/2)

		textPositionY := int32(rl.GetRenderHeight())/2 - fontSize/2
		rl.DrawText(strconv.Itoa(len(clickSlice))+click+"per second", textPositionX, textPositionY, fontSize, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
