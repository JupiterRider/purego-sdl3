package img

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

// [Animation] defines the animated image support.
//
// [Animation]: https://wiki.libsdl.org/SDL3_image/IMG_Animation
type Animation struct {
	W      int32         // The width of the frames.
	H      int32         // The height of the frames.
	Count  int32         // The number of frames.
	frames **sdl.Surface // An array of frames.
	delays *int32        // An array of frame delays, in milliseconds.
}

// Frames gets available frames.
func (a *Animation) Frames() []*sdl.Surface {
	return unsafe.Slice(a.frames, a.Count)
}

// Delays gets delays between frames.
func (a *Animation) Delays() []int32 {
	return unsafe.Slice(a.delays, a.Count)
}

// [Version] gets the version of the dynamically linked SDL_image library.
//
// [Version]: https://wiki.libsdl.org/SDL3_image/IMG_Version
func Version() (major, minor, patch int32) {
	version := imgVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}

// [FreeAnimation] disposes of an [Animation] and free its resources.
//
// [FreeAnimation]: https://wiki.libsdl.org/SDL3_image/IMG_FreeAnimation
func FreeAnimation(anim *Animation) {
	imgFreeAnimation(anim)
}

// [IsAVIF] detects AVIF image data on a readable/seekable [IOStream].
//
// [IsAVIF]: https://wiki.libsdl.org/SDL3_image/IMG_isAVIF
func IsAVIF(src *sdl.IOStream) bool {
	return imgIsAVIF(src)
}

// [IsBMP] detects BMP image data on a readable/seekable [IOStream].
//
// [IsBMP]: https://wiki.libsdl.org/SDL3_image/IMG_isBMP
func IsBMP(src *sdl.IOStream) bool {
	return imgIsBMP(src)
}

// [IsCUR] detects CUR image data on a readable/seekable [IOStream].
//
// [IsCUR]: https://wiki.libsdl.org/SDL3_image/IMG_isCUR
func IsCUR(src *sdl.IOStream) bool {
	return imgIsCUR(src)
}

// [IsGIF] detects GIF image data on a readable/seekable [IOStream].
//
// [IsGIF]: https://wiki.libsdl.org/SDL3_image/IMG_isGIF
func IsGIF(src *sdl.IOStream) bool {
	return imgIsGIF(src)
}

// [IsICO] detects ICO image data on a readable/seekable [IOStream].
//
// [IsICO]: https://wiki.libsdl.org/SDL3_image/IMG_isICO
func IsICO(src *sdl.IOStream) bool {
	return imgIsICO(src)
}

// [IsJPG] detects JPG image data on a readable/seekable [IOStream].
//
// [IsJPG]: https://wiki.libsdl.org/SDL3_image/IMG_isJPG
func IsJPG(src *sdl.IOStream) bool {
	return imgIsJPG(src)
}

// [IsJXL] detects JXL image data on a readable/seekable [IOStream].
//
// [IsJXL]: https://wiki.libsdl.org/SDL3_image/IMG_isJXL
func IsJXL(src *sdl.IOStream) bool {
	return imgIsJXL(src)
}

// [IsLBM] detects LBM image data on a readable/seekable [IOStream].
//
// [IsLBM]: https://wiki.libsdl.org/SDL3_image/IMG_isLBM
func IsLBM(src *sdl.IOStream) bool {
	return imgIsLBM(src)
}

// [IsPCX] detects PCX image data on a readable/seekable [IOStream].
//
// [IsPCX]: https://wiki.libsdl.org/SDL3_image/IMG_isPCX
func IsPCX(src *sdl.IOStream) bool {
	return imgIsPCX(src)
}

// [IsPNG] detects PNG image data on a readable/seekable [IOStream].
//
// [IsPNG]: https://wiki.libsdl.org/SDL3_image/IMG_isPNG
func IsPNG(src *sdl.IOStream) bool {
	return imgIsPNG(src)
}

// [IsPNM] detects PNM image data on a readable/seekable [IOStream].
//
// [IsPNM]: https://wiki.libsdl.org/SDL3_image/IMG_isPNM
func IsPNM(src *sdl.IOStream) bool {
	return imgIsPNM(src)
}

// [IsQOI] detects QOI image data on a readable/seekable [IOStream].
//
// [IsQOI]: https://wiki.libsdl.org/SDL3_image/IMG_isQOI
func IsQOI(src *sdl.IOStream) bool {
	return imgIsQOI(src)
}

// [IsSVG] detects SVG image data on a readable/seekable [IOStream].
//
// [IsSVG]: https://wiki.libsdl.org/SDL3_image/IMG_isSVG
func IsSVG(src *sdl.IOStream) bool {
	return imgIsSVG(src)
}

