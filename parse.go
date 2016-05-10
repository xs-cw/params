package params

import "reflect"

//import (
//	"fmt"
//	"reflect"
//)

func (t *Value) Parse(v interface{}) (err error) {
	//	defer func() {
	//		if x := recover(); x != nil {
	//			err = fmt.Errorf("value parse error: %v", x)
	//			//fmt.Println(err)
	//		}
	//	}()
	return t.parse(valueOf(v))
}

func (t *Value) parse(val reflect.Value) (err error) {
	t = t.elem()
	val = reflect.Indirect(val)
	typ := val.Type()
	// 如果 类型一样则调用默认 转换
	if t.Type() == typ {
		v, err := t.convert(typ)
		if err != nil {
			return err
		}
		val.Set(v)
		return nil
	}

	switch val.Kind() {
	case reflect.Slice:
		if t.Kind() == reflect.String {
			t.v, err = t.convert(typ)
			if err != nil {
				return err
			}
		}
		val.Set(reflect.MakeSlice(typ, t.Len(), t.Len()))
		for i := 0; i != t.Len(); i++ {
			t.Index(i).parse(val.Index(i))
		}
	case reflect.Array:
		if t.Kind() == reflect.String {
			t.v, err = t.convert(typ)
			if err != nil {
				return err
			}
		}
		l := val.Len()
		if l > t.Len() {
			l = t.Len()
		}
		for i := 0; i != l; i++ {
			t.Index(i).parse(val.Index(i))
		}
	case reflect.Map:
		if t.Kind() == reflect.String {
			t.v, err = t.convert(typ)
			if err != nil {
				return err
			}
		}
		val.Set(reflect.MakeMap(typ))
		ktyp := typ.Key()
		vtyp := typ.Elem()
		for _, i := range t.Keys() {
			kk, err := NewValue(i).convert(ktyp)
			if err != nil {
				continue
			}
			vk, err := t.Index(i).convert(vtyp)
			if err != nil {
				continue
			}
			val.SetMapIndex(kk, vk)
		}
	case reflect.Struct:
		if t.Kind() == reflect.String {
			t.v, err = t.convert(typ)
			if err != nil {
				return err
			}
		}
		for i := 0; i != val.NumField(); i++ {
			fk := val.Field(i)
			fname := typ.Field(i).Name
			d := t.Index(fname)
			d.parse(fk)
		}
	case reflect.Ptr:
		val.Set(reflect.New(typ.Elem()))
		err = t.parse(val.Elem())
	default:
		d, err := t.convert(val.Type())
		if err != nil {
			return err
		}
		val.Set(d)
	}
	return
}
