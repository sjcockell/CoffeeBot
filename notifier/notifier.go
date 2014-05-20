package notifier

import ("os"
		"log"
		"fmt"
		"encoding/xml"
		"strings"
        "code.google.com/p/goexmpp"
	   )

func Send(from xmpp.JID, to string, text string, messageType string, host string, port int) {
  var jid xmpp.JID = from
  var pw string = "notapassword"
  if jid.Domain == "" || pw == "" {
    os.Exit(2)
  }
	fmt.Println("Connecting...")
  c, err := xmpp.NewClientFromHost(&jid, pw, nil, host, port)
  if err != nil {
    log.Fatalf("NewClient(%v): %v", jid, err)
  } else {
		fmt.Println("Success?")
  }
  defer close(c.Out)
  fmt.Println("Starting session...")
  err = c.StartSession(true, &xmpp.Presence{})
  if err != nil {
    log.Fatalf("StartSession: %v", err)
  }
  fmt.Println("Getting roster...")
  roster := xmpp.Roster(c)
  fmt.Printf("%d roster entries:\n", len(roster))
  for i, entry := range roster {
    fmt.Printf("%d: %v\n", i, entry)
  }
  go func(ch <-chan xmpp.Stanza) {
    for obj := range ch {
      fmt.Printf("s: %v\n", obj)
    }
    fmt.Println("done reading")
  }(c.In)
  p := make([]byte, 1024)
  for {
    nr, _ := os.Stdin.Read(p)
    if nr == 0 {
      break
    }
    s := string(p)
    dec := xml.NewDecoder(strings.NewReader(s))
    t, err := dec.Token()
    if err != nil {
      fmt.Printf("token: %s\n", err)
      break
    }
    var se *xml.StartElement
    var ok bool
    if se, ok = t.(*xml.StartElement); !ok {
      fmt.Println("Couldn't find start element")
      break
    }
    var stan xmpp.Stanza
    switch se.Name.Local {
      case "iq":
        stan = &xmpp.Iq{}
      case "message":
        stan = &xmpp.Message{}
      case "presence":
        stan = &xmpp.Presence{}
      default:
        fmt.Println("Can't parse non-stanza.")
      continue
    }
    err = dec.Decode(stan)
    if err == nil {
      c.Out <- stan
    } else {
      fmt.Printf("Parse error: %v\n", err)
      break
    }
  }
  fmt.Println("done sending")
}
