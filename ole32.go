package main

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	modole32             = syscall.NewLazyDLL("ole32.dll")
	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procCoInitialize     = modole32.NewProc("CoInitialize")
)

func CoCreateInstance(clsid *GUID, outer *IUnknown, clsContext uint32, iid *GUID, object *uintptr) uint32 {
	r1, _, _ := syscall.Syscall6(
		procCoCreateInstance.Addr(),
		5,
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(outer)),
		uintptr(clsContext),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(object)),
		0)
	return uint32(r1)
}

func coInitialize() {
	r, _, err := procCoInitialize.Call(uintptr(0))
	if r < 0 {
		log.Fatal(r, err)
	}
}
