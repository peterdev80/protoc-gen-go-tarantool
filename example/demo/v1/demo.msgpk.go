// Code generated by protoc-gen-go-tarantool. DO NOT EDIT.

package pb

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

func (x *Ex) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err := enc.EncodeString(x.IdNum); err != nil {
		return err
	}
	switch y := x.Value.(type) {
	case *Ex_IData:
		if err := enc.EncodeInt64(y.IData); err != nil {
			return err
		}
	case *Ex_SData:
		if err := enc.EncodeString(y.SData); err != nil {
			return err
		}
	}
	return nil
}
func (x *Ex) DecodeMsgpack(dec *msgpack.Decoder) error {
	var err error
	var l int
	if l, err = dec.DecodeArrayLen(); err != nil {
		return err
	}
	if l != 2 {
		return fmt.Errorf("array len doesn't match: %d", l)
	}
	if x.IdNum, err = dec.DecodeString(); err != nil {
		return err
	}
	var valValue interface{}
	if err = dec.Decode(&valValue); err != nil {
		return err
	}
	switch a := valValue.(type) {
	case int8:
		x.Value = &Ex_IData{IData: int64(a)}
	case int16:
		x.Value = &Ex_IData{IData: int64(a)}
	case int32:
		x.Value = &Ex_IData{IData: int64(a)}
	case int64:
		x.Value = &Ex_IData{IData: int64(a)}
	case string:
		x.Value = &Ex_SData{SData: string(a)}
	}
	return nil
}
