package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"math"
)
func Greeting(w http.ResponseWriter, req *http.Request){
	CalculateLoop()
	io.WriteString(w, "<b>Code.education Rocks!!</b>\n")
}
func CalculateLoop()  {
	x := 0.00001;
	for i := 1;  i<=100000; i++ {
		x += math.Sqrt(x)
		fmt.Printf("Squrt is %g loop itertation %d \n",x, i)
	}

}
func main() {
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", Greeting)
	log.Fatal(s.ListenAndServe())
}

