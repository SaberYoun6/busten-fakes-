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
	"strconv"
)

type Node struct {
	parent *Node
	left *Node
	right *Node
	leaf bool
	code byte
	count int
	level int
}

type Huffman struct {
	root *Node
	dc bool //if true is DC.  defined by the high 4 bits of the 3rd byte after 0xffc4
	identifier int //0 - 3, defined by the low bits of the 3rd byte after 0xffc4
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

func printTree(n *Node) {
	fmt.Print("ID: " + strconv.Itoa(n.count))
	//fmt.Print(n.leaf)
	if n.parent != nil {
		fmt.Print(" parent ID is " + strconv.Itoa(n.parent.count))
	} else {
		fmt.Print(" this is root ")
	}
	fmt.Print(" level is " + strconv.Itoa(n.level))
	if n.leaf {
		fmt.Println(" Code is " + strconv.Itoa(int(n.code)))
	} else {
		fmt.Println()
	}
	if (n.left != nil) {
		printTree(n.left)
	}
	if (n.right != nil) {
		printTree(n.right)
	}
}

func newNode(l bool, c int, g int) *Node {
	n := new(Node)
	n.leaf = l
	n.count = c
	n.level = g
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
		n.setLeft(newNode(false, n.count * 2 + 1, n.level + 1))
		n.setRight(newNode(false, n.count * 2 + 2, n.level + 1))
	}
}

func nextRight(cur *Node, prev *Node) *Node {
	//fmt.Print("Cur ID is " + strconv.Itoa(cur.count))
	//if prev != nil {
	//	fmt.Println(" Prev ID is " + strconv.Itoa(cur.count))
	//} else {
	//	fmt.Println(" no prev")
	//}
	if cur.parent == nil && prev == nil {
		return leftmostChild(cur)
	} else if prev == nil {
		return nextRight(cur.parent, cur)
	} else if cur.left == prev || cur.parent == nil {
		return leftmostChild(cur.right)
	} else {
		return nextRight(cur.parent, cur)
	}
}

func leftmostChild(n *Node) *Node {
	//fmt.Println("Cur ID is " + strconv.Itoa(n.count))
	if n.left != nil {
		return leftmostChild(n.left)
	} else {
		return n
	}
}
