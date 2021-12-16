package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"strings"
)

func main() {
	fASCII := flag.Bool("ascii", false, "Output an ASCII map")
	fAGIF := flag.Bool("agif", false, "Output an animated GIF")
	fGIF := flag.Bool("gif", false, "Output a GIF of the final state")
	fScale := flag.Int("scale", 2, "Scale of output image")

	flag.Parse()

	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	m := newMap(input)

	var (
		imgs          []*image.Paletted
		imgFrameDelay []int
		imgOffset     int
		imgHeight     int
	)

	if *fAGIF {
		imgs = append(imgs, m.Image(imgOffset, 250, *fScale))
		imgFrameDelay = append(imgFrameDelay, 200)
		imgHeight = 250
	}
	for n := 0; m.Step(); n++ {
		if *fAGIF {
			if m.maxY <= 300 {
				imgs = append(imgs, m.Image(imgOffset, imgHeight, *fScale))
				imgFrameDelay = append(imgFrameDelay, 20)
			} else {
				if n > 1500 && n%8 == 0 && imgOffset <= m.maxY-imgHeight {
					imgOffset++
				}
				if n%8 == 0 {
					imgs = append(imgs, m.Image(imgOffset, imgHeight, *fScale))
					imgFrameDelay = append(imgFrameDelay, 8)
				}
			}
		}
	}
	if *fAGIF {
		imgs = append(imgs, m.Image(imgOffset, imgHeight, *fScale))
		imgFrameDelay = append(imgFrameDelay, 200)
	}

	switch {
	case *fASCII:
		fmt.Println(m)
	case *fAGIF:
		imgBoundsMax := imgs[0].Bounds().Max
		gif.EncodeAll(os.Stdout, &gif.GIF{
			Image: imgs,
			Delay: imgFrameDelay,
			Config: image.Config{
				ColorModel: imgs[0].Palette,
				Width:      imgBoundsMax.X,
				Height:     imgBoundsMax.Y,
			},
		})
	case *fGIF:
		img := m.Image(0, 0, *fScale)
		gif.Encode(os.Stdout, img, &gif.Options{NumColors: len(img.Palette)})
	default:
		var movingWater, stillWater int
		for c, v := range m.m {
			if c.y >= m.minY && c.y <= m.maxY {
				switch v {
				case '|':
					movingWater++
				case '~':
					stillWater++
				}
			}
		}
		fmt.Printf("Part 1: %d\n", stillWater+movingWater)
		fmt.Printf("Part 2: %d\n", stillWater)
	}
}

type Coord struct {
	x, y int
}

type Map struct {
	m           map[Coord]byte
	maxX, maxY  int
	minX, minY  int
	movingWater map[Coord]struct{}
	steps       int
}

func newMap(input []string) *Map {
	tap := Coord{500, 0}

	m := &Map{
		m:           map[Coord]byte{},
		movingWater: map[Coord]struct{}{tap: {}},
	}

	for _, line := range input {
		var a, b string
		var rowOrColumn, from, to int
		fmt.Sscanf(line, "%1s=%d, %1s=%d..%d", &a, &rowOrColumn, &b, &from, &to)
		switch a {
		case "x":
			for y := from; y <= to; y++ {
				m.m[Coord{rowOrColumn, y}] = '#'
			}
		case "y":
			for x := from; x <= to; x++ {
				m.m[Coord{x, rowOrColumn}] = '#'
			}
		}
	}

	for c := range m.m {
		m.minX = c.x
		m.maxX = c.x
		m.minY = c.y
		m.maxY = c.y
		break
	}

	for c := range m.m {
		switch {
		case c.x < m.minX:
			m.minX = c.x
		case c.x > m.maxX:
			m.maxX = c.x
		}
		switch {
		case c.y < m.minY:
			m.minY = c.y
		case c.y > m.maxY:
			m.maxY = c.y
		}
	}

	m.m[tap] = '+'

	return m
}

func (m *Map) Step() bool {
	var newStillWater []Coord

	convertToStillWater := func(from, to Coord) {
		for y := from.y; y <= to.y; y++ {
			for x := from.x; x <= to.x; x++ {
				newStillWater = append(newStillWater, Coord{x, y})
			}
		}
	}

	var newMovingWater []Coord

	for mw := range m.movingWater {
		switch down := (Coord{mw.x, mw.y + 1}); m.m[down] {
		case 0:
			if down.y <= m.maxY {
				newMovingWater = append(newMovingWater, down)
			}
		case '#', '~':
			if left := (Coord{mw.x - 1, mw.y}); m.m[left] == 0 {
				newMovingWater = append(newMovingWater, left)
			}
			if right := (Coord{mw.x + 1, mw.y}); m.m[right] == 0 {
				newMovingWater = append(newMovingWater, right)
			}
		}
		if left := (Coord{mw.x - 1, mw.y}); m.m[left] == '#' {
			var x int
			for x = mw.x; m.m[Coord{x, mw.y}] == '|'; x++ {
			}
			if m.m[Coord{x, mw.y}] == '#' {
				convertToStillWater(Coord{mw.x, mw.y}, Coord{x - 1, mw.y})
			}
		}
	}

	for _, w := range newMovingWater {
		m.m[w] = '|'
		m.movingWater[w] = struct{}{}
	}

	for _, w := range newStillWater {
		m.m[w] = '~'
		delete(m.movingWater, w)
	}

	m.steps++

	if len(newMovingWater) == 0 && len(newStillWater) == 0 {
		return false
	}

	return true
}

func (m *Map) Image(offset, height, scale int) *image.Paletted {
	sand := color.RGBA{0xcc, 0xcc, 0xcc, 0xff}
	clay := color.RGBA{0x30, 0x30, 0x30, 0xff}
	movingWater := color.RGBA{0x00, 0x99, 0xff, 0xff}
	stillWater := color.RGBA{0x00, 0x50, 0x99, 0xff}

	palette := []color.Color{
		sand, clay, movingWater, stillWater,
	}

	width := m.maxX + 1 - m.minX

	if height == 0 || height > m.maxY+1 {
		height = m.maxY + 1
	}

	img := image.NewPaletted(image.Rect(0, 0, scale*width, scale*height), palette)

	for y := offset; y < height+offset; y++ {
		for x := m.minX; x <= m.maxX; x++ {
			tile, ok := m.m[Coord{x, y}]
			if !ok {
				tile = '.'
			}
			var c color.Color
			switch tile {
			case '.':
				c = sand
			case '#':
				c = clay
			case '+', '|':
				c = movingWater
			case '~':
				c = stillWater
			}
			for sy := 0; sy < scale; sy++ {
				for sx := 0; sx < scale; sx++ {
					img.Set(scale*(x-m.minX)+sx, scale*(y-offset)+sy, c)
				}
			}
		}
	}

	return img
}

func (m *Map) String() string {
	var rows []string

	for y := 0; y <= m.maxY; y++ {
		row := make([]byte, m.maxX+1-m.minX)
		for x := m.minX; x <= m.maxX; x++ {
			var tile byte
			if t, ok := m.m[Coord{x, y}]; ok {
				tile = t
			} else {
				tile = '.'
			}
			row[x-m.minX] = tile
		}
		rows = append(rows, string(row))
	}

	return strings.Join(rows, "\n")
}
