package params

import (
	"fmt"
	"testing"

	"github.com/wzshiming/ffmt"
)

type A struct {
	Sss map[string]string
}

type TT struct {
	A string
	B int
}

func TestC(t *testing.T) {
	var tt = map[string]string{}
	v := NewValue(TT{"22", 11})
	v.Parse(&tt)
	fmt.Println(tt)
}
func TestB2(t *testing.T) {
	var tt = []TT{}
	v := NewValue(`[{"A":"222000000000000.5","B":50}]`)
	v.Parse(&tt)
	fmt.Println(tt)
}

func TestB(t *testing.T) {
	var tt = map[string][]TT{}
	v := NewValue(`{"dd":[{"A":"222000000000000.5","B":50}]}`)
	v.Parse(&tt)
	fmt.Println(tt)
}

func TestA(t *testing.T) {

	v := NewValue(nil)
	v.Set([]int{1, 23, 4})
	v.SetIndex(5, "12")
	ffmt.Puts(v.String())
	var i = []int64{}
	v.Parse(&i)
	ffmt.Puts(i)

	v1 := NewValue(nil)
	v1.Set(map[string]string{"Sss": `{"a":2220000000000000.5}`})
	v1.SetIndex("22", 32323.5)

	ffmt.Puts(v1.Index(22).Float64())

	var i1 = &map[int]int{}
	v1.Parse(&i1)
	ffmt.Puts(i1)

	var i2 *A
	v1.Parse(&i2)
	ffmt.Puts(i2)

	var i3 interface{}
	v1.Parse(&i3)
	ffmt.Puts(i3)

	//	v.SetIndex(5, "hehe")
	//	v.SetIndex(5, []string{"sdad", "111"})
	//	v.SetIndex(6, 4)

	//	v.SetIndex(10, "aaaa")
	//	v.SetIndex(10, nil)
	//	v.SetIndex(11, map[string]string{
	//		"aaa": "1111",
	//	})
	//	//fmt.Println(v.Index(5).Index(0))
	//	t.Log(v)
	//	var ss [][]string
	//	v.Parse(&ss)
	//	t.Log(ss)
	//	var ii []int
	//	v.Parse(&ii)
	//	t.Log(ii)
	//	t.Log(ss)
	//	t.Log(v.Interfaces())
	//	//v.MapIndex("nimei").Set("6666")
	//	t.Log(v.Slice(3, 8).String())
	//	t.Log(v.Index(3).Int())
	//	t.Log(v.Index(11).MapStringInt())

	//t.Log(v.Index(4).Bytes())
}
