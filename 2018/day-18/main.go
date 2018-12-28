package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"strings"
)

func main() {
	fAGIF := flag.Bool("agif", false, "Output an animated GIF")
	fASCII := flag.Bool("ascii", false, "Output an ASCII map")
	fDelay := flag.Int("delay", 8, "Frame delay of animated GIF")
	fScale := flag.Int("scale", 10, "Scale of output image")
	fSteps := flag.Int("steps", 10, "Number of steps to run when output is ASCII or GIF")
	flag.Parse()

	var area Area

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		area = append(area, scanner.Bytes())
	}

	var (
		imgs          []*image.Paletted
		imgFrameDelay []int
	)

	if *fAGIF {
		imgs = append(imgs, area.Image(*fScale))
		imgFrameDelay = append(imgFrameDelay, 200)
	}

	switch {
	case *fASCII:
		area = step(area, *fSteps)
		fmt.Println(area)
	case *fAGIF:
		for n := 0; n < *fSteps; n++ {
			area = step(area, 1)
			imgs = append(imgs, area.Image(*fScale))
			imgFrameDelay = append(imgFrameDelay, *fDelay)
		}
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
	default:
		areaAfter10Steps := step(area, 10)
		_, trees, lumberyards := areaAfter10Steps.Count()
		fmt.Printf("Part 1: %d\n", trees*lumberyards)

		areaAfter1000000000Steps := step(area, 1000000000)
		_, trees, lumberyards = areaAfter1000000000Steps.Count()
		fmt.Printf("Part 2: %d\n", trees*lumberyards)
	}
}

type Area [][]byte

func (a Area) String() string {
	var ss []string
	for y := range a {
		ss = append(ss, string(a[y]))
	}
	return strings.Join(ss, "\n")
}

func (a Area) Image(scale int) *image.Paletted {
	open := color.RGBA{0xcc, 0xcc, 0x99, 0xff}
	tree := color.RGBA{0x99, 0xcc, 0x00, 0xff}
	lumberyard := color.RGBA{0x30, 0x30, 0x30, 0xff}

	palette := []color.Color{open, tree, lumberyard}

	height, width := len(a), len(a[0])

	img := image.NewPaletted(image.Rect(0, 0, scale*width, scale*height), palette)

	for y := range a {
		for x := range a[0] {
			var c color.Color
			switch a[y][x] {
			case '.':
				c = open
			case '|':
				c = tree
			case '#':
				c = lumberyard
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

func (a Area) Count() (open, trees, lumberyards int) {
	for y := range a {
		for x := range a[y] {
			switch a[y][x] {
			case '.':
				open++
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	return open, trees, lumberyards
}

func (a Area) Hash() []byte {
	h := md5.New()
	for y := range a {
		h.Write(a[y])
	}
	return h.Sum(nil)
}

func step(a Area, steps int) Area {
	cur := make(Area, len(a))
	next := make(Area, len(a))
	for y := range a {
		cur[y] = make([]byte, len(a[y]))
		copy(cur[y], a[y])
		next[y] = make([]byte, len(a[y]))
		copy(next[y], a[y])
	}

	buf := []Area{cur, next}

	seen := map[string]int{string(cur.Hash()): 0}
	skippedAhead := false

	for n := 0; n < steps; n++ {
		cur = buf[n%2]
		next = buf[(n+1)%2]

		for y := range cur {
			for x := range cur[y] {
				_, trees, lumberyards := adjacent(cur, x, y)
				switch cur[y][x] {
				case '.':
					if trees >= 3 {
						next[y][x] = '|'
					} else {
						next[y][x] = '.'
					}
				case '|':
					if lumberyards >= 3 {
						next[y][x] = '#'
					} else {
						next[y][x] = '|'
					}
				case '#':
					if lumberyards >= 1 && trees >= 1 {
						next[y][x] = '#'
					} else {
						next[y][x] = '.'
					}
				}
			}
		}

		if !skippedAhead {
			nh := string(next.Hash())
			if nn, ok := seen[nh]; ok {
				odd := false
				if n%2 != 0 {
					odd = true
				}
				cycle := n - nn
				if cycle != 0 {
					n += cycle * ((steps - n) / cycle)
					if odd && n%2 == 0 {
						n--
					}
					skippedAhead = true
				}
			} else {
				seen[nh] = n
			}
		}
	}

	return next
}

func adjacent(a Area, x, y int) (open, trees, lumberyards int) {
	open = 8
	for yy := y - 1; yy <= y+1; yy++ {
		if yy < 0 || yy > len(a)-1 {
			continue
		}
		for xx := x - 1; xx <= x+1; xx++ {
			if xx < 0 || xx > len(a[yy])-1 {
				continue
			}
			if xx == x && yy == y {
				continue
			}
			switch a[yy][xx] {
			case '|':
				open--
				trees++
			case '#':
				open--
				lumberyards++
			}
		}
	}
	return open, trees, lumberyards
}
