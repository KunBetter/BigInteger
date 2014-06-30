// Operations
package BigInteger

func (bi *BigInteger) Add(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	positive := true
	if bi.Positive != other.Positive {
		if bi.Positive {
			return bi.Sub(other.Abs())
		} else {
			return other.Sub(bi.Abs())
		}
	} else {
		positive = bi.Positive
	}
	bLen := len(bi.Value)
	oLen := len(other.Value)
	minLen := min(bLen, oLen)
	maxLen := max(bLen, oLen)
	nbi := &BigInteger{
		Positive: positive,
		Value:    make([]int32, maxLen),
	}
	var Carry int32 = 0 //进位
	for i := 0; i < minLen; i++ {
		tV := bi.Value[i] + other.Value[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	for i := minLen; i < bLen; i++ {
		tV := bi.Value[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	for i := minLen; i < oLen; i++ {
		tV := other.Value[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	if Carry > 0 {
		nbi.Value = append(nbi.Value, Carry)
	}
	return nbi
}

func (bi *BigInteger) Multi(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	positive := true
	if bi.Positive != other.Positive {
		positive = false
	}
	bLen := len(bi.Value)
	oLen := len(other.Value)
	bufVLen := bLen + oLen - 1
	bufV := make([]int32, bufVLen)
	for j := 0; j < oLen; j++ {
		for i := 0; i < bLen; i++ {
			bufV[i+j] += other.Value[j] * bi.Value[i]
		}
	}
	nbi := &BigInteger{
		Positive: positive,
		Value:    make([]int32, bufVLen),
	}
	var Carry int32 = 0 //进位
	for i := 0; i < bufVLen; i++ {
		tV := bufV[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	if Carry > 0 {
		nbi.Value = append(nbi.Value, Carry)
	}
	return nbi
}

func (bi *BigInteger) Sub(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	if bi.Positive != other.Positive {
		if bi.Positive {
			return bi.Add(other.Abs())
		} else {
			nSub := other.Add(bi.Abs())
			nSub.Positive = false
			return nSub
		}
	}
	ba := bi.Abs()
	oa := other.Abs()
	if ba.GreaterAbs(oa) {
		tSub := ba.subAbs(oa)
		if !bi.Positive {
			tSub.Positive = false
		}
		return tSub
	} else {
		tSub := oa.subAbs(ba)
		if bi.Positive {
			tSub.Positive = false
		}
		return tSub
	}
}

func (bi *BigInteger) subAbs(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	bLen := len(bi.Value)
	oLen := len(other.Value)
	maxLen := max(bLen, oLen)
	nbi := &BigInteger{
		Positive: true,
		Value:    make([]int32, maxLen),
	}
	var Carry int32 = 0 //借位
	for i := 0; i < maxLen; i++ {
		tV := bi.Value[i] - other.Value[i] - Carry
		if tV < 0 {
			Carry = 1
			nbi.Value[i] = tV + HEX

		} else {
			nbi.Value[i] = tV
			Carry = 0
		}
	}
	return nbi
}

func (bi *BigInteger) Div(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	return nil
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
