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

// [IsAVIF] detects AVIF image data on a readable/seekable [sdl.IOStream].
//
// [IsAVIF]: https://wiki.libsdl.org/SDL3_image/IMG_isAVIF
func IsAVIF(src *sdl.IOStream) bool {
	return imgIsAVIF(src)
}

// [IsBMP] detects BMP image data on a readable/seekable [sdl.IOStream].
//
// [IsBMP]: https://wiki.libsdl.org/SDL3_image/IMG_isBMP
func IsBMP(src *sdl.IOStream) bool {
	return imgIsBMP(src)
}

// [IsCUR] detects CUR image data on a readable/seekable [sdl.IOStream].
//
// [IsCUR]: https://wiki.libsdl.org/SDL3_image/IMG_isCUR
func IsCUR(src *sdl.IOStream) bool {
	return imgIsCUR(src)
}

// [IsGIF] detects GIF image data on a readable/seekable [sdl.IOStream].
//
// [IsGIF]: https://wiki.libsdl.org/SDL3_image/IMG_isGIF
func IsGIF(src *sdl.IOStream) bool {
	return imgIsGIF(src)
}

// [IsICO] detects ICO image data on a readable/seekable [sdl.IOStream].
//
// [IsICO]: https://wiki.libsdl.org/SDL3_image/IMG_isICO
func IsICO(src *sdl.IOStream) bool {
	return imgIsICO(src)
}

// [IsJPG] detects JPG image data on a readable/seekable [sdl.IOStream].
//
// [IsJPG]: https://wiki.libsdl.org/SDL3_image/IMG_isJPG
func IsJPG(src *sdl.IOStream) bool {
	return imgIsJPG(src)
}

// [IsJXL] detects JXL image data on a readable/seekable [sdl.IOStream].
//
// [IsJXL]: https://wiki.libsdl.org/SDL3_image/IMG_isJXL
func IsJXL(src *sdl.IOStream) bool {
	return imgIsJXL(src)
}

// [IsLBM] detects LBM image data on a readable/seekable [sdl.IOStream].
//
// [IsLBM]: https://wiki.libsdl.org/SDL3_image/IMG_isLBM
func IsLBM(src *sdl.IOStream) bool {
	return imgIsLBM(src)
}

// [IsPCX] detects PCX image data on a readable/seekable [sdl.IOStream].
//
// [IsPCX]: https://wiki.libsdl.org/SDL3_image/IMG_isPCX
func IsPCX(src *sdl.IOStream) bool {
	return imgIsPCX(src)
}

// [IsPNG] detects PNG image data on a readable/seekable [sdl.IOStream].
//
// [IsPNG]: https://wiki.libsdl.org/SDL3_image/IMG_isPNG
func IsPNG(src *sdl.IOStream) bool {
	return imgIsPNG(src)
}

// [IsPNM] detects PNM image data on a readable/seekable [sdl.IOStream].
//
// [IsPNM]: https://wiki.libsdl.org/SDL3_image/IMG_isPNM
func IsPNM(src *sdl.IOStream) bool {
	return imgIsPNM(src)
}

// [IsQOI] detects QOI image data on a readable/seekable [sdl.IOStream].
//
// [IsQOI]: https://wiki.libsdl.org/SDL3_image/IMG_isQOI
func IsQOI(src *sdl.IOStream) bool {
	return imgIsQOI(src)
}

// [IsSVG] detects SVG image data on a readable/seekable [sdl.IOStream].
//
// [IsSVG]: https://wiki.libsdl.org/SDL3_image/IMG_isSVG
func IsSVG(src *sdl.IOStream) bool {
	return imgIsSVG(src)
}

// [IsTIF] detects TIFF image data on a readable/seekable [sdl.IOStream].
//
// [IsTIF]: https://wiki.libsdl.org/SDL3_image/IMG_isTIF
func IsTIF(src *sdl.IOStream) bool {
	return imgIsTIF(src)
}

// [IsWEBP] detects WEBP image data on a readable/seekable [sdl.IOStream].
//
// [IsWEBP]: https://wiki.libsdl.org/SDL3_image/IMG_isWEBP
func IsWEBP(src *sdl.IOStream) bool {
	return imgIsWEBP(src)
}

