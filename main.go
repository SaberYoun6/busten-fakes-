package main

import(
	"fmt"
	"io/ioutil"
	//"os/exec"
	//"strings"
	//"bytes"
	"strconv"
)

//import "github.com/dgryski/go-bitstream"

func main() {
	// this create two variable one of which is bitslice and they other ie err 
	bitslice, err :=  ioutil.ReadFile("alter.jpg")
	if err != nil {
		fmt.Print(err)
	}
	// this does the samething as a above
	//byteslice, err := ioutil.ReadFile("2017-7-24-065445.jpg")
	// this proved that if there is no error then the result will work
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(byteslice)
	//fmt.Print(bitslice)

	huff := parseJpeg(bitslice)

	//fmt.Print(huff)
	for _, h := range huff {
		fmt.Print("Is DC: ")
		fmt.Print(h.dc)
		fmt.Println(" table identity: " + strconv.Itoa(h.identifier))
		printTree(h.root)
	}
}




