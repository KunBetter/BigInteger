// example
package main

import (
	"fmt"
	"github.com/KunBetter/BigInteger"
)

func main() {
	a := "9912345657567657"
	b := "9934534534532234"
	bigA := BigInteger.BigInt(a)
	bigB := BigInteger.BigInt(b)
	//print
	fmt.Println("a = " + bigA.ToString())
	fmt.Println("b = " + bigB.ToString())
	//+
	AB := bigA.Add(bigB)
	fmt.Println("a + b = " + AB.ToString())
	//*
	AB = bigA.Multi(bigB)
	fmt.Println("a * b = " + AB.ToString())
}
