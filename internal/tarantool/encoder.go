package tarantool

import "google.golang.org/protobuf/compiler/protogen"

type encoderGen string

func (eg encoderGen) optionalEncode(g *protogen.GeneratedFile, encstr, goname string) {
	g.P("if err := enc.", encstr, "(", eg, ".", goname, "); err != nil {")
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
		encod, _, _ := typeCmd(kind(f))
		if of := f.Oneof; of != nil {
			if _, ok := mof[of.GoName]; !ok {

				g.P("", "switch y := x.", of.GoName, ".(type){")
				for _, ff := range of.Fields {
					encodOf, _, _ := typeCmd(kind(ff))
					g.P("case *", ff.GoIdent, ":")
					eny.optionalEncode(g, encodOf, ff.GoName)
				}
				g.P("}")
				mof[of.GoName] = struct{}{}
			}
		} else {
			if f.Desc.Cardinality().String() != "repeated" {
				enx.optionalEncode(g, encod, f.GoName)
			} else {
				enx.repeatedEncode(g, f.GoName)
			}
		}

	}
}
