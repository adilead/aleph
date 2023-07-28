package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/adilead/aleph/internal/graphics"
	// rg "github.com/gen2brain/raylib-go/raygui"
)

// TODO own shader type encapsulating rl.Shader for easy uniform handling
type canvasDimCallback func() (float32, float32)


type JuliaSimulator struct {
    shader *graphics.Shader
    target *rl.RenderTexture2D
    pointsOfInterest [][]float32
    c []float32
    offset []float32
    offsetSpeed []float32
    zoom float32
    canvasDim []float32
    getCanvasDim canvasDimCallback
    incrementSpeed int
}

func NewJuliaSimulator(getCanvasDIm canvasDimCallback) *JuliaSimulator {
    shader := graphics.NewShader("JuliaSetShader", "", "./shaders/julia.frag.glsl")
    target := rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))

    pointsOfInterest := [][]float32 {
        { -0.348827, 0.607167 },
        { -0.786268, 0.169728 },
        { -0.8, 0.156 },
        { 0.285, 0.0 },
        { -0.835, -0.2321 },
        { -0.70176, -0.3842 },
    }
    var zoom float32 = 1.0
    c := []float32{ pointsOfInterest[0][0], pointsOfInterest[0][1] }
    offset := []float32{ -1.0 * float32(rl.GetScreenWidth())/2, -1.0 * float32(rl.GetScreenHeight())/2 }
    offsetSpeed := []float32{0.0, 0.0}
    canvasDim := []float32{0.0, 0.0}
    canvasDim[0], canvasDim[1] = getCanvasDIm()

    shader.SetShaderValueFloat("zoom", zoom)
    shader.SetShaderValueV("offset", offset)
    shader.SetShaderValueV("screenDims", canvasDim[:])

    return &JuliaSimulator{
        shader: shader,
        target: &target,
        pointsOfInterest: pointsOfInterest,
        c: c,
        offset: offset,
        offsetSpeed: offsetSpeed,
        zoom: zoom,
        canvasDim: canvasDim,
        getCanvasDim: getCanvasDIm,
        incrementSpeed: 0,
    }
}

func (self *JuliaSimulator) DestroySimulation() {
    self.shader.Destroy()
}

func (self *JuliaSimulator) Update () {
    if rl.IsKeyPressed(rl.KeyOne) ||
        rl.IsKeyPressed(rl.KeyTwo) ||
        rl.IsKeyPressed(rl.KeyThree) ||
        rl.IsKeyPressed(rl.KeyFour) ||
        rl.IsKeyPressed(rl.KeyFive) ||
        rl.IsKeyPressed(rl.KeySix) {
        if rl.IsKeyPressed(rl.KeyOne) { 
            self.c[0] = self.pointsOfInterest[0][0]
            self.c[1] = self.pointsOfInterest[0][1]
        } else if rl.IsKeyPressed(rl.KeyTwo){
            self.c[0] = self.pointsOfInterest[1][0]
            self.c[1] = self.pointsOfInterest[1][1]
        } else if rl.IsKeyPressed(rl.KeyThree) {
            self.c[0] = self.pointsOfInterest[2][0]
            self.c[1] = self.pointsOfInterest[2][1]
        } else if rl.IsKeyPressed(rl.KeyFour) {
            self.c[0] = self.pointsOfInterest[3][0]
            self.c[1] = self.pointsOfInterest[3][1]
        } else if rl.IsKeyPressed(rl.KeyFive) {
            self.c[0] = self.pointsOfInterest[4][0]
            self.c[1] = self.pointsOfInterest[4][1]
        } else if rl.IsKeyPressed(rl.KeySix) {
            self.c[0] = self.pointsOfInterest[5][0]
            self.c[1] = self.pointsOfInterest[5][1]
        }

        self.shader.SetShaderValueV("c", self.c)        
    }


    if rl.IsKeyPressed(rl.KeyF1){ 
        self.canvasDim[0], self.canvasDim[1] = self.getCanvasDim()
        self.shader.SetShaderValueV("screenDims", self.canvasDim)

        rl.UnloadRenderTexture(*self.target)
        target := rl.LoadRenderTexture(int32(self.canvasDim[0]), int32(self.canvasDim[1]))
        self.target = &target

        self.offset = []float32{ -1.0 * self.canvasDim[0]/2, -1.0 * self.canvasDim[1]/2 }
        self.shader.SetShaderValueV("offset", self.offset)

        self.zoom = 1.0
        self.shader.SetShaderValueFloat("zoom", self.zoom)

        self.c = []float32{ self.pointsOfInterest[0][0], self.pointsOfInterest[0][1] }
        self.shader.SetShaderValueV("c", self.c)        
    }


    if rl.IsKeyPressed(rl.KeyRight) {
       self.incrementSpeed++
    } else if rl.IsKeyPressed(rl.KeyLeft) {
       self.incrementSpeed--
    }

    if rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsMouseButtonDown(rl.MouseRightButton) {
        if rl.IsMouseButtonDown(rl.MouseLeftButton){
            self.zoom += self.zoom*0.003
        }

        if rl.IsMouseButtonDown(rl.MouseRightButton){
            self.zoom -= self.zoom*0.003
        } 

        mousePos := rl.GetMousePosition()

        self.offsetSpeed[0] = mousePos.X - self.canvasDim[0]/2
        self.offsetSpeed[1] = mousePos.Y - self.canvasDim[1]/2

        // Slowly move camera to targetOffset
        self.offset[0] += rl.GetFrameTime()*self.offsetSpeed[0]*0.8
        self.offset[1] += rl.GetFrameTime()*self.offsetSpeed[1]*0.8
    } else { 
        self.offsetSpeed[0] = 0.0
        self.offsetSpeed[1] = 0.0
    }

    self.shader.SetShaderValueFloat("zoom", self.zoom)
    self.shader.SetShaderValueV("offset", self.offset)

    // Increment c value with time
    amount := rl.GetFrameTime()*float32(self.incrementSpeed)*0.0005
    self.c[0] += amount
    self.c[1] += amount

    self.shader.SetShaderValueV("c", self.c)
}

func (self *JuliaSimulator) render() {
    rl.ClearBackground(rl.Black)
    rl.BeginShaderMode(*self.shader.GetShader())
    rl.DrawTextureEx(self.target.Texture, rl.Vector2{X:0.0, Y:0.0}, 0.0, 1.0, rl.White)
    rl.EndShaderMode()
}

func (self *JuliaSimulator) GetRenderFunc() func() {
    return func() {
        self.render()
    }
}
