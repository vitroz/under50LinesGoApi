package main

import(

	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

)

var (
	addr = flag.String("addr", ":8083", "http service address")
	data map[string]string
)

func main(){
	flag.Parse()
	data = map[string]string{}
	r := httprouter.New() //Criando as rotas
	r.GET("/entry/:key", show)
	r.GET("/list", show)
	r.PUT("/entry/:key/:value", update)
	r.DELETE("/entry/:key", del)
	err := http.ListenAndServe(*addr, r) // Starting the server

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	k := p.ByName("key")
	if k == ""{
		fmt.Fprintf(w, "Read list: %v", data)
		return
	}
	fmt.Fprintf(w, "Read entry: data[%s] = %s", k, data[k])
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	k := p.ByName("key")
	v := p.ByName("value")

	data[k] = v

	fmt.Fprintf(w, "Updated: data[%s] = %s", k , data[k])
}

func del(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	k := p.ByName("key")
	value := data[k]

	if value == "" {
		fmt.Fprintf(w, "Entry does not exist")
		return
	}

	delete(data, k) 

	fmt.Fprintf(w, "Deleted Data! ")
}

