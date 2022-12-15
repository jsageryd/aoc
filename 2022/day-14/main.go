package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
	"strings"
)

func main() {
	fASCII := flag.Bool("ascii", false, "Output an ASCII map")
	fAGIF := flag.Bool("agif", false, "Output an animated GIF")
	fGIF := flag.Bool("gif", false, "Output a GIF of the final state")
	fFloor := flag.Bool("floor", false, "Whether cave has floor (part 2)")
	fScale := flag.Int("scale", 15, "Scale of output image")

	flag.Parse()

	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	switch {
	case *fASCII:
		makeASCII(input, *fFloor)
	case *fAGIF:
		makeAnimatedGIF(input, *fFloor, *fScale)
	case *fGIF:
		makeGIF(input, *fFloor, *fScale)
	default:
		fmt.Printf("Part 1: %d\n", part1(input))
		fmt.Printf("Part 2: %d\n", part2(input))
	}
}

func part1(input []string) int {
	cave := newCave(input, false)

	for cave.step() {
	}

	return cave.amountOfSand()
}

func part2(input []string) int {
	cave := newCave(input, true)

	for cave.step() {
	}

	return cave.amountOfSand()
}

type coord struct {
	x, y int
}

type Cave struct {
	grid                 map[coord]byte
	sandSource           coord
	topLeft, bottomRight coord
	grain                coord
	withFloor            bool
}

func newCave(input []string, withFloor bool) *Cave {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sandSource := coord{500, 0}

	cave := &Cave{
		grid:        map[coord]byte{sandSource: '+'},
		sandSource:  sandSource,
		topLeft:     sandSource,
		bottomRight: sandSource,
		grain:       sandSource,
		withFloor:   withFloor,
	}

	updateBounds := func(wall coord) {
		if wall.x < cave.topLeft.x {
			cave.topLeft.x = wall.x
		}
		if wall.x > cave.bottomRight.x {
			cave.bottomRight.x = wall.x
		}
		if wall.y < cave.topLeft.y {
			cave.topLeft.y = wall.y
		}
		if wall.y > cave.bottomRight.y {
			cave.bottomRight.y = wall.y
			if withFloor {
				cave.bottomRight.y += 2
			}
		}
	}

	for _, line := range input {
		var last coord

		for n, cStr := range strings.Split(line, " -> ") {
			var cur coord

			fmt.Sscanf(cStr, "%d,%d", &cur.x, &cur.y)

			updateBounds(cur)

			if n != 0 {
				fromX, toX := min(last.x, cur.x), max(last.x, cur.x)
				fromY, toY := min(last.y, cur.y), max(last.y, cur.y)

				for y := fromY; y <= toY; y++ {
					for x := fromX; x <= toX; x++ {
						cave.grid[coord{x, y}] = '#'
					}
				}
			}

			last = cur
		}
	}

	if withFloor {
		height := cave.bottomRight.y - sandSource.y
		cave.topLeft.x = sandSource.x - height - 1
		cave.bottomRight.x = sandSource.x + height + 1
	}

	return cave
}

// step moves forward one unit of time and returns true as long as sand flows.
func (c *Cave) step() bool {
	rest := true

	delete(c.grid, c.grain)

	for _, dxdy := range []coord{
		{0, 1}, {-1, 1}, {1, 1},
	} {
		dx, dy := dxdy.x, dxdy.y

		_, foundObstacle := c.grid[coord{c.grain.x + dx, c.grain.y + dy}]

		if c.withFloor {
			reachedFloor := c.grain.y >= c.bottomRight.y-1
			foundObstacle = foundObstacle || reachedFloor
		}

		if !foundObstacle {
			c.grain.x += dx
			c.grain.y += dy
			rest = false
			break
		}
	}

	if !c.withFloor {
		if c.grain.y > c.bottomRight.y {
			c.grain = c.sandSource
			return false
		}
	}

	c.grid[c.grain] = 'o'

	if rest {
		if c.grain == c.sandSource {
			return false
		}
		c.grain = c.sandSource
	}

	return true
}

func (c *Cave) amountOfSand() int {
	var grains int

	for _, t := range c.grid {
		if t == 'o' {
			grains++
		}
	}

	return grains
}

func (c *Cave) String() string {
	var b strings.Builder

	for y := c.topLeft.y; y <= c.bottomRight.y; y++ {
		for x := c.topLeft.x; x <= c.bottomRight.x; x++ {
			switch t := c.grid[coord{x, y}]; t {
			case '#', 'o':
				b.WriteByte(t)
			default:
				switch {
				case (coord{x, y}) == c.sandSource:
					b.WriteByte('+')
				case c.withFloor && y == c.bottomRight.y:
					b.WriteByte('#')
				default:
					b.WriteByte('.')
				}
			}
		}
		if y != c.bottomRight.y {
			b.WriteByte('\n')
		}
	}

	return b.String()
}

