// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementproxy

import "errors"

// <no value>

type NodeInfoCachedGet struct {
	SeqNo byte

	Properties1 struct {
		MaxAge byte
	}

	NodeId byte
}

func (cmd *NodeInfoCachedGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.MaxAge = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	return nil
}