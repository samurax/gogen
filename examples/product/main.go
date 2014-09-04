package main
import  "fmt"

const generic_Product_template=`
package main;
import "fmt";
import "encoding/gob"
import "os"
import "strings"

<once>
type       generic_Product  struct { 
Name       string
Code       uint64
Desc       string
}
</once>

func (pg *<data>) Print() {
     PrintProductDesc(pg.Product)
     fmt.Println("----------------------")
     s := fmt.Sprintf("%#v",*pg)
     s = strings.Replace(s,"{","\n    ",-1)
     s = strings.Replace(s,"}","",-1)
     s = strings.Replace(s,":","=",-1)
     s = strings.Replace(s,",","\n         ",-1)
     fmt.Println(s)
}

func  (pg *<data>) Save() error {
        filename:= fmt.Sprintf("prod%d.gob", pg.Product.Code )
        f , err := os.OpenFile (filename,os.O_WRONLY | os.O_TRUNC | os.O_CREATE,644)
        defer f.Close()
	if err == nil  {  encoder := gob.NewEncoder(f)
		          err = encoder.Encode(*pg) 
		       }
       return err 
}

func  (pg *<data>) Load() error  {
        filename:= fmt.Sprintf("prod%d.gob", pg.Product.Code )
	f , err := os.Open(filename) 
	if err == nil {  defer f.Close()
	                 decoder := gob.NewDecoder(f) 
			 err = decoder.Decode(pg)
	               }
	return err
}
`


func    PrintProductDesc (p generic_Product ) {
	fmt.Println("------------------------")
	fmt.Println("PRODUCT:",p.Name)
	fmt.Println("CODE:",p.Code)
	fmt.Println(p.Desc)
}


type    ProductElectronic struct {
        Product           generic_Product
        Price             float32
}

type    ProductTransport  struct {
        Product            generic_Product
        Price1             float32
        Price2             float32 
        Price3             float32 
	Tax1               float32
	Tax2               float32 
}

func main () {
	pe := ProductElectronic { Product: generic_Product {  "TV SONY",123,"flatron OLED TV"},
                                  Price : 2300 , 
                                 }
	err:= pe.Save()
	if err == nil { fmt.Println("SAVED OK")  
        } else        { fmt.Println("ERROR SAVING",err) }

        pe.Print()
	//---------------------------------
	pn := new (ProductElectronic)
	pn.Product.Code = 123
	err = pn.Load()
	if err != nil { 
	   fmt.Println("ERROR LLOADING")
	}
	pn.Print()
        //---------------------------------------------
	pt := ProductTransport {  Product:generic_Product { "BMW d118",1000,"Sport Car"},
	                          Price1 :26000 ,
				  Price2 :23000 ,
				  Price3 :21500 ,
				  Tax1   :16 ,
				  Tax2   :10 , 
			        }
        pt.Save()
	px :=  new (ProductTransport)
	px.Product.Code = 1000
	px.Load()
	px.Print()

}

//--------------------Compile using:   gogen && go build

