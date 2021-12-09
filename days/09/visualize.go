package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
)

var (
	InputFile = flag.String("input", "./input", "Path to puzzle input")
	HeightMap = flag.Bool("height-map", false, "Generate a height map")
	STL       = flag.Bool("stl", false, "Generates a block-based STL file")
)

func renderSTL(input []string) {
	filename := fmt.Sprintf("day09_%d.stl", time.Now().Unix())

	boxes := make([]sdf.SDF3, len(input)*len(input[0]))
	index := 0
	for y, row := range input {
		for x, cell := range row {
			num, err := strconv.Atoi(string(cell))
			if err != nil {
				fmt.Printf("error converting %d\n", cell)
				return
			}
			box2d := sdf.Box2D(sdf.V2{1, 1}, 0)
			// add one so level 0 has one unit
			height := float64(num+1) / 10
			box3d := sdf.Extrude3D(box2d, height)
			m := sdf.Translate3d(sdf.V3{float64(x), float64(y), height / 2})
			boxes[index] = sdf.Transform3D(box3d, m)
			index++
		}
	}

	render.RenderSTL(sdf.Union3D(boxes...), 100, filename)
}

func renderHeightMap(input []string) {
	image := image.NewNRGBA(image.Rect(0, 0, len(input[0]), len(input)))
	for y, row := range input {
		for x, cell := range row {
			num, err := strconv.Atoi(string(cell))
			if err != nil {
				fmt.Printf("error converting %d\n", cell)
				return
			}

			// 0 should be white and 9 very black
			// spread the gray values accors the whole range
			c := uint8(255.0 - (255.0 / 9.0 * float64(num-num%2)))
			image.Set(x, y,
				color.NRGBA{R: c, G: c, B: c, A: 255},
			)
		}
	}

	filename := fmt.Sprintf("day09_%d.png", time.Now().Unix())
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, image); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	flag.Parse()

	input := util.LoadString(*InputFile)
	if *HeightMap {
		renderHeightMap(input)
	}

	if *STL {
		renderSTL(input)
	}
}
