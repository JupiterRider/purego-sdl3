package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

const (
	PropSurfaceSDRWhitePointFloat    = "SDL.surface.SDR_white_point"
	PropSurfaceHDRHeadroomFloat      = "SDL.surface.HDR_headroom"
	PropSurfaceToneMapOperatorString = "SDL.surface.tonemap"
	PropSurfaceHotspotXNumber        = "SDL.surface.hotspot.x"
	PropSurfaceHotspotYNumber        = "SDL.surface.hotspot.y"
)

// [FlipMode] is a structure specifying the flip modes.
//
// [FlipMode]: https://wiki.libsdl.org/SDL3/SDL_FlipMode
type FlipMode uint32

const (
	FlipNone                  FlipMode                        = iota // Do not flip.
	FlipHorizontal                                                   // Flip horizontally.
	FlipVertical                                                     // Flip vertically.
	FlipHorizontalAndVertical = FlipHorizontal | FlipVertical        // Flip horizontally and vertically (not a diagonal flip).
)

// [ScaleMode] is a structure specifying the scaling modes.
//
// [ScaleMode]: https://wiki.libsdl.org/SDL3/SDL_ScaleMode
type ScaleMode int32

const (
	ScaleModeInvalid  ScaleMode = iota - 1
	ScaleModeNearest            // Nearest pixel sampling
	ScaleModeLinear             // Linear filtering
	ScaleModePixelArt           // Nearest pixel sampling with improved scaling for pixel art, available since SDL 3.4.0
)

// [SurfaceFlags] is a structure specifying the flags on an [Surfaces].
//
// [SurfaceFlags]: https://wiki.libsdl.org/SDL3/SDL_SurfaceFlags
type SurfaceFlags uint32

const (
	SurfacePreallocated SurfaceFlags = 1 << iota // Surface uses preallocated pixel memory.
	SurfaceLockNeeded                            // Surface needs to be locked to access pixels.
	SurfaceLocked                                // Surface is currently locked.
	SurfaceSIMDAligned                           // Surface uses pixel memory allocated with [aligned_alloc].
)

// [Surface] is a collection of pixels used in software blitting.
//
// [Surface]: https://wiki.libsdl.org/SDL3/SDL_Surface
type Surface struct {
	Flags    SurfaceFlags   // The flags of the surface, read-only.
	Format   PixelFormat    // The format of the surface, read-only.
	W        int32          // The width of the surface, read-only.
	H        int32          // The height of the surface, read-only.
	Pitch    int32          // The distance in bytes between rows of pixels, read-only.
	Pixels   unsafe.Pointer // A pointer to the pixels of the surface, the pixels are writeable if non-nil.
	Refcount int32          // Application reference count, used when freeing surface.
	Reserved unsafe.Pointer // Reserved for internal use.
}

func MustLock(surface *Surface) bool {
	if surface == nil {
		return InvalidParamError("surface")
	}
	return surface.Flags&SurfaceLockNeeded == SurfaceLockNeeded
}

// LoadBMPIO loads a BMP image from a seekable SDL data stream.
func LoadBMPIO(src *IOStream, closeio bool) *Surface {
	return sdlLoadBMPIO(src, closeio)
}

// [DestroySurface] frees a surface.
//
// [DestroySurface]: https://wiki.libsdl.org/SDL3/SDL_DestroySurface
func DestroySurface(surface *Surface) {
	sdlDestroySurface(surface)
}

// [AddSurfaceAlternateImage] adds an alternate version of a surface.
//
// [AddSurfaceAlternateImage]: https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage
func AddSurfaceAlternateImage(surface, image *Surface) bool {
	return sdlAddSurfaceAlternateImage(surface, image)
}

// [BlitSurface] performs a fast blit from the source surface to the destination surface with clipping.
//
// [BlitSurface]: https://wiki.libsdl.org/SDL3/SDL_BlitSurface
func BlitSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
	return sdlBlitSurface(src, srcrect, dst, dstrect)
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

