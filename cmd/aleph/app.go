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

    // self.simulator = simulation.NewJuliaSimulator(func() (float32, float32){return self.renderer.GetCanvasDim()} )
    self.simulator = simulation.NewGolSimulator(func() (float32, float32){return self.renderer.GetCanvasDim()} )

	for !rl.WindowShouldClose() {
        if rl.IsKeyPressed(rl.KeyF1){ 
            self.renderer.ToggleGui() // Toggle whether or not to show controls
        }
        self.update()
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

