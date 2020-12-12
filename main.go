package main

import (
	"fmt"
	"os"
	"os/signal"
)

func init() {
	// init db etc.
}

func main() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping server...")
	//more todo
}


//service <- sygnal od klienta i reagowanie <- controler
// model <- struct opisujacy dane, ale chcem miec też repo/store 
//
