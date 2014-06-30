// BigInteger
package BigInteger

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
