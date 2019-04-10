package main

import (
	"github.com/gonutz/w32"
	"syscall"
	"unsafe"
)

func main() {
	const text = "✓"
	if w32.OpenClipboard(0) {
		w32.EmptyClipboard()
		data := syscall.StringToUTF16(text)
		clipBuffer := w32.GlobalAlloc(w32.GMEM_DDESHARE, uint32(len(data)*2))
		w32.MoveMemory(
			w32.GlobalLock(clipBuffer),
			unsafe.Pointer(&data[0]),
			uint32(len(data)*2),
		)
		w32.GlobalUnlock(clipBuffer)
		w32.SetClipboardData(
			w32.CF_UNICODETEXT,
			w32.HANDLE(unsafe.Pointer(clipBuffer)),
		)
		w32.CloseClipboard()
	}
}
