// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package notifications

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureToastNotificationManagerForUser string = "rc(Windows.UI.Notifications.ToastNotificationManagerForUser;{79ab57f6-43fe-487b-8a7f-99567200ae94})"

type ToastNotificationManagerForUser struct {
	ole.IUnknown
}

func (impl *ToastNotificationManagerForUser) CreateToastNotifierWithId(applicationId string) (*ToastNotifier, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiToastNotificationManagerForUser))
	defer itf.Release()
	v := (*iToastNotificationManagerForUser)(unsafe.Pointer(itf))
	return v.CreateToastNotifierWithId(applicationId)
}

const (
	GUIDiToastNotificationManagerForUser      string = "79ab57f6-43fe-487b-8a7f-99567200ae94"
	SignatureiToastNotificationManagerForUser string = "{79ab57f6-43fe-487b-8a7f-99567200ae94}"
)

type iToastNotificationManagerForUser struct {
	ole.IInspectable
}

type iToastNotificationManagerForUserVtbl struct {
	ole.IInspectableVtbl

	CreateToastNotifier       uintptr
	CreateToastNotifierWithId uintptr
	GetHistory                uintptr
	GetUser                   uintptr
}

func (v *iToastNotificationManagerForUser) VTable() *iToastNotificationManagerForUserVtbl {
	return (*iToastNotificationManagerForUserVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iToastNotificationManagerForUser) CreateToastNotifierWithId(applicationId string) (*ToastNotifier, error) {
	var out *ToastNotifier
	applicationIdHStr, err := ole.NewHString(applicationId)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().CreateToastNotifierWithId,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(applicationIdHStr),    // in string
		uintptr(unsafe.Pointer(&out)), // out ToastNotifier
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
