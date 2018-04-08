/* This will take a byte slice, and then return an array of Huffman tables */

package main
/*
import (
	"fmt"
	"strconv"
)*/

func parseJpeg(in []byte) Huffman {
	var huff Huffman

	for i := range in {
		//fmt.Println("i is: " + strconv.Itoa(i))
		//If this bit and the next are a marker,
		if in[i] == 0xFF {
			i++

			if in[i] == 0xDA { //if the marker is for start of scan, we have gone through all the metadata
				break
			} else if in[i] == 0xC4 { //this is a Huffman table
				i, huff = parseHuffman(i, in)
				break
			}
		}
	}
	return huff

}

func parseHuffman(i int, in []byte) (int, Huffman) {
	h := new(Huffman)
	h.root = newNode(false, 0, 0)
	var nVals [16]int
	end := i + int(in[i+1] << 8 + in[i+2]) //the next two bytes are the length of the huffman table, including themselves
	if (in[i+3] >> 4) == 0 {
		h.dc = true
	}
	h.identifier = int(in[i+3] & 0x0F)
	i += 4

	total := 0
	for j := 0; j < 16; j++ {
		nVals[j] = int(in[i])
		total += nVals[j]
		i++
	}

	//fmt.Println("i is at " + strconv.Itoa(i) + " end is " + strconv.Itoa(end))
	//fmt.Println(nVals)
	sum := 0
	j := 0
	k := -1
	cur := h.root
	for ; i <= end; i++ {
		//fmt.Println("i is " + strconv.Itoa(i) + " j is " + strconv.Itoa(j) + " k is " + strconv.Itoa(k) + " sum is " + strconv.Itoa(sum))
		if j == total {
			break
		}
		if sum == j {
			k++
			if k == 16 {
				break
			}
			addLevel(h.root)
			//fmt.Println("Added level")
			sum += nVals[k]
		}

		if nVals[k] != 0 {
			cur = nextRight(cur, nil)
			cur.code = in[i]
			cur.leaf = true
			//fmt.Println("ID is " + strconv.Itoa(cur.count) + " and code is " + strconv.Itoa(int(cur.code)))
			j++
		}
	}

	return i, *h
}
