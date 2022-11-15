package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	Fun int    = 0
    Sred int   = 1
    Val int    = 2
)
const (
	Output int = 0
    Input  int = 1
    ValIn  int = 2
)

func main() {
	fmt.Println(os.Args[1:][0])
	file, err := os.Open(os.Args[1:][0])
	if err != nil{
		log.Fatal("error")
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)
	var code, inn string
    var i, sled, fun int

	for{
		n, err := file.Read(data)
		if err == io.EOF{
			break
		}

		code = " "+string(data[:n])
	}

	i = 0
	sled = Fun

	for true {
		if len(code)-1==i {
			break
		}
		i = i+1

		switch sled {
			case Fun:
				sled=Sred
                switch string(code[i]) {
					case "a":
						fun = Output
					case "b":
						fun = Input
					case "c":
						fun = ValIn
					default:
						fun = Output
				}
			case Sred:
				sled=Val
			case Val:
                sled=Fun
				switch fun {
					case Output:
						fmt.Println(string(code[i]))
					case Input:
						fmt.Scanf("%s\n", &inn)
					case ValIn:
						fmt.Println(inn)
				}
		}
	}

}
