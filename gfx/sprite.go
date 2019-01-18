package gfx

import (
    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/mathgl/mgl32"
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
    0.0, 1.0, 0.0, 1.0,
    1.0, 0.0, 1.0, 0.0,
    0.0, 0.0, 0.0, 0.0, 

    0.0, 1.0, 0.0, 1.0,
    1.0, 1.0, 1.0, 1.0,
    1.0, 0.0, 1.0, 0.0,
}

type FPoint struct {
    x float32
    y float32
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
    pos FPoint // this should probably go in another structure
}

func NewSprite(x, y, w, h int, program uint32, textures []Texture) Sprite {
    s := Sprite{
        textures: textures,
        program: program,
        rect: Rect{x, y, w, h},
		pos: FPoint{1.0, 1.0},
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

// get the model matrix for a sprite based on it's position, rotation and scale
func (s *Sprite) model() mgl32.Mat4 {
    model := mgl32.Translate3D(s.pos.x, s.pos.y, 1.0).Mul4(mgl32.HomogRotate3DZ(1.0))
    return model
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

    textureUniform := gl.GetUniformLocation(s.program, gl.Str("tex\x00"))
    gl.Uniform1i(textureUniform, 0)

    proj := mgl32.Ortho2D(0.0, 4.0, 4.0, 0.0) // TODO dont do this here!
    projLoc := gl.GetUniformLocation(s.program, gl.Str("proj\x00"))
    gl.UniformMatrix4fv(projLoc, 1, false,  &proj[0])

	model := s.model()
    modelLoc := gl.GetUniformLocation(s.program, gl.Str("model\x00"))
    gl.UniformMatrix4fv(modelLoc, 1, false,  &model[0])

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