// [IsXCF] detects XCF image data on a readable/seekable [sdl.IOStream].
//
// [IsXCF]: https://wiki.libsdl.org/SDL3_image/IMG_isXCF
func IsXCF(src *sdl.IOStream) bool {
	return imgIsXCF(src)
}

// [IsXPM] detects XPM image data on a readable/seekable [sdl.IOStream].
//
// [IsXPM]: https://wiki.libsdl.org/SDL3_image/IMG_isXPM
func IsXPM(src *sdl.IOStream) bool {
	return imgIsXPM(src)
}

// [IsXV] detects XV image data on a readable/seekable [sdl.IOStream].
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

// [LoadAnimation] loads an animation from an [sdl.IOStream].
//
// [LoadAnimation]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAnimation_IO
func LoadAnimationIO(src *sdl.IOStream, closeio bool) *Animation {
	return imgLoadAnimationIO(src, closeio)
}

// [LoadAnimationTypedIO] loads an animation from an [sdl.IOStream].
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

// [SaveAVIF] saves an [sdl.Surface] into a AVIF image file.
//
// [SaveAVIF]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF
func SaveAVIF(surface *sdl.Surface, file string, quality int32) bool {
	return imgSaveAVIF(surface, file, quality)
}

// [SaveAVIFIO] saves an [sdl.Surface] into AVIF image data, via an [sdl.IOStream].
//
// [SaveAVIFIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIF_IO
func SaveAVIFIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveAVIFIO(surface, dst, closeio, quality)
}

// [SaveJPG] saves an [sdl.Surface] into a JPEG image file.
//
// [SaveJPG]: https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG
func SaveJPG(surface *sdl.Surface, file string, quality int32) bool {
	return imgSaveJPG(surface, file, quality)
}

// [SaveJPGIO] saves an [sdl.Surface] into JPEG image data, via an [sdl.IOStream].
//
// [SaveJPGIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveJPG_IO
func SaveJPGIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveJPGIO(surface, dst, closeio, quality)
}

// [SavePNG] saves an [sdl.Surface] into a PNG image file.
//
// [SavePNG]: https://wiki.libsdl.org/SDL3_image/IMG_SavePNG
func SavePNG(surface *sdl.Surface, file string) bool {
	return imgSavePNG(surface, file)
}

// [SavePNGIO] saves an [sdl.Surface] into PNG image data, via an [sdl.IOStream].
//
// [SavePNGIO]: https://wiki.libsdl.org/SDL3_image/IMG_SavePNG_IO
func SavePNGIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSavePNGIO(surface, dst, closeio)
}

// [LoadGPUTexture] loads an image from a filesystem path into a GPU texture.
//
// Available since SDL_image 3.4.0.
//
// [LoadGPUTexture]: https://wiki.libsdl.org/SDL3_image/IMG_LoadGPUTexture
func LoadGPUTexture(device *sdl.GPUDevice, copyPass *sdl.GPUCopyPass, file string, width, height *int32) *sdl.GPUTexture {
	return imgLoadGPUTexture(device, copyPass, file, width, height)
}

// [LoadGPUTextureIO] loads an image from an SDL data source into a GPU texture.
//
// Available since SDL_image 3.4.0.
//
// [LoadGPUTextureIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadGPUTexture_IO
func LoadGPUTextureIO(device *sdl.GPUDevice, copyPass *sdl.GPUCopyPass, src *sdl.IOStream, closeio bool, width, height *int32) *sdl.GPUTexture {
	return imgLoadGPUTextureIO(device, copyPass, src, closeio, width, height)
}

// [LoadGPUTextureTypedIO] loads an image from an SDL data source into a GPU texture.
//
// Available since SDL_image 3.4.0.
//
// [LoadGPUTextureTypedIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadGPUTextureTyped_IO
func LoadGPUTextureTypedIO(device *sdl.GPUDevice, copyPass *sdl.GPUCopyPass, src *sdl.IOStream, closeio bool, _type string, width, height *int32) *sdl.GPUTexture {
	return imgLoadGPUTextureTypedIO(device, copyPass, src, closeio, _type, width, height)
}

