package gdax

type Message struct {
	Type          string           `json:"type"`
	ProductID     string           `json:"product_id"`
	ProductIDs    []string         `json:"product_ids"`
	TradeID       int              `json:"trade_id,number"`
	OrderID       string           `json:"order_id"`
	Sequence      int64            `json:"sequence,number"`
	MakerOrderId  string           `json:"maker_order_id"`
	TakerOrderId  string           `json:"taker_order_id"`
	Time          Time             `json:"time,string"`
	RemainingSize float64          `json:"remaining_size,string"`
	NewSize       float64          `json:"new_size,string"`
	OldSize       float64          `json:"old_size,string"`
	Size          float64          `json:"size,string"`
	Price         float64          `json:"price,string"`
	Side          string           `json:"side"`
	Reason        string           `json:"reason"`
	OrderType     string           `json:"order_type"`
	Funds         float64          `json:"funds,string"`
	NewFunds      float64          `json:"new_funds,string"`
	OldFunds      float64          `json:"old_funds,string"`
	Message       string           `json:"message"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      float64          `json:"last_size,string"`
	BestBid       float64          `json:"best_bid,string"`
	BestAsk       float64          `json:"best_ask,string"`
	Channels      []MessageChannel `json:"channels"`
}

type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}

type TickerMessage struct {
	Type      string  `json:"type"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price,string"`
	BestBid   float64 `json:"best_bid,string"`
	BestAsk   float64 `json:"best_ask,string"`
}
