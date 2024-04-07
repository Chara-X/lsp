package lsp

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/Chara-X/lsp/messages"
)

func ListenAndServe(handlers map[string]func(params any) any) {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(ScanMessage)
	for scanner.Scan() {
		go func() {
			var req messages.Request
			json.Unmarshal(scanner.Bytes(), &req)
			if res := handlers[req.Method](req.Params); res != nil {
				Send(messages.Response{JsonRPC: "2.0", ID: req.ID, Result: res})
			}
		}()
	}
}
