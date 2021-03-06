// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package associationv2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandSpecificGroupGet cc.CommandID = 0x0B

func init() {
	gob.Register(SpecificGroupGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x85),
		Command:      cc.CommandID(0x0B),
		Version:      2,
	}, NewSpecificGroupGet)
}

func NewSpecificGroupGet() cc.Command {
	return &SpecificGroupGet{}
}

// <no value>
type SpecificGroupGet struct {
}

func (cmd SpecificGroupGet) CommandClassID() cc.CommandClassID {
	return 0x85
}

func (cmd SpecificGroupGet) CommandID() cc.CommandID {
	return CommandSpecificGroupGet
}

func (cmd SpecificGroupGet) CommandIDString() string {
	return "ASSOCIATION_SPECIFIC_GROUP_GET"
}

func (cmd *SpecificGroupGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SpecificGroupGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