func makeASCII(input []string, floor bool) {
	cave := newCave(input, floor)

	for cave.step() {
	}

	fmt.Println(cave)
}

func makeAnimatedGIF(input []string, floor bool, scale int) {
	cave := newCave(input, floor)

	imgYOffset := 0
	imgHeight := 15

	imgs := []*image.Paletted{draw(cave, imgYOffset, imgHeight, scale)}
	imgFrameDelay := []int{200}

	var grainMaxY int

	for n := 0; cave.step(); n++ {
		if cave.bottomRight.y > 100 {
			if cave.grain.y < imgYOffset {
				continue
			}
			if cave.grain.y > grainMaxY {
				grainMaxY = cave.grain.y
			}
			if n%4 == 0 && grainMaxY > imgYOffset+imgHeight*3/4 &&
				imgYOffset <= cave.bottomRight.y-imgHeight {
				imgYOffset++
			}
			if n%2 != 0 {
				continue
			}
		}
		imgs = append(imgs, draw(cave, imgYOffset, imgHeight, scale))
		imgFrameDelay = append(imgFrameDelay, 1)
	}

	imgs = append(imgs, draw(cave, imgYOffset, imgHeight, scale))
	imgFrameDelay = append(imgFrameDelay, 200)
	disposalMethods := bytes.Repeat([]byte{gif.DisposalNone}, len(imgs))

	imgBoundsMax := imgs[0].Bounds().Max

	if err := optimizeImgs(imgs); err != nil {
		log.Fatalf("error optimizing images: %v", err)
	}

	gif.EncodeAll(os.Stdout, &gif.GIF{
		Image:    imgs,
		Delay:    imgFrameDelay,
		Disposal: disposalMethods,
		Config: image.Config{
			ColorModel: imgs[0].Palette,
			Width:      imgBoundsMax.X,
			Height:     imgBoundsMax.Y,
		},
	})
}

func makeGIF(input []string, floor bool, scale int) {
	cave := newCave(input, floor)

	for cave.step() {
	}

	img := draw(cave, 0, 0, scale)

	gif.Encode(os.Stdout, img, &gif.Options{NumColors: len(img.Palette)})
}

func draw(cave *Cave, yOffset, height, scale int) *image.Paletted {
	transparent := color.RGBA{0x00, 0x00, 0x00, 0x00}
	background := color.RGBA{0xcc, 0xcc, 0xcc, 0xff}
	wall := color.RGBA{0x30, 0x30, 0x30, 0xff}
	sand := color.RGBA{0xdd, 0x99, 0x22, 0xff}

	palette := []color.Color{
		transparent, background, wall, sand,
	}

	width := cave.bottomRight.x - cave.topLeft.x + 1

	maxHeight := cave.bottomRight.y - cave.topLeft.y + 1

	if height == 0 || height > maxHeight {
		height = maxHeight
	}

	img := image.NewPaletted(image.Rect(0, 0, scale*width, scale*height), palette)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tx, ty := cave.topLeft.x+x, cave.topLeft.y+y+yOffset

			tile := cave.grid[coord{tx, ty}]

			c := background

			switch tile {
			case '#':
				c = wall
			case 'o':
				c = sand
			}

			if cave.withFloor && ty == cave.bottomRight.y {
				c = wall
			}

			for sy := 0; sy < scale; sy++ {
				for sx := 0; sx < scale; sx++ {
					img.Set(scale*x+sx, scale*y+sy, c)
				}
			}
		}
	}

	return img
}

// optimizeImgs optimizes the given set of images for use in a GIF animation
// with disposal method "none", by replacing pixels that do not change between
// frames with transparent ones. This reduces the size of the resulting GIF.
func optimizeImgs(imgs []*image.Paletted) error {
	var transparent color.Color

	for _, c := range imgs[0].Palette {
		if _, _, _, a := c.RGBA(); a == 0x00 {
			transparent = c
			break
		}
	}

	if transparent == nil {
		return errors.New("palette contains no transparent color")
	}

	bounds := imgs[0].Bounds()

	for n := len(imgs) - 1; n > 0; n-- {
		prev, cur := imgs[n-1], imgs[n]
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if prev.At(x, y) == cur.At(x, y) {
					cur.Set(x, y, transparent)
				}
			}
		}
	}

	return nil
}
