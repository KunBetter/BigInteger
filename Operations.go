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
	bLen := int32(len(bi.Value))
	oLen := int32(len(other.Value))
	minLen := min(bLen, oLen)
	maxLen := max(bLen, oLen)
	nbi := &BigInteger{
		Positive: positive,
		Value:    make([]int32, maxLen),
	}
	var Carry int32 = 0 //进位
	var i int32 = 0
	for i = 0; i < minLen; i++ {
		tV := bi.Value[i] + other.Value[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	for i = minLen; i < bLen; i++ {
		tV := bi.Value[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	for i = minLen; i < oLen; i++ {
		tV := other.Value[i] + Carry
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
		tSub := ba.SubAbs(oa)
		if !bi.Positive {
			tSub.Positive = false
		}
		return tSub
	} else {
		tSub := oa.SubAbs(ba)
		if bi.Positive {
			tSub.Positive = false
		}
		return tSub
	}
}

func (bi *BigInteger) SubAbs(other *BigInteger) *BigInteger {
	if bi == nil || other == nil {
		return nil
	}
	bLen := int32(len(bi.Value))
	oLen := int32(len(other.Value))
	maxLen := max(bLen, oLen)
	nbi := &BigInteger{
		Positive: true,
		Value:    make([]int32, maxLen),
	}
	var i int32 = 0
	for i = 0; i < bLen-oLen; i++ {
		other.Value = append(other.Value, 0)
	}
	var Carry int32 = 0 //借位
	for i = 0; i < maxLen; i++ {
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

func (bi *BigInteger) MultiNum(num int32) *BigInteger {
	if bi == nil {
		return nil
	}
	bLen := len(bi.Value)
	bufV := make([]int32, bLen)
	for i := 0; i < bLen; i++ {
		bufV[i] += num * bi.Value[i]
	}
	nbi := &BigInteger{
		Positive: bi.Positive,
		Value:    make([]int32, bLen),
	}
	var Carry int32 = 0 //进位
	for i := 0; i < bLen; i++ {
		tV := bufV[i] + Carry
		nbi.Value[i] = tV % HEX
		Carry = tV / HEX
	}
	if Carry > 0 {
		nbi.Value = append(nbi.Value, Carry)
	}
	return nbi
}

func (bi *BigInteger) Div(other *BigInteger) (q, r *BigInteger) {
	if bi == nil || other == nil {
		return nil, nil
	}
	Q := []int32{}
	uLen := len(bi.Value)
	vLen := len(other.Value)
	U := make([]int32, uLen)
	for i := 0; i < uLen; i++ {
		U[i] = bi.Value[uLen-1-i]
	}
	V := make([]int32, vLen)
	for i := 0; i < vLen; i++ {
		V[i] = other.Value[vLen-1-i]
	}
	v := other
	for j := 0; j < uLen-vLen; j++ {
		x := (U[j]*HEX + U[j+1]) / V[0]
		y := x * V[1] / (V[0]*HEX + V[1])
		q := x - y
		p := v.MultiNum(q)
		Uj := BigIntSlice(U[j : j+vLen+1])
		if !p.LessEqualAbs(Uj) {
			p = p.Sub(v)
			if p.LessEqualAbs(Uj) {
				q--
			} else {
				p = p.Sub(v)
				q -= 2
			}
		}
		Ur := Uj.Sub(p)
		Q = append(Q, q)
		copy(U[j:], Ur.ToSlice())
	}
	positive := true
	if bi.Positive != other.Positive {
		positive = false
	}
	q = BigIntSlice(Q)
	q.Positive = positive
	r = BigIntSlice(U)
	r.Positive = bi.Positive
	return
}
