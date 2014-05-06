package notifier

import (
	"fmt"
	"github.com/mattn/go-xmpp"
	"log"
)


func Send(server string, username string, password string,
		notls bool, debug bool, session bool) {

	var talk *xmpp.Client
	var err error
	options := xmpp.Options{Host: server,
		User:     username,
		Password: password,
		NoTLS:    notls,
		Debug:    debug,
		Session:  session}

	talk, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}
}
