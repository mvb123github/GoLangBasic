package main

import (
	"GoLangBasic/say"

	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/julienschmidt/httprouter"
)

const applicationName = "App Const"

const (
	_ = iota // iota = 0 
	A = iota + 8 // iota = 1 
	B = A + iota // iota = ... 
	_ = iota
	D = iota
)

func MyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	w.WriteHeader(200)
	w.Write([]byte("Hello from go_program"))
}

func TypesGo() (int, bool, string) {
	return 80, false, "stringType"
}

func controlFlow() {
	a, b, c := 10, 20, 30

	if ((a > b) && (a > 7)) {
		fmt.Println("a")
	} else if (b > c) {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

	var text = "Admin"

	// type 1
	switch (text) {
		case "admin": 
			fmt.Println("Error not admin")
		case "Admin": 
			fmt.Println("Yeah - Admin!")	
			fallthrough	
		case "caseX":  	
			fmt.Println("Add case...")	
		default : 
			fmt.Println("Some login")	
	}	

	//	type 2
	switch {
		case text == "Admin" || text == "admin": 
			fmt.Println("ok")
		default : 
			fmt.Println("not ok")	
	}	

	// for
	for x := 1; x < 40; x++ {
		if (x%16 == 0) { continue }

		if ((x % 2) == 0) {
			fmt.Print(strconv.Itoa(x) + ",")	
		}	
		if (x%34 == 0) { break }
	}
	fmt.Println()	
}

func main() {
  var ai int = 4
	var bi bool = true

	const number = 70

	fmt.Printf("%T - %v\n", ai, ai)
	fmt.Printf("%T - %v\n", bi, bi)
	fmt.Printf("%T - %v\n", miniGlobalVar, miniGlobalVar) 
	fmt.Println(applicationName)
	result := number + ai
	fmt.Println(result)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(D)

	x1, _, _ := TypesGo()
	//x1, y1, z1 := TypesGo()
	fmt.Println(x1)
	//fmt.Println(y1)
	//fmt.Println(z1)

	fmt.Println(say.CallFromlSay())
	// fmt.Println(say.noVisible) - error because variable write with small first char that means not external!

	controlFlow()

	age := 2
	b := 3
	age = age + 16
	str := strconv.Itoa(age+b)
	fmt.Printf("hello, world\nResult = " + str + "\n")
	//fmt.Scanf("%d\n", &age)
	//fmt.Printf("New age = %d\n", age)
	//fmt.Scanf(" ");

	// router test: 
  log.Println("Read config")
	var cfg map[string]string
  cleanenv.ReadConfig("config.yml", &cfg)

	router := httprouter.New()
	router.GET("/", MyHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg["port"]), router))
}

var miniGlobalVar float32 = 100.97