package main

import (
	"Gopengl2/graphics"
	"Gopengl2/graphics/opengl"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := opengl.CreateWindow(800, 600, "test")
	graphics.Init(window)

	ro := graphics.CreateRenderObject("./Gopengl2/planets.png", 2)
	ro.CreateSquare(0, 0, 256, 0, 0, 32)

	for !window.ShouldClose() {
		graphics.Render()
	}

}