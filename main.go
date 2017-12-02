package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/bejohi/gococomp/helper"
	"github.com/bejohi/golbp/lbpCalc"
	"github.com/bejohi/golbp/model"
	"github.com/bejohi/gococomp/ccCalc"
)


// TODO [bejohi] Split up to different functions.
func main(){

	lenOfArgs := len(os.Args)

	if lenOfArgs < 3{
		panic("The usage of this program is: go run main.go <Path to image folder> <component radius> " +
			"[Keyword in images] [Log Path].")
	}

	fmt.Println("WELCOME TO THE Connected- Component- Analyzer.")

	imgFolderPath := os.Args[1]
	fmt.Println("Searching for images in: " + imgFolderPath)

	radius, radiusError := strconv.Atoi(os.Args[2])

	if radiusError != nil{
		panic("The radius must be an integer.")
	}
	fmt.Println("The radius is set to" + os.Args[2])

	var imgKeyword string

	if lenOfArgs >= 4 {
		imgKeyword = os.Args[3]
		fmt.Println("Only looking for images with " + imgKeyword + " in their name.")
	}

	logPath := "default.log"
	if lenOfArgs >= 5 {
		logPath  = os.Args[4]
		fmt.Println("Saving log under" + logPath)
	}
	helper.Log.LogPath = logPath

	files, err := ioutil.ReadDir(imgFolderPath)

	if err != nil {
		panic(err.Error())
	}

	imgEndings := []string{".jpg",".jpeg",".png"}

	imgFiles := make([]os.FileInfo,0)

	for _, file := range files {
		if imgKeyword != "" && !strings.Contains(file.Name(),imgKeyword){
			continue
		}

		for _, imgEnding := range imgEndings{
			if strings.Contains(file.Name(),imgEnding){
				imgFiles = append(imgFiles,file)
				break
			}
		}
	}

	helper.LogInfo(strconv.Itoa(len(imgFiles)) + " images where found.")

	for _, imgFile := range imgFiles{
		img, imgErr := lbpCalc.GetUniformImage(imgFolderPath + imgFile.Name(),model.EndOfEdgeUniform{})
		if imgErr != nil{
			helper.LogError(imgErr.Error())
			continue
		}
		height := img.Rect.Max.Y
		width := img.Rect.Max.X
		ccCount := ccCalc.CountConnectedComponents(img,radius)
		helper.LogInfo(imgFile.Name() + ";" + strconv.Itoa(radius) + ";" + strconv.Itoa(ccCount) + ";" + strconv.Itoa(height) + ";" + strconv.Itoa(width))
	}
}
