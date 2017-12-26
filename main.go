package main

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

//    0 1 2 3 4 5
// 0  0 0 0 0 0 0
// 1  0 0 1 0 0 0
// 2  0 0 1 0 0 1
// 3  0 1 0 1 0 1
// 4  1 1 0 1 1 0
// 5  1 1 1 1 1 1

func getRandomBit(r *rand.Rand, prob int, desired byte) byte {
	if r.Int()%100 < prob {
		return desired
	}

	if desired == 0 {
		return 1
	}

	return 0
}

// Not currently hoooked up to anything
func grayScale() {
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}

	i, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	rnd := rand.New(rand.NewSource(0))

	out := image.NewRGBA(image.Rect(0, 0, 128, 64))

	dump0 := ""
	dump1 := ""
	dump2 := ""
	dump3 := ""
	dump4 := ""
	dump5 := ""

	for p := 0; p < 8; p++ {
		for x := 0; x < 128; x++ {
			var v0 byte = 0
			var v1 byte = 0
			var v2 byte = 0
			var v3 byte = 0
			var v4 byte = 0
			var v5 byte = 0
			for o := 0; o < 8; o++ {
				y := p*8 + 7 - o
				c := i.At(x, y)

				var z0 byte = 0
				var z1 byte = 0
				var z2 byte = 0
				var z3 byte = 0
				var z4 byte = 0
				var z5 byte = 0
				r, g, b, _ := c.RGBA()

				fmt.Println(r)

				gray := (0.3 * float32(r) / 255) + (0.59 * float32(g) / 255) + (0.11 * float32(b) / 255)
				fmt.Println(gray)
				if gray > 255.0 {
					gray = 255.0
				}

				val := int(gray)
				val /= 42
				gray = float32(val * 42)

				var rgba color.RGBA
				rgba.A = 255
				rgba.R = uint8(gray)
				rgba.G = uint8(gray)
				rgba.B = uint8(gray)

				out.Set(x, y, rgba)

				switch val {
				case 0:
					z0 = getRandomBit(rnd, 16, 1)
					z1 = getRandomBit(rnd, 16, 1)
					z2 = getRandomBit(rnd, 16, 1)
					z3 = getRandomBit(rnd, 16, 1)
					z4 = getRandomBit(rnd, 16, 1)
					z5 = getRandomBit(rnd, 16, 1)
					break
				case 1:
					z0 = getRandomBit(rnd, 32, 1)
					z1 = getRandomBit(rnd, 32, 1)
					z2 = getRandomBit(rnd, 32, 1)
					z3 = getRandomBit(rnd, 32, 1)
					z4 = getRandomBit(rnd, 32, 1)
					z5 = getRandomBit(rnd, 32, 1)
					break
				case 2:
					z0 = getRandomBit(rnd, 48, 1)
					z1 = getRandomBit(rnd, 48, 1)
					z2 = getRandomBit(rnd, 48, 1)
					z3 = getRandomBit(rnd, 48, 1)
					z4 = getRandomBit(rnd, 48, 1)
					z5 = getRandomBit(rnd, 48, 1)
					break
				case 3:
					z0 = getRandomBit(rnd, 64, 1)
					z1 = getRandomBit(rnd, 64, 1)
					z2 = getRandomBit(rnd, 64, 1)
					z3 = getRandomBit(rnd, 64, 1)
					z4 = getRandomBit(rnd, 64, 1)
					z5 = getRandomBit(rnd, 64, 1)
				case 4:
					z0 = getRandomBit(rnd, 80, 1)
					z1 = getRandomBit(rnd, 80, 1)
					z2 = getRandomBit(rnd, 80, 1)
					z3 = getRandomBit(rnd, 80, 1)
					z4 = getRandomBit(rnd, 80, 1)
					z5 = getRandomBit(rnd, 80, 1)
				case 5:
					z0 = getRandomBit(rnd, 96, 1)
					z1 = getRandomBit(rnd, 96, 1)
					z2 = getRandomBit(rnd, 96, 1)
					z3 = getRandomBit(rnd, 96, 1)
					z4 = getRandomBit(rnd, 96, 1)
					z5 = getRandomBit(rnd, 96, 1)
					break
				}

				v0 = (v0 << 1) | z0
				v1 = (v1 << 1) | z1
				v2 = (v2 << 1) | z2
				v3 = (v3 << 1) | z3
				v4 = (v4 << 1) | z4
				v5 = (v5 << 1) | z5

				//    0 1 2 3 4 5
				// 0  0 0 0 0 0 0
				// 1  0 0 1 0 0 0
				// 2  0 0 1 0 0 1
				// 3  0 1 0 1 0 1
				// 4  1 1 0 1 1 0
				// 5  1 1 1 1 1 1

			}
			dump0 += "0x"
			dump0 += hex.EncodeToString([]byte{v0})
			dump0 += ", "

			dump1 += "0x"
			dump1 += hex.EncodeToString([]byte{v1})
			dump1 += ", "

			dump2 += "0x"
			dump2 += hex.EncodeToString([]byte{v2})
			dump2 += ", "

			dump3 += "0x"
			dump3 += hex.EncodeToString([]byte{v3})
			dump3 += ", "

			dump4 += "0x"
			dump4 += hex.EncodeToString([]byte{v4})
			dump4 += ", "

			dump5 += "0x"
			dump5 += hex.EncodeToString([]byte{v5})
			dump5 += ", "
		}
	}

	f, _ := os.Create("out.png")

	png.Encode(f, out)

	ioutil.WriteFile("dump0.bin", []byte(dump0), 0644)
	ioutil.WriteFile("dump1.bin", []byte(dump1), 0644)
	ioutil.WriteFile("dump2.bin", []byte(dump2), 0644)
	ioutil.WriteFile("dump3.bin", []byte(dump3), 0644)
	ioutil.WriteFile("dump4.bin", []byte(dump4), 0644)
	ioutil.WriteFile("dump5.bin", []byte(dump5), 0644)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Invalid Arguments")
	}

	output := os.Args[2]

	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	i, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	pages := i.Bounds().Size().Y / 8
	width := i.Bounds().Size().X

	dump := "static byte const _" + output + "[] = {"
	for p := 0; p < pages; p++ {
		for x := 0; x < width; x++ {
			var v byte = 0
			for o := 0; o < 8; o++ {
				y := p*8 + 7 - o
				c := i.At(x, y)

				var z byte = 0
				r, g, b, _ := c.RGBA()
				if r > 5 || g > 5 || b > 5 {
					z = 1
				}

				v = (v << 1) | z
			}
			dump += "0x"
			dump += hex.EncodeToString([]byte{v})

			if p != (pages-1) || x != (width-1) {
				dump += ", "
			}
		}
	}

	dump += "};"

	ioutil.WriteFile(output+".txt", []byte(dump), 0644)
}
