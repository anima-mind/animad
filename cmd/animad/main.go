// animad — Anima server runtime (daemon / REST API).
// Blueprint: https://github.com/anima-mind/anima (spec doc 03, perfil server §A.5).
//
// Estado: pre-implementación. Este binario es el esqueleto del daemon;
// el plan de implementación del perfil server se escribirá en el blueprint
// (hermano del doc 04) antes de la Fase 0 de este runtime.
package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.0.1"

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("animad %s\n", version)
		return
	}

	fmt.Fprintln(os.Stderr, "animad: daemon not implemented yet — blueprint at https://github.com/anima-mind/anima")
	os.Exit(1)
}
