// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandCommandsSupportedGet cc.CommandID = 0x02

func init() {
	gob.Register(CommandsSupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x98),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewCommandsSupportedGet)
}

func NewCommandsSupportedGet() cc.Command {
	return &CommandsSupportedGet{}
}

// <no value>
type CommandsSupportedGet struct {
}

func (cmd CommandsSupportedGet) CommandClassID() cc.CommandClassID {
	return 0x98
}

func (cmd CommandsSupportedGet) CommandID() cc.CommandID {
	return CommandCommandsSupportedGet
}

func (cmd CommandsSupportedGet) CommandIDString() string {
	return "SECURITY_COMMANDS_SUPPORTED_GET"
}

func (cmd *CommandsSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *CommandsSupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
