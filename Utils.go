// Utils
package BigInteger

const HEX = 10000 //Hexadecimal 进制

func min(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func MaxBigIntAbs(a, b *BigInteger) *BigInteger {
	aLen := len(a.Value)
	bLen := len(b.Value)
	if aLen > bLen {
		return a
	} else if aLen < bLen {
		return b
	} else {
		for i := aLen - 1; i >= 0; i-- {
			if a.Value[i] > b.Value[i] {
				return a
			} else if a.Value[i] < b.Value[i] {
				return b
			}
		}
	}
	return a
}

func MinBigIntAbs(a, b *BigInteger) *BigInteger {
	aLen := len(a.Value)
	bLen := len(b.Value)
	if aLen > bLen {
		return b
	} else if aLen < bLen {
		return a
	} else {
		for i := aLen - 1; i >= 0; i-- {
			if a.Value[i] > b.Value[i] {
				return b
			} else if a.Value[i] < b.Value[i] {
				return a
			}
		}
	}
	return a
}

func MaxBigInt(a, b *BigInteger) *BigInteger {
	if a.Positive != b.Positive {
		if a.Positive {
			return a
		} else {
			return b
		}
	} else {
		if a.Positive {
			return MaxBigIntAbs(a, b)
		} else {
			return MinBigIntAbs(a, b)
		}
	}
}

func MinBigInt(a, b *BigInteger) *BigInteger {
	if a.Positive != b.Positive {
		if a.Positive {
			return b
		} else {
			return a
		}
	} else {
		if a.Positive {
			return MinBigIntAbs(a, b)
		} else {
			return MaxBigIntAbs(a, b)
		}
	}
}
