// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecificv2

import "fmt"

type CommandID byte

const (
	CommandGet CommandID = 0x04

	CommandReport CommandID = 0x05

	CommandDeviceSpecificGet CommandID = 0x06

	CommandDeviceSpecificReport CommandID = 0x07
)

func (c CommandID) String() string {
	switch c {

	case CommandGet:
		return "MANUFACTURER_SPECIFIC_GET"

	case CommandReport:
		return "MANUFACTURER_SPECIFIC_REPORT"

	case CommandDeviceSpecificGet:
		return "DEVICE_SPECIFIC_GET"

	case CommandDeviceSpecificReport:
		return "DEVICE_SPECIFIC_REPORT"

	default:
		return fmt.Sprintf("Unknown (0x%X)", byte(c))
	}
}