package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
	"unicode"
)

type option struct {
	file     string
	exported bool
	withtest bool
}
var opt = &option{}

func (opt *option) init() {
	flag.StringVar(&opt.file, "file", "", "specify file name")
	flag.StringVar(&opt.file, "f", "", "same -file")

	flag.BoolVar(&opt.exported, "exporeted", false, "only exported function")
	flag.BoolVar(&opt.exported, "e", false, "same -exported")

	flag.BoolVar(&opt.withtest, "test", false, "output with test")
	flag.BoolVar(&opt.withtest, "t", false, "same -test")

	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("invalid args: ", flag.Args())
	}
}

func init() {
	log.SetPrefix("expfunc: ")
	log.SetFlags(log.Lshortfile)

	opt.init()
}

func tmplTest(fn string) string {
	return `func Test` + strings.Title(fn) + `(t *testing.T) {
	t.Fatal()
}`
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, opt.file, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	var funcs []string
	for _, d := range f.Decls {
		switch td := d.(type) {
		case *ast.FuncDecl:
			var recv string
			if td.Recv != nil {
				recv = td.Recv.List[0].Names[0].Name + "_"
			}
			if opt.exported {
				if unicode.IsUpper(rune(td.Name.Name[0])) {
					funcs = append(funcs, recv+td.Name.Name)
				}
				continue
			}
			funcs = append(funcs, recv+td.Name.Name)
		}
	}

	if opt.withtest {
		for _, fn := range funcs {
			fmt.Println(tmplTest(fn))
		}
	} else {
		for _, fn := range funcs {
			fmt.Println(fn)
		}
	}
}
