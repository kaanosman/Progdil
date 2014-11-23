package main

import (
"fmt"
"math/rand"
"time"
"flag"
"strconv"
)

func seedAndReturnRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func randomFrom(source []string) string {
	return source[seedAndReturnRandom(len(source))]
}

func main() {

 flag.Parse()
 dil := flag.Arg(0)
 adet,err := strconv.Atoi(flag.Arg(1));
  
 isim:= []string {"kedi","insan","deniz","masa","dolap","bilgisayar",
			"kapi","pencere","hafta","gece","gunduz","marti","kazak",
			"sarisin","sahil","merdiven","ahmet","yol","araba","agac",
			"yaprak","dal","adam","konuk","ahmet","koca","kadin"}


 sifat:= []string {"deli", "dolu" , "yaramaz","iyi","guzel","hizli","yavas","yakisikli","mert",
			"cimri","kel","masmavi","ucan","kac","yarim","bu","su"}
 name:= []string {"joe", "mike" , "caitlyn","graves","soraka"}

 adjactive:= []string {"happy", "sad" , "beautiful"}



if (dil=="tr") {
   for i := 0 ; i < adet ; i++ {
fmt.Println(randomFrom(sifat),randomFrom(isim))
  }  
}

if (dil=="en") {
   for i := 0 ; i < adet ; i++ { 

fmt.Println(randomFrom(adjactive),randomFrom(name))
fmt.Println(err)
  }
}










}
