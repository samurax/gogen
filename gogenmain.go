package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CheckGoErrors(file string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, file, nil, parser.DeclarationErrors)
	if err != nil {
		return err
	}
	return nil
}

func GAssert(cond bool, errmsg string) {
	if cond == false {
		log.Fatal("ERROR: ", errmsg)
	}
}
func GEAssert(err error) {
	if err != nil {
		log.Fatal("ERROR :", err)
	}
}

const message_createdir = "Please create directory %s with templates"
const message_definitionerr = "Data definition error at %s"
const file_configdir = "gogen.d"
const file_output = "inkogen.go"

const ch_comma = ","
const ch_pointer = "*"

type TypeConf struct {
	name    string
	datadef map[string]string
}

type TypesStatusInfo struct {
	msg         string
	structname  string
	skipcomment bool
	proctype    int
	typesdef    map[string]TypeConf
	curtype     *TypeConf
}

var Sinfo *TypesStatusInfo

func InitStatus() {
	Sinfo = new(TypesStatusInfo)
	Sinfo.typesdef = make(map[string]TypeConf, 0)
}

func tokenizeChar(ch string, token string) []string {
	ks := strings.Split(token, ch)
	nks := make([]string, 0)
	for _, t := range ks {
		s := strings.TrimSpace(t)
		if s == "" {
			continue
		}
		nks = append(nks, s)
	}
	return nks
}

func tokenizeSpec(okeys []string) []string {
	keys := make([]string, 0)
	for _, k := range okeys {
		ke := strings.TrimSpace(k)
		if ke == "" {
			continue
		}
		comma := tokenizeChar(ch_comma, ke)
		for _, kp := range comma {
			keys = append(keys, kp)
		}
	}
	return keys
}

func tokenizeCodeline(textline string) []string {
	ks := strings.Split(textline, " ")
	keys := make([]string, 0)
	for _, k := range ks {
		cl := strings.TrimSpace(k)
		if cl == "" {
			continue
		}
		keys = append(keys, cl)
	}
	return tokenizeSpec(keys)
}

func findGenericsDecl(url string , templin *TempFiles ) error {

	if FC.verbose > 0 {
		fmt.Println("--", url)
	}

	f, err := os.Open(url)
	GEAssert(err)
	defer f.Close()
	scan := bufio.NewScanner(f)
	comeat := CommentEater{}
	for scan.Scan() {
		codeline := scan.Text()
		keys := tokenizeCodeline(codeline)
		l := len(keys)
		if   l == 0 { continue }
		keyword := keys[0]
		codeline = comeat.Eat(codeline)
		switch {
		case codeline == "":
			continue
		case keyword == "const" && l > 1:
                ReadInlineCode(keys[1],codeline,scan,templin)
                        continue
		case codeline == "" || l == 0:
			continue
		case Sinfo.proctype == 0 && keyword == "type" && l >= 4 && keys[2] == "struct":
			structname := keys[1]
			datadef := make(map[string]string)
			nt := TypeConf{structname, datadef}
			Sinfo.typesdef[structname] = nt
			Sinfo.curtype = &nt
			Sinfo.proctype++
			continue
		case keyword == "type":
		case Sinfo.proctype > 0 && keyword == "}":
			Sinfo.proctype--
			fallthrough
		case Sinfo.proctype == 0:
			continue
		case Sinfo.curtype == nil:
			continue
		}

		nkeys := strings.Split(strings.TrimSpace(codeline)," ")
		l = len(nkeys)
		typename := strings.TrimSpace(nkeys[l-1])
		prev := ""
		for _, k := range nkeys[0:l-1] {
			k = strings.TrimSpace(k)
			if k == "*" {
				Sinfo.curtype.datadef[prev] = "*" + typename
				continue
			}

			switch  {
			case k=="*":
				Sinfo.curtype.datadef[prev] = "*" + typename
				continue
			case k=="[":
				Sinfo.curtype.datadef[prev] = "[]" + typename
				continue
			case k=="]", k=="" , typename == "":
				continue
			}
			Sinfo.curtype.datadef[k] = typename
			prev = k
		}
	}
	return nil
}

func Init() int {
	templin := TempFiles{}
	templin.Init()
	
	fls, derr := ioutil.ReadDir(".")
	if derr != nil {
		log.Fatal(derr)
	}
	outfile, err := os.OpenFile(FC.outputfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 640)
	GEAssert(err)
	writer := bufio.NewWriter(outfile)
	cw := OpenBuffers(writer)

	for _, f := range fls {
		if f.Name() == FC.outputfile {
			continue
		}
		p := strings.Split(f.Name(), ".")
		ext := p[len(p)-1]
		if ext == "go" {
			err := CheckGoErrors(f.Name())
			GEAssert(err)
			InitStatus()
			findGenericsDecl(f.Name(),&templin)
			CheckGenerics(f.Name(), Sinfo.typesdef, cw,&templin )
		}
	}

	exitcode := CloseBuffers(cw)
	templin.Remove()

	if exitcode == 0 {
		cmd := exec.Command("gofmt", "-w", FC.outputfile)
		err := cmd.Start()
		GEAssert(err)
	}

	return exitcode
}

func main() {
	templatedir := flag.String("d", "gogen.d", "Template directory")
	templateext := flag.String("x", ".ge", "Template extension")
	outputfile := flag.String("o", "inkogen.go", "Output file")
	verbose := flag.Int("v", 0, "Verbose")
	flag.Parse()
	FC.templatedir = *templatedir
	FC.templateext = *templateext
	FC.outputfile = *outputfile
	FC.verbose = *verbose
	args := flag.Args()
	startdir := "."
	if len(args) > 0 {
		startdir = args[0]
		err := os.Chdir(startdir)
		GEAssert(err)
	}
	exitcode := Init()
	os.Exit(exitcode)
}