// [ConvertPixels] copies a block of pixels of one format to another format.
//
// [ConvertPixels]: https://wiki.libsdl.org/SDL3/SDL_ConvertPixels
func ConvertPixels(width int32, height int32, srcFormat PixelFormat, src unsafe.Pointer, srcPitch int32, dstFormat PixelFormat, dst unsafe.Pointer, dstPitch int32) bool {
	return sdlConvertPixels(width, height, srcFormat, src, srcPitch, dstFormat, dst, dstPitch)
}

// [ConvertPixelsAndColorspace] copies a block of pixels of one format and colorspace to another format and colorspace.
//
// [ConvertPixelsAndColorspace]: https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace
func ConvertPixelsAndColorspace(width int32, height int32, srcFormat PixelFormat, srcColorspace Colorspace, srcProperties PropertiesID, src unsafe.Pointer, srcPitch int32, dstFormat PixelFormat, dstColorspace Colorspace, dstProperties PropertiesID, dst unsafe.Pointer, dstPitch int32) bool {
	return sdlConvertPixelsAndColorspace(width, height, srcFormat, srcColorspace, srcProperties, src, srcPitch, dstFormat, dstColorspace, dstProperties, dst, dstPitch)
}

// [ConvertSurface] copies an existing surface to a new surface of the specified format.
//
// [ConvertSurface]: https://wiki.libsdl.org/SDL3/SDL_ConvertSurface
func ConvertSurface(surface *Surface, format PixelFormat) *Surface {
	return sdlConvertSurface(surface, format)
}

// [ConvertSurfaceAndColorspace] copies an existing surface to a new surface of the specified format and colorspace.
//
// [ConvertSurfaceAndColorspace]: https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace
func ConvertSurfaceAndColorspace(surface *Surface, format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) *Surface {
	return sdlConvertSurfaceAndColorspace(surface, format, palette, colorspace, props)
}

// [CreateSurface] allocates a new surface with a specific pixel format.
//
// [CreateSurface]: https://wiki.libsdl.org/SDL3/SDL_CreateSurface
func CreateSurface(width int32, height int32, format PixelFormat) *Surface {
	return sdlCreateSurface(width, height, format)
}

// [CreateSurfaceFrom] allocates a new surface with a specific pixel format and existing pixel data.
//
// [CreateSurfaceFrom]: https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom
func CreateSurfaceFrom(width int32, height int32, format PixelFormat, pixels unsafe.Pointer, pitch int32) *Surface {
	return sdlCreateSurfaceFrom(width, height, format, pixels, pitch)
}

// [CreateSurfacePalette] creates a palette and associate it with a surface.
//
// [CreateSurfacePalette]: https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette
func CreateSurfacePalette(surface *Surface) *Palette {
	return sdlCreateSurfacePalette(surface)
}

// [DuplicateSurface] creates a new surface identical to the existing surface.
//
// [DuplicateSurface]: https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface
func DuplicateSurface(surface *Surface) *Surface {
	return sdlDuplicateSurface(surface)
}

// [FillSurfaceRect] performs a fast fill of a rectangle with a specific color.
//
// [FillSurfaceRect]: https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect
func FillSurfaceRect(dst *Surface, rect *Rect, color uint32) bool {
	return sdlFillSurfaceRect(dst, rect, color)
}

// func FillSurfaceRects(dst *Surface, rects *Rect, count int32, color uint32) bool {
//	return sdlFillSurfaceRects(dst, rects, count, color)
// }

// [FlipSurface] flips a surface vertically or horizontally.
//
// [FlipSurface]: https://wiki.libsdl.org/SDL3/SDL_FlipSurface
func FlipSurface(surface *Surface, flip FlipMode) bool {
	return sdlFlipSurface(surface, flip)
}

// [GetSurfaceAlphaMod] gets the additional alpha value used in blit operations.
//
// [GetSurfaceAlphaMod]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod
func GetSurfaceAlphaMod(surface *Surface, alpha *uint8) bool {
	return sdlGetSurfaceAlphaMod(surface, alpha)
}

