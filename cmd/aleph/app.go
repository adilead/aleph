package main

import (
	"flag"
	"github.com/adilead/aleph/internal/graphics"
	"github.com/adilead/aleph/internal/simulation"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type foo func()

type App struct {
    args *flag.FlagSet
    renderer *graphics.Renderer
    simulator simulation.Simulator
}

func NewApp(argv []string)(*App){
    app := App{} 

    // flag begin
    app.args = flag.NewFlagSet("args", flag.ExitOnError)
    // TODO define flags here
    app.args.Parse(argv)
    //flag end
    
    rl.SetTraceLog(rl.LogNone)


    return &app 
}

//main loop renders julia set
func (self *App) Run() error {
    screenWidth := 1920
    screenHeight := 1080
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Aleph")
	defer rl.CloseWindow()

    self.renderer = graphics.NewRenderer(int32(screenWidth), int32(screenHeight))
    defer self.renderer.Deinit()

    self.simulator = simulation.NewJuliaSimulator(func() (float32, float32){return self.renderer.GetCanvasDim()} )

    // pointsOfInterest := [][]float32 {
    //     { -0.348827, 0.607167 },
    //     { -0.786268, 0.169728 },
    //     { -0.8, 0.156 },
    //     { 0.285, 0.0 },
    //     { -0.835, -0.2321 },
    //     { -0.70176, -0.3842 },
    // }

    // c := []float32{ pointsOfInterest[0][0], pointsOfInterest[0][1] }
    // self.renderer.SetShaderValueV("c", c)        

    // var zoom float32 = 1.0
    // self.renderer.SetShaderValueFloat("zoom", zoom)

    // offset := []float32{ -1.0 * float32(rl.GetScreenWidth())/2, -1.0 * float32(rl.GetScreenHeight())/2 }
    // self.renderer.SetShaderValueV("offset", offset)

    // screenDims := []float32{0.0, 0.0}
    // screenDims[0], screenDims[1] = self.renderer.GetCanvasDim()
    // self.renderer.SetShaderValueV("screenDims", screenDims)

    // offsetSpeed := []float32{0.0, 0.0}


    // pause := false
    // incrementSpeed := 0

	for !rl.WindowShouldClose() {
        if rl.IsKeyPressed(rl.KeyF1){ 
            self.renderer.ToggleGui() // Toggle whether or not to show controls
        }

        self.update()

        // if rl.IsKeyPressed(rl.KeyOne) ||
        //     rl.IsKeyPressed(rl.KeyTwo) ||
        //     rl.IsKeyPressed(rl.KeyThree) ||
        //     rl.IsKeyPressed(rl.KeyFour) ||
        //     rl.IsKeyPressed(rl.KeyFive) ||
        //     rl.IsKeyPressed(rl.KeySix) {
        //     if rl.IsKeyPressed(rl.KeyOne) { 
        //         c[0] = pointsOfInterest[0][0]
        //         c[1] = pointsOfInterest[0][1]
        //     } else if rl.IsKeyPressed(rl.KeyTwo){
        //         c[0] = pointsOfInterest[1][0]
        //         c[1] = pointsOfInterest[1][1]
        //     } else if rl.IsKeyPressed(rl.KeyThree) {
        //         c[0] = pointsOfInterest[2][0]
        //         c[1] = pointsOfInterest[2][1]
        //     } else if rl.IsKeyPressed(rl.KeyFour) {
        //         c[0] = pointsOfInterest[3][0]
        //         c[1] = pointsOfInterest[3][1]
        //     } else if rl.IsKeyPressed(rl.KeyFive) {
        //         c[0] = pointsOfInterest[4][0]
        //         c[1] = pointsOfInterest[4][1]
        //     } else if rl.IsKeyPressed(rl.KeySix) {
        //         c[0] = pointsOfInterest[5][0]
        //         c[1] = pointsOfInterest[5][1]
        //     }

        //     self.renderer.SetShaderValueV("c", c)        
        // }

        // if rl.IsKeyPressed(rl.KeySpace) {
        //     pause = !pause;                 // Pause animation (c change)
        // }

        // if rl.IsKeyPressed(rl.KeyF1){ 
        //     self.renderer.ToggleGui() // Toggle whether or not to show controls

        //     screenDims[0], screenDims[1] = self.renderer.GetCanvasDim()
        //     self.renderer.SetShaderValueV("screenDims", screenDims)

        //     offset = []float32{ -1.0 * screenDims[0]/2, -1.0 * screenDims[1]/2 }
        //     self.renderer.SetShaderValueV("offset", offset)

        //     zoom = 1.0
        //     self.renderer.SetShaderValueFloat("zoom", zoom)

        //     c = []float32{ pointsOfInterest[0][0], pointsOfInterest[0][1] }
        //     self.renderer.SetShaderValueV("c", c)        
        // }


        // if !pause {
        //     if rl.IsKeyPressed(rl.KeyRight) {
        //        incrementSpeed++
        //     } else if rl.IsKeyPressed(rl.KeyLeft) {
        //        incrementSpeed--
        //     }

        //     if rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsMouseButtonDown(rl.MouseRightButton) {
        //         if rl.IsMouseButtonDown(rl.MouseLeftButton){
        //             zoom += zoom*0.003
        //         }

        //         if rl.IsMouseButtonDown(rl.MouseRightButton){
        //             zoom -= zoom*0.003
        //         } 

        //         mousePos := rl.GetMousePosition()

        //         offsetSpeed[0] = mousePos.X - screenDims[0]/2
        //         offsetSpeed[1] = mousePos.Y - screenDims[1]/2

        //         // Slowly move camera to targetOffset
        //         offset[0] += rl.GetFrameTime()*offsetSpeed[0]*0.8
        //         offset[1] += rl.GetFrameTime()*offsetSpeed[1]*0.8
        //     } else { 
        //         offsetSpeed[0] = 0.0
        //         offsetSpeed[1] = 0.0
        //     }

        //     self.renderer.SetShaderValueFloat("zoom", zoom)
        //     self.renderer.SetShaderValueV("offset", offset)

        //     // Increment c value with time
        //     amount := rl.GetFrameTime()*float32(incrementSpeed)*0.0005
        //     c[0] += amount
        //     c[1] += amount

        //     self.renderer.SetShaderValueV("c", c)
        // }
        self.renderer.Render(self.simulator.GetRenderFunc())
    }

    return nil
}

func (self *App) Deinit() {
    // self.renderer.Deinit()
}

func (self *App) processInput() {
    //Process input in the gui from the LAST run iteration
}

func (self *App) update() {
    self.simulator.Update()
}

