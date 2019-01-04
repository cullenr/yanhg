package main

import (
	"github.com/cullenr/yanhg/gfx"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	_ "image/png"
	"log"
)

const windowWidth = 800
const windowHeight = 600

var vertexShader = `
#version 330

layout (location = 0) in vec2 vert;
layout (location = 1) in vec2 _uv;

out vec2 uv;

void main()
{
    uv = _uv;
    gl_Position = vec4(vert.x, vert.y, 0.0, 1.0);
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

	sprite := gfx.NewSprite(program, []uint32{texture})

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		sprite.Draw()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
