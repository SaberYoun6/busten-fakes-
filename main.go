package main
import "fmt"
import "io/ioutil"
// this function is here to see if any error will show up
func Check(e error) error
{
	// this is used to see if nothing isn't found and then return a nil if there are things found
	if err != nil {
		return e
	}
}


func main(){
	// this create two variable one of which is bitslice and they other ie err 
	bitslice, err :=  ioutil.ReadFile("/home/winnethepooh/Pictures/alter.jpg")
	Check(err)
	// this does the samething as a above
	byteslice, err:= ioutil.ReadFile("/home/winnethepooh/Pictures/2017-17-24-0645445.jpg")
	// this proved that if there is no error then the result will work
	Check(err)
	s  := make([]byte,len(bitslice))
	sc := make([]byte,len(byteslice))
	for i :=0 ; i <= len(bitslice){
		s[i] 
		for j :=  0; j <=len(byteslice){	
			if  {

			}
			else if {

			}
			else {
			}


		}

	}
	fmt.Print(bitslice)
}




