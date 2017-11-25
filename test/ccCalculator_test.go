package test

import (
	"testing"
	"github.com/bejohi/gococomp/model"
	"github.com/bejohi/gococomp/ccCalc"
)

func getMockBoolMatrix(height int, width int) *[][]bool{
	matrix := make([][]bool,height)
	for i := 0; i < width; i++{
		matrix[i] = make([]bool, width)
	}
	return &matrix
}

func TestCalcPixelDistance_WithXDistance0_ResultShouldBe5(t *testing.T){
	// Arrange
	pix1 := model.LbpPixel{0,0}
	pix2 := model.LbpPixel{0,5}

	// Act
	distance := ccCalc.CalcPixelDistance(&pix1,&pix2)

	// Assert
	if distance != 5{
		t.Error("Distance should be 5.")
	}
}

func TestCalcPixelDistance_WithYDistance0_ResultShouldBe5(t *testing.T){
	// Arrange
	pix1 := model.LbpPixel{5,0}
	pix2 := model.LbpPixel{0,0}

	// Act
	distance := ccCalc.CalcPixelDistance(&pix1,&pix2)

	// Assert
	if distance != 5{
		t.Error("Distance should be 5.")
	}
}

func TestCalcPixelDistance_WithYandXDistance_ResultShouldBe7(t *testing.T){
	// Arrange
	pix1 := model.LbpPixel{5,0}
	pix2 := model.LbpPixel{0,5}

	// Act
	distance := ccCalc.CalcPixelDistance(&pix1,&pix2)

	// Assert
	if distance != 7{
		t.Error("Distance should be 7.")
	}
}

func TestGetRectangleAroundPixelByRadius_WithPixelAtBorder_ResultShouldBeInRange(t *testing.T){
	// Arrange
	width := 10
	height := 5
	pixel := model.LbpPixel{0,0}
	radius := 30

	// Arrange
	resultRect := ccCalc.GetRectangleAroundPixelByRadius(&pixel,radius,width,height)

	// Assert
	if resultRect.Left != 0 || resultRect.Top != 0 || resultRect.Right != 9 || resultRect.Bottom != 4{
		t.Errorf("","The rectangle border was out of range:",
			resultRect.Left, resultRect.Top, resultRect.Right,resultRect.Bottom)
	}

}