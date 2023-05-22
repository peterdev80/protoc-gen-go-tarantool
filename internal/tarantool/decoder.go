package tarantool

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

type decodeGen string

// decoder скалярного типа
func (dg decodeGen) decodeType(g *protogen.GeneratedFile, fieldName string) {
	g.P("if x.", fieldName, ",err= dec.", dg, "(); err != nil {")
	g.P("return err")
	g.P("} ")
}

// decoder структуры
func (dg decodeGen) decodeMessage(g *protogen.GeneratedFile, fieldName string) {
	g.P("if err := dec.Decode(&x.", fieldName, "); err != nil {")
	g.P("return err")
	g.P("}")
}

// декодер слайса
func (dg decodeGen) slcDecodeMessage(g *protogen.GeneratedFile, fieldName string) {
	g.P("l, err = dec.DecodeArrayLen()")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P("x.", fieldName, "= make([]", dg, ",l)")
	g.P("for n := 0; n < l; n++ {")
	g.P("if err := dec.Decode(&x.", fieldName, "[n]); err != nil {")
	g.P("return err")
	g.P("}")
	g.P("}")
}

// Алгоритм декодирования
func decode(g *protogen.GeneratedFile, msg *protogen.Message) {
	ln := lenGen(msg)
	g.P("var err error")
	g.P("var l int")
	g.P("if l, err = dec.DecodeArrayLen(); err != nil {")
	g.P("return err")
	g.P("}")
	g.P("if l !=", ln, " {")
	g.P("return fmt.Errorf(\"array len doesn't match: %d\", l)")
	g.P("}")
	mof := map[string]struct{}{}
	for _, f := range msg.Fields {
		if f.Desc.IsMap() {
			mapdecode(g, f)
			continue
		}
		if f.Enum != nil {
			// если enum и slice, то декодер по типу enum
			if f.Desc.Cardinality().String() == "repeated" {
				cst := decodeGen(f.Enum.GoIdent.GoName)
				cst.slcDecodeMessage(g, f.GoName)

				continue
			}
			decodeEnum(g, f)
			continue

		}
		// обработка oneof
		if f.Oneof != nil {
			name := f.Oneof.GoName
			if _, ok := mof[name]; ok {
				continue
			}
			decodeOneof(g, f, name)
			mof[f.Oneof.GoName] = struct{}{}

			continue
		}
		_, decod, gotype := typeCmd(kind(f))
		cst := decodeGen(decod)
		if f.Desc.Cardinality().String() == "repeated" {
			cst = decodeGen(gotype)
			cst.slcDecodeMessage(g, f.GoName)

			continue
		}
		if f.Message != nil {
			cst = decodeGen(f.Message.GoIdent.GoName)
			cst.decodeMessage(g, f.GoName)
		} else {
			cst.decodeType(g, f.GoName)
		}
	}
}

func mapdecode(g *protogen.GeneratedFile, f *protogen.Field) {
	k := f.Desc.MapValue().Kind().String()
	last := MapCmd.Last(k)
	in := fmt.Sprint("i", f.GoName)
	mn := fmt.Sprint("x.", f.GoName)
	g.P(mn, "=map[", MapCmd.Last(f.Desc.MapKey().Kind().String()), "]",
		last, "{}")
	g.P("dec.SetMapDecoder(func(dec *msgpack.Decoder) (interface{}, error) {")
	g.P("return dec.DecodeUntypedMap()")
	g.P("})")

	g.P("var ", in, " interface{}")
	g.P("err = dec.Decode(&", in, ")")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P("for key, value := range ", in, ".(map[interface{}]interface{}) {")
	g.P("switch a:= value.(type) {")
	t, ok := MapCmd[k]
	if !ok {
		last = k
		g.P("case ", k, ":")

	}
	for _, vl := range t {
		g.P("case ", vl, ":")
		g.P(mn, "[key.(string)]=", last, "(a)")
	}
	g.P("}")
	g.P("}")
}

// декодер enum
func decodeEnum(g *protogen.GeneratedFile, f *protogen.Field) {
	g.P("en", f.GoName, ",err:=dec.DecodeInt32()")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P(" x.", f.GoName, "=", f.Enum.GoIdent, "(", "en", f.GoName, ")")
}

// декодер oneof
func decodeOneof(g *protogen.GeneratedFile, f *protogen.Field, name string) {
	g.P("var val", name, " interface{}")
	g.P("if err=dec.Decode(&val", name, ");err!=nil{")
	g.P("return err")
	g.P("}")

	g.P("switch a:=val", name, ".(type) {")
	for _, ff := range f.Oneof.Fields {
		k := kind(ff)

		if f.Message != nil {
			k = f.Message.GoIdent.GoName
		}
		last := MapCmd.Last(k)
		t, ok := MapCmd[k]
		if !ok {
			last = k
			g.P("case ", k, ":")
			g.P("x.", name, "=&", ff.GoIdent, "{", ff.GoName, ":", "&a}")
		}
		for _, vl := range t {
			g.P("case ", vl, ":")
			g.P("x.", name, "=&", ff.GoIdent, "{", ff.GoName, ":", last, "(a)}")
		}

	}
	g.P("}")
}
