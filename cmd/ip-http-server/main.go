package main
import (
    "fmt"
    "net/http"
   // "strings"
	//"log"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type regAddress struct {
	RegAddress string
}

var regAddresses []regAddress

func find(addr string) bool {
	fmt.Println("in find")
	for _, value := range regAddresses {
		if addr == value.RegAddress {
			fmt.Println("find")
			return true
		}
	}
	fmt.Println("no find")
	return false 
}



func register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("in register")
	bFind := find(p.ByName("name"))
	if bFind == false {
		fmt.Println("no find")
		var getAddress regAddress
		getAddress.RegAddress = p.ByName("name")
		regAddresses = append(regAddresses,getAddress)
		fmt.Println(regAddresses)
	}
	
	print(w,r,p)
}

func print(w http.ResponseWriter, r *http.Request,p httprouter.Params) {
	// for _, value := range regAddresses {
	// 	fmt.Fprintf(w,"%s\n",value.RegAddress);
	// } 
	fmt.Println(regAddresses)
	data, err := json.Marshal(regAddresses)
	if err != nil {
        fmt.Println("Json marshaling failed：", err)
	}
	fmt.Fprintf(w,"%s\n",string(data));
}

func delete(w http.ResponseWriter, r *http.Request,p httprouter.Params) {
	for index, value := range regAddresses {
		if p.ByName("name") == value.RegAddress {
			regAddresses = append(regAddresses[:index], regAddresses[index+1:]...)
			break;
		}
	} 
	print(w,r,p)
}

func main() {
	mux := httprouter.New()
	mux.GET("/add/:name",register)
	mux.GET("/print",print)
	mux.GET("/delete/:name",delete)
	server := http.Server{
		Addr:"127.0.0.1:12560",
		Handler:mux,
	}
    server.ListenAndServe()
}
