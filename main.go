package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"errors"
	"io/ioutil"
	"encoding/json"
	"os/signal"
	_ "expvar"
	"./application"
)


var cfg application.Config


func main() {

	cfgFilePath := flag.String("cfg", "", "Path to configuration file")
	onlyValidateCfg := flag.Bool("validate", false, "validate configuration file and stop")

	flag.Parse()

	err := openConfig(*cfgFilePath)
	if err != nil {
		fmt.Println("Unable to open application config: " + err.Error())
		os.Exit(2)
	}
	if *onlyValidateCfg {
		fmt.Println("Configuration file valid")
		os.Exit(0)
	}

	app, err := application.NewApplication(cfg)
	if err != nil {
		fmt.Println("Fatal error: " + err.Error())
		os.Exit(1)
	}
	defer app.Stop()

	//при принудительном завершении проиложения defer не работает
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, syscall.SIGTERM)
		<-sigchan
		app.Stop()
		os.Exit(0)
	}()

	go func() {
		for {
			sigchan := make(chan os.Signal)
			signal.Notify(sigchan, syscall.SIGUSR1)
			<-sigchan
			app.ReopenLogs()
		}
	}()

	err = app.Start()
	if err != nil {
		fmt.Println("Fatal error: " + err.Error())
		os.Exit(1)
	}
}

func openConfig(filename string) error {
	if filename == "" {
		return errors.New("please, specify configuration file path: -cfg=PATH")
	}

	file, e := ioutil.ReadFile(filename)
	if e != nil {
		return e
	}

	if e := json.Unmarshal(file, &cfg); e != nil {
		return e
	}
	if e := cfg.Validate(); e != nil {
		return e
	}

	return nil
}