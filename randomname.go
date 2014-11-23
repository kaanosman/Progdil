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
  
 isim:= []string {"Kedi",
                  "Insan",
                  "Deniz",
                  "Masa",
                  "Dolap",
                  "Bilgisayar",
		  "Kapi",
                  "Pencere",
                  "Hafta",
                  "Gece",
                  "Gunduz",
                  "Marti",
                  "Kazak",
	          "Perde",
                  "Sahil",
                  "Merdiven",
                  "Ahmet",
                  "Yol",
                  "Araba",
                  "Agac"}

 sifat:= []string {"Deli",
                   "Dolu" ,
                   "Yaramaz",
                   "Zayif",
                   "Guzel",
                   "Hizli",
                   "Yavas",
                   "Yakisikli",
                   "Mert",
		   "Cimri",
                   "Kel",
                   "Masmavi",
                   "Ucan",
                   "Kirik",
                   "Kesik",
                   "Sevimli",
                   "Kötü",
                   "Iyi",
                   "Uzun",
                   "Kisa"}
                   
 name:= []string {"Door",
                  "Mike" ,
                  "Table",
                  "Caitlyn",
                  "Chair",
                  "Stairs",
                  "Joe",
                  "Graves",
                  "Soraka",
                  "Window",
                  "House",
                  "Towel",
                  "Addison",
                  "Isabella",
                  "Andrew",
                  "Louis",
                  "Paul",
                  "Victor",
                  "Lamp",
                  "Mirror"}

 adjactive:= []string {"Angry",
                       "Amphibious",
                       "Bad",
                       "Beautiful", 
                       "Big",
                       "Careless",
                       "Difficult",
		       "Dry",
                       "Elastic",
                       "Empty",
                       "First",
                       "Great",
                       "Happy",
                       "Heavy",
                       "Natural",
                       "Nervous",
                       "Productive",
                       "Sharp",
                       "Small",
                       "Visible"}  
 


if (dil=="tr") {
   for i := 0 ; i < adet ; i++ {
       fmt.Println(randomFrom(sifat),randomFrom(isim))
  }  
}

if (dil=="en") {
   for i := 0 ; i < adet ; i++ { 
       fmt.Println(randomFrom(adjactive),randomFrom(name))
  }
  fmt.Println(err)
}


}
