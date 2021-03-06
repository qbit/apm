// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs gen_defs.go

package apm

type ApmPowerInfo struct {
	Battery_state uint8
	Ac_state      uint8
	Battery_life  uint8
	Spare1        uint8
	Minutes_left  uint32
	Spare2        [6]uint32
}

const ApmIOCGetPower = 0x40204103

const ApmIOCSuspend = 0x20004102

const ApmIOCStandby = 0x20004102
