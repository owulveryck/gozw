// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package manufacturerspecificv2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandReport cc.CommandID = 0x05

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x72),
		Command:      cc.CommandID(0x05),
		Version:      2,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	ManufacturerId uint16

	ProductTypeId uint16

	ProductId uint16
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x72
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "MANUFACTURER_SPECIFIC_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductTypeId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ManufacturerId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ProductTypeId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ProductId)
		payload = append(payload, buf...)
	}

	return
}
