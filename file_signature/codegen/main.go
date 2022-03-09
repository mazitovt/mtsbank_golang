package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"html/template"
	"log"
	"mtsbank_golang/file_signature/packing"
	"mtsbank_golang/file_signature/signature"
	"os"
	"reflect"
	"strings"
)

type tpl struct {
	FieldName string
}

var (
	intTpl = template.Must(template.New("intTpl").Parse(`
	// {{.FieldName}}
	var sizeUintRaw uint64
	binary.Read(r, binary.BigEndian, &sizeUintRaw)
	in.{{.FieldName}} = uint(sizeUintRaw)

`))
	strTpl = template.Must(template.New("strTpl").Parse(`
	// {{.FieldName}}
	data, e = packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	in.{{.FieldName}} = string(data)
`))
	sliceTpl = template.Must(template.New("sliceTpl").Parse(`
	// {{.FieldName}}
	data, e = packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	in.{{.FieldName}} = data
`))

	timeTpl = template.Must(template.New("timeTpl").Parse(`
	// {{ .FieldName }}
	data, e := packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	var t time.Time
	e = t.GobDecode(data)
	if e != nil {
		return e
	}

	in.{{ .FieldName }} = t
	
`))
)

func test() {

	sig := signature.NewSignatureSha256FromFile()
	sig2 := signature.NewSignatureSha256FromFile()

	sig.ParseString("2022-02-27T21:02:09+05:00::11::sourceName.txt====sign====a����3\u001B+���g}�H�(\u0016:A�5)ވ�r1���\u0004")

	sig.SetSizeUint(55)

	b, _ := packing.PackSignature(sig)

	e := sig2.Unpack(b.Bytes())

	fmt.Printf("error: %v", e)

	fmt.Println(sig2)
	fmt.Println(sig)

	fmt.Println(sig.Equal(sig2))

}

func main() {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(out, "package ", node.Name.Name)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `
import (
	"bytes"
	"encoding/binary"
	"mtsbank_golang/file_signature/packing"
	"time"
)`)

	for _, f := range node.Decls {
		g, ok := f.(*ast.GenDecl)
		if !ok {
			fmt.Printf("SKIP %#T is not ast.GenDecl\n", f)
			continue
		}
		for _, spec := range g.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				fmt.Printf("SKIP %T is not ast.TypeSpec\n", spec)
				continue
			}

			currStruct, ok := currType.Type.(*ast.StructType)
			if !ok {
				fmt.Printf("SKIP %T is not ast.StructType\n", currStruct)
				continue
			}

			if g.Doc == nil {
				fmt.Printf("SKIP struct %v doesnt have comments\n", currType.Name.Name)
				continue
			}

			needCodegen := false
			for _, comment := range g.Doc.List {
				needCodegen = needCodegen || strings.HasPrefix(comment.Text, "// cgen: binpack")
			}

			if !needCodegen {
				fmt.Printf("SKIP struct %v doesnt have cgen mark\n", currType.Name.Name)
				continue
			}

			fmt.Printf(" proccess struct %s\n", currType.Name.Name)
			fmt.Printf("\tstrat generating unpack method\n")
			fmt.Fprintln(out, "func (in *"+currType.Name.Name+") Unpack(data []byte) error {")
			fmt.Fprintln(out, `
	r := bytes.NewReader(data)
	orderByte := binary.BigEndian`)

			for _, field := range currStruct.Fields.List {
				if field.Tag != nil {
					tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
					if tag.Get("cgen") == "-" {
						fmt.Printf("property %v skiped mark\n", field.Names[0].Name)
						continue
					}
				}

				fieldName := field.Names[0].Name

				var buf bytes.Buffer
				printer.Fprint(&buf, fset, field.Type)
				fieldType := buf.String()

				fmt.Printf("\tgenerating code for field %s.%s\n", fieldType, fieldName)

				switch fieldType {
				case "uint":
					intTpl.Execute(out, tpl{fieldName})
				case "string":
					strTpl.Execute(out, tpl{fieldName})
				case "[]byte":
					sliceTpl.Execute(out, tpl{fieldName})
				case "time.Time":
					timeTpl.Execute(out, tpl{fieldName})
					fmt.Printf("its time")
				default:
					fmt.Printf("unsupported %s\n", fieldType)
					//continue
				}
			}

			fmt.Fprintln(out, "\treturn nil")
			fmt.Fprintln(out, "}")
			fmt.Fprintln(out)

		}

	}

}
