package main

import ("fmt"
  "reminder"
  "news"
)

func main() {
  reminder := reminder.Reminder()
  titles, links := news.News()
  fmt.Println(reminder)
  fmt.Println(titles[0])
  fmt.Println(links[1])
}
