package simulation

import (
	// rg "github.com/gen2brain/raylib-go/raygui"
)

type Simulator interface {
    Update()
    GetRenderFunc() func()
    DestroySimulation()
}
