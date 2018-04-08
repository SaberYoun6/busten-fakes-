/*Table start marker: 0xFFC4
  Next 2 bytes are length of table, including themselves
  Next 4 bits: 0 - DC table, 1 - AC table
  Next 4 bits: table identifier 0 - 3
  Next 16 bytes are values for each bit length
	ie 0x000203 has 0 1 bit identifiers, 2 2 bit identifiers,
		and 3 3 bit identifiers
  Bytes till end describe the codes, the number of bytes are the sum of 
	the number of the identifiers
*/

/*
  huffman.go
  
  Represents a huffman table

*/

package main

import (
	"github.com/dgryski/go-bitstream"
	"fmt"
)

type Node struct {
	parent *Node
	left *Node
	right *Node
	leaf bool
	code byte
	count int
}

type Huffman struct {
	root *Node
	dc bool //if true is DC.  defined by the high 4 bits of the 3rd byte after 0xffc4
	identifier int //0 - 3, defined by the low bits of the 3rd byte after 0xffc4
	highest int
}

//takes a bitreader, then returns the next code and the bitreader
func (h *Huffman) getCode(imageData *bitstream.BitReader) byte {
	n := h.root

	for {
		b, err := imageData.ReadBit()

		if err != nil {
			fmt.Println(err)
		}

		if b == bitstream.Zero {
			n = n.left
		} else {
			n = n.right
		}

		if n.leaf {
			break
		}
	}

	return n.code
}

func newNode(l bool, c int) *Node {
	n := new(Node)
	n.leaf = l
	n.count = c
	return n
}

func (n *Node) setRight(c *Node) *Node {
	n.right = c
	c.parent = n
	return n.right
}

func (n *Node) setLeft(c *Node) *Node {
	n.left = c
	c.parent = n
	return n.left
}

func addLevel(n *Node) {
	if n.left != nil {
		addLevel(n.left)
	}
	if n.right != nil {
		addLevel(n.right)
	}
	if !n.leaf && n.left == nil && n.right == nil {
		n.setLeft(newNode(false, n.count * 2 + 1))
		n.setRight(newNode(false, n.count * 2 + 2))
	}
}

func nextRight(cur *Node, prev *Node) *Node {
	if prev == nil {
		return nextRight(cur.parent, cur)
	} else if cur.left == prev {
		return leftmostChild(cur.right)
	} else {
		return nextRight(cur.parent, cur)
	}
}

func leftmostChild(n *Node) *Node {
	if n.left != nil {
		return leftmostChild(n.left)
	} else {
		return n
	}
}
