package params

import "reflect"

func (t *Value) Bool() bool {
	var i bool
	t.Parse(&i)
	return i
}

func (t *Value) Int64() int64 {
	var i int64
	t.Parse(&i)
	return i
}

func (t *Value) Int() int {
	return int(t.Int64())
}

func (t *Value) Int32() int32 {
	return int32(t.Int64())
}

func (t *Value) Int16() int16 {
	return int16(t.Int64())
}

func (t *Value) Int8() int8 {
	return int8(t.Int64())
}

func (t *Value) Uint64() uint64 {
	var i uint64
	t.Parse(&i)
	return i
}

func (t *Value) Uint() uint {
	return uint(t.Uint64())
}

func (t *Value) Uint32() uint32 {
	return uint32(t.Uint64())
}

func (t *Value) Uint16() uint16 {
	return uint16(t.Uint64())
}

func (t *Value) Uint8() uint8 {
	return uint8(t.Uint64())
}

func (t *Value) Byte() byte {
	return byte(t.Uint64())
}

func (t *Value) Rune() rune {
	return rune(t.Uint64())
}

func (t *Value) String() string {
	var i string
	t.Parse(&i)
	return i
}

func (t *Value) Float64() float64 {
	var i float64
	t.Parse(&i)
	return i
}

func (t *Value) Float32() float32 {
	return float32(t.Float64())
}

func (t *Value) Complex128() complex128 {
	var i complex128
	t.Parse(&i)
	return i
}

func (t *Value) Complex64() complex64 {
	return complex64(t.Complex128())
}

func (t *Value) Interface() interface{} {
	if !t.IsValid() {
		return nil
	}
	return t.v.Interface()
}

// 类型判断
func (t *Value) IsValid() bool {
	return t.v.IsValid()
}

func (t *Value) IsNil() bool {
	return t.v.IsNil()
}

func (t *Value) IsString() bool {
	return t.Kind() == reflect.String
}

func (t *Value) IsInt() bool {
	switch t.Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return true
	}
	return false
}

func (t *Value) IsUint() bool {
	switch t.Kind() {
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		return true
	}
	return false
}

func (t *Value) IsFloat() bool {
	switch t.Kind() {
	case reflect.Float64, reflect.Float32:
		return true
	}
	return false
}

func (t *Value) IsComplex() bool {
	switch t.Kind() {
	case reflect.Complex128, reflect.Complex64:
		return true
	}
	return false
}

func (t *Value) IsArray() bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

func (t *Value) IsMap() bool {
	return t.Kind() == reflect.Map
}

func (t *Value) IsBool() bool {
	return t.Kind() == reflect.Bool
}

// 类型
func (t *Value) Kind() reflect.Kind {
	//t.elem()
	return t.v.Kind()
}

func (t *Value) Type() reflect.Type {
	//t.elem()
	if !t.IsValid() {
		return reflect.Type(nil)
	}
	return t.v.Type()
}
