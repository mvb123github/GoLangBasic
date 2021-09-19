package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/julienschmidt/httprouter"
)

func MyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	w.WriteHeader(200)
	w.Write([]byte("Hello from go_program"))
}

func main() {
	age := 2
	b := 3
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