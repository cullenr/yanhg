package main

import (
	"github.com/cullenr/yanhg/gfx"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	_ "image/png"
	"log"
)

const windowWidth = 128
const windowHeight = 128

var vertexShader = `
#version 330
// used to convert tile coordinates to screen space coordinates (out: -1 to +1)
layout (location = 0) in vec2 vert;
layout (location = 1) in vec2 _uv;

out vec2 uv;

uniform mat4 proj;
uniform mat4 model;

void main()
{
    uv  = _uv;
    gl_Position = proj * model * vec4(vert.xy, 0.0, 1.0);
}
` + "\x00"

var fragmentShader = `
#version 330

out vec4 color;

in vec2 uv;

uniform sampler2D tex;

void main()
{
    color = texture(tex, uv);
}
` + "\x00"

func main() {
	window, err := gfx.InitWindow(windowWidth, windowHeight)
	if err != nil {
		log.Fatalln(err)
	}
	defer gfx.Destroy()
	// TODO : these will be loaded when the level loads
	program, err := gfx.ProgramFromSource(vertexShader, fragmentShader)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO : these will be loaded when the level loads
	texture, err := gfx.LoadTexture("square.png")
	if err != nil {
		log.Fatalln(err)
	}

	sprite := gfx.NewSprite(
		32, 32, 32, 32,
		program,
		[]gfx.Texture{texture},
	)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		sprite.Draw()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
