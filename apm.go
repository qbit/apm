/*
Package apm allows users to query OpenBSD's power management system.
*/
package apm

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// HasBatt returns true if the state is any state but '4'. 4 explicitly means
// a battery is absent. See https://man.openbsd.org/man8/amd64/apm.8 for more
// information.
func (a *ApmPowerInfo) HasBatt() bool {
	return a.Battery_state != 0x4
}

// Charging returns true of the device is actively charging.
func (a *ApmPowerInfo) Charging() bool {
	return a.Ac_state == 0x1
}

// Get queries the system for power information.
func (a *ApmPowerInfo) Get() error {
	f, err := os.OpenFile("/dev/apm", syscall.O_RDONLY, 0444)

	if err != nil {
		return err
	}

	defer f.Close()

	_, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(f.Fd()),
		uintptr(ApmIOCGetPower),
		uintptr(unsafe.Pointer(a)))

	if errno != 0 {
		return errno
	}

	return nil
}

// NewPwrInfo returns an ApmPowerInfo instance, prepopulated with the current
// state.
func NewPwrInfo() (*ApmPowerInfo, error) {
	d := ApmPowerInfo{}
	err := d.Get()
	if err != nil {
		return nil, err
	}

	return &d, nil
}
