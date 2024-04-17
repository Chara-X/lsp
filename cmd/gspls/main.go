package main

import (
	"github.com/Chara-X/lsp"
	"go.lsp.dev/protocol"
)

func main() {
	lsp.ListenAndServe(map[string]func(params any) any{
		"initialize": func(params any) any {
			return protocol.InitializeResult{Capabilities: protocol.ServerCapabilities{}}
		},
		"initialized": func(params any) any {
			lsp.Notify("window/showMessage", protocol.ShowMessageParams{Type: protocol.MessageTypeInfo, Message: "initialized"})
			return nil
		},
	})
}
