package main
import "fmt"
import "io/ioutil"
//todo: add this back in import "https://github.com/dgryski/go-bitstream"

func main(){
	// this create two variable one of which is bitslice and they other ie err 
	bitslice, err :=  ioutil.ReadFile("/home/winnethepooh/Pictures/alter.jpg")
	if err != nil {
		fmt.Print(err)
	}
	// this does the samething as a above
	byteslice, err:= ioutil.ReadFile("/home/winnethepooh/Pictures/2017-17-24-0645445.jpg")
	// this proved that if there is no error then the result will work
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(byteslice)
	fmt.Print(bitslice)
}




