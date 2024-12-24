//go:build darwin
// +build darwin

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Foundation

#include <AppKit/AppKit.h>
#include <stdlib.h>

void copyFileToClipboardInObjC(const char* cpath) {
    @autoreleasepool {
        NSString* nsPath = [NSString stringWithUTF8String:cpath];
        NSURL* fileURL = [NSURL fileURLWithPath:nsPath];

        NSPasteboard* pasteboard = [NSPasteboard generalPasteboard];
        [pasteboard clearContents];

        [pasteboard writeObjects:@[fileURL]];
    }
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func copyToClipboardInObjC(path string) error {
	cpath := C.CString(path)
	if cpath == nil {
		return errors.New("failed to allocate C string")
	}
	defer C.free(unsafe.Pointer(cpath))

	C.copyFileToClipboardInObjC(cpath)
	return nil
}
