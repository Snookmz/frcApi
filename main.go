package main

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/cli"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/core"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/handler"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/logger"
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/loggerAbstraction"
	"fmt"
	"github.com/rs/cors"
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

	router := h.GenerateRoutes(h.NsAuthenticationMw)

	handlers := cors.Default().Handler(router)
	handlers = c.Handler(handlers)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",i.Api.Port), handlers))


}