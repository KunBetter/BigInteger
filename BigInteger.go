// BigInteger
package BigInteger

import (
	"fmt"
)

type BigInteger struct {
	Positive bool
	Value    []int32
}

func BigInt(bis string) *BigInteger {
	positive := true
	buf := []byte(bis)
	bLen := len(buf)
	if bLen <= 0 {
		return nil
	}
	if buf[0] == '-' || buf[0] == '+' {
		if buf[0] == '-' {
			positive = false
		}
		bLen--
		buf = buf[1:]
	}
	vLen := bLen / 4
	hLen := bLen % 4
	if hLen > 0 {
		vLen++
	}
	bi := &BigInteger{
		Positive: positive,
		Value:    make([]int32, vLen),
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

func BigIntSlice(vs []int32) *BigInteger {
	vLen := len(vs)
	v := make([]int32, vLen)
	for i := 0; i < vLen; i++ {
		v[i] = vs[vLen-1-i]
	}
	nbi := &BigInteger{
		Positive: true,
		Value:    v,
	}
	return nbi
}

func (bi *BigInteger) ToSlice() []int32 {
	vLen := len(bi.Value)
	v := make([]int32, vLen)
	for i := 0; i < vLen; i++ {
		v[i] = bi.Value[vLen-1-i]
	}
	return v
}

func (bi *BigInteger) Equal(other *BigInteger) bool {
	if bi.Positive != other.Positive {
		return false
	}
	bLen := len(bi.Value)
	oLen := len(other.Value)
	if bLen != oLen {
		return false
	} else {
		for i := bLen - 1; i >= 0; i-- {
			if bi.Value[i] != other.Value[i] {
				return false
			}
		}
	}
	return true
}

func (bi *BigInteger) Abs() *BigInteger {
	nbi := &BigInteger{
		Positive: true,
		Value:    bi.Value,
	}
	return nbi
}

func (a *BigInteger) GreaterAbs(b *BigInteger) bool {
	aLen := len(a.Value)
	bLen := len(b.Value)
	if aLen > bLen {
		return true
	} else if aLen < bLen {
		return false
	} else {
		for i := aLen - 1; i >= 0; i-- {
			if a.Value[i] > b.Value[i] {
				return true
			} else if a.Value[i] < b.Value[i] {
				return false
			}
		}
	}
	return false
}

func (a *BigInteger) LessEqualAbs(b *BigInteger) bool {
	aLen := len(a.Value)
	bLen := len(b.Value)
	if aLen > bLen {
		return false
	} else if aLen < bLen {
		return true
	} else {
		for i := aLen - 1; i >= 0; i-- {
			if a.Value[i] < b.Value[i] {
				return true
			} else if a.Value[i] > b.Value[i] {
				return false
			}
		}
	}
	return true
}

func (bi *BigInteger) ToString() string {
	if bi == nil {
		return "this big int is illegal!"
	}
	bLen := len(bi.Value)
	for i := bLen - 1; i >= 0; i-- {
		if bi.Value[i] == 0 {
			bLen--
		} else {
			break
		}
	}
	str := ""
	if !bi.Positive {
		str += "-"
	}
	if bLen > 0 {
		str += fmt.Sprintf("%d", bi.Value[bLen-1])
	}
	for i := bLen - 2; i >= 0; i-- {
		str += fmt.Sprintf("%04d", bi.Value[i])
	}
	return str
}
