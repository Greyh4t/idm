package idm

import (
	"fmt"
	"syscall"
	"unicode/utf16"
)

func ErrMsg(errno int) string {
	// Ask windows for the remaining errors
	flags := syscall.FORMAT_MESSAGE_FROM_SYSTEM | syscall.FORMAT_MESSAGE_ARGUMENT_ARRAY | syscall.FORMAT_MESSAGE_IGNORE_INSERTS
	b := make([]uint16, 300)

	n, err := syscall.FormatMessage(uint32(flags), 0, uint32(errno), 0, b, nil)
	if err != nil {
		return fmt.Sprintf("error %d (FormatMessage failed with: %v)", errno, err)
	}

	// trim terminating \r and \n
	for ; n > 0 && (b[n-1] == '\n' || b[n-1] == '\r'); n-- {
	}

	return string(utf16.Decode(b[:n]))
}
