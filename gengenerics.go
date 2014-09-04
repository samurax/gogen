package main

import (
	"bufio"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

const templatedir = "gogen.d"
const templateext = ".ge"
const fileseparator = "/"
const outputfile = "inkogen.go"

type FilesConfig struct {
	templatedir string
	templateext string
	outputfile  string
	verbose     int
}

var FC = FilesConfig{templatedir, templateext, outputfile, 0}

type tgeneric struct {
	gentype string
	datadef map[string]string
	stname  string
}

func isGenericName(name string) bool {
	generictype := "generic"
	l := len(name)
	if l > 8 && strings.ToLower(name[0:8]) == generictype + "_" {
		return true
	}
	return false
}

func GWarn(err error) bool {
	if err != nil {
		log.Println("WARN:", err)
		return true
	}
	return false
}

type CodeWriter struct {
	W           *bufio.Writer
	errors      int
	lines       int
	packages    map[string]bool
	oncecode    map[string]bool
	code        string
	packagename string
}

func OpenBuffers(w *bufio.Writer) *CodeWriter {
	cw := new(CodeWriter)
	cw.W = w
	cw.packages = make(map[string]bool)
	cw.oncecode = make(map[string]bool)
	cw.packagename = "main"
	return cw
}

func CloseBuffers(cw *CodeWriter) int {
	msg := " saved in: " + FC.outputfile
	switch {
	case cw.errors > 0:
		msg = fmt.Sprintf("ERRORS:%d", cw.errors)
		fmt.Println(msg)
		return -1
	case FC.verbose > 0:
		fmt.Println("     lines:", cw.lines, msg)
	}
	endl := "\n"
	cw.W.WriteString("package " + cw.packagename + endl)
	for p, _ := range cw.packages {
		cw.W.WriteString("import " + p + endl)
	}
	cw.W.WriteString(cw.code)
	cw.W.Flush()
	return 0
}

func outputTokenize(textline string) []string {
	ks := strings.Split(textline, " ")
	keys := make([]string, 0)
	for _, k := range ks {
		cl := strings.TrimSpace(k)
		if cl == "" {
			continue
		}
		keys = append(keys, cl)
	}
	return keys
}

func GenerateCode(structname string, template string, declin string, typesdef map[string]TypeConf, cw *CodeWriter , templin * TempFiles ) {

	filename :=  template + FC.templateext
	tmplurl := FC.templatedir + fileseparator + template + FC.templateext

        _ , serr :=  os.Stat ( filename )
        if  serr == nil {  tmplurl = filename }
        inlinefile := templin.tfiles[template]
        if inlinefile != nil {tmplurl = inlinefile.Name()}

	f, err := os.Open(tmplurl)
	if GWarn(err) {
		log.Println("FILE ERROR:", tmplurl)
		cw.errors++
		return }

	fset := token.NewFileSet()
	pr, perr := parser.ParseFile(fset, tmplurl, nil, parser.ImportsOnly)
	if GWarn(perr) {
		cw.errors++
		log.Println("FILE:", tmplurl)
		return
	}
	scan := bufio.NewScanner(f)
	for _, s := range pr.Imports {
		cw.packages[s.Path.Value] = true
	}

	datapatt := "<data>"
	consimport := true
	once := false
	comeater := CommentEater{}
	for scan.Scan() {
		codeline := scan.Text()
		codeline = strings.Replace(codeline, datapatt, structname, -1)
		codeline = comeater.Eat(codeline)
		keys := outputTokenize(codeline)
		l := len(keys)
		if l == 0 {
			continue
		}
		keyword := keys[0]
		switch {
		case keyword == "</once>":
			once = false
			continue
		case keyword == "<once>":
			odone := cw.oncecode[template]
			if odone {
				once = true
			}
			cw.oncecode[template] = true
			continue
		case once:
			continue
		case keyword == "import":
			if strings.Index(codeline, ")") >= 0 {
				consimport = false
			}
			continue
		case consimport && (keyword == "func" || keyword == "type"):
			consimport = false
		case consimport:
			if strings.Index(codeline, ")") >= 0 {
				consimport = false
			}
			continue
		}

		cw.lines++
		outputline := strings.Join(keys, " ")
		cw.code += outputline + "\n"
	}
}

func CheckGenerics(gofile string, typesdef map[string]TypeConf, cw *CodeWriter,templin *TempFiles) {
	generictype := false
	for declin, data := range typesdef {
		for vname, vtype := range data.datadef {
			if isGenericName(vtype) {
				generictype = true
				templ := strings.Split(vtype, "_")
				template := templ[1]
				structname := vname
				if structname == template { structname = declin  }
				GenerateCode(structname , template, declin, typesdef, cw , templin )
			}
		}
	}
	if !generictype {
		return
	}
	for name, data := range typesdef {
		if FC.verbose > 0 {
			fmt.Printf("  %s TYPE:%20s\n", gofile, name)
			fmt.Println("---------------------------------------------")
		}

		for vname, vtype := range data.datadef {
			msg := ""
			if isGenericName(vtype) {
				msg = "<----<<<"
			}
			if FC.verbose > 0 {
				fmt.Printf("   %16s   %s %s\n", vname, vtype, msg)
			}

		}
	}
}
