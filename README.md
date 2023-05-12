# GoClickSpeedTest
This program shows how fast you click using the Raylib library 

## How it Works
This is an explanation of the program for my computer science class

### Imports

```go
import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/slices"
)
```
 - strconv: converting numbers to strings to show them on screen.
 - time: fetching the time when a click happens to remove clicks that are older than a second
 - raylib: for the gui
 - slices: manipulate the slices to remove the clicks that are older than a second

### Code

```go
func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "Click Test")
	var clickSlice []time.Time
	var rightClickSlice []time.Time
```
Initialize the GUI and create variables
```go
	for !rl.WindowShouldClose() {
		currentTime := time.Now()
		rl.BeginDrawing()
```
Draw the frame and get the current time in case a click happens
```go
		if rl.IsMouseButtonPressed(0) {
			clickSlice = append(clickSlice, currentTime)
		}
		if rl.IsMouseButtonPressed(1) {
			rightClickSlice = append(rightClickSlice, currentTime)
		}
```
Add the current time to the slice if a mouse button is pressed
```go
		now := time.Now()
		if len(clickSlice) > 0 {
			for i := 0; i <= len(clickSlice); i++ {
				if now.Sub(clickSlice[0]) > 1000000000 {
					clickSlice = slices.Delete(clickSlice, 0, 1)
				}
			}
		}
		if len(rightClickSlice) > 0 {
			for i := 0; i <= len(rightClickSlice); i++ {
				if now.Sub(rightClickSlice[0]) > 1000000000 {
					rightClickSlice = slices.Delete(rightClickSlice, 0, 1)
				}
			}
		}
```
Remove cicks that are over a second old from the slice
```go
		rl.ClearBackground(rl.RayWhite)
		fontSize := int32(rl.GetRenderWidth() / 40)
		textWidth := rl.MeasureText(strconv.Itoa(len(clickSlice))+" | "+strconv.Itoa(len(rightClickSlice))+" clicks per second", fontSize)
		textPositionX := int32(rl.GetRenderWidth()/2 - int(textWidth)/2)

		textPositionY := int32(rl.GetRenderHeight())/2 - fontSize/2
		rl.DrawText(strconv.Itoa(len(clickSlice))+" | "+strconv.Itoa(len(rightClickSlice))+" clicks per second", textPositionX, textPositionY, fontSize, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
```
Draw the window with the clickspeed on screen
