package broker

import (
	"bytes"

	mqtt "github.com/mochi-mqtt/server/v2"
)

type LifecycleHookOptions struct {
	started chan bool
}

type LifecycleHook struct {
	mqtt.HookBase
	config *LifecycleHookOptions
}

func (h *LifecycleHook) Init(config any) error {
	if _, ok := config.(*LifecycleHookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	if config == nil {
		config = new(LifecycleHookOptions)
	}

	h.config = config.(*LifecycleHookOptions)

	return nil
}

func (h *LifecycleHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnStarted,
	}, []byte{b})
}

func (h *LifecycleHook) OnStarted() {
	h.config.started <- true
}
