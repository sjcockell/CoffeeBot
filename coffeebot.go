package main

import ("fmt"
  "github.com/sjcockell/CoffeeBot/reminder"
  "github.com/sjcockell/CoffeeBot/news"
  "github.com/sjcockell/CoffeeBot/notifier"
  "code.google.com/p/goexmpp"
)

func main() {
  reminder := reminder.Reminder()
  titles, links := news.News()
  fmt.Println(reminder)
  fmt.Println(titles[0])
  fmt.Println(links[1])
  from := xmpp.JID{Node: "CoffeeBot", Domain: "example.com"}
  notifier.Send(from,
      "target@example.com",
      "This is a test",
      "normal",
      "example.com",
      5222)
}
