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
// TODO make a standard header for shaders including these constants

// viewportWidth / tilesize
const int   tilePixelRes    = 32;
const vec2  viewPixelRes    = vec2(128, 128);
const vec2  viewTileRes     = viewPixelRes / tilePixelRes; // vec2(4, 4);
const vec2  viewTileOffset  = viewTileRes / 2;  // vec2(2, 2);

// used to convert tile coordinates to screen space coordinates (out: -1 to +1)
const vec2  viewTileSize  = tilePixelRes / viewPixelRes; // vec2(0.5, 0.5);

// TODO : _is this needed? we have pos for screen pos and no MVP, seems like 
// we should be sending points and using GEOM shader or some other billboarding
layout (location = 0) in vec2 vert;
layout (location = 1) in vec2 _uv;

// the postition of the quad
const vec2 u_pos = vec2(1, 2);
// the dimensions of the quad
const vec2 u_dim = vec2(1, 1);

out vec2 uv;
out vec2 dim;

// takes a point in tile space and translates to screen space where tile space
// is bottom left (0, 0) top right is some arbatrary integer like (25, 20)
// screen space in the range of (-1 to +1) where bottom left is (-1, -1)
vec2 project(vec2 v) {
    vec2 p = u_pos.xy;
    p.y = viewTileRes.y - p.y; // invert y so we have a top right system

    return v * viewTileSize + u_pos / viewTileRes * 2 - 1;
}

void main()
{
    uv  = _uv;
    dim = u_dim;
    gl_Position = vec4(project(vert), 0.0, 1.0);
}
` + "\x00"

var fragmentShader = `
#version 330

out vec4 color;

in vec2 uv;
in vec2 dim;
in vec2 pos;

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
