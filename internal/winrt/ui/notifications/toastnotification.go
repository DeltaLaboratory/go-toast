// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package notifications

import (
	"syscall"
	"unsafe"

	"git.sr.ht/~jackmordaunt/go-toast/internal/winrt/data/xml/dom"
	"github.com/go-ole/go-ole"
)

const SignatureToastNotification string = "rc(Windows.UI.Notifications.ToastNotification;{997e2675-059e-4e60-8b06-1760917c8b80})"

func CreateToastNotification(content *dom.XmlDocument) (*ToastNotification, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.UI.Notifications.ToastNotification", ole.NewGUID(GUIDiToastNotificationFactory))
	if err != nil {
		return nil, err
	}
	v := (*iToastNotificationFactory)(unsafe.Pointer(inspectable))

	var out *ToastNotification
	hr, _, _ := syscall.SyscallN(
		v.VTable().CreateToastNotification,
		0,                                // this is a static func, so there's no this
		uintptr(unsafe.Pointer(content)), // in dom.XmlDocument
		uintptr(unsafe.Pointer(&out)),    // out ToastNotification
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

type ToastNotification struct {
	ole.IUnknown
}

const (
	GUIDiToastNotification      string = "997e2675-059e-4e60-8b06-1760917c8b80"
	SignatureiToastNotification string = "{997e2675-059e-4e60-8b06-1760917c8b80}"
)

type iToastNotification struct {
	ole.IInspectable
}

type iToastNotificationVtbl struct {
	ole.IInspectableVtbl

	GetContent        uintptr
	SetExpirationTime uintptr
	GetExpirationTime uintptr
	AddDismissed      uintptr
	RemoveDismissed   uintptr
	AddActivated      uintptr
	RemoveActivated   uintptr
	AddFailed         uintptr
	RemoveFailed      uintptr
}

func (v *iToastNotification) VTable() *iToastNotificationVtbl {
	return (*iToastNotificationVtbl)(unsafe.Pointer(v.RawVTable))
}

const (
	GUIDiToastNotificationFactory      string = "04124b20-82c6-4229-b109-fd9ed4662b53"
	SignatureiToastNotificationFactory string = "{04124b20-82c6-4229-b109-fd9ed4662b53}"
)

type iToastNotificationFactory struct {
	ole.IInspectable
}

type iToastNotificationFactoryVtbl struct {
	ole.IInspectableVtbl

	CreateToastNotification uintptr
}

func (v *iToastNotificationFactory) VTable() *iToastNotificationFactoryVtbl {
	return (*iToastNotificationFactoryVtbl)(unsafe.Pointer(v.RawVTable))
}
