package readpic 

import (
	"io"
	"image/jpeg/huffman.go"
)
func Checks(e error)
{
	if( e != nil)
	{
		panic(e)
	}
}

func ConvertImageToBit(w io.Writer,r io.Reader ) error {
	img, err := jpeg.Decode(r)
	Checks(err)
	bit, err := Bit.BitWriter(img)
	Checks(err)
	return bit 
}
