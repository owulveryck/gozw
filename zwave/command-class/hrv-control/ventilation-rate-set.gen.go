// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package hrvcontrol

// <no value>

type HrvControlVentilationRateSet struct {
	VentilationRate byte
}

func ParseHrvControlVentilationRateSet(payload []byte) HrvControlVentilationRateSet {
	val := HrvControlVentilationRateSet{}

	i := 2

	val.VentilationRate = payload[i]
	i++

	return val
}