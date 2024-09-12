package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"

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
	grid   [][]GridSquare
	stdDev float64
}

type GridSquare struct {
	live bool
	r    uint8
	g    uint8
	b    uint8
}

func Init() *Game {
	g := make([][]GridSquare, WIDTH)
	for x := range g {
		g[x] = make([]GridSquare, HEIGHT)
	}

	return &Game{grid: g}
}

func (g *Game) Update() error {
	randX := GaussianRandom(float64(WIDTH/2), g.stdDev)
	randY := GaussianRandom(float64(HEIGHT/2), g.stdDev)

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.stdDev++
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.stdDev--
	}

	if int(randX) < WIDTH && int(randY) < HEIGHT && randY > 0 && randX > 0 {
		g.grid[int(randX)][int(randY)].live = true
		g.grid[int(randX)][int(randY)].r = uint8(GaussianRandom(255/2, g.stdDev))
		g.grid[int(randX)][int(randY)].g = uint8(GaussianRandom(255/2, g.stdDev))
		g.grid[int(randX)][int(randY)].b = uint8(GaussianRandom(255/2, g.stdDev))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Gaussian Splatter - stdDev: "+strconv.FormatFloat(g.stdDev, 'f', 0, 64))

	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j].live == true {
				vector.DrawFilledCircle(screen, float32(i), float32(j), 3, color.RGBA{g.grid[i][j].r, g.grid[i][j].g, g.grid[i][j].b, 10}, true)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	g := Init()

	ebiten.SetWindowSize(WIDTH*WINDOW_SCALE, HEIGHT*WINDOW_SCALE)
	ebiten.SetWindowTitle("Gaussian Distribution Splatter")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
