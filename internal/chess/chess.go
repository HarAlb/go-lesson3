package chess

import (
	"fmt"
	"strings"
)

func Generate(countLines int) {
	myArray := [][]string{}
    frameCount := countLines + 2
	frame :=  GetFrame(frameCount)
	for i := 0; i < countLines; i++ {
		innerArray := []string{}
		mod := i % 2
		for j := 0; j < frameCount; j++ {
			var symbol string;
			switch mod {
			case 1:
			    symbol = " "
			    	
			default:
				symbol = "#"
			}
			
			if j == 0 || j == frameCount-1 {
				innerArray = append(innerArray, "*")
				continue
			}
			
			if j % 2 == 0 {
				if symbol == "#"{
					symbol = " "
				}else{
					symbol = "#"
				}
				
			}
			innerArray = append(innerArray, symbol)
		}
		myArray = append(myArray, innerArray)
	}
	Render(frame, myArray)
}

func Render(frame string, myArray [][]string) {
	fmt.Println(frame)
	for _,item := range myArray {
		fmt.Println(strings.Join(item, ""))
	}
	fmt.Println(frame)
}

func GetFrame(FrameCount int) string {
	frame := "";
	for i := 0; i < FrameCount; i++ {
		frame += "*"
	}
	return frame
}