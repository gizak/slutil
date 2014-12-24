package slutil

import "reflect"

func CopyIntSl(sl []int) []int {
	nsl := make([]int, len(sl))
	copy(nsl, sl)
	return nsl
}

func CopyFloatSl(sl []float64) []float64 {
	nsl := make([]float64, len(sl))
	copy(nsl, sl)
	return nsl
}

func Copy(sl interface{}) interface{} {
	v := reflect.ValueOf(sl)
	if v.Kind() != reflect.Slice {
		panic("input is not a slice")
	}
	nsl := reflect.MakeSlice(v.Type(), v.Len(), v.Len())

	reflect.Copy(nsl, v)
	return nsl.Interface()
}

func TotalInt(sl []int) int {
	sum := 0
	for _, v := range sl {
		sum += v
	}
	return sum
}

func TotalFloat(sl []float64) float64 {
	sum := 0.0
	for _, v := range sl {
		sum += v
	}
	return sum
}

func parseTableArgs(args []int) (min, step, max, n int) {
	min = 0
	step = 1
	max = args[0]
	if len(args) == 2 {
		max = args[1]
		min = args[0]
	}

	if len(args) == 3 {
		min = args[0]
		step = args[1]
		max = args[2]
	}

	n = (max - min) / step
	return
}

// args: max -> 0..<max
//       min, max -> min<=..<max
//       min, step, max -> min<=..<max
func Table(f interface{}, args ...int) interface{} {
	min, step, _, n := parseTableArgs(args)
	// f chk needed
	typ := reflect.SliceOf(reflect.TypeOf(f).Out(0))
	nsl := reflect.MakeSlice(typ, n, n)
	fv := reflect.ValueOf(f)
	for i := 0; i < n; i++ {
		ii := min + step*i
		nsl.Index(i).Set(fv.Call([]reflect.Value{reflect.ValueOf(ii)})[0])
	}
	return nsl.Interface()
}

func TableInt(f func(int) int, args ...int) []int {
	min, step, _, n := parseTableArgs(args)
	if f == nil {
		f = func(i int) int { return i }
	}

	sl := make([]int, n)
	for i := 0; i < n; i++ {
		sl[i] = f(min + i*step)
	}
	return sl
}

func TableFloat64(f func(int) float64, args ...int) []float64 {
	min, step, _, n := parseTableArgs(args)

	sl := make([]float64, n)
	for i := 0; i < n; i++ {
		sl[i] = f(min + i*step)
	}
	return sl
}
