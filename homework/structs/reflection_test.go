package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyGamePerson struct {
	hasHouse bool   `myJson:"has_house"`
	hasGun   bool   `myJson:"has_gun"`
	respect  uint8  `myJson:"respect"`
	mana     uint16 `myJson:"mana"`
	x        int32  `myJson:"x"`
	name     string `myJson:"name"`
	hits     []int  `myJson:"hits"`
}

type jsonBuilder struct {
	strings.Builder
}

func (b *jsonBuilder) writeField(fieldName string) (err error) {
	err = b.WriteByte('"')
	if err != nil {
		return
	}

	_, err = b.WriteString(fieldName)
	if err != nil {
		return
	}

	err = b.WriteByte('"')
	if err != nil {
		return
	}

	err = b.WriteByte(':')
	if err != nil {
		return
	}

	return
}

func (b *jsonBuilder) writeBoolVal(val bool) (err error) {
	if val {
		_, err = b.WriteString("true")
	} else {
		_, err = b.WriteString("false")
	}

	return
}

func (b *jsonBuilder) writeIntVal(val int64) (err error) {
	_, err = b.WriteString(strconv.FormatInt(val, 10))
	return
}

func (b *jsonBuilder) writeUintVal(val uint64) (err error) {
	_, err = b.WriteString(strconv.FormatUint(val, 10))
	return
}

func (b *jsonBuilder) writeStringVal(str string) (err error) {
	err = b.WriteByte('"')
	if err != nil {
		return
	}

	_, err = b.WriteString(str)
	if err != nil {
		return
	}

	err = b.WriteByte('"')
	if err != nil {
		return
	}

	return
}

func (b *jsonBuilder) writeFieldWithValOfBool(fieldName string, val bool) (err error) {
	err = b.writeField(fieldName)
	if err != nil {
		return
	}

	err = b.writeBoolVal(val)
	return
}

func (b *jsonBuilder) writeFieldWithValOfInt(fieldName string, val int64) (err error) {
	err = b.writeField(fieldName)
	if err != nil {
		return
	}

	err = b.writeIntVal(val)
	return
}

func (b *jsonBuilder) writeFieldWithValOfUint(fieldName string, val uint64) (err error) {
	err = b.writeField(fieldName)
	if err != nil {
		return
	}

	err = b.writeUintVal(val)
	return
}

func (b *jsonBuilder) writeFieldWithValOfString(fieldName string, val string) (err error) {
	err = b.writeField(fieldName)
	if err != nil {
		return
	}

	err = b.writeStringVal(val)
	return
}

func (b *jsonBuilder) addSlice(fieldName string, rVal reflect.Value) (err error) {
	b.writeField(fieldName)
	b.WriteByte('[')
	n := rVal.Len()
LOOP:
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		rValElem := rVal.Index(i)
		switch rValElem.Kind() {
		case reflect.Bool:
			err = b.writeBoolVal(rValElem.Bool())
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			err = b.writeIntVal(rValElem.Int())
		case reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr:
			err = b.writeUintVal(rValElem.Uint())
		case reflect.String:
			err = b.writeStringVal(rValElem.String())
		default:
			err = fmt.Errorf("unsupported slice element of kind `%s`", rValElem.Kind())
			break LOOP
		}
	}
	b.WriteByte(']')

	return
}

func (p MyGamePerson) ToJson() (_ string, err error) {
	b := jsonBuilder{}
	b.WriteByte('{')

	rType := reflect.TypeOf(p)
	rVal := reflect.ValueOf(p)

LOOP:
	for i := 0; i < rType.NumField(); i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		rField := rType.Field(i)
		fName := rField.Tag.Get("myJson")
		if fName == "" {
			continue
		}

		switch rField.Type.Kind() {
		case reflect.Bool:
			err = b.writeFieldWithValOfBool(fName, rVal.Field(i).Bool())
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			err = b.writeFieldWithValOfInt(fName, rVal.Field(i).Int())
		case reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr:
			err = b.writeFieldWithValOfUint(fName, rVal.Field(i).Uint())
		case reflect.String:
			err = b.writeFieldWithValOfString(fName, rVal.Field(i).String())
		case reflect.Slice:
			err = b.addSlice(fName, rVal.Field(i))
		default:
			err = fmt.Errorf("unsupported struct field of kind `%s`", rField.Type.Kind())
			break LOOP
		}

	}

	if err != nil {
		return
	}

	b.WriteByte('}')
	return b.String(), nil
}

func TestReflection(t *testing.T) {

	tests := []struct {
		input MyGamePerson
		json  string
		err   error
	}{
		{
			input: MyGamePerson{
				hasHouse: true,
				respect:  1,
				mana:     2,
				name:     "test",
				hits:     []int{1, 2, 3},
			},
			json: `{"has_house":true,"has_gun":false,"respect":1,"mana":2,"x":0,"name":"test","hits":[1,2,3]}`,
		},
		{
			input: MyGamePerson{},
			json:  `{"has_house":false,"has_gun":false,"respect":0,"mana":0,"x":0,"name":"","hits":[]}`,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test#%d", idx), func(t *testing.T) {
			json, err := tt.input.ToJson()
			assert.Equal(t, tt.json, json)
			assert.Equal(t, tt.err, err)
		})
	}
}
