package main

import "strings"
import "log"
import "io/ioutil"
import "path/filepath"
import "fmt"

type generic_dumper bool

func (d *Htmldata) Dump(title string) {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("__Htmldata_________" + title)
	str = strings.Replace(str, ":", "=", -1)
	str = strings.Replace(str, "{", "\n....... ", -1)
	str = strings.Replace(str, ",", "\n.......", -1)
	str = strings.Replace(str, "}", "\n", -1)
	fmt.Println(str)
}
func (d *Godata) Dump(title string) {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("__Godata_________" + title)
	str = strings.Replace(str, ":", "=", -1)
	str = strings.Replace(str, "{", "\n....... ", -1)
	str = strings.Replace(str, ",", "\n.......", -1)
	str = strings.Replace(str, "}", "\n", -1)
	fmt.Println(str)
}

type generic_dirscan struct {
	Ext   string
	Base  string
	Dir   string
	Url   string
	IsDir bool
}

func GenAssert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func (gen *Search_html) Scan(url string, prefix string) {
	fls, err := ioutil.ReadDir(url)
	GenAssert(err)
	parsedir := func(url string) {
		gen.Search_html.Ext = filepath.Ext(url)
		gen.Search_html.Dir = filepath.Dir(url)
		gen.Search_html.Base = filepath.Base(url)
		gen.Search_html.Url = url
	}
	parsedir(url)
	gen.Search_html.IsDir = true
	gen.DirInfo(prefix)
	for _, f := range fls {
		nurl := url + "/" + f.Name()
		parsedir(nurl)
		isdir := f.IsDir()
		gen.Search_html.IsDir = isdir
		gen.DirInfo(prefix)
		if isdir {
			gen.Scan(nurl, prefix+"-----")
		}
	}
}
func (gen *Search_go) Scan(url string, prefix string) {
	fls, err := ioutil.ReadDir(url)
	GenAssert(err)
	parsedir := func(url string) {
		gen.Search_go.Ext = filepath.Ext(url)
		gen.Search_go.Dir = filepath.Dir(url)
		gen.Search_go.Base = filepath.Base(url)
		gen.Search_go.Url = url
	}
	parsedir(url)
	gen.Search_go.IsDir = true
	gen.DirInfo(prefix)
	for _, f := range fls {
		nurl := url + "/" + f.Name()
		parsedir(nurl)
		isdir := f.IsDir()
		gen.Search_go.IsDir = isdir
		gen.DirInfo(prefix)
		if isdir {
			gen.Scan(nurl, prefix+"-----")
		}
	}
}
