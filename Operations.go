// Operations
package BigInteger

func (bi *BigInteger) Add(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	bLen := len(bi.Value)
	oLen := len(other.Value)
	minLen := min(bLen, oLen)
	maxLen := max(bLen, oLen)
	nbi := &BigInteger{
		Value: make([]int32, maxLen),
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
		Value: make([]int32, bufVLen),
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
