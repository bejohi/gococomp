package ccCalc

import (
	"github.com/bejohi/gococomp/model"
	"math"
	"image"
	"fmt"
)


func CreateConnectedComponentImg(uniformImg *image.Gray, radius int) model.ConnectedComponentImg{
	
	ccImg := model.ConnectedComponentImg{}
	ccImg.Height = uniformImg.Rect.Max.Y
	ccImg.Width = uniformImg.Rect.Max.X

	ccImg.ComponentMatrix = make([][]model.ConnectedComponent, ccImg.Height)
	for y := 0; y < ccImg.Height; y++{
		fmt.Println(y)
		ccImg.ComponentMatrix[y] = make([]model.ConnectedComponent,ccImg.Width)
		for x := 0; x < ccImg.Width; x++{
			centerPixel := model.LbpPixel{X:x,Y:y}
			component := GetAllUniformPixelInRadius(uniformImg,ccImg.Height, ccImg.Width,centerPixel,radius)
			ccImg.Count += len((*component.Pixels))
			ccImg.ComponentMatrix[y][x] = component
		}
	}
	fmt.Println("---------------")
	return ccImg
}

func GetAllUniformPixelInRadius(uniformImg *image.Gray, imgHeight int, imgWidth int, centerPixel model.LbpPixel, radius int) model.ConnectedComponent{

	// We use this rectangle to roughly calculate the radius around our pixel.
	// Inside the for loop we then calculate the euclidean distance, to know exactly if the pixel is in range.
	roughRect := GetRectangleAroundPixelByRadius(centerPixel,radius,imgWidth,imgHeight)

	listOfUniformPixel := []model.LbpPixel{}

	for y := roughRect.Top; y <= roughRect.Bottom; y++{
		for x := roughRect.Left; x <= roughRect.Right;x++{
			pixel := model.LbpPixel{x,y}
			if pixel.Equals(centerPixel){
				continue
			}
			if CalcPixelDistance(&pixel,&centerPixel) > radius{
				continue
			}
			if uniformImg.GrayAt(x,y).Y > 0{
				listOfUniformPixel = append(listOfUniformPixel,pixel)
			}
		}
	}
	return model.ConnectedComponent{&listOfUniformPixel}
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