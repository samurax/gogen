package main

import "strings"
import "fmt"
import "encoding/gob"
import "os"

type generic_Product struct {
	Name string
	Code uint64
	Desc string
}

func (pg *ProductElectronic) Print() {
	PrintProductDesc(pg.Product)
	fmt.Println("----------------------")
	s := fmt.Sprintf("%#v", *pg)
	s = strings.Replace(s, "{", "\n ", -1)
	s = strings.Replace(s, "}", "", -1)
	s = strings.Replace(s, ":", "=", -1)
	s = strings.Replace(s, ",", "\n ", -1)
	fmt.Println(s)
}
func (pg *ProductElectronic) Save() error {
	filename := fmt.Sprintf("prod%d.gob", pg.Product.Code)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 644)
	defer f.Close()
	if err == nil {
		encoder := gob.NewEncoder(f)
		err = encoder.Encode(*pg)
	}
	return err
}
func (pg *ProductElectronic) Load() error {
	filename := fmt.Sprintf("prod%d.gob", pg.Product.Code)
	f, err := os.Open(filename)
	if err == nil {
		defer f.Close()
		decoder := gob.NewDecoder(f)
		err = decoder.Decode(pg)
	}
	return err
}
func (pg *ProductTransport) Print() {
	PrintProductDesc(pg.Product)
	fmt.Println("----------------------")
	s := fmt.Sprintf("%#v", *pg)
	s = strings.Replace(s, "{", "\n ", -1)
	s = strings.Replace(s, "}", "", -1)
	s = strings.Replace(s, ":", "=", -1)
	s = strings.Replace(s, ",", "\n ", -1)
	fmt.Println(s)
}
func (pg *ProductTransport) Save() error {
	filename := fmt.Sprintf("prod%d.gob", pg.Product.Code)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 644)
	defer f.Close()
	if err == nil {
		encoder := gob.NewEncoder(f)
		err = encoder.Encode(*pg)
	}
	return err
}
func (pg *ProductTransport) Load() error {
	filename := fmt.Sprintf("prod%d.gob", pg.Product.Code)
	f, err := os.Open(filename)
	if err == nil {
		defer f.Close()
		decoder := gob.NewDecoder(f)
		err = decoder.Decode(pg)
	}
	return err
}
