// +build ignore

package apm

// #include <sys/types.h>
// #include <machine/apmvar.h>
import "C"

// ApmPowerInfo from C
type ApmPowerInfo C.struct_apm_power_info

// ApmIOCGetPower our getpower var
const ApmIOCGetPower = C.APM_IOC_GETPOWER

// ApmIOCSuspend fires off a suspend
const ApmIOCSuspend = C.APM_IOC_SUSPEND

// ApmIOCStandby puts system in standby
const ApmIOCStandby = C.APM_IOC_SUSPEND
