// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package powerlevel

// <no value>

type PowerlevelSet struct {
	PowerLevel byte

	Timeout byte
}

func ParsePowerlevelSet(payload []byte) PowerlevelSet {
	val := PowerlevelSet{}

	i := 2

	val.PowerLevel = payload[i]
	i++

	val.Timeout = payload[i]
	i++

	return val
}