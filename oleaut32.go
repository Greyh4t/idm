package main

import (
	"syscall"
	"unsafe"
)

var (
	modoleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	procSysAllocString = modoleaut32.NewProc("SysAllocString")
	procSysFreeString  = modoleaut32.NewProc("SysFreeString")
	procSysStringLen   = modoleaut32.NewProc("SysStringLen")
)

func SysAllocString(psz *uint16) *uint16 {
	r1, _, _ := syscall.Syscall(procSysAllocString.Addr(), 1, uintptr(unsafe.Pointer(psz)), 0, 0)
	return (*uint16)(unsafe.Pointer(r1))
}

func SysFreeString(bstrString *uint16) {
	syscall.Syscall(procSysFreeString.Addr(), 1, uintptr(unsafe.Pointer(bstrString)), 0, 0)
}

func SysStringLen(bstr *uint16) uint32 {
	r1, _, _ := syscall.Syscall(procSysStringLen.Addr(), 1, uintptr(unsafe.Pointer(bstr)), 0, 0)
	return uint32(r1)
}
