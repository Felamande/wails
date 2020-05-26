package webview

/*
#include <windows.h>
*/
import "C"

import (
	"unsafe"
)

//export _webviewDispatchGoCallback
func _webviewDispatchGoCallback(index unsafe.Pointer) {
	// fmt.Printf("thread %v enter _webviewDispatchGoCallback\n", GetCurrentThreadId())
	var f func()
	m.Lock()
	f = fns[uintptr(index)]
	delete(fns, uintptr(index))
	m.Unlock()
	f()
}

//export _webviewExternalInvokeCallback
func _webviewExternalInvokeCallback(w unsafe.Pointer, data unsafe.Pointer) {
	// fmt.Println("thread %v enter _webviewExternalInvokeCallback", GetCurrentThreadId())

	m.Lock()
	var (
		cb ExternalInvokeCallbackFunc
		wv WebView
	)
	for wv, cb = range cbs {
		if wv.(*webview).w == w {
			break
		}
	}
	m.Unlock()
	cb(wv, C.GoString((*C.char)(data)))
}
