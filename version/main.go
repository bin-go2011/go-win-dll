package main

import (
	"log"
	"syscall"

	"golang.org/x/sys/windows"
)

func main() {
	h, err := windows.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Panic("LoadLibrary", err)
	}
	defer windows.FreeLibrary(h)
	proc, err := windows.GetProcAddress(h, "GetVersion")
	if err != nil {
		log.Panic("GetProcAddress", err)
	}
	r, _, _ := syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
	major := byte(r)
	minor := uint8(r >> 8)
	build := uint16(r >> 16)
	print("windows version ", major, ".", minor, " (Build ", build, ")\n")
}
