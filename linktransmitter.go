package idm

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

type ICIDMLinkTransmitter2Vtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	SendLinkToIDM  uintptr
	SendLinkToIDM2 uintptr
	SendLinksArray uintptr
}

type ICIDMLinkTransmitter2 struct {
	vtbl *ICIDMLinkTransmitter2Vtbl
}

func (obj *ICIDMLinkTransmitter2) Release() uint32 {
	r1, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)

	return uint32(r1)
}

func (obj *ICIDMLinkTransmitter2) SendLinkToIDM(url, referer, cookies, data, user, password,
	localPath, localFileName *uint16, flags int32) uint32 {
	r1, _, _ := syscall.Syscall12(
		obj.vtbl.SendLinkToIDM,
		10,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(url)),
		uintptr(unsafe.Pointer(referer)),
		uintptr(unsafe.Pointer(cookies)),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(user)),
		uintptr(unsafe.Pointer(password)),
		uintptr(unsafe.Pointer(localPath)),
		uintptr(unsafe.Pointer(localFileName)),
		uintptr(flags),
		0,
		0,
	)

	return uint32(r1)
}

func (obj *ICIDMLinkTransmitter2) SendLinkToIDM2(url, referer, cookies, data, user, password,
	localPath, localFileName *uint16, flags int32, reserved1, reserved2 *win.VARIANT) uint32 {
	r1, _, _ := syscall.Syscall12(
		obj.vtbl.SendLinkToIDM2,
		12,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(url)),
		uintptr(unsafe.Pointer(referer)),
		uintptr(unsafe.Pointer(cookies)),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(user)),
		uintptr(unsafe.Pointer(password)),
		uintptr(unsafe.Pointer(localPath)),
		uintptr(unsafe.Pointer(localFileName)),
		uintptr(flags),
		uintptr(unsafe.Pointer(reserved1)),
		uintptr(unsafe.Pointer(reserved2)),
	)

	return uint32(r1)
}

func (obj *ICIDMLinkTransmitter2) SendLinksArray(location *uint16, pLinksArray *win.VARIANT) uint32 {
	r1, _, _ := syscall.Syscall(
		obj.vtbl.SendLinksArray,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(location)),
		uintptr(unsafe.Pointer(pLinksArray)),
	)

	return uint32(r1)
}
