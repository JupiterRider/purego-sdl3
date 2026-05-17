//go:build windows

package shared

import (
	"fmt"
	"syscall"
)

func Load(name string) (uintptr, error) {
	handle, err := syscall.LoadLibrary(name)
	if err != nil {
		err = fmt.Errorf("%s: error loading library: %w", name, err)
	}
	return uintptr(handle), err
}

func Get(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}
	return addr
}
