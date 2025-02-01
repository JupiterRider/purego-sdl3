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

// func AddSurfaceAlternateImage(surface *Surface, image *Surface) bool {
//	return sdlAddSurfaceAlternateImage(surface, image)
// }

// BlitSurface performs a fast blit from the source surface to the destination surface with clipping.
func BlitSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
	return sdlBlitSurface(src, srcrect, dst, dstrect)
}

// Blit performs a fast blit from the source surface to the destination surface with clipping.
func (src *Surface) Blit(srcrect *Rect, dst *Surface, dstrect *Rect) bool {
	return BlitSurface(src, srcrect, dst, dstrect)
}

// func BlitSurface9Grid(src *Surface, srcrect *Rect, left_width int32, right_width int32, top_height int32, bottom_height int32, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) bool {
//	return sdlBlitSurface9Grid(src, srcrect, left_width, right_width, top_height, bottom_height, scale, scaleMode, dst, dstrect)
// }

// func BlitSurfaceScaled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) bool {
//	return sdlBlitSurfaceScaled(src, srcrect, dst, dstrect, scaleMode)
// }

// func BlitSurfaceTiled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
//	return sdlBlitSurfaceTiled(src, srcrect, dst, dstrect)
// }

// func BlitSurfaceTiledWithScale(src *Surface, srcrect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) bool {
//	return sdlBlitSurfaceTiledWithScale(src, srcrect, scale, scaleMode, dst, dstrect)
// }

// func BlitSurfaceUnchecked(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
//	return sdlBlitSurfaceUnchecked(src, srcrect, dst, dstrect)
// }

// func BlitSurfaceUncheckedScaled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) bool {
//	return sdlBlitSurfaceUncheckedScaled(src, srcrect, dst, dstrect, scaleMode)
// }

// func ClearSurface(surface *Surface, r float32, g float32, b float32, a float32) bool {
//	return sdlClearSurface(surface, r, g, b, a)
// }

// func ConvertPixels(width int32, height int32, src_format PixelFormat, src unsafe.Pointer, src_pitch int32, dst_format PixelFormat, dst unsafe.Pointer, dst_pitch int32) bool {
//	return sdlConvertPixels(width, height, src_format, src, src_pitch, dst_format, dst, dst_pitch)
// }

// func ConvertPixelsAndColorspace(width int32, height int32, src_format PixelFormat, src_colorspace Colorspace, src_properties PropertiesID, src unsafe.Pointer, src_pitch int32, dst_format PixelFormat, dst_colorspace Colorspace, dst_properties PropertiesID, dst unsafe.Pointer, dst_pitch int32) bool {
//	return sdlConvertPixelsAndColorspace(width, height, src_format, src_colorspace, src_properties, src, src_pitch, dst_format, dst_colorspace, dst_properties, dst, dst_pitch)
// }

// func ConvertSurface(surface *Surface, format PixelFormat) *Surface {
//	return sdlConvertSurface(surface, format)
// }

// func ConvertSurfaceAndColorspace(surface *Surface, format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) *Surface {
//	return sdlConvertSurfaceAndColorspace(surface, format, palette, colorspace, props)
// }

// CreateSurface allocates a new surface with a specific pixel format.
func CreateSurface(width int32, height int32, format PixelFormat) *Surface {
	return sdlCreateSurface(width, height, format)
}

// func CreateSurfaceFrom(width int32, height int32, format PixelFormat, pixels unsafe.Pointer, pitch int32) *Surface {
//	return sdlCreateSurfaceFrom(width, height, format, pixels, pitch)
// }

// func CreateSurfacePalette(surface *Surface) *Palette {
//	return sdlCreateSurfacePalette(surface)
// }

// func DuplicateSurface(surface *Surface) *Surface {
//	return sdlDuplicateSurface(surface)
// }

// func FillSurfaceRect(dst *Surface, rect *Rect, color uint32) bool {
//	return sdlFillSurfaceRect(dst, rect, color)
// }

// func FillSurfaceRects(dst *Surface, rects *Rect, count int32, color uint32) bool {
//	return sdlFillSurfaceRects(dst, rects, count, color)
// }

// func FlipSurface(surface *Surface, flip FlipMode) bool {
//	return sdlFlipSurface(surface, flip)
// }

// func GetSurfaceAlphaMod(surface *Surface, alpha *uint8) bool {
//	return sdlGetSurfaceAlphaMod(surface, alpha)
// }

// func GetSurfaceBlendMode(surface *Surface, blendMode *BlendMode) bool {
//	return sdlGetSurfaceBlendMode(surface, blendMode)
// }

// func GetSurfaceClipRect(surface *Surface, rect *Rect) bool {
//	return sdlGetSurfaceClipRect(surface, rect)
// }

// func GetSurfaceColorKey(surface *Surface, key *uint32) bool {
//	return sdlGetSurfaceColorKey(surface, key)
// }

// func GetSurfaceColorMod(surface *Surface, r *uint8, g *uint8, b *uint8) bool {
//	return sdlGetSurfaceColorMod(surface, r, g, b)
// }

// func GetSurfaceColorspace(surface *Surface) Colorspace {
//	return sdlGetSurfaceColorspace(surface)
// }

