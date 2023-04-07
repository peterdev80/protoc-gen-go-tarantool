package tarantool

import "google.golang.org/protobuf/compiler/protogen"

type encoderGen string

func (eg encoderGen) optionalEncode(g *protogen.GeneratedFile, f *protogen.Field) {
	encstr, _, _ := typeCmd(kind(f))
	if f.Enum != nil {
		g.P("if err := enc.", encstr, "(int32(", eg, ".", f.GoName, ".Number())); err != nil {")
	} else {
		g.P("if err := enc.", encstr, "(", eg, ".", f.GoName, "); err != nil {")
	}

	g.P("return err")
	g.P("}")
}

func (eg encoderGen) repeatedEncode(g *protogen.GeneratedFile, goname string) {
	g.P("if err := enc.Encode(", eg, ".", goname, "); err != nil {")
	g.P("return err")
	g.P("}")
}

// encoder формирование EncodeMsgpack.
func encoder(g *protogen.GeneratedFile, msg *protogen.Message) {
	var (
		eny encoderGen = "y"
		enx encoderGen = "x"
	)

	// исключаем из len one off
	ln := lenGen(msg)

	g.P("if err := enc.EncodeArrayLen(", ln, "); err != nil {")
	g.P("return err")
	g.P("}")
	mof := map[string]struct{}{}

	for _, f := range msg.Fields {
		if of := f.Oneof; of != nil {
			if _, ok := mof[of.GoName]; !ok {
				g.P("", "switch y := x.", of.GoName, ".(type){")
				for _, ff := range of.Fields {

					g.P("case *", ff.GoIdent, ":")
					eny.optionalEncode(g, ff)
				}
				g.P("}")
				mof[of.GoName] = struct{}{}
			}
		} else {
			if f.Desc.Cardinality().String() != "repeated" {
				enx.optionalEncode(g, f)
			} else {
				enx.repeatedEncode(g, f.GoName)
			}
		}
	}
}
