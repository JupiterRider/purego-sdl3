//go:build !windows

package shared

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

func Load(name string) (uintptr, error) {
	localName := fmt.Sprintf(".%s%s", string(os.PathSeparator), name)
	if p, err := purego.Dlopen(localName, purego.RTLD_LAZY); err == nil {
		return p, nil
	}
	return purego.Dlopen(name, purego.RTLD_LAZY)
}

func Get(lib uintptr, name string) uintptr {
	addr, err := purego.Dlsym(lib, name)
	if err != nil {
		panic(err)
	}
	return addr
}