// [GetClipboardImage] gets the image currently in the clipboard.
//
// Available since SDL_image 3.4.0.
//
// [GetClipboardImage]: https://wiki.libsdl.org/SDL3_image/IMG_GetClipboardImage
func GetClipboardImage() *sdl.Surface {
	return imgGetClipboardImage()
}

// [IsANI] detects ANI animated cursor data on a readable/seekable [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [IsANI]: https://wiki.libsdl.org/SDL3_image/IMG_IsANI
func IsANI(src *sdl.IOStream) bool {
	return imgIsANI(src)
}

// [Save] saves an [sdl.Surface] into an image file.
//
// Available since SDL_image 3.4.0.
//
// [Save]: https://wiki.libsdl.org/SDL3_image/IMG_Save
func Save(surface *sdl.Surface, file string) bool {
	return imgSave(surface, file)
}

// [SaveTypedIO] saves an [sdl.Surface] into formatted image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveTypedIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveTyped_IO
func SaveTypedIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, _type string) bool {
	return imgSaveTypedIO(surface, dst, closeio, _type)
}

// [SaveBMP] saves an [sdl.Surface] into a BMP image file.
//
// Available since SDL_image 3.4.0.
//
// [SaveBMP]: https://wiki.libsdl.org/SDL3_image/IMG_SaveBMP
func SaveBMP(surface *sdl.Surface, file string) bool {
	return imgSaveBMP(surface, file)
}

// [SaveBMPIO] saves an [sdl.Surface] into BMP image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveBMPIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveBMP_IO
func SaveBMPIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveBMPIO(surface, dst, closeio)
}

// [SaveCUR] saves an [sdl.Surface] into a CUR image file.
//
// Available since SDL_image 3.4.0.
//
// [SaveCUR]: https://wiki.libsdl.org/SDL3_image/IMG_SaveCUR
func SaveCUR(surface *sdl.Surface, file string) bool {
	return imgSaveCUR(surface, file)
}

// [SaveCURIO] saves an [sdl.Surface] into CUR image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveCURIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveCUR_IO
func SaveCURIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveCURIO(surface, dst, closeio)
}

// [SaveGIF] saves an [sdl.Surface] into a GIF image file.
//
// Available since SDL_image 3.4.0.
//
// [SaveGIF]: https://wiki.libsdl.org/SDL3_image/IMG_SaveGIF
func SaveGIF(surface *sdl.Surface, file string) bool {
	return imgSaveGIF(surface, file)
}

// [SaveGIFIO] saves an [sdl.Surface] into GIF image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveGIFIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveGIF_IO
func SaveGIFIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveGIFIO(surface, dst, closeio)
}

// [SaveICO] saves an [sdl.Surface] into a ICO image file.
//
// Available since SDL_image 3.4.0.
//
// [SaveICO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveICO
func SaveICO(surface *sdl.Surface, file string) bool {
	return imgSaveICO(surface, file)
}

// [SaveICOIO] saves an [sdl.Surface] into ICO image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveICOIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveICO_IO
func SaveICOIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveICOIO(surface, dst, closeio)
}

// [SaveTGA] saves an [sdl.Surface] into a TGA image file.
//
// Available since SDL_image 3.4.0.
//
// [SaveTGA]: https://wiki.libsdl.org/SDL3_image/IMG_SaveTGA
func SaveTGA(surface *sdl.Surface, file string) bool {
	return imgSaveTGA(surface, file)
}

// [SaveTGAIO] saves an [sdl.Surface] into TGA image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveTGAIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveTGA_IO
func SaveTGAIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveTGAIO(surface, dst, closeio)
}

// [SaveWEBP] saves an [sdl.Surface] into a WEBP image file.
//
// Available since SDL_image 3.4.0.
//
// [SaveWEBP]: https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBP
func SaveWEBP(surface *sdl.Surface, file string, quality float32) bool {
	return imgSaveWEBP(surface, file, quality)
}

// [SaveWEBPIO] saves an [sdl.Surface] into WEBP image data, via an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveWEBPIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBP_IO
func SaveWEBPIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality float32) bool {
	return imgSaveWEBPIO(surface, dst, closeio, quality)
}

