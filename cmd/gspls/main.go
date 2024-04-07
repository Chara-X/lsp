package main

import (
	"github.com/Chara-X/lsp"
	"github.com/Chara-X/lsp/messages"
	"go.lsp.dev/protocol"
)

func main() {
	lsp.ListenAndServe(map[string]func(params any) any{
		"initialize": func(params any) any {
			return protocol.InitializeResult{Capabilities: protocol.ServerCapabilities{}}
		},
		"initialized": func(params any) any {
			notify("window/showMessage", protocol.ShowMessageParams{Type: protocol.MessageTypeInfo, Message: "initialized"})
			return nil
		},
	})
}
func notify(name string, params any) {
	lsp.Send(messages.Request{JsonRPC: "2.0", Method: name, Params: params})
}
