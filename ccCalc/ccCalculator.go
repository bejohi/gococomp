package ccCalc

import (
	"github.com/bejohi/gococomp/model"
	"math"
	"image"
)


func CountConnectedComponents(uniformImg *image.Gray, radius int) int{
	height := uniformImg.Rect.Max.Y
	width := uniformImg.Rect.Max.X
	count := 0
	for y := 0; y < height; y++{
		for x := 0; x < width; x++{
			centerPixel := model.LbpPixel{X:x,Y:y}
			count += GetAllUniformPixelInRadius(uniformImg,height,width,centerPixel,radius)
		}
	}
	return count
}

func GetAllUniformPixelInRadius(uniformImg *image.Gray, imgHeight int, imgWidth int, centerPixel model.LbpPixel, radius int) int{

	// We use this rectangle to roughly calculate the radius around our pixel.
	// Inside the for loop we then calculate the euclidean distance, to know exactly if the pixel is in range.
	roughRect := GetRectangleAroundPixelByRadius(centerPixel,radius,imgWidth,imgHeight)
	count := 0

	for y := roughRect.Top; y <= roughRect.Bottom; y++{
		for x := roughRect.Left; x <= roughRect.Right;x++{
			pixel := model.LbpPixel{x,y}
			if pixel.Equals(centerPixel) || CalcPixelDistance(&pixel,&centerPixel) > radius{
				continue
			}
			if uniformImg.GrayAt(x,y).Y >= 255{
				count++
			}
		}
	}
	return count
}

// getRectangleAroundPixelByRadius creates a rectangle around a given pixel, which is definitely in range of a matrix.
// Therefor matrixWidth and matrixHeight are provided.
func GetRectangleAroundPixelByRadius(pixel model.LbpPixel, radius int, matrixWidth int, matrixHeight int) model.SidesRect {
	left := pixel.X - radius
	right := pixel.X + radius
	top := pixel.Y - radius
	bottom := pixel.Y + radius

	if left < 0{
		left = 0
	} else if left > matrixWidth -1{
		left = matrixWidth -1
	}

	if right < 0{
		right = 0
	} else if right > matrixWidth -1 {
		right = matrixWidth -1
	}

	if top < 0{
		top = 0
	} else if right > matrixHeight -1 {
		top = matrixHeight -1
	}

	if bottom < 0 {
		bottom = 0
	} else if bottom > matrixHeight -1 {
		bottom = matrixHeight -1
	}

	return model.SidesRect{left,top,right,bottom}
}

// CalcPixelDistance calculates the euclidean distance between to pixels in the same matrix.
// Therefore the Pythagorean theorem is used.
func CalcPixelDistance(pix1 *model.LbpPixel, pix2 *model.LbpPixel) int{
	leg1 := math.Abs(float64(pix1.X - pix2.X))
	leg2 := math.Abs(float64(pix1.Y - pix2.Y))
	hypotenuse := math.Sqrt(leg1 * leg1 + leg2 * leg2)
	return int(hypotenuse)
}