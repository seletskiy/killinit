package main

var signals map[string]int = map[string]int{
	"ABRT":   0x6,
	"ALRM":   0xe,
	"BUS":    0xa,
	"CHLD":   0x14,
	"CONT":   0x13,
	"EMT":    0x7,
	"FPE":    0x8,
	"HUP":    0x1,
	"ILL":    0x4,
	"INFO":   0x1d,
	"INT":    0x2,
	"IO":     0x17,
	"IOT":    0x6,
	"KILL":   0x9,
	"PIPE":   0xd,
	"PROF":   0x1b,
	"QUIT":   0x3,
	"SEGV":   0xb,
	"STOP":   0x11,
	"SYS":    0xc,
	"TERM":   0xf,
	"THR":    0x20,
	"TRAP":   0x5,
	"TSTP":   0x12,
	"TTIN":   0x15,
	"TTOU":   0x16,
	"URG":    0x10,
	"USR1":   0x1e,
	"USR2":   0x1f,
	"VTALRM": 0x1a,
	"WINCH":  0x1c,
	"XCPU":   0x18,
	"XFSZ":   0x19,
}
