package messages

type Response struct {
	JsonRPC string         `json:"jsonrpc"`
	ID      any            `json:"id"`
	Result  any            `json:"result"`
	Error   *ResponseError `json:"error"`
}
