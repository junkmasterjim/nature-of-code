package main

import (
	"image/color"
	"log"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const WIDTH int = 320
const HEIGHT int = 240
const WINDOW_SCALE int = 4

// GaussianRandom returns a random number from a Gaussian distribution
// with the given mean and standard deviation.
func GaussianRandom(mean, stdDev float64) float64 {
	// Box-Muller transform
	u1, u2 := rand.Float64(), rand.Float64()
	z0 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z0*stdDev + mean
}

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
	stepDist := int(GaussianRandom(5, 2.5))

	dir := rand.IntN(4)
	switch dir {
	case 0:
		if w.x-stepDist >= 0 {
			w.x -= stepDist
		}
	case 1:
		if w.y-stepDist >= 0 {
			w.y -= stepDist
		}
	case 2:
		if w.x+stepDist < WIDTH {
			w.x += stepDist
		}
	case 3:
		if w.y+stepDist < HEIGHT {
			w.y += stepDist
		}
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
	ebitenutil.DebugPrint(screen, "Gaussian Walker")

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
	ebiten.SetWindowTitle("Gaussian Walker")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
