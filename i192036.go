package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global Variables
var list []*block
var list2 []*Block_chain
var id int = 0
var blockid int = 0
var first int
var second string
var b *Block_chain

type block struct {
	transaction  string
	nonce        int
	previousHash string
}

type Block_chain struct {
	previousHash string
	currentHash  string
	str1         string
}

// type genesis struct {
// 	transaction string
// 	nonce       int
// }

// func GenesisBlock(transaction string, nonce int) *genesis {

// 	s := new(genesis)
// 	s.transaction = transaction
// 	s.nonce = nonce
// 	var genesis string = transaction + strconv.Itoa(nonce)
// 	var output = CalculateHash(genesis)
// 	fmt.Print(output)

// 	return s

// }

// -------------------------------------------------------------------------------------------------------------------------------
func NewBlock(transaction string, nonce int, previousHash string) *block {

	s := new(block)
	s.transaction = transaction
	s.nonce = nonce
	var current_hash string
	var prev_hash string
	var temp string

	if id == 0 {
		s.previousHash = ""
		var res1 string = s.transaction + strconv.Itoa(s.nonce) + s.previousHash
		current_hash = CalculateHash(res1)
		b = addBlock(s.previousHash, current_hash, res1)

	} else {
		blockid--

		var res string = s.transaction + strconv.Itoa(s.nonce) + s.previousHash
		// fmt.Printf("String to calculate Hash hehe :  %s\n", res)

		current_hash = CalculateHash(res)
		temp = list[blockid].transaction + strconv.Itoa(list[blockid].nonce) + list[blockid].previousHash
		prev_hash = CalculateHash(temp)
		b = addBlock(prev_hash, current_hash, res)
		blockid++

	}

	// fmt.Printf("NEW BLOCK HAMARa Bhhehe :  %s\n", *b)

	list2 = append(list2, b)

	id++
	blockid++
	list = append(list, s)
	return s

}

// -------------------------------------------------------------------------------------------------------------------------------
func addBlock(prev_Hash string, curr_Hash string, str string) *Block_chain {
	b := new(Block_chain)
	b.previousHash = prev_Hash
	b.currentHash = curr_Hash
	b.str1 = str
	return b
}

// -------------------------------------------------------------------------------------------------------------------------------
func DisplayBlock() {

	for i := 0; i < len(list); i++ {
		no := i
		if i == 0 {

			fmt.Printf("%s Genesis Block  %d %s\n", strings.Repeat("=", 32), no, strings.Repeat("=", 32))
			fmt.Println("\nTransaction:\t", list[i].transaction, "\nNonce:\t", list[i].nonce, "\nCurrent Hash:\t", list2[i].currentHash)

		} else {
			fmt.Printf("%s Block  %d %s\n", strings.Repeat("=", 35), no, strings.Repeat("=", 35))
			fmt.Println("\nTransaction:\t", list[i].transaction, "\nNonce:\t", list[i].nonce, "\nCurrent Hash:\t", list2[i].currentHash, "\nPrevious hash:\t", list2[i].previousHash)
		}
	}

}

// -------------------------------------------------------------------------------------------------------------------------------
func ChangeBlock() {

	fmt.Printf("Enter the Block you want to change transaction of ")
	fmt.Scanln(&first)

	if first <= len(list) {
		fmt.Println("Enter Your Transaction ")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		second = in.Text()
		list[first].transaction = second
		var str = list[first].transaction + strconv.Itoa(list[first].nonce) + list2[first].previousHash
		var out = CalculateHash(str)
		list2[first].currentHash = out
		DisplayBlock()
	} else {
		fmt.Printf("Following Block Doesn't exist ")
	}

}

// -------------------------------------------------------------------------------------------------------------------------------
func VerifyChain() {

	for i := 0; i < len(list2)-1; i++ {
		if list2[i].currentHash != list2[i+1].previousHash {
			fmt.Printf("Change Detected \n")
			fmt.Printf("Block changed at index :  %d\n", i)
			break

		} else {
			fmt.Printf("Block not changed at index :  %d\n", i)

		}
	}

}

// -------------------------------------------------------------------------------------------------------------------------------
var CalculateHash = func(stringToHash string) string {
	// fmt.Printf("String to calculate Hash :  %s\n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

// -------------------------------------------------------------------------------------------------------------------------------
func main() {
	// fmt.Print(id)

	NewBlock("Alice how are you", 4, "") //Genesis Block whatever u put as input in third argument it will take as null
	NewBlock(" Hi bob I'm Good", 9, list2[id-1].previousHash)
	NewBlock(" Hi John I'm Good too", 6, list2[id-1].previousHash)
	NewBlock(" Bye Guys", 9, list2[id-1].previousHash)
	// NewBlock(" Hi bob I'm Good3", 4, "list2[id-1].previousHash");
	// NewBlock(" Hi bob I'm Good4", 4, "list2[id-1].previousHash");
	DisplayBlock()
	ChangeBlock()
	VerifyChain()

}
