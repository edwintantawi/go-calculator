package stack

type stringStack []string

func NewStrings() stringStack {
	return stringStack{}
}

func (s *stringStack) Push(d string) {
	*s = append(*s, d)
}

func (s *stringStack) Pop() string {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

type floatStack []float64

func NewFloats() floatStack {
	return floatStack{}
}

func (s *floatStack) Push(d float64) {
	*s = append(*s, d)
}

func (s *floatStack) Pop() float64 {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}
