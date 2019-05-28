package main

import "fmt"
import "net/http"
import "time"

func handlerIndex(w http.ResponseWriter, r *http.Request){
	var massage = ("welcome")
	w.Write([]byte(massage))
}
func handlerHallo (w http.ResponseWriter, r *http.Request){
	var massage = ("hallo dunia")
	w.Write([]byte(massage))
}

func main(){
	http.HandleFunc("/",handlerIndex)
	http.HandleFunc("/hallo",handlerHallo)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n",address)

	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10


	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}
}