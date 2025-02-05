//go:build !linux

package device

import (
	"github.com/Mostbesep/EscapeRoute/wireguard/conn"
	"github.com/Mostbesep/EscapeRoute/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