// [LoadANIAnimationIO] loads an ANI animation directly from an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [LoadANIAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadANIAnimation_IO
func LoadANIAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadANIAnimationIO(src)
}

// [LoadAPNGAnimationIO] loads an APNG animation directly from an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [LoadAPNGAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAPNGAnimation_IO
func LoadAPNGAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadAPNGAnimationIO(src)
}

// [LoadAVIFAnimationIO] loads an AVIF animation directly from an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [LoadAVIFAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_LoadAVIFAnimation_IO
func LoadAVIFAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadAVIFAnimationIO(src)
}

// [SaveAnimation] saves an animation to a file.
//
// Available since SDL_image 3.4.0.
//
// [SaveAnimation]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAnimation
func SaveAnimation(anim *Animation, file string) bool {
	return imgSaveAnimation(anim, file)
}

// [SaveAnimationTypedIO] saves an animation to an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveAnimationTypedIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAnimationTyped_IO
func SaveAnimationTypedIO(anim *Animation, dst *sdl.IOStream, closeio bool, _type string) bool {
	return imgSaveAnimationTypedIO(anim, dst, closeio, _type)
}

// [SaveANIAnimationIO] saves an animation in ANI format to an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveANIAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveANIAnimation_IO
func SaveANIAnimationIO(anim *Animation, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveANIAnimationIO(anim, dst, closeio)
}

// [SaveAPNGAnimationIO] saves an animation in APNG format to an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveAPNGAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAPNGAnimation_IO
func SaveAPNGAnimationIO(anim *Animation, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveAPNGAnimationIO(anim, dst, closeio)
}

// [SaveAVIFAnimationIO] saves an animation in AVIF format to an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveAVIFAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveAVIFAnimation_IO
func SaveAVIFAnimationIO(anim *Animation, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveAVIFAnimationIO(anim, dst, closeio, quality)
}

// [SaveGIFAnimationIO] saves an animation in GIF format to an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveGIFAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveGIFAnimation_IO
func SaveGIFAnimationIO(anim *Animation, dst *sdl.IOStream, closeio bool) bool {
	return imgSaveGIFAnimationIO(anim, dst, closeio)
}

// [SaveWEBPAnimationIO] saves an animation in WEBP format to an [sdl.IOStream].
//
// Available since SDL_image 3.4.0.
//
// [SaveWEBPAnimationIO]: https://wiki.libsdl.org/SDL3_image/IMG_SaveWEBPAnimation_IO
func SaveWEBPAnimationIO(anim *Animation, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveWEBPAnimationIO(anim, dst, closeio, quality)
}

// [CreateAnimatedCursor] creates an animated cursor from an animation.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimatedCursor]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimatedCursor
func CreateAnimatedCursor(anim *Animation, hotX, hotY int32) *sdl.Cursor {
	return imgCreateAnimatedCursor(anim, hotX, hotY)
}

// [AnimationEncoder] is an object representing the encoder context.
//
// Available since SDL_image 3.4.0.
//
// [AnimationEncoder]: https://wiki.libsdl.org/SDL3_image/IMG_AnimationEncoder
type AnimationEncoder struct{}

// [CreateAnimationEncoder] creates an encoder to save a series of images to a file.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimationEncoder]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationEncoder
func CreateAnimationEncoder(file string) *AnimationEncoder {
	return imgCreateAnimationEncoder(file)
}

// [CreateAnimationEncoderIO] creates an encoder to save a series of images to an IOStream.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimationEncoderIO]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationEncoder_IO
func CreateAnimationEncoderIO(dst *sdl.IOStream, closeio bool, _type string) *AnimationEncoder {
	return imgCreateAnimationEncoderIO(dst, closeio, _type)
}

// [CreateAnimationEncoderWithProperties] creates an animation encoder with the specified properties.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimationEncoderWithProperties]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationEncoderWithProperties
func CreateAnimationEncoderWithProperties(props sdl.PropertiesID) *AnimationEncoder {
	return imgCreateAnimationEncoderWithProperties(props)
}

