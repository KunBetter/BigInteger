BigInteger大数运算
==========
目前实现了**不限位数**大数的加法和乘法。

Multi算法
-----
http://en.wikipedia.org/wiki/Multiplication_algorithm#Fast_multiplication_algorithms_for_large_inputs  
**Fourier transform methods傅立叶变换**

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
结果:
a = 9912345657567657
b = 9934534534532234
a + b = 19846880192099891
a * b = 98474540253326514286676702355738
```