package hmstt

const (
	MQ_CHANNEL_HMSTT = "hmstt_channel"
	PREFIX_HMSTT     = "hmstt"

	STATE_OFF = "off"
	STATE_ON  = "on"

	STATE_SWITCH_1 = "switch_1"
	STATE_SWITCH_2 = "switch_2"
	STATE_SWITCH_3 = "switch_3"
	STATE_SWITCH_4 = "switch_4"
)

var (
	SWITCH_STATES = []string{STATE_SWITCH_1, STATE_SWITCH_2, STATE_SWITCH_3, STATE_SWITCH_4}
)
