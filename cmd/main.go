package main

import (
	"fmt"

	"github.com/HarAlb/go-lesson3/internal/chess"
)

func main(){
	var countLines int;
	for {
		fmt.Print("Enter count of lines please: ")
		fmt.Scanln(&countLines)
	
		chess.Generate(countLines)
	}
	
}