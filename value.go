package values

import "reflect"

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
func (t *Value) Set(v interface{}) *Value {
	return t.SetValue(valueOf(v))
}

func (t *Value) elem() *Value {
	for t.v.Kind() == reflect.Ptr || t.v.Kind() == reflect.Interface {
		t.v = t.v.Elem()
	}
	return t
}

func (t *Value) SetValue(v reflect.Value) *Value {
	t.v = v
	return t
}

func (t *Value) SetMapIndex(k, v interface{}) {
	if !t.IsMap() {
		var i map[interface{}]interface{}
		typ := typeOf(i)
		t.v = reflect.MakeMap(typ)
	}
	typ := t.v.Type()
	kv := reflect.New(typ.Key())
	vv := reflect.New(typ.Elem())
	NewValue(k).parse(kv)
	NewValue(v).parse(vv)
	t.v.SetMapIndex(kv.Elem(), vv.Elem())
	return
}

func (t *Value) SetIndex(i int, v interface{}) *Value {
	if !t.IsArray() {
		var tt []interface{}
		typ := typeOf(tt)
		t.v = reflect.MakeSlice(typ, i+1, i+1)
	}
	if t.Len() <= i {
		t.v = reflect.AppendSlice(t.v, reflect.MakeSlice(t.v.Type(), i-t.Len()+1, i-t.Len()+1))
	}
	d := t.v.Index(i)
	vv := reflect.New(d.Type())
	NewValue(v).parse(vv)
	d.Set(vv.Elem())
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

// 类型判断
func (t *Value) IsValid() bool {
	return t.v.IsValid()
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