const (
	PropAnimationEncoderCreateFilenameString            = "SDL_image.animation_encoder.create.filename"
	PropAnimationEncoderCreateIostreamPointer           = "SDL_image.animation_encoder.create.iostream"
	PropAnimationEncoderCreateIostreamAutocloseBoolean  = "SDL_image.animation_encoder.create.iostream.autoclose"
	PropAnimationEncoderCreateTypeString                = "SDL_image.animation_encoder.create.type"
	PropAnimationEncoderCreateQualityNumber             = "SDL_image.animation_encoder.create.quality"
	PropAnimationEncoderCreateTimebaseNumeratorNumber   = "SDL_image.animation_encoder.create.timebase.numerator"
	PropAnimationEncoderCreateTimebaseDenominatorNumber = "SDL_image.animation_encoder.create.timebase.denominator"

	PropAnimationEncoderCreateAvifMaxThreadsNumber       = "SDL_image.animation_encoder.create.avif.max_threads"
	PropAnimationEncoderCreateAvifKeyframeIntervalNumber = "SDL_image.animation_encoder.create.avif.keyframe_interval"
	PropAnimationEncoderCreateGifUseLutBoolean           = "SDL_image.animation_encoder.create.gif.use_lut"
)

// [AddAnimationEncoderFrame] adds a frame to an animation encoder.
//
// Available since SDL_image 3.4.0.
//
// [AddAnimationEncoderFrame]: https://wiki.libsdl.org/SDL3_image/IMG_AddAnimationEncoderFrame
func AddAnimationEncoderFrame(encoder *AnimationEncoder, surface *sdl.Surface, duration uint64) bool {
	return imgAddAnimationEncoderFrame(encoder, surface, duration)
}

// [CloseAnimationEncoder] closes an animation encoder, finishing any encoding.
//
// Available since SDL_image 3.4.0.
//
// [CloseAnimationEncoder]: https://wiki.libsdl.org/SDL3_image/IMG_CloseAnimationEncoder
func CloseAnimationEncoder(encoder *AnimationEncoder) bool {
	return imgCloseAnimationEncoder(encoder)
}

// [AnimationDecoderStatus] is an enum representing the status of an animation decoder.
//
// Available since SDL_image 3.4.0.
//
// [AnimationDecoderStatus]: https://wiki.libsdl.org/SDL3_image/IMG_AnimationDecoderStatus
type AnimationDecoderStatus int32

const (
	DecoderStatusInvalid  AnimationDecoderStatus = iota - 1 // The decoder is invalid.
	DecoderStatusOk                                         // The decoder is ready to decode the next frame.
	DecoderStatusFailed                                     // The decoder failed to decode a frame, call [sdl.GetError] for more information.
	DecoderStatusComplete                                   // No more frames available.
)

// [AnimationDecoder] is an object representing animation decoder.
//
// Available since SDL_image 3.4.0.
//
// [AnimationDecoder]: https://wiki.libsdl.org/SDL3_image/IMG_AnimationDecoder
type AnimationDecoder struct{}

// [CreateAnimationDecoder] creates a decoder to read a series of images from a file.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimationDecoder]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationDecoder
func CreateAnimationDecoder(file string) *AnimationDecoder {
	return imgCreateAnimationDecoder(file)
}

// [CreateAnimationDecoderIO] creates a decoder to read a series of images from an IOStream.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimationDecoderIO]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationDecoder_IO
func CreateAnimationDecoderIO(src *sdl.IOStream, closeio bool, _type string) *AnimationDecoder {
	return imgCreateAnimationDecoderIO(src, closeio, _type)
}

// [CreateAnimationDecoderWithProperties] creates an animation decoder with the specified properties.
//
// Available since SDL_image 3.4.0.
//
// [CreateAnimationDecoderWithProperties]: https://wiki.libsdl.org/SDL3_image/IMG_CreateAnimationDecoderWithProperties
func CreateAnimationDecoderWithProperties(props sdl.PropertiesID) *AnimationDecoder {
	return imgCreateAnimationDecoderWithProperties(props)
}

