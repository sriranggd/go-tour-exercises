package main

import (
	"fmt"
	"math"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	fmt.Println("From Sqrtv1")
	fmt.Println(Sqrtv1(320))
	fmt.Println("\n\nFrom Sqrtv2")
	fmt.Println(Sqrtv2(320))
	fmt.Println("\n\nFrom the math package")
	fmt.Println(math.Sqrt(320))

	fmt.Println("\n\nPic exercise output-1")
	pic.Show(Pic1)

	fmt.Println("\n\nPic exercise output-2")
	pic.Show(Pic2)

	fmt.Println("\n\nPic exercise output-3")
	pic.Show(Pic3)

	fmt.Println("\n\nWord count exercise output")
	wc.Test(WordCount)
}
