package graphics

import (
    "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	rg "github.com/gen2brain/raylib-go/raygui"
)

type Renderer struct {
    shader *rl.Shader
    target *rl.RenderTexture2D
    guiToggle bool
    canvasOutline rl.Rectangle
    guiOutline rl.Rectangle
}


func NewRenderer(windowWidth int32, windowHeight int32) (*Renderer) {
    shader := rl.LoadShader("", "./shaders/julia.frag.glsl");
    target := rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))

    canvasOutline := rl.Rectangle{X: 0, Y: 0, Width: float32(windowWidth), Height: float32(windowHeight)}

    rl.SetTargetFPS(60)
    return &Renderer{shader:&shader, 
        target:&target, 
        guiToggle: false, 
        canvasOutline: canvasOutline,
        guiOutline: rl.Rectangle{X: 0, Y: float32(windowHeight)/2, Width: float32(windowWidth), Height: float32(windowHeight)/2},
    }
}

func (self *Renderer) SetShaderValueFloat(name string, value float32){
    loc := rl.GetShaderLocation(*self.shader, name)
    v := make([]float32, 1)
    v[0] = value
    rl.SetShaderValue(*self.shader, loc, v, rl.ShaderUniformFloat)
}

func (self *Renderer) SetShaderValueV(name string, value []float32){
    loc := rl.GetShaderLocation(*self.shader, name)
    var uniformType rl.ShaderUniformDataType
    if len(value) == 2 {
        uniformType = rl.ShaderUniformVec2
    } else if len(value) == 3 {
        uniformType = rl.ShaderUniformVec3
    } else if len(value) == 4 {
        uniformType = rl.ShaderUniformVec4
    }
    rl.SetShaderValue(*self.shader, loc, value, uniformType)
}

func (self *Renderer) Render(renderGui bool) {
    rl.BeginTextureMode(*self.target)
    rl.ClearBackground(rl.Black)
    rl.DrawRectangle(0,0,int32(rl.GetScreenWidth()),int32(rl.GetScreenHeight()), rl.Black)
    rl.EndTextureMode()

    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    rl.BeginShaderMode(*self.shader)
    rl.DrawTextureEx(self.target.Texture, rl.Vector2{X:0.0, Y:0.0}, 0.0, 1.0, rl.White)
    rl.EndShaderMode()

    if self.guiToggle {
        self.renderGui()
    }

    rl.EndDrawing()
}

func (self *Renderer) ToggleGui() {
    self.guiToggle = !self.guiToggle
    if self.guiToggle {
        self.canvasOutline.Height = float32(rl.GetScreenHeight())/2
    } else {
        self.canvasOutline.Height = float32(rl.GetScreenHeight())
    }
    rl.UnloadRenderTexture(*self.target)
    target := rl.LoadRenderTexture(int32(self.canvasOutline.Width), int32(self.canvasOutline.Height))
    self.target = &target
}

func (self *Renderer) GetCanvasDim() (float32, float32){
    return self.canvasOutline.Width, self.canvasOutline.Height
}

func (self *Renderer) renderGui() {
    rl.DrawText("Press Mouse buttons right/left to zoom in/out and move", 10, 15, 10, rl.RayWhite);
    rl.DrawText("Press KEY_F1 to toggle these controls", 10, 30, 10, rl.RayWhite);
    rl.DrawText("Press KEYS [1 - 6] to change point of interest", 10, 45, 10, rl.RayWhite);
    rl.DrawText("Press KEY_LEFT | KEY_RIGHT to change speed", 10, 60, 10, rl.RayWhite);
    rl.DrawText("Press KEY_SPACE to pause movement animation", 10, 75, 10, rl.RayWhite);

    //TODO More GUI rendering here
    closeGui := rg.WindowBox(self.guiOutline, "Window")
    if closeGui {
        self.ToggleGui()
    }
}

func (self *Renderer) Deinit() {
    rl.UnloadShader(*self.shader)
    rl.UnloadRenderTexture(*self.target)
}

func Test() {
    fmt.Println("Hello from graphics")
}
