// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchtogglebinary

import "encoding/gob"

func init() {
	gob.Register(Set{})
}

// <no value>
type Set struct {
}

func (cmd Set) CommandClassID() byte {
	return 0x28
}

func (cmd Set) CommandID() byte {
	return byte(CommandSet)
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	return
}