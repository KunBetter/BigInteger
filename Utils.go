// Utils
package BigInteger

const HEX = 10000 //Hexadecimal 进制

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxBigInt(a, b *BigInteger) *BigInteger {
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

func MinBigInt(a, b *BigInteger) *BigInteger {
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
