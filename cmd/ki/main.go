package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/neonnetwork/ki/pkg/structure"
	"image/color"
	"log"
)

var (
	NodeRoot *structure.BinaryTreeNode[int] = structure.NewBinaryTreeNode(1)
	NodeCurr *structure.BinaryTreeNode[int] = NodeRoot

	NodeVal = 2
)

func init() {

}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	
	rl.InitWindow(1280, 720, "Ki [raylib]")
	defer rl.CloseWindow()
	
	rl.SetTargetFPS(120)
	rl.HideCursor()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyF) {
			rl.ToggleFullscreen()
		}
		
		if rl.IsKeyPressed(rl.KeyX) {
			NodeCurr = NodeCurr.AddRight(NodeVal)
			NodeVal++
		}

		render()
	}
}

func render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	rl.DrawFPS(8, 8)

	renderNode(make([]int, 0), NodeRoot)

	rl.DrawRectangle(rl.GetMouseX()-8, rl.GetMouseY()-8, 16, 16, rl.White)
}

func renderNode(path []int, node *structure.BinaryTreeNode[int]) {
	pos, size := []int{rl.GetScreenWidth(), rl.GetScreenHeight()}, []int{rl.GetScreenWidth(), rl.GetScreenHeight()}

	for index, value := range path {
		switch value {
		case -1:
			{
				break
			}
		case +1:
			{
				if (index % 2) == 0 {
					pos[0] /= 2
				} else {
					pos[1] /= 2
				}

				break
			}
		}
	}

	rel := []int{rl.GetScreenWidth(), rl.GetScreenHeight()}
	rel[0] -= pos[0]
	rel[1] -= pos[1]

	c := color.RGBA{
		R: byte((node.Value() * 25) % 255),
		G: byte((node.Value() * 25) % 255),
		B: byte((node.Value() * 25) % 255),
		A: 255,
	}

	log.Println(pos)
	
	rl.DrawRectangle(
		int32(rel[0]), int32(rel[1]),
		int32(size[0]), int32(size[1]),
		c)

	node.Left().
		IfPresent(func(value *structure.BinaryTreeNode[int]) { renderNode(append(path, -1), value) })
	node.Right().
		IfPresent(func(value *structure.BinaryTreeNode[int]) { renderNode(append(path, +1), value) })
}
