package main

import ("fmt"
  "github.com/sjcockell/CoffeeBot/reminder"
  "github.com/sjcockell/CoffeeBot/news"
  "github.com/sjcockell/CoffeeBot/notifier"
)

func main() {
  reminder := reminder.Reminder()
  titles, links := news.News()
  fmt.Println(reminder)
  fmt.Println(titles[0])
  fmt.Println(links[1])
  notifier.Send(from='CoffeeBot@bsu-srv.ncl.ac.uk',
      ['sjcockell@bsu-srv.ncl.ac.uk'],
      'This is a test',
      'normal')
}
