package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
)

type mFunc func(z complex128) color.Color

func main() {
	http.HandleFunc("/mandelbrot", mandelbrot)
	http.HandleFunc("/mandelbrot64", mandelbrot64)
	log.Fatal(http.ListenAndServe("172.26.0.3:19700", nil))
}

func mandelbrot(w http.ResponseWriter, r *http.Request) {
	mandelbrotRunner(w, "mandelbrot")
}
func mandelbrot64(w http.ResponseWriter, r *http.Request) {
	mandelbrotRunner(w, "mandelbrot64")
}

func mandelbrotRunner(w io.Writer, fName string) {
	var fn mFunc
	switch fName {
	case "mandelbrot64":
		fn = mb64
	default:
		fn = mb
	}
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/height*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, fn(z))
		}
	}
	png.Encode(w, img)
}

func mb(z complex128) color.Color {
	const iterations = 200
	const contrast = 10

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{(n * n / 6) % 255, (2*n + 40) % 255, (contrast * n) % 255, 255}
		}
	}
	return color.Black
}

func mb64(z complex128) color.Color {
	const iterations = 200
	const contrast = 20

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v - v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.RGBA{(n * n) % 255, ((2*n + 128) / 5) % 255, uint8(contrast*math.Tan(float64(n/10))) % 255, 255}
		}
	}
	return color.Black
}

func mandelbrotBigFloat() {}

func mandelbrotBigRat() {}
