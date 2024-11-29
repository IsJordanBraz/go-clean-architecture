package event

import "time"

type OrderCreated struct {
	Name     string
	Payload  interface{}
	DateTime time.Time
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name:     "OrderCreated",
		DateTime: time.Now(),
	}
}

func (o *OrderCreated) GetName() string {
	return o.Name
}

func (o *OrderCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *OrderCreated) GetDateTime() time.Time {
	return o.DateTime
}

func (o *OrderCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}
