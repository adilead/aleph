package graphics

import rl "github.com/gen2brain/raylib-go/raylib"
type Shader struct {
    shader rl.Shader
    name string
}

func NewShader(name string, vspath string, fspath string) *Shader{
    return &Shader{
        shader: rl.LoadShader(vspath, fspath),
        name: name,
    }
}

func (self *Shader) SetShaderValueFloat(name string, value float32){
    loc := rl.GetShaderLocation(self.shader, name)
    v := make([]float32, 1)
    v[0] = value
    rl.SetShaderValue(self.shader, loc, v, rl.ShaderUniformFloat)
}

func (self *Shader) SetShaderValueV(name string, value []float32){
    loc := rl.GetShaderLocation(self.shader, name)
    var uniformType rl.ShaderUniformDataType
    if len(value) == 2 {
        uniformType = rl.ShaderUniformVec2
    } else if len(value) == 3 {
        uniformType = rl.ShaderUniformVec3
    } else if len(value) == 4 {
        uniformType = rl.ShaderUniformVec4
    }
    rl.SetShaderValue(self.shader, loc, value, uniformType)
}

func (self *Shader) Destroy() {
    rl.UnloadShader(self.shader)
}

func (self *Shader) GetShader() *rl.Shader {
    return &self.shader
}
