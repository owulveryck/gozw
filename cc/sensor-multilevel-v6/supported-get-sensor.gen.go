// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensormultilevelv6

import (
	"encoding/gob"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandSupportedGetSensor cc.CommandID = 0x01

func init() {
	gob.Register(SupportedGetSensor{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x31),
		Command:      cc.CommandID(0x01),
		Version:      6,
	}, NewSupportedGetSensor)
}

func NewSupportedGetSensor() cc.Command {
	return &SupportedGetSensor{}
}

// <no value>
type SupportedGetSensor struct {
}

func (cmd SupportedGetSensor) CommandClassID() cc.CommandClassID {
	return 0x31
}

func (cmd SupportedGetSensor) CommandID() cc.CommandID {
	return CommandSupportedGetSensor
}

func (cmd SupportedGetSensor) CommandIDString() string {
	return "SENSOR_MULTILEVEL_SUPPORTED_GET_SENSOR"
}

func (cmd *SupportedGetSensor) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SupportedGetSensor) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
