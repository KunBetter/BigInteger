BigInteger大数运算
==========
实现了**不限位数**大数的加、减、乘、除运算。

Multi算法
-----
http://en.wikipedia.org/wiki/Multiplication_algorithm#Fast_multiplication_algorithms_for_large_inputs  
**Fourier transform methods傅立叶变换**

Div算法
-----
http://kunbetter.github.io/

Installation
-----
go get github.com/KunBetter/BigInteger

Usage
-----
```go
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
结果:
a = -91119932345657567657
b = 345345322348
a + b = -91119932000312245309
a - b = -91119932691002890005
a * b = -31467842408239064460532284098636
a / b = -263851647, r = -260391860501
a == b ? false
max(a, b) = 345345322348
min(a, b) = -91119932345657567657
```