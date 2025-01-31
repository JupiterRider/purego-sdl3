package sdl

import "unsafe"

type SurfaceFlags uint32

const (
	SurfacePreallocated SurfaceFlags = 1 << iota
	SurfaceLockNeeded
	SurfaceLocked
	SurfaceSimdAligned
)

// Surface is a collection of pixels used in software blitting.
type Surface struct {
	Flags    SurfaceFlags
	Format   PixelFormat
	W        int32
	H        int32
	Pitch    int32
	Pixels   unsafe.Pointer
	Refcount int32
	Reserved unsafe.Pointer
}

// LoadBMPIO loads a BMP image from a seekable SDL data stream.
func LoadBMPIO(src *IOStream, closeio bool) *Surface {
	return sdlLoadBMPIO(src, closeio)
}

// DestroySurface frees a surface.
func DestroySurface(surface *Surface) {
	sdlDestroySurface(surface)
}

// Destroy frees a surface.
func (surface *Surface) Destroy() {
	DestroySurface(surface)
}
