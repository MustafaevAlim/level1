package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	aBig := new(big.Int)
	bBig := new(big.Int)
	aBig.SetString(a, 10)
	bBig.SetString(b, 10)

	resAdd := new(big.Int).Add(aBig, bBig)
	resDiv := new(big.Int).Div(aBig, bBig)
	resMul := new(big.Int).Mul(aBig, bBig)
	resSub := new(big.Int).Sub(aBig, bBig)

	fmt.Println("Add: ", resAdd)
	fmt.Println("Div: ", resDiv)
	fmt.Println("Mul: ", resMul)
	fmt.Println("Sub: ", resSub)
}
