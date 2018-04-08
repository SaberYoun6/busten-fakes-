/* This will take a byte slice, and then return an array of Huffman tables */

package huffman

func parseJpeg(in []byte) Huffman {
	var huff Huffman

	for i := range in {
		//If this bit and the next are a marker,
		if in[i] == 0xFF {
			i++

			if in[i] == 0xDA { //if the marker is for start of scan, we have gone through all the metadata
				break
			} else if in[i] == 0xC4 { //this is a Huffman table
				i, huff = parseHuffman(i, in)
				return huff
			}
		}
	}

}

func parseHuffman(i int, in []byte) (int, Huffman) {
	h := new(Huffman)
	h.root = newNode(false, 0)
	var nVals [16]int
	end = i + in[i+1] + in[i+2] //the next two bytes are the length of the huffman table, including themselves
	if (in[i+3] >> 4) == 0 {
		h.dc = true
	}
	h.identifier = in[i+3] & 0x0F
	i += 4

	for j := 0; j < 16; j++ {
		nVals[j] = in[i]
		i++
	}

	sum := nVals[0]
	j := 0
	k := 0
	var cur Node
	for ; i <= end; i++ {
		if sum == j {
			addLevel(h.root)
			k++
			sum += nVals[k]
		}

		if nVals[k] != 0 {
			cur = nextRight(h.root, nil)
			cur.code = in[i]
			cur.leaf = true
			j++
		}
	}


	return i, h
}
