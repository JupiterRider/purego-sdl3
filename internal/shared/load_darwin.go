//go:build darwin

package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ebitengine/purego"
)

func Load(name string) (uintptr, error) {
	candidates := darwinLibraryCandidates(name)

	var lastErr error
	for _, candidate := range candidates {
		if candidate != name {
			info, err := os.Stat(candidate)
			if err != nil {
				lastErr = err
				continue
			}
			if info.IsDir() {
				lastErr = fmt.Errorf("%s is a directory", candidate)
				continue
			}
		}

		p, err := purego.Dlopen(candidate, purego.RTLD_LAZY)
		if err == nil {
			return p, nil
		}
		lastErr = err
	}

	return 0, lastErr
}

func darwinLibraryCandidates(name string) []string {
	candidates := make([]string, 0, 5)

	candidates = append(candidates, filepath.Join(".", name))

	if exe, err := os.Executable(); err == nil {
		if resolvedExe, err := filepath.EvalSymlinks(exe); err == nil {
			exe = resolvedExe
		}

		exeDir := filepath.Dir(exe)
		candidates = append(candidates, filepath.Join(exeDir, name))

		if bundleRoot, ok := darwinBundleRootFromExecutable(exe); ok {
			candidates = append(candidates, filepath.Join(bundleRoot, "Contents", "MacOS", name))
			candidates = append(candidates, filepath.Join(bundleRoot, "Contents", "Frameworks", name))
		}
	}

	candidates = append(candidates, name)

	return uniqueStrings(candidates)
}

func darwinBundleRootFromExecutable(exe string) (string, bool) {
	exe = filepath.Clean(exe)
	parts := strings.Split(exe, string(os.PathSeparator))

	for i := len(parts) - 1; i >= 0; i-- {
		if strings.HasSuffix(parts[i], ".app") {
			root := strings.Join(parts[:i+1], string(os.PathSeparator))
			if strings.HasPrefix(exe, string(os.PathSeparator)) {
				root = string(os.PathSeparator) + root
			}
			return root, true
		}
	}

	return "", false
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	out := make([]string, 0, len(values))

	for _, value := range values {
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}

	return out
}

func Get(lib uintptr, name string) uintptr {
	addr, err := purego.Dlsym(lib, name)
	if err != nil {
		panic(err)
	}
	return addr
}
