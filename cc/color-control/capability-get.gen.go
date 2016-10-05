// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package colorcontrol

import (
	"encoding/gob"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandCapabilityGet cc.CommandID = 0x01

func init() {
	gob.Register(CapabilityGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x33),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewCapabilityGet)
}

func NewCapabilityGet() cc.Command {
	return &CapabilityGet{}
}

// <no value>
type CapabilityGet struct {
}

func (cmd CapabilityGet) CommandClassID() cc.CommandClassID {
	return 0x33
}

func (cmd CapabilityGet) CommandID() cc.CommandID {
	return CommandCapabilityGet
}

func (cmd CapabilityGet) CommandIDString() string {
	return "CAPABILITY_GET"
}

func (cmd *CapabilityGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *CapabilityGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
