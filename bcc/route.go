package bcc

// RouteRule define route
type RouteRule struct {
	RouteRuleID        string `json:"routeRuleId"`
	RouteTableID       string `json:"routeTableId"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	NexthopID          string `json:"nexthopId"`
	NexthopType        string `json:"nexthopType"`
	Description        string `json:"description"`
}
