// BigInteger
package BigInteger

import (
	"fmt"
)

type BigInteger struct {
	Value []int32
}

func BigInt(bis string) *BigInteger {
	buf := []byte(bis)
	bLen := len(buf)
	vLen := bLen / 4
	hLen := bLen % 4
	if hLen > 0 {
		vLen++
	}
	bi := &BigInteger{
		Value: make([]int32, vLen),
	}
	var tValue int32 = 0
	vLen--
	for i := 0; i < hLen; i++ {
		e := int32(buf[i] - 48)
		if e < 1 || e > 9 {
			return nil
		}
		tValue = tValue*10 + e
	}
	if hLen > 0 {
		bi.Value[vLen] = tValue
		vLen--
	}
	tValue = 0
	span := 0
	for i := hLen; i < bLen; i++ {
		tValue = tValue*10 + int32(buf[i]-48)
		span++
		if span == 4 {
			bi.Value[vLen] = tValue
			vLen--
			tValue = 0
			span = 0
		}
	}
	return bi
}

func (bi *BigInteger) ToString() string {
	if bi == nil {
		return "this big int is illegal!"
	}
	bLen := len(bi.Value)
	str := ""
	if bLen > 0 {
		str += fmt.Sprintf("%d", bi.Value[bLen-1])
	}
	for i := bLen - 2; i >= 0; i-- {
		str += fmt.Sprintf("%04d", bi.Value[i])
	}
	return str
}
