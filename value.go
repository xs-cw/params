package params

import (
	//encode "encoding/json"
	"fmt"
	"reflect"
)

func valueOf(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func typeOf(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

type Value struct {
	v reflect.Value
}

func newValue(v reflect.Value) *Value {
	return &Value{v}
}

func NewValue(v interface{}) *Value {
	va := &Value{}
	va.Set(v)
	return va
}

// 数据修改

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

	//	var tv interface{}
	//	if typ == typeOf(tv) {
	//		return t.v, nil
	//	}
	//	if t.Kind() == reflect.String {
	//		if typ.Kind() != reflect.String {
	//			v := reflect.New(typ)
	//			var i interface{}
	//			err := encode.Unmarshal([]byte(t.v.String()), &i)
	//			if err != nil {
	//				NewValue(i).parse(v)
	//				return v.Elem(), nil
	//			}
	//		}
	//	} else {
	//		if typ.Kind() == reflect.String {
	//			b, err := encode.Marshal(t.Interface())
	//			if err != nil {
	//				return valueOf(string(b)), nil
	//			}
	//		}
	//	}
	//	else if (t.Kind() == reflect.Map && typ.Kind() != reflect.Map) || (t.Kind() != reflect.Map && typ.Kind() == reflect.Map) {
	//		//		err = fmt.Errorf("value convert error: %v to %v", t.Kind(), typ.Kind())
	//		//		return
	//	}
	// 这里 还要 加 interface 转 map[string]string  或者 []string 这种转换
	return t.v.Convert(typ), nil
}

func (t *Value) elem() *Value {
	for t.v.Kind() == reflect.Ptr || t.v.Kind() == reflect.Interface {
		t.v = t.v.Elem()
	}
	return t
}

func (t *Value) Keys() (ret []interface{}) {
	switch t.Kind() {
	case reflect.Map:
		for _, k := range t.v.MapKeys() {
			ret = append(ret, k.Interface())
		}
	case reflect.Slice, reflect.Array:
		l := t.Len()
		for i := 0; i != l; i++ {
			ret = append(ret, i)
		}
	case reflect.Struct:
		for i := 0; i != t.v.NumField(); i++ {
			t.v.FieldByNameFunc(func(name string) bool {
				ret = append(ret, name)
				return false
			})
		}
	}
	return
}

func (t *Value) Len() int {
	switch t.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return t.v.Len()
	case reflect.Struct:
		return t.Type().FieldAlign()
	}
	return 0
}

func (t *Value) index(key *Value) *Value {
	val := valueOf(nil)
	switch t.Kind() {
	case reflect.Map:
		temp, err := key.convert(t.v.Type().Key())
		if err != nil {
			return newValue(val)
		}
		val = t.v.MapIndex(temp)
	case reflect.Slice, reflect.Array:
		val = t.v.Index(key.Int())
	case reflect.Struct:
		val = t.v.FieldByName(key.String())
	}
	return newValue(val)
}

func (t *Value) Index(key interface{}) *Value {
	return t.index(NewValue(key))
}

func (t *Value) setIndex(key *Value, val *Value) {
	switch t.Kind() {
	case reflect.Map:
		temp1, err := key.convert(t.v.Type().Key())
		if err != nil {
			return
		}
		temp2, err := key.convert(t.v.Type().Elem())
		if err != nil {
			return
		}
		t.v.SetMapIndex(temp1, temp2)
	case reflect.Slice, reflect.Array:
		i := key.Int()
		if t.Len() <= i {
			t.v = reflect.AppendSlice(t.v, reflect.MakeSlice(t.v.Type(), i-t.Len()+1, i-t.Len()+1))
		}
		d := t.v.Index(i)
		vv := reflect.New(d.Type())
		val.parse(vv)
		d.Set(vv.Elem())
	case reflect.Struct:
		val.parse(t.index(key).v)
	}
	return
}

func (t *Value) SetIndex(key interface{}, val interface{}) {
	t.setIndex(NewValue(key), NewValue(val))
	return
}

func (t *Value) Set(v interface{}) *Value {
	return t.set(valueOf(v))
}

func (t *Value) set(v reflect.Value) *Value {
	t.v = v
	return t
}

func (t *Value) Slice(i, j int) *Value {
	if j > t.Len() {
		j = t.Len()
	}
	if i > j || i < 0 {
		return NewValue(nil)
	}
	return newValue(t.v.Slice(i, j))
}

func (t *Value) Indirect() {
	t.v = reflect.Indirect(t.v)
}