// [IsTIF] detects TIFF image data on a readable/seekable [IOStream].
//
// [IsTIF]: https://wiki.libsdl.org/SDL3_image/IMG_isTIF
func IsTIF(src *sdl.IOStream) bool {
	return imgIsTIF(src)
}

// [IsWEBP] detects WEBP image data on a readable/seekable [IOStream].
//
// [IsWEBP]: https://wiki.libsdl.org/SDL3_image/IMG_isWEBP
func IsWEBP(src *sdl.IOStream) bool {
	return imgIsWEBP(src)
}

// [IsXCF] detects XCF image data on a readable/seekable [IOStream].
//
// [IsXCF]: https://wiki.libsdl.org/SDL3_image/IMG_isXCF
func IsXCF(src *sdl.IOStream) bool {
	return imgIsXCF(src)
}

// [IsXPM] detects XPM image data on a readable/seekable [IOStream].
//
// [IsXPM]: https://wiki.libsdl.org/SDL3_image/IMG_isXPM
func IsXPM(src *sdl.IOStream) bool {
	return imgIsXPM(src)
}

// [IsXV] detects XV image data on a readable/seekable [IOStream].
//
// [IsXV]: https://wiki.libsdl.org/SDL3_image/IMG_isXV
func IsXV(src *sdl.IOStream) bool {
	return imgIsXV(src)
}

// [Load] loads an image from a filesystem path into a software surface.
//
// [Load]: https://wiki.libsdl.org/SDL3_image/IMG_Load
func Load(file string) *sdl.Surface {
	return imgLoad(file)
}

// [Load] loads an image from an SDL data source into a software surface.
//
// [Load]: https://wiki.libsdl.org/SDL3_image/IMG_Load_IO
func LoadIO(src *sdl.IOStream, closeio bool) *sdl.Surface {
	return imgLoadIO(src, closeio)
}

// [LoadAnimation] loads an animation from a file.
//
// [LoadAnimation]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation
func LoadAnimation(file string) *Animation {
	return imgLoadAnimation(file)
}

// [LoadAnimation] loads an animation from an [IOStream].
//
// [LoadAnimation]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO
func LoadAnimationIO(src *sdl.IOStream, closeio bool) *Animation {
	return imgLoadAnimationIO(src, closeio)
}

// [LoadAnimationTypedIO] loads an animation from an [IOStream].
//
// [LoadAnimationTypedIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimationTyped_IO
func LoadAnimationTypedIO(src *sdl.IOStream, closeio bool, format string) *Animation {
	return imgLoadAnimationTypedIO(src, closeio, format)
}

// [LoadAVIFIO] loads a AVIF image directly.
//
// [LoadAVIFIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIF_IO
func LoadAVIFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadAVIFIO(src)
}

// [LoadBMPIO] loads a BMP image directly.
//
// [LoadBMPIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadBMP_IO
func LoadBMPIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadBMPIO(src)
}

// [LoadCURIO] loads a CUR image directly.
//
// [LoadCURIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadCUR_IO
func LoadCURIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadCURIO(src)
}

// [LoadGIFIO] loads a GIF image directly.
//
// [LoadGIFIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadGIF_IO
func LoadGIFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadGIFIO(src)
}

// [LoadGIFAnimationIO] loads a GIF animation directly.
//
// [LoadGIFAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadGIFAnimation_IO
func LoadGIFAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadGIFAnimationIO(src)
}

// [LoadICOIO] loads a ICO image directly.
//
// [LoadICOIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadICO_IO
func LoadICOIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadICOIO(src)
}

// [LoadJPGIO] loads a JPG image directly.
//
// [LoadJPGIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadJPG_IO
func LoadJPGIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadJPGIO(src)
}

// [LoadJXLIO] loads a JXL image directly.
//
// [LoadJXLIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadJXL_IO
func LoadJXLIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadJXLIO(src)
}

// [LoadLBMIO] loads a LBM image directly.
//
// [LoadLBMIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadLBM_IO
func LoadLBMIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadLBMIO(src)
}

// [LoadPCXIO] loads a PCX image directly.
//
// [LoadPCXIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadPCX_IO
func LoadPCXIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadPCXIO(src)
}

// [LoadPNGIO] loads a PNG image directly.
//
// [LoadPNGIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadPNG_IO
func LoadPNGIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadPNGIO(src)
}

// [LoadPNMIO] loads a PNM image directly.
//
// [LoadPNMIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadPNM_IO
func LoadPNMIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadPNMIO(src)
}

// [LoadQOIIO] loads a QOI image directly.
//
// [LoadQOIIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadQOI_IO
func LoadQOIIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadQOIIO(src)
}

