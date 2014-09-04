package main

type Product struct {
      Name     string 
      Code     string
}

type    Stock struct {
Name     string
file     string
Prds   []Product
}

type    StockGroup struct {
	file      string
        Name      string
        Stocks []  Stock
}

type        GenericPersist   struct {
Stock       generic_persist
StockGroup  generic_persist
}

type        GenericNameDontCare   struct {
Stock       generic_dumper
StockGroup  generic_dumper
}


func main () {
	stock1 := Stock {Name:"Stock1",file:"stock1.gob",
	                Prds:[]Product{ {"Laptop. Tochino 30400","0003404040"},
			                 {"TV felipo R455","00049590404"},
					 {"Multimedia Player","94995999"}, 
				       },
			 }
	stock1.Save()
	stock1.Dump("saved sock1")

        stock2 :=  Stock { file:"stock1.gob"  }
	stock2.Load()
	stock2.Dump("loaded stock2" )

	stock3 := Stock {Name:"Stock3",file:"stock1.gob",
	                Prds:[]Product{ {"Led Lamp. Jion","00303404040"},
			                 {"Raspberry","000445590404"},
				       },
			 }
	 stocks := new(StockGroup)
	 stocks.Stocks = []Stock{stock1,stock2,stock3}
	 stocks.file="teststocks.gob"
	 stocks.Save()
	 stocks.Dump("saved stocks")

         newstocks := new (StockGroup)
	 newstocks.file = stocks.file
         newstocks.Load()
         newstocks.Dump("loaded newstock")
}
