// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package alarmv2

// <no value>

type AlarmTypeSupportedReport struct {
	NumberOfBitMasks byte

	V1Alarm bool

	BitMask byte
}

func ParseAlarmTypeSupportedReport(payload []byte) AlarmTypeSupportedReport {
	val := AlarmTypeSupportedReport{}

	i := 2

	val.NumberOfBitMasks = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		val.V1Alarm = true
	} else {
		val.V1Alarm = false
	}

	i += 1

	val.BitMask = payload[i]
	i++

	return val
}