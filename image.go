https://go-tour-jp.appspot.com/moretypes/23

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var image [][]uint8

	for i := 0; i < dy; i++ {
		image = append(image, make([]uint8, dx))
		for j := 0; j < dx; j++ {
			//image[i][j] = uint8((i+j)/2)
			image[i][j] = uint8(j)

		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
