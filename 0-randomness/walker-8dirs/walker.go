package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const WIDTH int = 320
const HEIGHT int = 240
const WINDOW_SCALE int = 4

type Game struct {
	w     *Walker
	grid  [][]bool
	count int
}

type Walker struct {
	x, y int
}

func (w *Walker) Show(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(w.x), float32(w.y), 2, color.RGBA{128, 128, 128, 255}, false)
}

func (w *Walker) Step() {
	rx := rand.IntN(3) - 1
	ry := rand.IntN(3) - 1

	if w.x+rx >= 0 && w.x+rx < WIDTH {
		w.x += rx
	}
	if w.y+ry >= 0 && w.y+ry < HEIGHT {
		w.y += ry
	}
}

func (g *Game) Update() error {
	g.count++
	if g.count >= 3 {
		g.count = 0
	}

	if g.count == 0 {
		g.grid[g.w.x][g.w.y] = true

		g.w.Step()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "8 Direction Walker")

	for x := range g.grid {
		for y := range g.grid[x] {
			if g.grid[x][y] == true {
				vector.DrawFilledRect(screen, float32(x), float32(y), 1, 1, color.RGBA{100, 100, 100, 255}, false)
			}
		}
	}

	g.w.Show(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func Init() *Game {
	walker := &Walker{x: WIDTH / 2, y: HEIGHT / 2}

	grid := make([][]bool, WIDTH)
	for x := range grid {
		grid[x] = make([]bool, HEIGHT)
	}

	g := &Game{w: walker, grid: grid}
	return g
}

func main() {
	g := Init()

	ebiten.SetWindowSize(WIDTH*WINDOW_SCALE, HEIGHT*WINDOW_SCALE)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
