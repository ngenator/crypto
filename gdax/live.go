package gdax

import (
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn *ws.Conn
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Watch(feedIDs ...string) {
	var Dialer ws.Dialer

	conn, _, err := Dialer.Dial(FEED_URL, nil)
	if err != nil {
		log.Error(err)
	}

	c.conn = conn

	var prices map[string]float64

	defer c.conn.Close()

	feeds := NewFeeds(feedIDs)

	go func() {
		subscribe := Message{
			Type: "subscribe",
			Channels: []MessageChannel{
				{
					Name:       "ticker",
					ProductIDs: feeds.IDs(),
				},
			},
		}

		log.Info("Subscribing to ", feeds.IDs())

		if err := c.conn.WriteJSON(subscribe); err != nil {
			log.Error(err)
		}
	}()

	log.Info("Watching feed...")

	message := TickerMessage{}
	for true {
		if err := c.conn.ReadJSON(&message); err != nil {
			log.Error(err)
			break
		}

		if prices[message.ProductID] != message.Price {
			prices[message.ProductID] = message.Price
			log.Printf("[%s] Price: $%f", message.ProductID, message.Price)
		}
	}
}
