package gdax

import (
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const FEED_URL = "wss://ws-feed.gdax.com"

type Feed struct {
	ID      string
	Updates chan *TickerMessage
}

func NewFeed(id string) Feed {
	return Feed{
		ID: id,
		Updates: make(chan *TickerMessage, 1),
	}
}

type Feeds map[string]Feed

func NewFeeds(feedIDs []string) Feeds {
	var feeds = make(Feeds)

	for _, feedID := range feedIDs {
		feeds[feedID] = NewFeed(feedID)
	}

	return feeds
}

func (f Feeds) IDs() []string {
	var feeds []string
	for key := range f {
		feeds = append(feeds, key)
	}
	return feeds
}

type FeedWatcher struct {
	conn *ws.Conn
	Feeds Feeds
}

func NewFeedWatcher() *FeedWatcher {
	var Dialer ws.Dialer

	conn, _, err := Dialer.Dial(FEED_URL, nil)
	if err != nil {
		log.Error(err)
	}

	return &FeedWatcher{
		conn: conn,
		Feeds: make(Feeds),
	}
}

func (f *FeedWatcher) Watch(feedIDs ...string) {

	go func() {
		subscribe := Message{
			Type: "subscribe",
			Channels: []MessageChannel{
				{
					Name:       "ticker",
					ProductIDs: feedIDs,
				},
			},
		}

		log.Info("Subscribing to ", feedIDs)

		if err := f.conn.WriteJSON(subscribe); err != nil {
			log.Error(err)
		}
	}()

	log.Info("Watching feeds...")

	var prices = make(map[string]float64)

	message := TickerMessage{}
	for true {
		if err := f.conn.ReadJSON(&message); err != nil {
			log.Error(err)
			break
		}

		if prices[message.ProductID] != message.Price {
			prices[message.ProductID] = message.Price
			log.Printf("[%s] Price: $%f", message.ProductID, message.Price)
		}
	}
}
