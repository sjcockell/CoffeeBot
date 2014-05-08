package notifier

import ("code.google.com/p/goexmpp")

func Send(from string, to []string, text string, messageType string) {
	var jid xmpp.JID = from
  var pw string = "notapassword"
  if jid.Domain == "" || *pw == "" {
    os.Exit(2)
  }
  c, err := xmpp.NewClient(&jid, *pw, nil)
  if err != nil {
    log.Fatalf("NewClient(%v): %v", jid, err)
  }
  defer close(c.Out)
}
