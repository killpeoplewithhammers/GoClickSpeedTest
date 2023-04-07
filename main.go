package main

import (
	"fmt"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/slices"
)

func main() {
	rl.InitWindow(800, 450, "Click Test")
	startTime := time.Now()
	fmt.Println(startTime)

	var clickSlice []time.Time

	for !rl.WindowShouldClose() {
		currentTime := time.Now()
		rl.BeginDrawing()
		if rl.IsMouseButtonPressed(0) {
			clickSlice = append(clickSlice, currentTime)
		}
		now := time.Now()
		if len(clickSlice) > 0 {
			for i := 0; i <= len(clickSlice); i++ {
				if now.Sub(clickSlice[0]) > 1000000000 {
					slices.Delete(clickSlice, 0, 1)
				}
			} // */
		}
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(strconv.Itoa(len(clickSlice))+" clicks per second", 300, 200, 20, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
