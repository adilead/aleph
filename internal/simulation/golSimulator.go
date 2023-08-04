package simulation

import (
    "time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GolSimulator struct {
    board [][]uint //state of each cell
    counts [][]uint //how many dead/alive neighbors a cell has
    canvasDim []float32
    offset []int
    Width int
    Height int
    cellSize int
    updateIntervalInMilli int64 //in nanoSec
    lastUpdate time.Time
    lastHoveredCell [2]int
    pause bool
    
    getCanvasDim canvasDimCallback
}

func NewGolSimulator(getCanvasDim canvasDimCallback) *GolSimulator {
    n := 100
    m := 100
    board := make([][]uint, n)
    counts := make([][]uint, n)
    rows_b := make([]uint, n*m)
    rows_c := make([]uint, n*m)
    for i := 0; i < n; i++ {
        board[i] = rows_b[i*m : (i+1)*m]
        counts[i] = rows_c[i*m : (i+1)*m]
    }

    canvasDim := []float32{0.0, 0.0}
    canvasDim[0], canvasDim[1] = getCanvasDim()
    cellSize := int(canvasDim[0]) / 100

    // offset of camera into the board
    offset := []int{-1 * int(canvasDim[0]/2 - float32(cellSize*m)/2), -1 * int(canvasDim[1]/2 - float32(cellSize*n)/2)}

    return &GolSimulator{
        board: board,
        counts: counts,
        canvasDim: canvasDim,
        Width: m,
        Height: n,
        offset: offset,
        cellSize: cellSize,
        getCanvasDim: getCanvasDim,
        updateIntervalInMilli: 1000, //default one sec
        lastUpdate: time.Now(),
        lastHoveredCell: [2]int{-1, -1},
        pause: false,
    } 
}

func (self *GolSimulator) DestroySimulation() {

}

func (self *GolSimulator) Update() {

    if rl.IsKeyPressed(rl.KeyF1){ 
        self.canvasDim[0], self.canvasDim[1] = self.getCanvasDim()
        //TODO resize the board
    }

    if rl.IsKeyPressed(rl.KeyP) {
        self.pause = !self.pause
    }

    if rl.IsMouseButtonDown(rl.MouseLeftButton) {
        mousePos := rl.GetMousePosition()
        v := rl.Vector2Add(mousePos, rl.NewVector2(float32(self.offset[0]), float32(self.offset[1])))
        v = rl.Vector2DivideV(v, rl.NewVector2(float32(self.cellSize), float32(self.cellSize)))
        v = rl.Vector2Add(v, rl.NewVector2(1.0, 1.0))
        x, y := int(v.X), int(v.Y)

        //toggle the cell
        if x < self.Width-1 && y < self.Height-1 && (self.lastHoveredCell[0] != x || self.lastHoveredCell[1] != y) {
            self.board[y][x] = (self.board[y][x] + 1) % 2
            self.lastHoveredCell[0] = x
            self.lastHoveredCell[1] = y
        }
        self.pause = true
    } 

    if rl.IsMouseButtonReleased(rl.MouseLeftButton){
        self.lastHoveredCell[0] = -1
        self.lastHoveredCell[1] = -1
    }

    now := time.Now().UnixMilli()
    if now - self.lastUpdate.UnixMilli() < self.updateIntervalInMilli || self.pause {
        return
    }

    // count the neighbors
    for y := 1; y < int(self.Height)-1; y++ {
        for x := 1; x < int(self.Width)-1; x++ {
            self.counts[y][x] = 
            self.board[y-1][x-1] + self.board[y-1][x] + self.board[y-1][x+1] + 
            self.board[y][x-1] + self.board[y][x+1] + 
            self.board[y+1][x-1] + self.board[y+1][x] + self.board[y+1][x+1]
        }
    }

    //update the cells
    for y := 1; y < int(self.Height)-1; y++ {
        for x := 1; x < int(self.Width)-1; x++ {

            if self.board[y][x] == 0 {
                if self.counts[y][x] == 3 {
                    self.board[y][x] = 1 //Birth
                }
            } else {
                if self.counts[y][x] < 2 || self.counts[y][x] > 3 {
                    self.board[y][x] = 0 //starvation and overpopulation
                }
            }
        }
    }
    self.lastUpdate = time.Now()
}

func (self *GolSimulator) render() {
    for y := 1; y < self.Height-1; y++ {
        for x := 1; x < self.Width-1; x++ {
            if self.board[y][x] == 1 {
                //TODO Don't render if out of canvas
                rl.DrawRectangle(int32((x-1)*self.cellSize - self.offset[0]), 
                int32((y-1)*self.cellSize - self.offset[1]), 
                int32(self.cellSize), 
                int32(self.cellSize), 
                rl.White)
            }
        }
    }
}

func (self *GolSimulator) GetRenderFunc() func() {
    return func() {
        self.render()
    }
}
