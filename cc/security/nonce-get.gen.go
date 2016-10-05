// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"

	"gitlab.com/helioslabs/gozw/cc"
)

const CommandNonceGet cc.CommandID = 0x40

func init() {
	gob.Register(NonceGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x98),
		Command:      cc.CommandID(0x40),
		Version:      1,
	}, NewNonceGet)
}

func NewNonceGet() cc.Command {
	return &NonceGet{}
}

// <no value>
type NonceGet struct {
}

func (cmd NonceGet) CommandClassID() cc.CommandClassID {
	return 0x98
}

func (cmd NonceGet) CommandID() cc.CommandID {
	return CommandNonceGet
}

func (cmd NonceGet) CommandIDString() string {
	return "SECURITY_NONCE_GET"
}

func (cmd *NonceGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *NonceGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
