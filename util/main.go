package main

import(
	"flag"
	"fmt"
)

func main()  {

	var str = flag.String("reverse", "", "string for reversing")

	flag.Parse()

	runes := []rune(*str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Print(string(runes))
}