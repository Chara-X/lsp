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
			Notify("window/showMessage", protocol.ShowMessageParams{Type: protocol.MessageTypeInfo, Message: "initialized"})
			return nil
		},
	})
}
func Notify(name string, params any) {
	lsp.SendMessage(messages.Request{JsonRPC: "2.0", Method: name, Params: params})
}
