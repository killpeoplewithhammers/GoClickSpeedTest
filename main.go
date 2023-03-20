package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "BOUNCY TEXT")

	rl.SetTargetFPS(120)

	titleHeight := int32(200)
	up := false

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if up {
			titleHeight += 5
		}
		if !up {
			titleHeight -= 5
		}
		if titleHeight > 300 {
			up = false
		} else if titleHeight < 100 {
			up = true
		}

		rl.DrawText("this should go up and down", 240, int32(titleHeight), 20, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
