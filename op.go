package interval

type BinaryOpFunc func(float64, float64, bool, bool) float64

func makeBinaryOpFunc(f func(float64, float64) float64) BinaryOpFunc {
	return func(val0, val1 float64, has0, has1 bool) float64 {
		if has0 && has1 {
			return f(val0, val1)
		}
		if has1 {
			return val1
		}
		return val0
	}
}
