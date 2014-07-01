// example
package main

import (
	"fmt"
	"github.com/KunBetter/BigInteger"
)

func main() {
	a := "-91119932345657567657"
	b := "345345322348"
	bigA := BigInteger.BigInt(a)
	bigB := BigInteger.BigInt(b)
	//print
	fmt.Println("a = " + bigA.ToString())
	fmt.Println("b = " + bigB.ToString())
	//'+'
	AB := bigA.Add(bigB)
	fmt.Println("a + b = " + AB.ToString())
	//'-'
	AB = bigA.Sub(bigB)
	fmt.Println("a - b = " + AB.ToString())
	//'*'
	AB = bigA.Multi(bigB)
	fmt.Println("a * b = " + AB.ToString())
	//'/'
	q, r := bigA.Div(bigB)
	fmt.Printf("a / b = %v, r = %v\n", q.ToString(), r.ToString())
	//'=='
	fmt.Printf("a == b ? %v\n", bigA.Equal(bigB))
	//'max'
	fmt.Printf("max(a, b) = %v\n", BigInteger.MaxBigInt(bigA, bigB).ToString())
	//'min'
	fmt.Printf("min(a, b) = %v\n", BigInteger.MinBigInt(bigA, bigB).ToString())
}
