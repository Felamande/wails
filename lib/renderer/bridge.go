package renderer

import (
	bridge "github.com/Felamande/wails/lib/renderer/bridge"
)

// NewBridge returns a new Bridge struct
func NewBridge() *bridge.Bridge {
	return &bridge.Bridge{}
}
