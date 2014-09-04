package main

import "strings"
import "encoding/gob"
import "os"
import "fmt"

type generic_persist bool

func (c *Stock) Load() error {
	f, err := os.Open(c.file)
	defer f.Close()
	if err == nil {
		decoder := gob.NewDecoder(f)
		err = decoder.Decode(c)
	}
	return err
}
func (c *Stock) Save() error {
	fo, err := os.OpenFile(c.file, os.O_WRONLY|os.O_CREATE, 640)
	if err != nil {
		return err
	}
	defer fo.Close()
	encoder := gob.NewEncoder(fo)
	err = encoder.Encode(*c)
	return err
}
func (c *StockGroup) Load() error {
	f, err := os.Open(c.file)
	defer f.Close()
	if err == nil {
		decoder := gob.NewDecoder(f)
		err = decoder.Decode(c)
	}
	return err
}
func (c *StockGroup) Save() error {
	fo, err := os.OpenFile(c.file, os.O_WRONLY|os.O_CREATE, 640)
	if err != nil {
		return err
	}
	defer fo.Close()
	encoder := gob.NewEncoder(fo)
	err = encoder.Encode(*c)
	return err
}

type generic_dumper bool

func (d *Stock) Dump(title string) {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("__Stock_________" + title)
	str = strings.Replace(str, ":", "=", -1)
	str = strings.Replace(str, "{", "\n....... ", -1)
	str = strings.Replace(str, ",", "\n.......", -1)
	str = strings.Replace(str, "}", "\n", -1)
	fmt.Println(str)
}
func (d *StockGroup) Dump(title string) {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("__StockGroup_________" + title)
	str = strings.Replace(str, ":", "=", -1)
	str = strings.Replace(str, "{", "\n....... ", -1)
	str = strings.Replace(str, ",", "\n.......", -1)
	str = strings.Replace(str, "}", "\n", -1)
	fmt.Println(str)
}
