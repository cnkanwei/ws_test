package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

var wsConnClose bool

var pingStr = "{\"type\":\"ping\"}"

func main() {

	u := ""
	p := ""
	o := ""
	flag.StringVar(&u,"uri","/echo","uri")
	flag.StringVar(&p,"protocol","ws","协议")
	flag.StringVar(&o,"origin","http://127.0.0.1:8080","域名")
	flag.Parse()

	reHttp,_:=regexp.Compile("^(http|https)")
	url_ := strings.Trim(reHttp.ReplaceAllString(o,p),"/") + u

	ws,err := websocket.Dial(url_,p,o)
	if err != nil {
		panic(err)
	}

	wsConnClose = false

	defer func() {
		err := ws.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	go ping(ws)

	go doReceive(ws)

	go solveStdin(ws)

	h := make(chan bool,1)
	<- h
}

func solveStdin(ws *websocket.Conn){
	for{
		if wsConnClose == true{
			return
		}
		f:=bufio.NewReader(os.Stdin)
		var Input string
		Input , err := f.ReadString('\n')
		Input = strings.Trim(Input,"\n")
		if err != nil{
			log.Println(err)
		}
		send(ws,[]byte(Input))
	}
}

func ping(ws *websocket.Conn){
	for {
		if wsConnClose {
			return
		}
		send(ws,[]byte(pingStr))
		time.Sleep(time.Second*2)
	}
}

func send(ws *websocket.Conn,msg []byte)  {
	_,err := ws.Write(msg)
	if err != nil {
		log.Println(err)
	}
	if string(msg)!=pingStr{
		log.Println("send:"+string(msg))
	}
}

func doReceive(ws *websocket.Conn)  {
	message := make([]byte,520)
	for{
		n,err := ws.Read(message)
		if err != nil && err != io.EOF{
			log.Println(err)
			wsConnClose = true
			return
		}
		if n ==0 {
			log.Println("Server closed")
			wsConnClose = true
			return
		}
		if string(message[:n])!=pingStr{
			log.Println("Receive:"+string(message[:n]))
		}

	}
}