// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package notifications

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const (
	GUIDiToastNotificationManagerStatics5      string = "d6f5f569-d40d-407c-8989-88cab42cfd14"
	SignatureiToastNotificationManagerStatics5 string = "{d6f5f569-d40d-407c-8989-88cab42cfd14}"
)

type iToastNotificationManagerStatics5 struct {
	ole.IInspectable
}

type iToastNotificationManagerStatics5Vtbl struct {
	ole.IInspectableVtbl

	GetDefault uintptr
}

func (v *iToastNotificationManagerStatics5) VTable() *iToastNotificationManagerStatics5Vtbl {
	return (*iToastNotificationManagerStatics5Vtbl)(unsafe.Pointer(v.RawVTable))
}

func GetDefault() (*ToastNotificationManagerForUser, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.UI.Notifications.ToastNotificationManager", ole.NewGUID(GUIDiToastNotificationManagerStatics5))
	if err != nil {
		return nil, err
	}
	v := (*iToastNotificationManagerStatics5)(unsafe.Pointer(inspectable))

	var out *ToastNotificationManagerForUser
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDefault,
		0,                             // this is a static func, so there's no this
		uintptr(unsafe.Pointer(&out)), // out ToastNotificationManagerForUser
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
