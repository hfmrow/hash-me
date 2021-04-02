// main.go

/*
	CLI version
	Copyright ©2020 H.F.M - Hash Me v1.0 github.com/hfmrow/hash-me
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"log"
	"os"
)

func main() {

	var err error

	/* Init & read options file */
	opt = new(MainOpt)
	opt.Init()

	err = chkCmdLine()

	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}

	doIt()
}
