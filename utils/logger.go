package utils

import "log"

// Log is used to log message consistently
func Log(m interface{}) {
	log.Printf("\n\n\t%v\n", m)
}
