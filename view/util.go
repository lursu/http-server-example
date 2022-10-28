// in the real world these would probably be autogenerated protobufs
package view

import (
	"time"
)

type HealthCheck struct {
	Uptime time.Duration
	Msg    string
}

type Failure struct {
	Error string
	Msg   string
}