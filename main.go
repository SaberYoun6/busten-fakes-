package main
import(
	"fmt"
	"io/ioutil"
	"os/exec"
	"stringst "
)

//import "github.com/dgryski/go-bitstream"

func main(){
	// this create two variable one of which is bitslice and they other ie err 
	//bitslice, err :=  ioutil.ReadFile("alter.jpg")
	//if err != nil {
	//	fmt.Print(err)
	//}
	// this does the samething as a above
	byteslice, err := ioutil.ReadFile("2017-7-24-065445.jpg")
	// this proved that if there is no error then the result will work
	if err != nil {
		fmt.Print(err)
	}
	cmd:=exec.Command("exiftool","0-9","g-p")
	cmd.Stdin=ioutil.ReadFile("2017-7-24-065445.jpg")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("the  metadata should read%q\n",out.
	//fmt.Print(byteslice)
	//fmt.Print(bitslice)

	huff := parseJpeg(byteslice)

	//fmt.Print(huff)
	printTree(huff.root)
}




