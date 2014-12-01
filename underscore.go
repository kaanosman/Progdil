		
package main

import (
	"fmt"
	"strings"
	"flag"       	
	)

func underscore() string { 
 
    flag.Parse()
    word := flag.Arg(0) 
 
    list := strings.Split(word, "")
 
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

    return(strings.Join(list, ""))

}

func main() {
    fmt.Println(underscore())
}
