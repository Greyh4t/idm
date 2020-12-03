package idm

import (
	"fmt"
	"unsafe"

	"github.com/lxn/win"
)

const (
	FlagNormal     Flag = iota // show confirmations dialogs;
	FlagSlience                // do not show any confirmations dialogs;
	FlagAddToQueue             // add to queue only, do not start downloading.
)

type Flag int

type Link struct {
	URL           string
	Referer       string
	Cookies       string
	UserAgent     string
	PostData      string
	Username      string
	Password      string
	LocalPath     string // where to save a file on your computer
	LocalFilename string // file name to save with
	Flags         Flag
}

type IDMLinkTransmitter2 struct {
	obj *ICIDMLinkTransmitter2
}

func NewIDMLinkTransmitter2() (*IDMLinkTransmitter2, error) {
	win.CoInitializeEx(nil, win.COINIT_MULTITHREADED)

	var object *ICIDMLinkTransmitter2
	hr := win.CoCreateInstance(
		CIDMLinkTransmitterCLSID,
		nil,
		win.CLSCTX_LOCAL_SERVER,
		ICIDMLinkTransmitter2IID,
		(*unsafe.Pointer)(unsafe.Pointer(&object)),
	)

	if hr != 0 {
		return nil, fmt.Errorf("CoCreateInstance failed, %s", ErrMsg(int(hr)))
	}

	return &IDMLinkTransmitter2{
		obj: object,
	}, nil
}

func (lt *IDMLinkTransmitter2) SendLinkToIDM(links ...Link) error {
	if len(links) == 0 {
		return fmt.Errorf("links can't be empty")
	}

	for _, link := range links {
		err := lt.sendLinkToIDM(link)
		if err != nil {
			return err
		}
	}

	return nil
}

func (lt *IDMLinkTransmitter2) sendLinkToIDM(link Link) error {
	if lt.obj == nil {
		return fmt.Errorf("IDMLinkTransmitter2::SendLinkToIDM is nil")
	}

	if len(link.URL) == 0 {
		return fmt.Errorf("url can't be empty")
	}

	var (
		bstrs                []*uint16
		bURL                 *uint16
		bReferer             *uint16
		bCookies             *uint16
		bData                *uint16
		bUsername            *uint16
		bPassword            *uint16
		bLocalPath           *uint16
		bLocalFilename       *uint16
		reserved1, reserved2 win.VARIANT
	)

	bURL = win.StringToBSTR(link.URL)
	bstrs = append(bstrs, bURL)
	if len(link.Referer) > 0 {
		bReferer = win.StringToBSTR(link.Referer)
		bstrs = append(bstrs, bReferer)
	}

	if len(link.Cookies) > 0 {
		bCookies = win.StringToBSTR(link.Cookies)
		bstrs = append(bstrs, bCookies)
	}

	if len(link.PostData) > 0 {
		bData = win.StringToBSTR(link.PostData)
		bstrs = append(bstrs, bData)
	}

	if len(link.Username) > 0 {
		bUsername = win.StringToBSTR(link.Username)
		bstrs = append(bstrs, bUsername)
	}

	if len(link.Password) > 0 {
		bPassword = win.StringToBSTR(link.Password)
		bstrs = append(bstrs, bPassword)
	}

	if len(link.LocalPath) > 0 {
		bLocalPath = win.StringToBSTR(link.LocalPath)
		bstrs = append(bstrs, bLocalPath)
	}

	if len(link.LocalFilename) > 0 {
		bLocalFilename = win.StringToBSTR(link.LocalFilename)
		bstrs = append(bstrs, bLocalFilename)
	}

	if len(link.UserAgent) > 0 {
		bUserAgent := win.StringToBSTR(link.UserAgent)
		bstrs = append(bstrs, bUserAgent)
		reserved1.SetBSTR(bUserAgent)
	}

	hr := lt.obj.SendLinkToIDM2(bURL, bReferer, bCookies, bData, bUsername,
		bPassword, bLocalPath, bLocalFilename, int32(link.Flags), &reserved1, &reserved2)

	lt.freeString(bstrs)

	if int32(hr) != 0 {
		return fmt.Errorf("IDMLinkTransmitter2::SendLinkToIDM failed, %s", ErrMsg(int(hr)))
	}

	return nil
}

func (lt *IDMLinkTransmitter2) freeString(bstrs []*uint16) {
	for _, bstr := range bstrs {
		win.SysFreeString(bstr)
	}
}

func (lt *IDMLinkTransmitter2) Release() {
	lt.obj.Release()

	win.CoUninitialize()
}
