// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package basictariffinfo

import "encoding/binary"

// <no value>

type BasicTariffInfoReport struct {
	TotalNoImportRates byte

	Dual bool

	E1CurrentRateInUse byte

	E1RateConsumptionRegister uint32

	E1TimeForNextRateHours byte

	E1TimeForNextRateMinutes byte

	E1TimeForNextRateSeconds byte

	E2CurrentRateInUse byte

	E2RateConsumptionRegister uint32
}

func ParseBasicTariffInfoReport(payload []byte) BasicTariffInfoReport {
	val := BasicTariffInfoReport{}

	i := 2

	val.TotalNoImportRates = (payload[i] & 0x0F)

	if payload[i]&0x80 == 0x80 {
		val.Dual = true
	} else {
		val.Dual = false
	}

	i += 1

	val.E1CurrentRateInUse = (payload[i] & 0x0F)

	i += 1

	val.E1RateConsumptionRegister = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	val.E1TimeForNextRateHours = payload[i]
	i++

	val.E1TimeForNextRateMinutes = payload[i]
	i++

	val.E1TimeForNextRateSeconds = payload[i]
	i++

	val.E2CurrentRateInUse = (payload[i] & 0x0F)

	i += 1

	val.E2RateConsumptionRegister = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	return val
}