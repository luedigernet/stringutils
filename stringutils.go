package stringutils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func SayHello(name string)string{

	log.Debug("Enter SayHello method")
	return fmt.Sprintf("Hello % how are you?\n",name)
	log.Debug("Leaving SayHello method")
	return ""
}