// func GetSurfaceImages(surface *Surface, count *int32) **Surface {
//	return sdlGetSurfaceImages(surface, count)
// }

// func GetSurfacePalette(surface *Surface) *Palette {
//	return sdlGetSurfacePalette(surface)
// }

// func GetSurfaceProperties(surface *Surface) PropertiesID {
//	return sdlGetSurfaceProperties(surface)
// }

// func LoadBMP(file string) *Surface {
//	return sdlLoadBMP(file)
// }

// LockSurface sets up a surface for directly accessing the pixels.
func LockSurface(surface *Surface) bool {
	return sdlLockSurface(surface)
}

// Lock sets up a surface for directly accessing the pixels.
func (surface *Surface) Lock() bool {
	return LockSurface(surface)
}

// func MapSurfaceRGB(surface *Surface, r uint8, g uint8, b uint8) uint32 {
//	return sdlMapSurfaceRGB(surface, r, g, b)
// }

// func MapSurfaceRGBA(surface *Surface, r uint8, g uint8, b uint8, a uint8) uint32 {
//	return sdlMapSurfaceRGBA(surface, r, g, b, a)
// }

// func PremultiplyAlpha(width int32, height int32, src_format PixelFormat, src unsafe.Pointer, src_pitch int32, dst_format PixelFormat, dst unsafe.Pointer, dst_pitch int32, linear bool) bool {
//	return sdlPremultiplyAlpha(width, height, src_format, src, src_pitch, dst_format, dst, dst_pitch, linear)
// }

// func PremultiplySurfaceAlpha(surface *Surface, linear bool) bool {
//	return sdlPremultiplySurfaceAlpha(surface, linear)
// }

// func ReadSurfacePixel(surface *Surface, x int32, y int32, r *uint8, g *uint8, b *uint8, a *uint8) bool {
//	return sdlReadSurfacePixel(surface, x, y, r, g, b, a)
// }

// func ReadSurfacePixelFloat(surface *Surface, x int32, y int32, r *float32, g *float32, b *float32, a *float32) bool {
//	return sdlReadSurfacePixelFloat(surface, x, y, r, g, b, a)
// }

// func RemoveSurfaceAlternateImages(surface *Surface)  {
//	sdlRemoveSurfaceAlternateImages(surface)
// }

// func SaveBMP(surface *Surface, file string) bool {
//	return sdlSaveBMP(surface, file)
// }

// func SaveBMP_IO(surface *Surface, dst *IOStream, closeio bool) bool {
//	return sdlSaveBMP_IO(surface, dst, closeio)
// }

// func ScaleSurface(surface *Surface, width int32, height int32, scaleMode ScaleMode) *Surface {
//	return sdlScaleSurface(surface, width, height, scaleMode)
// }

// func SetSurfaceAlphaMod(surface *Surface, alpha uint8) bool {
//	return sdlSetSurfaceAlphaMod(surface, alpha)
// }

// func SetSurfaceBlendMode(surface *Surface, blendMode BlendMode) bool {
//	return sdlSetSurfaceBlendMode(surface, blendMode)
// }

// func SetSurfaceClipRect(surface *Surface, rect *Rect) bool {
//	return sdlSetSurfaceClipRect(surface, rect)
// }

// func SetSurfaceColorKey(surface *Surface, enabled bool, key uint32) bool {
//	return sdlSetSurfaceColorKey(surface, enabled, key)
// }

// func SetSurfaceColorMod(surface *Surface, r uint8, g uint8, b uint8) bool {
//	return sdlSetSurfaceColorMod(surface, r, g, b)
// }

// func SetSurfaceColorspace(surface *Surface, colorspace Colorspace) bool {
//	return sdlSetSurfaceColorspace(surface, colorspace)
// }

// func SetSurfacePalette(surface *Surface, palette *Palette) bool {
//	return sdlSetSurfacePalette(surface, palette)
// }

// func SetSurfaceRLE(surface *Surface, enabled bool) bool {
//	return sdlSetSurfaceRLE(surface, enabled)
// }

// func SurfaceHasAlternateImages(surface *Surface) bool {
//	return sdlSurfaceHasAlternateImages(surface)
// }

// func SurfaceHasColorKey(surface *Surface) bool {
//	return sdlSurfaceHasColorKey(surface)
// }

// func SurfaceHasRLE(surface *Surface) bool {
//	return sdlSurfaceHasRLE(surface)
// }

// UnlockSurface releases a surface after directly accessing the pixels.
func UnlockSurface(surface *Surface) {
	sdlUnlockSurface(surface)
}

// Unlock releases a surface after directly accessing the pixels.
func (surface *Surface) Unlock() {
	UnlockSurface(surface)
}

// func WriteSurfacePixel(surface *Surface, x int32, y int32, r uint8, g uint8, b uint8, a uint8) bool {
//	return sdlWriteSurfacePixel(surface, x, y, r, g, b, a)
// }

// func WriteSurfacePixelFloat(surface *Surface, x int32, y int32, r float32, g float32, b float32, a float32) bool {
//	return sdlWriteSurfacePixelFloat(surface, x, y, r, g, b, a)
// }
