package serialapi

import (
	"encoding/binary"
	"errors"

	"gitlab.com/helioslabs/gozw/frame"
	"gitlab.com/helioslabs/gozw/protocol"
	"gitlab.com/helioslabs/gozw/session"
)

func (s *Layer) MemoryGetID() (homeID uint32, nodeID byte, err error) {

	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionID: protocol.FnMemoryGetID,
		HasReturn:  true,
		ReturnCallback: func(err error, ret *frame.Frame) bool {
			done <- ret
			return false
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return 0, 0, errors.New("Error getting home/node id")
	}

	homeID = binary.BigEndian.Uint32(ret.Payload[1:5])
	nodeID = ret.Payload[5]

	return
}
