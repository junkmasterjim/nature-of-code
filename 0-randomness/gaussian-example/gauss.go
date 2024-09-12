package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// GaussianRandom returns a random number from a Gaussian distribution
// with the given mean and standard deviation.
func GaussianRandom(mean, stdDev float64) float64 {
	// Box-Muller transform
	u1, u2 := rand.Float64(), rand.Float64()
	z0 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z0*stdDev + mean
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Example usage of GaussianRandom
	x := GaussianRandom(320, 60)
	y := GaussianRandom(240, 40) // Center vertically

	// Draw a circle at (x, y)
	vector.DrawFilledCircle(screen, float32(x), float32(y), 5, color.White, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gaussian Distribution Example")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
