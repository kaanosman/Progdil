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

func randomFrom(list []string) string {
	return list[seedAndReturnRandom(len(list))]
}

func check(list []string, s string) bool {
   for i := 0 ; i < len(list) ; i++ {  
     if list[i] == s {
        return true
     } else {
        return false
     } 
  }
}

func main() {
 var checklist []string
 
 flag.Parse()
 lang := flag.Arg(0)
 count,err := strconv.Atoi(flag.Arg(1));
  
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
 


if (lang == "tr") {
   for (count > 0) {
     name := (randomFrom(sifat),randomFrom(isim))
     if (check(checklist, name) == false) {
        checklist = append(checklist, name)
        fmt.Println(name)
        count--
     }
  }  
}

if (lang == "en") {
   for (count > 0) { 
     name := (randomFrom(adjactive)+" "+randomFrom(name))  
     if (check(checklist,name) == false) {
         checklist = append(checklist, name)
         fmt.Println(name)
         count--
}
  }
}
 fmt.Println(err)
}
