		
package main

import (
	"fmt"
	"strings"
	"flag"       	
	)

func main() { 
 
    flag.Parse()
    arg := flag.Arg(0) 
 
    s := strings.Split(arg, "")
 
    first := 0
    last := len(s)-1
    
	for s[first] == "_" {
            first++
        }
	for s[last] == "_" {
            last--
        }
        for i := first ; i < last ; i++ {
            if s[i] == "_" {
               s[i] = " "
            }
        }

    fmt.Println(strings.Join(s, ""))

}
