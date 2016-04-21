package values

import (
	encode "encoding/json"
	"fmt"

	"reflect"
)

func (t *Value) convert(typ reflect.Type) (ret reflect.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			//ret = reflect.New(typ).Elem()
			err = fmt.Errorf("value convert error: %v", x)
		}
	}()
	if t.Type() == typ {
		return t.v, nil
	}
	if t.Kind() == reflect.String && typ.Kind() != reflect.String {
		v := reflect.New(typ)
		var i interface{}
		//		var t = encode.NewDecoder(bytes.NewBuffer([]byte(t.v.String())))
		//		t.UseNumber()
		//		t.Decode(&i)
		encode.Unmarshal([]byte(t.v.String()), &i)
		NewValue(i).parse(v)
		return v.Elem(), nil
	} else if t.v.Kind() != reflect.String && typ.Kind() == reflect.String {
		b, _ := encode.Marshal(t.Interface())
		return valueOf(string(b)), nil
	}
	//	else if (t.Kind() == reflect.Map && typ.Kind() != reflect.Map) || (t.Kind() != reflect.Map && typ.Kind() == reflect.Map) {
	//		//		err = fmt.Errorf("value convert error: %v to %v", t.Kind(), typ.Kind())
	//		//		return
	//	}

	return t.v.Convert(typ), nil
}

func (t *Value) MapIndex(k interface{}) *Value {
	typ := t.v.Type()
	kv := reflect.New(typ.Key())
	NewValue(k).parse(kv)
	return newValue(t.v.MapIndex(kv.Elem()))
}

func (t *Value) Index(i int) *Value {
	return newValue(t.v.Index(i))
}

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

func (t *Value) Len() int {
	return t.v.Len()
}