// [GetSurfaceBlendMode] gets the blend mode used for blit operations.
//
// [GetSurfaceBlendMode]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode
func GetSurfaceBlendMode(surface *Surface, blendMode *BlendMode) bool {
	return sdlGetSurfaceBlendMode(surface, blendMode)
}

// [GetSurfaceClipRect] gets the clipping rectangle for a surface.
//
// [GetSurfaceClipRect]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect
func GetSurfaceClipRect(surface *Surface, rect *Rect) bool {
	return sdlGetSurfaceClipRect(surface, rect)
}

// [GetSurfaceColorKey] gets the color key (transparent pixel) for a surface.
//
// [GetSurfaceColorKey]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey
func GetSurfaceColorKey(surface *Surface, key *uint32) bool {
	return sdlGetSurfaceColorKey(surface, key)
}

// [GetSurfaceColorMod] gets the additional color value multiplied into blit operations.
//
// [GetSurfaceColorMod]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod
func GetSurfaceColorMod(surface *Surface, r *uint8, g *uint8, b *uint8) bool {
	return sdlGetSurfaceColorMod(surface, r, g, b)
}

// [GetSurfaceColorspace] gets the colorspace used by a surface.
//
// [GetSurfaceColorspace]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace
func GetSurfaceColorspace(surface *Surface) Colorspace {
	return sdlGetSurfaceColorspace(surface)
}

// [GetSurfaceImages] gets an array including all versions of a surface.
//
// [GetSurfaceImages]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages
func GetSurfaceImages(surface *Surface) []*Surface {
	var count int32
	surfaces := sdlGetSurfaceImages(surface, &count)
	defer Free(unsafe.Pointer(surfaces))
	return mem.Copy(surfaces, count)
}

// [GetSurfacePalette] gets the palette used by a surface.
//
// [GetSurfacePalette]: https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette
func GetSurfacePalette(surface *Surface) *Palette {
	return sdlGetSurfacePalette(surface)
}

// [GetSurfaceProperties] gets the properties associated with a surface.
//
// [GetSurfaceProperties]: https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties
func GetSurfaceProperties(surface *Surface) PropertiesID {
	return sdlGetSurfaceProperties(surface)
}

// [LoadBMP] loads a BMP image from a file.
//
// [LoadBMP]: https://wiki.libsdl.org/SDL3/SDL_LoadBMP
func LoadBMP(file string) *Surface {
	return sdlLoadBMP(file)
}

// [LockSurface] sets up a surface for directly accessing the pixels.
//
// [LockSurface]: https://wiki.libsdl.org/SDL3/SDL_LockSurface
func LockSurface(surface *Surface) bool {
	return sdlLockSurface(surface)
}

// [MapSurfaceRGB] maps an RGB triple to an opaque pixel value for a surface.
//
// [MapSurfaceRGB]: https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB
func MapSurfaceRGB(surface *Surface, r uint8, g uint8, b uint8) uint32 {
	return sdlMapSurfaceRGB(surface, r, g, b)
}

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

// [RemoveSurfaceAlternateImages] removes all alternate versions of a surface.
//
// [RemoveSurfaceAlternateImages]: https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages
func RemoveSurfaceAlternateImages(surface *Surface) {
	sdlRemoveSurfaceAlternateImages(surface)
}

// [SaveBMP] saves a surface to a file in BMP format.
//
// [SaveBMP]: https://wiki.libsdl.org/SDL3/SDL_SaveBMP
func SaveBMP(surface *Surface, file string) bool {
	return sdlSaveBMP(surface, file)
}

func SaveBMPIO(surface *Surface, dst *IOStream, closeio bool) bool {
	return sdlSaveBMPIO(surface, dst, closeio)
}

