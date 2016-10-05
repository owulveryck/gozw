// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatsetpointv3

import (
	"encoding/gob"
	"errors"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandCapabilitiesReport cc.CommandID = 0x0A

func init() {
	gob.Register(CapabilitiesReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x43),
		Command:      cc.CommandID(0x0A),
		Version:      3,
	}, NewCapabilitiesReport)
}

func NewCapabilitiesReport() cc.Command {
	return &CapabilitiesReport{}
}

// <no value>
type CapabilitiesReport struct {
	Properties1 struct {
		SetpointType byte
	}

	Properties2 struct {
		Size byte

		Scale byte

		Precision byte
	}

	MinValue []byte

	Properties3 struct {
		Size byte

		Scale byte

		Precision byte
	}

	Maxvalue []byte
}

func (cmd CapabilitiesReport) CommandClassID() cc.CommandClassID {
	return 0x43
}

func (cmd CapabilitiesReport) CommandID() cc.CommandID {
	return CommandCapabilitiesReport
}

func (cmd CapabilitiesReport) CommandIDString() string {
	return "THERMOSTAT_SETPOINT_CAPABILITIES_REPORT"
}

func (cmd *CapabilitiesReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.SetpointType = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.Size = (payload[i] & 0x07)

	cmd.Properties2.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties2.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[1+2] >> 0) & 0x07
		cmd.MinValue = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties3.Size = (payload[i] & 0x07)

	cmd.Properties3.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties3.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[3+2] >> 0) & 0x07
		cmd.Maxvalue = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *CapabilitiesReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.SetpointType) & byte(0x0F)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.Size) & byte(0x07)

		val |= (cmd.Properties2.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties2.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.MinValue != nil && len(cmd.MinValue) > 0 {
		payload = append(payload, cmd.MinValue...)
	}

	{
		var val byte

		val |= (cmd.Properties3.Size) & byte(0x07)

		val |= (cmd.Properties3.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties3.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.Maxvalue != nil && len(cmd.Maxvalue) > 0 {
		payload = append(payload, cmd.Maxvalue...)
	}

	return
}
