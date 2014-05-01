package news

import ( "fmt"
  "github.com/ungerik/go-rss"
)

func News() ([10]string, [10]string) {
  var titles = [10]string{}
  var links = [10]string{}
  channel, err := rss.Read("https://news.google.com/news/feeds?q=coffee+science&output=rss")
  if err != nil {

  } else {
    for i:=0;i<len(channel.Item);i++ {
      titles[i] = channel.Item[i].Title
      links[i] = channel.Item[i].Link
      fmt.Println(channel.Item[i].Title)
      fmt.Println(channel.Item[i].PubDate)
    }
  }
  return titles, links
}
