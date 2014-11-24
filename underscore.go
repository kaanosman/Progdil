		
package main

import (
	"fmt"
	"strings"
	"flag"       	
	)

func underscore() { 
 
    flag.Parse()
    word := flag.Arg(0) 
 
    list := strings.Split(arg, "")
 
    first := 0
    last := len(list)-1
    
	for list[first] == "_" {  //while döngüsü
            first++
        }
	for list[last] == "_" {   //while döngüsü
            last--
        }
        for i := first ; i < last ; i++ {
            if list[i] == "_" {
               list[i] = " "
            }
        }

    fmt.Println(strings.Join(list, ""))

}

func main() {
    underscore()
}
