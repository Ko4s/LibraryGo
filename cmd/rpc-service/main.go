package main

import (
	"fmt"
	configreader "library/pkg/config_reader"
	mongostorage "library/pkg/storage/mongo_storage"
	"os"
	"os/signal"
)

const (
	configName = "config"
	configType = "yaml"
	configPath = "."
)

func init() {
	// init db etc.
}

func main() {

	cr, err := configreader.NewConfigReader(configName, configType, configPath)

	if err != nil {
		panic(err)
	}

	_, err = mongostorage.NewStorage(cr)

	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping server...")
	//more todo
}

//service <- sygnal od klienta i reagowanie <- controler
// model <- struct opisujacy dane, ale chcem miec też repo/store
//
