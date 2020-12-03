package idm

import (
	"syscall"

	"github.com/lxn/win"
)

var (
	CIDMLinkTransmitterCLSID = NewREFCLSID("{AC746233-E9D3-49CD-862F-068F7B7CCCA4}")
	ICIDMLinkTransmitterIID  = NewREFIID("{4BD46AAE-C51F-4BF7-8BC0-2E86E33D1873}")
	ICIDMLinkTransmitter2IID = NewREFIID("{94D09862-1875-4FC9-B434-91CF25C840A1}")
)

func NewREFCLSID(guid string) win.REFCLSID {
	clsid := win.CLSID(*NewGUID(guid))
	return win.REFCLSID(&clsid)
}

func NewREFIID(guid string) win.REFIID {
	iid := win.IID(*NewGUID(guid))
	return win.REFIID(&iid)
}

func NewGUID(guid string) *syscall.GUID {
	d := []byte(guid)
	var d1, d2, d3, d4a, d4b []byte

	switch len(d) {
	case 38:
		if d[0] != '{' || d[37] != '}' {
			return nil
		}
		d = d[1:37]
		fallthrough
	case 36:
		if d[8] != '-' || d[13] != '-' || d[18] != '-' || d[23] != '-' {
			return nil
		}
		d1 = d[0:8]
		d2 = d[9:13]
		d3 = d[14:18]
		d4a = d[19:23]
		d4b = d[24:36]
	case 32:
		d1 = d[0:8]
		d2 = d[8:12]
		d3 = d[12:16]
		d4a = d[16:20]
		d4b = d[20:32]
	default:
		return nil
	}

	var g syscall.GUID
	var ok1, ok2, ok3, ok4 bool
	g.Data1, ok1 = decodeHexUint32(d1)
	g.Data2, ok2 = decodeHexUint16(d2)
	g.Data3, ok3 = decodeHexUint16(d3)
	g.Data4, ok4 = decodeHexByte64(d4a, d4b)
	if ok1 && ok2 && ok3 && ok4 {
		return &g
	}
	return nil
}

func decodeHexUint32(src []byte) (value uint32, ok bool) {
	var b1, b2, b3, b4 byte
	var ok1, ok2, ok3, ok4 bool
	b1, ok1 = decodeHexByte(src[0], src[1])
	b2, ok2 = decodeHexByte(src[2], src[3])
	b3, ok3 = decodeHexByte(src[4], src[5])
	b4, ok4 = decodeHexByte(src[6], src[7])
	value = (uint32(b1) << 24) | (uint32(b2) << 16) | (uint32(b3) << 8) | uint32(b4)
	ok = ok1 && ok2 && ok3 && ok4
	return
}

func decodeHexUint16(src []byte) (value uint16, ok bool) {
	var b1, b2 byte
	var ok1, ok2 bool
	b1, ok1 = decodeHexByte(src[0], src[1])
	b2, ok2 = decodeHexByte(src[2], src[3])
	value = (uint16(b1) << 8) | uint16(b2)
	ok = ok1 && ok2
	return
}

func decodeHexByte64(s1 []byte, s2 []byte) (value [8]byte, ok bool) {
	var ok1, ok2, ok3, ok4, ok5, ok6, ok7, ok8 bool
	value[0], ok1 = decodeHexByte(s1[0], s1[1])
	value[1], ok2 = decodeHexByte(s1[2], s1[3])
	value[2], ok3 = decodeHexByte(s2[0], s2[1])
	value[3], ok4 = decodeHexByte(s2[2], s2[3])
	value[4], ok5 = decodeHexByte(s2[4], s2[5])
	value[5], ok6 = decodeHexByte(s2[6], s2[7])
	value[6], ok7 = decodeHexByte(s2[8], s2[9])
	value[7], ok8 = decodeHexByte(s2[10], s2[11])
	ok = ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8
	return
}

func decodeHexByte(c1, c2 byte) (value byte, ok bool) {
	var n1, n2 byte
	var ok1, ok2 bool
	n1, ok1 = decodeHexChar(c1)
	n2, ok2 = decodeHexChar(c2)
	value = (n1 << 4) | n2
	ok = ok1 && ok2
	return
}

func decodeHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}
