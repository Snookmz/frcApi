package cli

import (
	"bitbucket.org/turbosoftnetworks/netscope-api-v2/core"
	"flag"
	"fmt"
	"log"
	"os"
)

func ReadFlags () {
	config := flag.String("c", "config.ini", "location of the configuration file")
	version := flag.Bool("v", false, "Print version information")
	debug := flag.Int("d", 0, "Debug level: 1 ERROR, 2 ERROR|WARNING, 3 ERROR|WARNING|INFO, 4 ERROR|WARNING|INFO|TRACE")
	help := flag.Bool("h", false, "Command line options")
	//interactive := flag.Bool("i", false, "enter CLI tool in interactive mode")

	flag.Parse()

	if *help {
		fmt.Println(`Command line options:`)
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *config != "" {
		core.CONFIG = *config
	}


	if *version {
		fmt.Println("Version: ", core.VERSION)
		os.Exit(0)
	}


	if *debug > 0 {
		core.DEBUG = *debug
		_, err := fmt.Fprintln(os.Stderr)
		if err != nil {
			log.Fatalf("error returned from FPrintln: %s\n", err.Error())
		}

		_, err = fmt.Fprintln(os.Stderr, "Debug level: ", *debug)
		if err != nil {
			log.Fatalf("error returned from Fprintln: %s\n", err.Error())
		}

		_, err = fmt.Fprint(os.Stderr, "showing ")
			if err != nil {
			log.Fatalf("error returned from Fprintln: %s\n", err.Error())
		}

		for i := *debug; i <= 4; i++ {
			if *debug > i {
				break
			}
			_, err = fmt.Fprint(os.Stderr, core.Lognames[i], " ")
			if err != nil {
				log.Fatalf("error returned from Fprintln: %s\n", err.Error())
			}
		}
		_, err = fmt.Fprint (os.Stderr, "and ANY level (DELETEME) logs\n")
		if err != nil {
			log.Fatalf("error returned from Fprintln: %s\n", err.Error())
		}
	}


	// must be after initialisation functions.
	//if *interactive {
	//	core.ReadConfig("./config.ini")
	//	Interactive()
	//	os.Exit(0)
	//}


}

