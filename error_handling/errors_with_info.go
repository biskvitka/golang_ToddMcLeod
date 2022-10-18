package main

import (
	"errors"
	"log"
)

func main(){
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root to negative number")
	}
	return 43, nil
}