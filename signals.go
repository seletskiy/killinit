package main

var signals map[string]int = map[string]int{
	"ABRT":   0x6,
	"ALRM":   0xe,
	"BUS":    0x7,
	"CHLD":   0x11,
	"CLD":    0x11,
	"CONT":   0x12,
	"FPE":    0x8,
	"HUP":    0x1,
	"ILL":    0x4,
	"INT":    0x2,
	"IO":     0x1d,
	"IOT":    0x6,
	"KILL":   0x9,
	"PIPE":   0xd,
	"POLL":   0x1d,
	"PROF":   0x1b,
	"PWR":    0x1e,
	"QUIT":   0x3,
	"SEGV":   0xb,
	"STKFLT": 0x10,
	"STOP":   0x13,
	"SYS":    0x1f,
	"TERM":   0xf,
	"TRAP":   0x5,
	"TSTP":   0x14,
	"TTIN":   0x15,
	"TTOU":   0x16,
	"UNUSED": 0x1f,
	"URG":    0x17,
	"USR1":   0xa,
	"USR2":   0xc,
	"VTALRM": 0x1a,
	"WINCH":  0x1c,
	"XCPU":   0x18,
	"XFSZ":   0x19,
}
