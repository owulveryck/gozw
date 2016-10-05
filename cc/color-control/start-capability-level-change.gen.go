// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package colorcontrol

import (
	"encoding/gob"
	"errors"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandStartCapabilityLevelChange cc.CommandID = 0x06

func init() {
	gob.Register(StartCapabilityLevelChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x33),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewStartCapabilityLevelChange)
}

func NewStartCapabilityLevelChange() cc.Command {
	return &StartCapabilityLevelChange{}
}

// <no value>
type StartCapabilityLevelChange struct {
	Properties1 struct {
		Res1 byte

		IgnoreStartState bool

		Updown bool

		Res2 bool
	}

	CapabilityId byte

	StartState byte
}

func (cmd StartCapabilityLevelChange) CommandClassID() cc.CommandClassID {
	return 0x33
}

func (cmd StartCapabilityLevelChange) CommandID() cc.CommandID {
	return CommandStartCapabilityLevelChange
}

func (cmd StartCapabilityLevelChange) CommandIDString() string {
	return "START_CAPABILITY_LEVEL_CHANGE"
}

func (cmd *StartCapabilityLevelChange) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Res1 = (payload[i] & 0x1F)

	cmd.Properties1.IgnoreStartState = payload[i]&0x20 == 0x20

	cmd.Properties1.Updown = payload[i]&0x40 == 0x40

	cmd.Properties1.Res2 = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CapabilityId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartState = payload[i]
	i++

	return nil
}

func (cmd *StartCapabilityLevelChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.Res1) & byte(0x1F)

		if cmd.Properties1.IgnoreStartState {
			val |= byte(0x20) // flip bits on
		} else {
			val &= ^byte(0x20) // flip bits off
		}

		if cmd.Properties1.Updown {
			val |= byte(0x40) // flip bits on
		} else {
			val &= ^byte(0x40) // flip bits off
		}

		if cmd.Properties1.Res2 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.CapabilityId)

	payload = append(payload, cmd.StartState)

	return
}
