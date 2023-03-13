package core

import "fmt"

var (
	ErrorServerStopped = fmt.Errorf("server stopped")
	ErrorProtocolIsNil = fmt.Errorf("protocol is nil")
)
