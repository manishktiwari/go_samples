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
}
