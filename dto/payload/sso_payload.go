package payload

type SSOPayload struct {
	Ticket     string `json:"ticket"`
	ServiceUrl string `json:"service_url"`
}
