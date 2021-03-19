package data

import "math"

var datas =make([]string, 0, math.MaxInt16)

func Add(s string) string{
	datas = append(datas, s)
	return s
}