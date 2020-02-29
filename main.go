package main

import (
	"fmt"
	"github.com/rs/cors"
	"github.com/snookmz/frcApi/cli"
	"github.com/snookmz/frcApi/core"
	"github.com/snookmz/frcApi/handler"
	"github.com/snookmz/frcApi/logger"
	"github.com/snookmz/frcApi/loggerAbstraction"
	"log"
	"net/http"
	"os"
)

func main() {

	cli.ReadFlags()

	var iniFile string
	if core.CONFIG == "" {
		iniFile = "./config.ini"
	} else {
		iniFile = core.CONFIG
	}

	i, err := core.ReadConfig(iniFile)
	if err != nil {
		log.Panicf("error returned from core.ReadConfig(%s): %s\n", iniFile, err.Error())
	}

	var h handler.Handler
	logServer := logger.NewLogger(core.DEBUG, i.Log.File)
	h.Logger = loggerAbstraction.NewRegisteredLogger(logServer)
	h.Dir = i.Dir
	h.DropBox.AccessToken = i.DropBox.AccessToken
	h.DropBox.Path = i.DropBox.Path
	h.DropBox.Url = i.DropBox.Url

	if core.DEBUG >= 1 {
		_, _ = fmt.Fprintf(os.Stderr, "Listening on port: %v\n", i.Api.Port)
	}

	h.Logger.LPrintf(2, "Started, listening on port: %v\n\n", i.Api.Port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "X-Requested-With", "Authorization", "Id-token"},
		OptionsPassthrough: true,
		Debug: false,
	})

	router := h.GenerateRoutes()

	handlers := cors.Default().Handler(router)
	handlers = c.Handler(handlers)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",i.Api.Port), handlers))


}