package cache

import "log"

var verbose = false

func EnableVerboseLog() {
	verbose = true
}

func verboseLogf(format string, v ...interface{}) {
	if verbose {
		log.Printf(format, v...)
	}
}
