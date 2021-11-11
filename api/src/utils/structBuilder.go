package utils

import "reflect"

type StructBuilderType struct {
	field []reflect.StructField
}

func StructBuilder() *StructBuilderType {
	return &StructBuilderType{}
}

func (b *StructBuilderType) AddField(fname string, ftype reflect.Type) {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: fname,
			Type: ftype,
		})
}

type Struct struct {
	strct reflect.Type
	index map[string]int
}

func (b *StructBuilderType) Build() Struct {
	strct := reflect.StructOf(b.field)
	index := make(map[string]int)
	for i := 0; i < strct.NumField(); i++ {
		index[strct.Field(i).Name] = i
	}

	return Struct{strct, index}
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (s *Struct) NewInstance() *Instance {
	instance := reflect.New(s.strct).Elem()
	return &Instance{instance, s.index}
}

func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

func (i *Instance) SetInt(name string, value int) {
	i.Field(name).SetInt(int64(value))
}

func (i *Instance) SetFloat(name string, value float64) {
	i.Field(name).SetFloat(value)
}

func (i *Instance) Value() interface{} {
	return i.internal.Interface()
}

func (i *Instance) Pointer() interface{} {
	return i.internal.Addr().Interface()
}
