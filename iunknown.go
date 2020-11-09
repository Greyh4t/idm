package main

import (
	"syscall"
	"unsafe"
)

type IUnknown struct {
	Vtbl *IUnknownVtbl
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func (iunknown *IUnknown) QueryInterface(iid *GUID, object *uintptr) uint32 {
	r1, _, _ := syscall.Syscall(
		iunknown.Vtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(iunknown)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(object)))
	return uint32(r1)
}

func (iunknown *IUnknown) AddRef() uint32 {
	r1, _, _ := syscall.Syscall(iunknown.Vtbl.AddRef, 1, uintptr(unsafe.Pointer(iunknown)), 0, 0)
	return uint32(r1)
}

func (iunknown *IUnknown) Release() uint32 {
	r1, _, _ := syscall.Syscall(iunknown.Vtbl.Release, 1, uintptr(unsafe.Pointer(iunknown)), 0, 0)
	return uint32(r1)
}
