package main

import (
	"github.com/Mostbesep/EscapeRoute/proxy/pkg/mixed"
)

func main() {
	proxy := mixed.NewProxy()
	_ = proxy.ListenAndServe()
}
