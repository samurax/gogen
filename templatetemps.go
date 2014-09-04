package main

import  "io/ioutil"
import  "os"


type  TempFiles    struct {
          tfiles  map[string]*os.File 
}

func (tt *TempFiles ) Remove ()  {
          for _ , v:= range tt.tfiles {
            os.Remove(v.Name())
          }
}

func (tt *TempFiles) Init () {  
         tt.tfiles= make(map[string]*os.File)
}

func (tt * TempFiles) OpenFile (id string ) (* os.File , error ) {
          file,err := ioutil.TempFile("/tmp/",id)
          if  err == nil  { tt.tfiles[id] = file } 
       return file,err
}