// [LoadSizedSVGIO] loads an SVG image, scaled to a specific size.
//
// [LoadSizedSVGIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadSizedSVG_IO
func LoadSizedSVGIO(src *sdl.IOStream, width int32, height int32) *sdl.Surface {
	return imgLoadSizedSVGIO(src, width, height)
}

// [LoadSVGIO] loads a SVG image directly.
//
// [LoadSVGIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadSVG_IO
func LoadSVGIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadSVGIO(src)
}

// [LoadTexture] loads an image from a filesystem path into a GPU texture.
//
// [LoadTexture]: https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture
func LoadTexture(renderer *sdl.Renderer, file string) *sdl.Texture {
	return imgLoadTexture(renderer, file)
}

// [LoadTextureIO] loads an image from an SDL data source into a GPU texture.
//
// [LoadTextureIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadTexture_IO
func LoadTextureIO(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
	return imgLoadTextureIO(renderer, src, closeio)
}

// [LoadTextureTypedIO] loads an image from an SDL data source into a GPU texture.
//
// [LoadTextureTypedIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadTextureTyped_IO
func LoadTextureTypedIO(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, format string) *sdl.Texture {
	return imgLoadTextureTypedIO(renderer, src, closeio, format)
}

// [LoadTGAIO] loads a TGA image directly.
//
// [LoadTGAIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadTGA_IO
func LoadTGAIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadTGAIO(src)
}

// [LoadTIFIO] loads a TIFF image directly.
//
// [LoadTIFIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadTIF_IO
func LoadTIFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadTIFIO(src)
}

// [LoadTypedIO] loads an image from an SDL data source into a software surface.
//
// [LoadTypedIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadTyped_IO
func LoadTypedIO(src *sdl.IOStream, closeio bool, format string) *sdl.Surface {
	return imgLoadTypedIO(src, closeio, format)
}

// [LoadWEBPIO] loads a WEBP image directly.
//
// [LoadWEBPIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBP_IO
func LoadWEBPIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadWEBPIO(src)
}

// [LoadWEBPAnimationIO] loads a WEBP animation directly.
//
// [LoadWEBPAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadWEBPAnimation_IO
func LoadWEBPAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadWEBPAnimationIO(src)
}

// [LoadXCFIO] loads a XCF image directly.
//
// [LoadXCFIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadXCF_IO
func LoadXCFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadXCFIO(src)
}

// [LoadXPMIO] loads a XPM image directly.
//
// [LoadXPMIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadXPM_IO
func LoadXPMIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadXPMIO(src)
}

// [LoadXVIO] loads a XV image directly.
//
// [LoadXVIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadXV_IO
func LoadXVIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadXVIO(src)
}

// [ReadXPMFromArray] loads an XPM image from a memory array.
//
// [ReadXPMFromArray]: https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray
func ReadXPMFromArray(xpm **byte) *sdl.Surface {
	return imgReadXPMFromArray(xpm)
}

// [ReadXPMFromArrayToRGB888] loads an XPM image from a memory array.
//
// [ReadXPMFromArrayToRGB888]: https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888
func ReadXPMFromArrayToRGB888(xpm **byte) *sdl.Surface {
	return imgReadXPMFromArrayToRGB888(xpm)
}

// [SaveAVIF] saves an [Surface] into a AVIF image file.
//
// [SaveAVIF]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF
func SaveAVIF(surface *sdl.Surface, file string, quality int32) bool {
	return imgSaveAVIF(surface, file, quality)
}

// [SaveAVIFIO] saves an [Surface] into AVIF image data, via an SDL_IOStream.
//
// [SaveAVIFIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO
func SaveAVIFIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveAVIFIO(surface, dst, closeio, quality)
}

// [SaveJPG] saves an [Surface] into a JPEG image file.
//
// [SaveJPG]: https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG
func SaveJPG(surface *sdl.Surface, file string, quality int32) bool {
	return imgSaveJPG(surface, file, quality)
}

// [SaveJPGIO] saves an [Surface] into JPEG image data, via an SDL_IOStream.
//
// [SaveJPGIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO
func SaveJPGIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveJPGIO(surface, dst, closeio, quality)
}

// [SavePNG] saves an [Surface] into a PNG image file.
//
// [SavePNG]: https://wiki.libsdl.org/SDL3_image/IMG_SavePNG
func SavePNG(surface *sdl.Surface, file string) bool {
	return imgSavePNG(surface, file)
}

// [SavePNGIO] saves an [Surface] into PNG image data, via an SDL_IOStream.
//
// [SavePNGIO]: https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO
func SavePNGIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSavePNGIO(surface, dst, closeio)
}
