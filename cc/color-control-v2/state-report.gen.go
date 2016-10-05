// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package colorcontrolv2

import (
	"encoding/gob"
	"errors"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandStateReport cc.CommandID = 0x04

func init() {
	gob.Register(StateReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x33),
		Command:      cc.CommandID(0x04),
		Version:      2,
	}, NewStateReport)
}

func NewStateReport() cc.Command {
	return &StateReport{}
}

// <no value>
type StateReport struct {
	CapabilityId byte

	State byte
}

func (cmd StateReport) CommandClassID() cc.CommandClassID {
	return 0x33
}

func (cmd StateReport) CommandID() cc.CommandID {
	return CommandStateReport
}

func (cmd StateReport) CommandIDString() string {
	return "STATE_REPORT"
}

func (cmd *StateReport) UnmarshalBinary(data []byte) error {
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

	cmd.CapabilityId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.State = payload[i]
	i++

	return nil
}

func (cmd *StateReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.CapabilityId)

	payload = append(payload, cmd.State)

	return
}
