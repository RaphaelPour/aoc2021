package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RaphaelPour/aoc2021/util"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <path puzzle input>\n", os.Args[0])
		return
	}

	input := util.LoadString(os.Args[1])
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

	filename := fmt.Sprintf("image_%d.png", time.Now().Unix())
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, image); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("rendered %s to %s\n", os.Args[1], filename)
}