// [ScaleSurface] creates a new surface identical to the existing surface, scaled to the desired size.
//
// [ScaleSurface]: https://wiki.libsdl.org/SDL3/SDL_ScaleSurface
func ScaleSurface(surface *Surface, width int32, height int32, scaleMode ScaleMode) *Surface {
	return sdlScaleSurface(surface, width, height, scaleMode)
}

// [SetSurfaceAlphaMod] sets an additional alpha value used in blit operations.
//
// [SetSurfaceAlphaMod]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod
func SetSurfaceAlphaMod(surface *Surface, alpha uint8) bool {
	return sdlSetSurfaceAlphaMod(surface, alpha)
}

// [SetSurfaceBlendMode] sets the blend mode used for blit operations.
//
// [SetSurfaceBlendMode]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode
func SetSurfaceBlendMode(surface *Surface, blendMode BlendMode) bool {
	return sdlSetSurfaceBlendMode(surface, blendMode)
}

// [SetSurfaceClipRect] sets the clipping rectangle for a surface.
//
// [SetSurfaceClipRect]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect
func SetSurfaceClipRect(surface *Surface, rect *Rect) bool {
	return sdlSetSurfaceClipRect(surface, rect)
}

// [SetSurfaceColorKey] sets the color key (transparent pixel) in a surface.
//
// [SetSurfaceColorKey]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey
func SetSurfaceColorKey(surface *Surface, enabled bool, key uint32) bool {
	return sdlSetSurfaceColorKey(surface, enabled, key)
}

// [SetSurfaceColorMod] sets an additional color value multiplied into blit operations.
//
// [SetSurfaceColorMod]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod
func SetSurfaceColorMod(surface *Surface, r uint8, g uint8, b uint8) bool {
	return sdlSetSurfaceColorMod(surface, r, g, b)
}

// [SetSurfaceColorspace] sets the colorspace used by a surface.
//
// [SetSurfaceColorspace]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace
func SetSurfaceColorspace(surface *Surface, colorspace Colorspace) bool {
	return sdlSetSurfaceColorspace(surface, colorspace)
}

// [SetSurfacePalette] sets the palette used by a surface.
//
// [SetSurfacePalette]: https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette
func SetSurfacePalette(surface *Surface, palette *Palette) bool {
	return sdlSetSurfacePalette(surface, palette)
}

// [SetSurfaceRLE] sets the RLE acceleration hint for a surface.
//
// [SetSurfaceRLE]: https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE
func SetSurfaceRLE(surface *Surface, enabled bool) bool {
	return sdlSetSurfaceRLE(surface, enabled)
}

// [SurfaceHasAlternateImages] returns whether a surface has alternate versions available.
//
// [SurfaceHasAlternateImages]: https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages
func SurfaceHasAlternateImages(surface *Surface) bool {
	return sdlSurfaceHasAlternateImages(surface)
}

// [SurfaceHasColorKey] returns whether the surface has a color key.
//
// [SurfaceHasColorKey]: https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey
func SurfaceHasColorKey(surface *Surface) bool {
	return sdlSurfaceHasColorKey(surface)
}

// [SurfaceHasRLE] returns whether the surface is RLE enabled.
//
// [SurfaceHasRLE]: https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE
func SurfaceHasRLE(surface *Surface) bool {
	return sdlSurfaceHasRLE(surface)
}

// [UnlockSurface] releases a surface after directly accessing the pixels.
//
// [UnlockSurface]: https://wiki.libsdl.org/SDL3/SDL_UnlockSurface
func UnlockSurface(surface *Surface) {
	sdlUnlockSurface(surface)
}

// func WriteSurfacePixel(surface *Surface, x int32, y int32, r uint8, g uint8, b uint8, a uint8) bool {
//	return sdlWriteSurfacePixel(surface, x, y, r, g, b, a)
// }

// func WriteSurfacePixelFloat(surface *Surface, x int32, y int32, r float32, g float32, b float32, a float32) bool {
//	return sdlWriteSurfacePixelFloat(surface, x, y, r, g, b, a)
// }