const (
	PropAnimationDecoderCreateFilenameString            = "SDL_image.animation_decoder.create.filename"
	PropAnimationDecoderCreateIostreamPointer           = "SDL_image.animation_decoder.create.iostream"
	PropAnimationDecoderCreateIostreamAutocloseBoolean  = "SDL_image.animation_decoder.create.iostream.autoclose"
	PropAnimationDecoderCreateTypeString                = "SDL_image.animation_decoder.create.type"
	PropAnimationDecoderCreateTimebaseNumeratorNumber   = "SDL_image.animation_decoder.create.timebase.numerator"
	PropAnimationDecoderCreateTimebaseDenominatorNumber = "SDL_image.animation_decoder.create.timebase.denominator"

	PropAnimationDecoderCreateAvifMaxThreadsNumber           = "SDL_image.animation_decoder.create.avif.max_threads"
	PropAnimationDecoderCreateAvifAllowIncrementalBoolean    = "SDL_image.animation_decoder.create.avif.allow_incremental"
	PropAnimationDecoderCreateAvifAllowProgressiveBoolean    = "SDL_image.animation_decoder.create.avif.allow_progressive"
	PropAnimationDecoderCreateGifTransparentColorIndexNumber = "SDL_image.animation_encoder.create.gif.transparent_color_index"
	PropAnimationDecoderCreateGifNumColorsNumber             = "SDL_image.animation_encoder.create.gif.num_colors"
)

// [GetAnimationDecoderProperties] gets the properties of an animation decoder.
//
// Available since SDL_image 3.4.0.
//
// [GetAnimationDecoderProperties]: https://wiki.libsdl.org/SDL3_image/IMG_GetAnimationDecoderProperties
func GetAnimationDecoderProperties(decoder *AnimationDecoder) sdl.PropertiesID {
	return imgGetAnimationDecoderProperties(decoder)
}

const (
	PropMetadataIgnorePropsBoolean = "SDL_image.metadata.ignore_props"
	PropMetadataDescriptionString  = "SDL_image.metadata.description"
	PropMetadataCopyrightString    = "SDL_image.metadata.copyright"
	PropMetadataTitleString        = "SDL_image.metadata.title"
	PropMetadataAuthorString       = "SDL_image.metadata.author"
	PropMetadataCreationTimeString = "SDL_image.metadata.creation_time"
	PropMetadataFrameCountNumber   = "SDL_image.metadata.frame_count"
	PropMetadataLoopCountNumber    = "SDL_image.metadata.loop_count"
)

// [GetAnimationDecoderFrame] gets the next frame in an animation decoder.
//
// Available since SDL_image 3.4.0.
//
// [GetAnimationDecoderFrame]: https://wiki.libsdl.org/SDL3_image/IMG_GetAnimationDecoderFrame
func GetAnimationDecoderFrame(decoder *AnimationDecoder, frame **sdl.Surface, duration *uint64) bool {
	return imgGetAnimationDecoderFrame(decoder, frame, duration)
}

// [GetAnimationDecoderStatus] gets the decoder status indicating the current state of the decoder.
//
// Available since SDL_image 3.4.0.
//
// [GetAnimationDecoderStatus]: https://wiki.libsdl.org/SDL3_image/IMG_GetAnimationDecoderStatus
func GetAnimationDecoderStatus(decoder *AnimationDecoder) AnimationDecoderStatus {
	return imgGetAnimationDecoderStatus(decoder)
}

// [ResetAnimationDecoder] resets an animation decoder.
//
// Available since SDL_image 3.4.0.
//
// [ResetAnimationDecoder]: https://wiki.libsdl.org/SDL3_image/IMG_ResetAnimationDecoder
func ResetAnimationDecoder(decoder *AnimationDecoder) bool {
	return imgResetAnimationDecoder(decoder)
}

// [CloseAnimationDecoder] closes an animation decoder, finishing any decoding.
//
// Available since SDL_image 3.4.0.
//
// [CloseAnimationDecoder]: https://wiki.libsdl.org/SDL3_image/IMG_CloseAnimationDecoder
func CloseAnimationDecoder(decoder *AnimationDecoder) bool {
	return imgCloseAnimationDecoder(decoder)
}
