package main

import (
	"encoding/json"
	"errors"
	_ "expvar"
	"flag"
	"fmt"
	"github.com/sda0/eth_scanner/api"
	"github.com/sda0/eth_scanner/application"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var cfgScanner application.Config
var cfgApi api.Config

func main() {
	cfgFilePath := flag.String("cfg", "", "Path to configuration file")
	onlyValidateCfg := flag.Bool("validate", false, "validate configuration file and stop")

	flag.Parse()

	err := openConfig(*cfgFilePath)
	if err != nil {
		fmt.Println("Unable to open config: " + err.Error())
		os.Exit(2)
	}

	if *onlyValidateCfg {
		fmt.Println("Configuration file valid")
		os.Exit(0)
	}

	app, err := application.NewApplication(cfgScanner)
	if err != nil {
		fmt.Println("Fatal error: " + err.Error())
		os.Exit(1)
	}
	defer app.Stop()

	apiServer, err := api.NewServer(cfgApi)
	if err != nil {
		fmt.Println("Fatal error: " + err.Error())
		os.Exit(1)
	}
	defer apiServer.Stop()

	go func() {
		err = app.Start()
		if err != nil {
			fmt.Println("Fatal error: " + err.Error())
			os.Exit(1)
		}
	}()

	//при принудительном завершении defer не работает
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
		<-sigchan
		app.Stop()
		apiServer.Stop()
		os.Exit(0)
	}()

	err = apiServer.Start()
	if err != nil {
		fmt.Println("Fatal error: " + err.Error())
		os.Exit(1)
	}
}

func openConfig(filename string) error {
	if filename == "" {
		return errors.New("please, specify configuration file path: -cfg=PATH")
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &cfgScanner); err != nil {
		return err
	}
	if err := cfgScanner.Validate(); err != nil {
		return err
	}

	if err := json.Unmarshal(file, &cfgApi); err != nil {
		return err
	}
	if err := cfgApi.Validate(); err != nil {
		return err
	}

	return nil
}
