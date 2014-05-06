package notifier

import (
	"github.com/mattn/go-xmpp"
	"log"
)


func Send(server string, username string, password string,
		notls bool, debug bool, session bool, to string, message string) {

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
	} else {
			talk.Send(xmpp.Chat{Remote: to, Type: "normal", Text: message})
	}
}
