package slack

const (
	ResponseTypeEphemeral = "ephemeral"
	ResponseTypeInChannel = "in_channel"
)

type Response struct {
	Text string `json:"text"`
	Type string `json:"response_type"`
}
