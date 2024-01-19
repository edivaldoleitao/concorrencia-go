package main

import (
	"fmt"
	"math/rand"
	"time"
  "strconv"
)


func enviamsg(ch chan string) {
  
  for {
    r := rand.Intn(10)
    valor :=  strconv.Itoa(r)
    ch <- valor
    fmt.Println("enviou mensagem: ",valor)
  }
 
}

func recebemsg(ch chan string) {
  
  for {
    str := <- ch
    fmt.Println("Recebeu mensagem: ",str)
    
  }
}


func main() {
  
  msgs := make(chan string)

  go enviamsg(msgs)
  go recebemsg(msgs)
  r := rand.Intn(10)
  time.Sleep(time.Duration(r) * time.Millisecond)
}
