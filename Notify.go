package lsp

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Chara-X/lsp/messages"
)

func Notify(name string, params any) {
	send(messages.Request{JsonRPC: "2.0", Method: name, Params: params})
}
func send[T messages.Request | messages.Response](message T) {
	var body, _ = json.Marshal(message)
	outMutex.Lock()
	defer outMutex.Unlock()
	os.Stdout.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(body), body))
}
