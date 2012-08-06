package main

import(
	"bufio"
	"nimble-cube/core"
	"io"
	"flag"
	"os"
	"fmt"
)


func main(){
	flag.Parse()
	if flag.NArg() == 0{
		read(os.Stdin)
	}
	for _,arg:=range flag.Args(){
		in, err := os.Open(arg)
		core.Fatal(err)
		read(bufio.NewReader(in))
		in.Close()
	}
}

func read(in io.Reader){
	v := core.Read(in)
	fmt.Println(v)
}
