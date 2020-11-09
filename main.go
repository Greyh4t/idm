package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

const (
	CLSCTX_LOCAL_SERVER = 0x00000004
)

var (
	clsid = NewGUID("{AC746233-E9D3-49CD-862F-068F7B7CCCA4}")
	iid   = NewGUID("{94D09862-1875-4FC9-B434-91CF25C840A1}")
)

type ICIDMLinkTransmitter2 struct {
	ICIDMLinkTransmitter
}

type ICIDMLinkTransmitter struct {
	IUnknown
}

type ICIDMLinkTransmitter2Vtbl struct {
	ICIDMLinkTransmitterVtbl
	SendLinkToIDM2 uintptr
	SendLinksArray uintptr
}

type ICIDMLinkTransmitterVtbl struct {
	IUnknownVtbl
	SendLinkToIDM uintptr
}

/*
   ( ['in'], BSTR, 'bstrUrl' ),
   ( ['in'], BSTR, 'bstrReferer' ),
   ( ['in'], BSTR, 'bstrCookies' ),
   ( ['in'], BSTR, 'bstrData' ),
   ( ['in'], BSTR, 'bstrUser' ),
   ( ['in'], BSTR, 'bstrPassword' ),
   ( ['in'], BSTR, 'bstrLocalPath' ),
   ( ['in'], BSTR, 'bstrLocalFileName' ),
   ( ['in'], c_int, 'lFlags' )),
*/
func (icIDMLinkT *ICIDMLinkTransmitter) SendLinkToIDM(url, referer, cookies, data, user, password,
	localPath, localFileName *uint16, flags int32) uint32 {
	vtbl := (*ICIDMLinkTransmitter2Vtbl)(unsafe.Pointer(icIDMLinkT.Vtbl))
	r1, _, _ := syscall.Syscall9(
		vtbl.SendLinkToIDM,
		9,
		uintptr(unsafe.Pointer(url)),
		uintptr(unsafe.Pointer(referer)),
		uintptr(unsafe.Pointer(cookies)),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(user)),
		uintptr(unsafe.Pointer(password)),
		uintptr(unsafe.Pointer(localPath)),
		uintptr(unsafe.Pointer(localFileName)),
		uintptr(flags),
	)
	return uint32(r1)
}

type IDMLinkTransmitter struct {
	object *ICIDMLinkTransmitter2
}

func (idmLinkT *IDMLinkTransmitter) SendLinkToIDM(url, referer, cookies, data, user, password,
	localPath, localFileName string, flags int) error {
	if idmLinkT.object == nil {
		return fmt.Errorf("IDMLinkTransmitter::SendLinkToIDM is nil")
	}

	urlRaw := SysAllocString(syscall.StringToUTF16Ptr(url))
	refererRaw := SysAllocString(syscall.StringToUTF16Ptr(referer))
	cookiesRaw := SysAllocString(syscall.StringToUTF16Ptr(cookies))
	dataRaw := SysAllocString(syscall.StringToUTF16Ptr(data))
	userRaw := SysAllocString(syscall.StringToUTF16Ptr(user))
	passwordRaw := SysAllocString(syscall.StringToUTF16Ptr(password))
	localPathRaw := SysAllocString(syscall.StringToUTF16Ptr(localPath))
	localFileNameRaw := SysAllocString(syscall.StringToUTF16Ptr(localFileName))
	defer func() {
		SysFreeString(urlRaw)
		SysFreeString(refererRaw)
		SysFreeString(cookiesRaw)
		SysFreeString(dataRaw)
		SysFreeString(userRaw)
		SysFreeString(passwordRaw)
		SysFreeString(localPathRaw)
		SysFreeString(localFileNameRaw)
	}()

	if hr := idmLinkT.object.SendLinkToIDM(urlRaw, refererRaw, cookiesRaw, dataRaw, userRaw,
		passwordRaw, localPathRaw, localFileNameRaw, int32(flags)); int32(hr) < 0 {
		return fmt.Errorf("IDMLinkTransmitter::SendLinkToIDM failed, %v", hr)
	}

	return nil
}

func NewIDMLinkTransmitter() (*IDMLinkTransmitter, error) {
	var object uintptr
	hr := CoCreateInstance(
		clsid,
		nil,
		CLSCTX_LOCAL_SERVER,
		iid,
		&object,
	)
	if hr < 0 {
		return nil, fmt.Errorf("CoCreateInstance failed, %v", hr)
	}
	return &IDMLinkTransmitter{
		object: (*ICIDMLinkTransmitter2)(unsafe.Pointer(object)),
	}, nil
}

func main() {
	coInitialize()

	idmLinkT, err := NewIDMLinkTransmitter()
	if err != nil {
		log.Fatal(err)
	}

	err = idmLinkT.SendLinkToIDM(
		"http://www.example.com",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		0,
	)
	if err != nil {
		log.Fatal(err)
	}
}
