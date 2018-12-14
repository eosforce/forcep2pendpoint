package main

import (
	"fmt"

	"flag"
	"net/http"
	//"strings"
	"bytes"
	"github.com/eosforce/relay/p2p"
	"encoding/json"
)

var url = flag.String("url","http://127.0.0.1:12560/print","the url to get peer ")
var listeningAddress = flag.String("listening-address", "0.0.0.0:19808", "address on with the relay will listen")
var showLog = flag.Bool("v", false, "show detail log")

type regAddress struct {
	RegAddress string
}

func GetPeerAddress() string{
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", *url, nil)
	response, _ := client.Do(reqest)

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	defer response.Body.Close()
	//s := buf.String() 
	//bodybyte := readAll(response.Body, bytes.MinRead)
	var RegAddresses []regAddress
    // movies2 := make([]Movie, 10)
    if err3 := json.Unmarshal(buf.Bytes(), &RegAddresses); err3 !=nil{
        fmt.Println("transfer json error")
    }
	if len(RegAddresses) > 0 {
		return RegAddresses[0].RegAddress
	}
	return "127.0.0.1:18088"
}

func main() {
	flag.Parse()

	if *showLog {
		p2p.EnableP2PLogging()
	}
	defer p2p.SyncLogger()
	peer_address := GetPeerAddress() 
	relay := p2p.NewRelay(*listeningAddress,peer_address)
	
	relay.RegisterHandler(p2p.StringLoggerHandler)
	fmt.Println(relay.Start())
}
