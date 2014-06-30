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
	a := "9952345657567657"
	b := "9934534534532234"
	bigA := BigInteger.BigInt(a)
	bigB := BigInteger.BigInt(b)
	//print
	fmt.Println("a = " + bigA.ToString())
	fmt.Println("b = " + bigB.ToString())
	//'+'
	AB := bigA.Add(bigB)
	fmt.Println("a + b = " + AB.ToString())
	//'*'
	AB = bigA.Multi(bigB)
	fmt.Println("a * b = " + AB.ToString())
	//'-'
	AB = bigA.Sub(bigB)
	fmt.Println("a - b = " + AB.ToString())
	//'/'
	AB = bigA.Div(bigB)
	fmt.Println("a / b = " + AB.ToString())
	//'=='
	fmt.Printf("a == b ? %v\n", bigA.Equal(bigB))
	//'max'
	fmt.Printf("max(a, b) = %v\n", BigInteger.MaxBigInt(bigA, bigB).ToString())
	//'min'
	fmt.Printf("min(a, b) = %v\n", BigInteger.MinBigInt(bigA, bigB).ToString())
}
结果:
a = 9952345657567657
b = 9934534534532234
a + b = 19886880192099891
a * b = 98871921634707803646676702355738
a - b = 17811123035423
a / b = this big int is illegal!
a == b ? false
max(a, b) = 9952345657567657
min(a, b) = 9934534534532234
```