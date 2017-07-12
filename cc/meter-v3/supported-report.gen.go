// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package meterv3

import (
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gozwave/gozw/cc"
)

const CommandSupportedReport cc.CommandID = 0x04

func init() {
	gob.Register(SupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x32),
		Command:      cc.CommandID(0x04),
		Version:      3,
	}, NewSupportedReport)
}

func NewSupportedReport() cc.Command {
	return &SupportedReport{}
}

// <no value>
type SupportedReport struct {
	Properties1 struct {
		MeterType byte

		MeterReset bool
	}

	ScaleSupported byte
}

func (cmd SupportedReport) CommandClassID() cc.CommandClassID {
	return 0x32
}

func (cmd SupportedReport) CommandID() cc.CommandID {
	return CommandSupportedReport
}

func (cmd SupportedReport) CommandIDString() string {
	return "METER_SUPPORTED_REPORT"
}

func (cmd *SupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.Properties1) %d<=%d", len(payload), i)
	}

	cmd.Properties1.MeterType = (payload[i] & 0x1F)

	cmd.Properties1.MeterReset = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return fmt.Errorf("slice index out of bounds (.ScaleSupported) %d<=%d", len(payload), i)
	}

	cmd.ScaleSupported = payload[i]
	i++

	return nil
}

func (cmd *SupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.MeterType) & byte(0x1F)

		if cmd.Properties1.MeterReset {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ScaleSupported)

	return
}