package gfx

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var glTextures = []uint32{
    gl.TEXTURE0,
    gl.TEXTURE1,
    gl.TEXTURE2,
    gl.TEXTURE3,
    gl.TEXTURE4,
    gl.TEXTURE5,
    gl.TEXTURE6,
    gl.TEXTURE7,
}

// TODO : use geometry shader for funz
var squareVerts = []float32{
    //X     Y    U    V
    -1.0, -1.0, 1.0, 0.0,
    1.0, -1.0, 0.0, 0.0,
    -1.0, 1.0, 1.0, 1.0,
    1.0, -1.0, 0.0, 0.0,
    1.0, 1.0, 0.0, 1.0,
    -1.0, 1.0, 1.0, 1.0,
}

type Rect struct {
    x int
    y int
    w int
    h int
}

type Texture struct {
    glid uint32
    w int
    h int
}

type Sprite struct {
    vao uint32
    vbo uint32
    program uint32
    textures []Texture
    rect Rect
}

func NewSprite(x, y, w, h int, program uint32, textures []Texture) Sprite {
    s := Sprite{
        textures: textures,
        program: program,
        rect: Rect{x, y, w, h},
    }

	gl.UseProgram(s.program)

	gl.GenVertexArrays(1, &s.vao)
	gl.BindVertexArray(s.vao)

	gl.GenBuffers(1, &s.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(squareVerts) * float32Size, gl.Ptr(squareVerts), gl.STATIC_DRAW)

	vertAttr := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttr)
	gl.VertexAttribPointer(vertAttr, 3, gl.FLOAT, false, 4*float32Size, gl.PtrOffset(0))

	uvAttr := uint32(gl.GetAttribLocation(program, gl.Str("_uv\x00")))
	gl.EnableVertexAttribArray(uvAttr)
	gl.VertexAttribPointer(uvAttr, 2, gl.FLOAT, false, 4 * float32Size, gl.PtrOffset(2 * float32Size))

    return s
}

func (s *Sprite) Draw() {
	gl.UseProgram(s.program)
	gl.BindVertexArray(s.vao)
    // maybe we should iterate over all textures in glTextures and bind to 0 if
    // there are no matching textures in sprites textures array? this will
    // unbind the texture if its not used.
    for i, texture := range s.textures {
        gl.ActiveTexture(glTextures[i])
        gl.BindTexture(gl.TEXTURE_2D, texture.glid)
    }
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
