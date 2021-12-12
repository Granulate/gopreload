//+build linux,go1.8

package gopreload

import (
	"log"
	"os"
	"plugin"
	"runtime"
	"strings"
)

func init() {
	gpv := os.Getenv("GO_PRELOAD")
	if gpv == "" {
		return
	}

	for _, elem := range strings.Split(gpv, ":") {
		elem = strings.ReplaceAll(elem, "{version}", runtime.Version())
		_, err := plugin.Open(elem)
		if err != nil {
			log.Printf("%v from GO_PRELOAD cannot be loaded: %v", elem, err)
		}
	}
}
