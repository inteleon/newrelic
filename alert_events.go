package newrelic

type Event struct {
	Id            int    `json:"id,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	Product       string `json:"product,omitempty"`
	EntityType    string `json:"entity_type,omitempty"`
	EntityGroupId int    `json:"entity_group_id,omitempty"`
	EntityId      int    `json:"entity_id,omitempty"`
	Priority      string `json:"priority,omitempty"`
	Description   string `json:"description,omitempty"`
	Timestamp     int    `json:"timestamp,omitempty"`
	IncidentId    int    `json:"incident_id"`
}

type EventFilter struct {
	Product       string `json:"product,omitempty"`
	EntityType    string `json:"entity_type,omitempty"`
	EntityGroupId int    `json:"entity_group_id,omitempty"`
	EntityId      int    `json:"entity_id,omitempty"`
	EventType     string `json:"event_type,omitempty"`
}

type EventOptions struct {
	Filter EventFilter `json:"filter,omitempty"`
	Page   int         `json:"page,omitempty"`
}

type EventResponse struct {
	RecentEvents []Event `json:"recent_events,omitempty"`
}

func (c *Client) GetEvents(options *EventOptions) (*EventResponse, error) {
	events := &EventResponse{}
	err := c.doRequest("GET", "/alerts_events.json", nil, events)
	if err != nil {
		return nil, err
	}
	return events, nil
}