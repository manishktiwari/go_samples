package main

import (
	"fmt"

	"com.github.manishktiwari/go-demo1/demo1"
	"com.github.manishktiwari/go-demo1/pathseparator"
	"rsc.io/quote"
)

func main() {
	fmt.Println(demo1.BestUSer())
	fmt.Println(quote.Go())

	var dir, file string

	dir, file = pathseparator.SplitPath("dir/css/fileName.css")

	fmt.Println("directory : ", dir)
	fmt.Println("filename : ", file)

	// how to discard  a value:
	_, file = pathseparator.SplitPath("dir/css1/fileName1.css")
	fmt.Println("filename : ", file)

	// short declaration : how to directly use values without declaring separately
	_, fileName := pathseparator.SplitPath("dir/shortdeclaration/shortdeclaration.css")

	fmt.Println("filename : ", fileName)

	// Type Conversion
	speed := 100 // int
	force := 4.5 // float64

	// speed = speed * force // it will give compliation error
	speed = speed * int(force) // converts float to int with trimming off the fractional part and hence result here will be 400 not 4.5
	fmt.Println(speed)
	fmt.Println(force) // will print 4.5, here conversion creates a new value and doesn't change the actual variable

	// speed = float64(speed) * force
	// here even if RHS is possible but the output of float64(speed) * force  is a flot64 value which can be assigned to speed variable which is int

	speed = 100                         // again setting the speed to 100
	speed = int(float64(speed) * force) // it will be calculated as 450 now
	fmt.Println(speed)
}
