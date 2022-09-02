package myutils

func RangeSSN[T int | float64](start, step T, num int) []T {
	array := make([]T, num)
	array[0] = start
	for i := 1; i < num; i++ {
		array[i] = array[i-1] + step
	}
	return array
}

func RangeSN[T int | float64](start T, num int) []T {
	return RangeSSN(start, 1, num)
}

func Range(num int) []int {
	return RangeSSN(0, 1, num)
}

func RangeSEN[T int | float64](start, end T, num int) []float64 {
	step := float64(end) / float64(num-1)
	return RangeSSN(float64(start), step, num)
}

func RangeSE(start, end int) []int {
	num := end - start
	return RangeSSN(start, 1, num)
}
