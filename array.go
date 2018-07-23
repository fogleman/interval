package interval

import "math"

type Array struct {
	Intervals []Interval
}

func NewArray() *Array {
	return &Array{}
}

func NewArrayWithInterval(start, end int, value float64) *Array {
	var a Array
	a.Intervals = append(a.Intervals, Interval{start, end, value})
	return &a
}

func (a *Array) Set(start, end int, value float64) *Array {
	f := func(val0, val1 float64, has0, has1 bool) float64 {
		if has1 {
			return val1
		}
		return val0
	}
	b := NewArrayWithInterval(start, end, value)
	return a.BinaryOp(b, f)
}

func (a *Array) Min(start, end int, value float64) *Array {
	f := makeBinaryOpFunc(math.Min)
	b := NewArrayWithInterval(start, end, value)
	return a.BinaryOp(b, f)
}

func (a *Array) Max(start, end int, value float64) *Array {
	f := makeBinaryOpFunc(math.Max)
	b := NewArrayWithInterval(start, end, value)
	return a.BinaryOp(b, f)
}

func (a *Array) BinaryOp(b *Array, f BinaryOpFunc) *Array {
	vs0 := a.Intervals
	vs1 := b.Intervals
	n0 := len(vs0)
	n1 := len(vs1)
	var i0, i1, j0, j1 int
	var intervals []Interval
	for i0 < n0 && i1 < n1 {
		v0 := vs0[i0]
		v1 := vs1[i1]
		start0 := v0.Start
		start1 := v1.Start
		end0 := v0.End
		end1 := v1.End
		value0 := v0.Value
		value1 := v1.Value
		// where are we?
		p0 := start0 + j0
		p1 := start1 + j1
		if p0 == p1 {
			// both intervals are present
			p := p0
			// which ends first?
			q := minInt(end0, end1)
			// interval from p to q
			value := f(value0, value1, true, true)
			intervals = appendInterval(intervals, p, q, value)
			// update "pointers"
			d := q - p
			j0 += d
			j1 += d
			if j0 >= end0-start0 {
				j0 = 0
				i0++
			}
			if j1 >= end1-start1 {
				j1 = 0
				i1++
			}
		} else {
			// only one span is present, which one?
			if p0 < p1 {
				// distance to end of span0
				d0 := end0 - p0
				// distance to beginning of span1
				d1 := start1 - p0
				// do the min distance
				d := minInt(d0, d1)
				q := p0 + d
				// span from p0 to q
				value := f(value0, 0, true, false)
				intervals = appendInterval(intervals, p0, q, value)
				// update "pointers"
				j0 += d
				if j0 >= end0-start0 {
					j0 = 0
					i0++
				}
			} else {
				// distance to end of span1
				d1 := end1 - p1
				// distance to beginning of span0
				d0 := start0 - p1
				// do the min distance
				d := minInt(d0, d1)
				q := p1 + d
				// span from p1 to q
				value := f(0, value1, false, true)
				intervals = appendInterval(intervals, p1, q, value)
				// update "pointers"
				j1 += d
				if j1 >= end1-start1 {
					j1 = 0
					i1++
				}
			}
		}
	}
	for i0 < n0 {
		v := vs0[i0]
		value := f(v.Value, 0, true, false)
		intervals = appendInterval(intervals, v.Start+j0, v.End, value)
		j0 = 0
		i0++
	}
	for i1 < n1 {
		v := vs1[i1]
		value := f(0, v.Value, false, true)
		intervals = appendInterval(intervals, v.Start+j1, v.End, value)
		j1 = 0
		i1++
	}
	return &Array{intervals}
}

func appendInterval(intervals []Interval, start, end int, value float64) []Interval {
	n := len(intervals)
	if n > 0 {
		back := &intervals[n-1]
		if back.Value == value && back.End == start {
			back.End = end
			return intervals
		}
	}
	return append(intervals, Interval{start, end, value})
}
