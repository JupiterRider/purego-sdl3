package img

import (
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

var (
	imgVersion                  func() int32
	imgLoadTypedIO              func(*sdl.IOStream, bool, string) *sdl.Surface
	imgLoad                     func(string) *sdl.Surface
	imgLoadIO                   func(*sdl.IOStream, bool) *sdl.Surface
	imgLoadTexture              func(*sdl.Renderer, string) *sdl.Texture
	imgLoadTextureIO            func(*sdl.Renderer, *sdl.IOStream, bool) *sdl.Texture
	imgLoadTextureTypedIO       func(*sdl.Renderer, *sdl.IOStream, bool, string) *sdl.Texture
	imgIsAVIF                   func(*sdl.IOStream) bool
	imgIsICO                    func(*sdl.IOStream) bool
	imgIsCUR                    func(*sdl.IOStream) bool
	imgIsBMP                    func(*sdl.IOStream) bool
	imgIsGIF                    func(*sdl.IOStream) bool
	imgIsJPG                    func(*sdl.IOStream) bool
	imgIsJXL                    func(*sdl.IOStream) bool
	imgIsLBM                    func(*sdl.IOStream) bool
	imgIsPCX                    func(*sdl.IOStream) bool
	imgIsPNG                    func(*sdl.IOStream) bool
	imgIsPNM                    func(*sdl.IOStream) bool
	imgIsSVG                    func(*sdl.IOStream) bool
	imgIsQOI                    func(*sdl.IOStream) bool
	imgIsTIF                    func(*sdl.IOStream) bool
	imgIsXCF                    func(*sdl.IOStream) bool
	imgIsXPM                    func(*sdl.IOStream) bool
	imgIsXV                     func(*sdl.IOStream) bool
	imgIsWEBP                   func(*sdl.IOStream) bool
	imgLoadAVIFIO               func(*sdl.IOStream) *sdl.Surface
	imgLoadICOIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadCURIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadBMPIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadGIFIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadJPGIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadJXLIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadLBMIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadPCXIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadPNGIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadPNMIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadSVGIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadQOIIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadTGAIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadTIFIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadXCFIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadXPMIO                func(*sdl.IOStream) *sdl.Surface
	imgLoadXVIO                 func(*sdl.IOStream) *sdl.Surface
	imgLoadWEBPIO               func(*sdl.IOStream) *sdl.Surface
	imgLoadSizedSVGIO           func(*sdl.IOStream, int32, int32) *sdl.Surface
	imgReadXPMFromArray         func(**byte) *sdl.Surface
	imgReadXPMFromArrayToRGB888 func(**byte) *sdl.Surface
	imgSaveAVIF                 func(*sdl.Surface, string, int32) bool
	imgSaveAVIFIO               func(*sdl.Surface, *sdl.IOStream, bool, int32) bool
	imgSavePNG                  func(*sdl.Surface, string) bool
	imgSavePNGIO                func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveJPG                  func(*sdl.Surface, string, int32) bool
	imgSaveJPGIO                func(*sdl.Surface, *sdl.IOStream, bool, int32) bool
	imgLoadAnimation            func(string) *Animation
	imgLoadAnimationIO          func(*sdl.IOStream, bool) *Animation
	imgLoadAnimationTypedIO     func(*sdl.IOStream, bool, string) *Animation
	imgFreeAnimation            func(*Animation)
	imgLoadGIFAnimationIO       func(*sdl.IOStream) *Animation
	imgLoadWEBPAnimationIO      func(*sdl.IOStream) *Animation

	imgLoadGPUTexture                       func(*sdl.GPUDevice, *sdl.GPUCopyPass, string, *int32, *int32) *sdl.GPUTexture
	imgLoadGPUTextureIO                     func(*sdl.GPUDevice, *sdl.GPUCopyPass, *sdl.IOStream, bool, *int32, *int32) *sdl.GPUTexture
	imgLoadGPUTextureTypedIO                func(*sdl.GPUDevice, *sdl.GPUCopyPass, *sdl.IOStream, bool, string, *int32, *int32) *sdl.GPUTexture
	imgGetClipboardImage                    func() *sdl.Surface
	imgIsANI                                func(*sdl.IOStream) bool
	imgSave                                 func(*sdl.Surface, string) bool
	imgSaveTypedIO                          func(*sdl.Surface, *sdl.IOStream, bool, string) bool
	imgSaveBMP                              func(*sdl.Surface, string) bool
	imgSaveBMPIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveCUR                              func(*sdl.Surface, string) bool
	imgSaveCURIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveGIF                              func(*sdl.Surface, string) bool
	imgSaveGIFIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveICO                              func(*sdl.Surface, string) bool
	imgSaveICOIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveTGA                              func(*sdl.Surface, string) bool
	imgSaveTGAIO                            func(*sdl.Surface, *sdl.IOStream, bool) bool
	imgSaveWEBP                             func(*sdl.Surface, string, float32) bool
	imgSaveWEBPIO                           func(*sdl.Surface, *sdl.IOStream, bool, float32) bool
	imgLoadANIAnimationIO                   func(*sdl.IOStream) *Animation
	imgLoadAPNGAnimationIO                  func(*sdl.IOStream) *Animation
	imgLoadAVIFAnimationIO                  func(*sdl.IOStream) *Animation
	imgSaveAnimation                        func(*Animation, string) bool
	imgSaveAnimationTypedIO                 func(*Animation, *sdl.IOStream, bool, string) bool
	imgSaveANIAnimationIO                   func(*Animation, *sdl.IOStream, bool) bool
	imgSaveAPNGAnimationIO                  func(*Animation, *sdl.IOStream, bool) bool
	imgSaveAVIFAnimationIO                  func(*Animation, *sdl.IOStream, bool, int32) bool
	imgSaveGIFAnimationIO                   func(*Animation, *sdl.IOStream, bool) bool
	imgSaveWEBPAnimationIO                  func(*Animation, *sdl.IOStream, bool, int32) bool
	imgCreateAnimatedCursor                 func(*Animation, int32, int32) *sdl.Cursor
	imgCreateAnimationEncoder               func(string) *AnimationEncoder
	imgCreateAnimationEncoderIO             func(*sdl.IOStream, bool, string) *AnimationEncoder
	imgCreateAnimationEncoderWithProperties func(sdl.PropertiesID) *AnimationEncoder
	imgAddAnimationEncoderFrame             func(*AnimationEncoder, *sdl.Surface, uint64) bool
	imgCloseAnimationEncoder                func(*AnimationEncoder) bool
	imgCreateAnimationDecoder               func(string) *AnimationDecoder
	imgCreateAnimationDecoderIO             func(*sdl.IOStream, bool, string) *AnimationDecoder
	imgCreateAnimationDecoderWithProperties func(sdl.PropertiesID) *AnimationDecoder
	imgGetAnimationDecoderProperties        func(*AnimationDecoder) sdl.PropertiesID
	imgGetAnimationDecoderFrame             func(*AnimationDecoder, **sdl.Surface, *uint64) bool
	imgGetAnimationDecoderStatus            func(*AnimationDecoder) AnimationDecoderStatus
	imgResetAnimationDecoder                func(*AnimationDecoder) bool
	imgCloseAnimationDecoder                func(*AnimationDecoder) bool
)

func init() {
	runtime.LockOSThread()

	var filename string
	switch runtime.GOOS {
	case "linux", "freebsd":
		filename = "libSDL3_image.so.0"
	case "windows":
		filename = "SDL3_image.dll"
	case "darwin":
		filename = "libSDL3_image.dylib"
	}

	lib, err := shared.Load(filename)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&imgFreeAnimation, lib, "IMG_FreeAnimation")
	purego.RegisterLibFunc(&imgIsAVIF, lib, "IMG_isAVIF")
	purego.RegisterLibFunc(&imgIsBMP, lib, "IMG_isBMP")
	purego.RegisterLibFunc(&imgIsCUR, lib, "IMG_isCUR")
	purego.RegisterLibFunc(&imgIsGIF, lib, "IMG_isGIF")
	purego.RegisterLibFunc(&imgIsICO, lib, "IMG_isICO")
	purego.RegisterLibFunc(&imgIsJPG, lib, "IMG_isJPG")
	purego.RegisterLibFunc(&imgIsJXL, lib, "IMG_isJXL")
	purego.RegisterLibFunc(&imgIsLBM, lib, "IMG_isLBM")
	purego.RegisterLibFunc(&imgIsPCX, lib, "IMG_isPCX")
	purego.RegisterLibFunc(&imgIsPNG, lib, "IMG_isPNG")
	purego.RegisterLibFunc(&imgIsPNM, lib, "IMG_isPNM")
	purego.RegisterLibFunc(&imgIsQOI, lib, "IMG_isQOI")
	purego.RegisterLibFunc(&imgIsSVG, lib, "IMG_isSVG")
	purego.RegisterLibFunc(&imgIsTIF, lib, "IMG_isTIF")
	purego.RegisterLibFunc(&imgIsWEBP, lib, "IMG_isWEBP")
	purego.RegisterLibFunc(&imgIsXCF, lib, "IMG_isXCF")
	purego.RegisterLibFunc(&imgIsXPM, lib, "IMG_isXPM")
	purego.RegisterLibFunc(&imgIsXV, lib, "IMG_isXV")
	purego.RegisterLibFunc(&imgLoad, lib, "IMG_Load")
	purego.RegisterLibFunc(&imgLoadIO, lib, "IMG_Load_IO")
	purego.RegisterLibFunc(&imgLoadAnimation, lib, "IMG_LoadAnimation")
	purego.RegisterLibFunc(&imgLoadAnimationIO, lib, "IMG_LoadAnimation_IO")
	purego.RegisterLibFunc(&imgLoadAnimationTypedIO, lib, "IMG_LoadAnimationTyped_IO")
	purego.RegisterLibFunc(&imgLoadAVIFIO, lib, "IMG_LoadAVIF_IO")
	purego.RegisterLibFunc(&imgLoadBMPIO, lib, "IMG_LoadBMP_IO")
	purego.RegisterLibFunc(&imgLoadCURIO, lib, "IMG_LoadCUR_IO")
	purego.RegisterLibFunc(&imgLoadGIFIO, lib, "IMG_LoadGIF_IO")
	purego.RegisterLibFunc(&imgLoadGIFAnimationIO, lib, "IMG_LoadGIFAnimation_IO")
	purego.RegisterLibFunc(&imgLoadICOIO, lib, "IMG_LoadICO_IO")
	purego.RegisterLibFunc(&imgLoadJPGIO, lib, "IMG_LoadJPG_IO")
	purego.RegisterLibFunc(&imgLoadJXLIO, lib, "IMG_LoadJXL_IO")
	purego.RegisterLibFunc(&imgLoadLBMIO, lib, "IMG_LoadLBM_IO")
	purego.RegisterLibFunc(&imgLoadPCXIO, lib, "IMG_LoadPCX_IO")
	purego.RegisterLibFunc(&imgLoadPNGIO, lib, "IMG_LoadPNG_IO")
	purego.RegisterLibFunc(&imgLoadPNMIO, lib, "IMG_LoadPNM_IO")
	purego.RegisterLibFunc(&imgLoadQOIIO, lib, "IMG_LoadQOI_IO")
	purego.RegisterLibFunc(&imgLoadSizedSVGIO, lib, "IMG_LoadSizedSVG_IO")
	purego.RegisterLibFunc(&imgLoadSVGIO, lib, "IMG_LoadSVG_IO")
	purego.RegisterLibFunc(&imgLoadTexture, lib, "IMG_LoadTexture")
	purego.RegisterLibFunc(&imgLoadTextureIO, lib, "IMG_LoadTexture_IO")
	purego.RegisterLibFunc(&imgLoadTextureTypedIO, lib, "IMG_LoadTextureTyped_IO")
	purego.RegisterLibFunc(&imgLoadTGAIO, lib, "IMG_LoadTGA_IO")
	purego.RegisterLibFunc(&imgLoadTIFIO, lib, "IMG_LoadTIF_IO")
	purego.RegisterLibFunc(&imgLoadTypedIO, lib, "IMG_LoadTyped_IO")
	purego.RegisterLibFunc(&imgLoadWEBPIO, lib, "IMG_LoadWEBP_IO")
	purego.RegisterLibFunc(&imgLoadWEBPAnimationIO, lib, "IMG_LoadWEBPAnimation_IO")
	purego.RegisterLibFunc(&imgLoadXCFIO, lib, "IMG_LoadXCF_IO")
	purego.RegisterLibFunc(&imgLoadXPMIO, lib, "IMG_LoadXPM_IO")
	purego.RegisterLibFunc(&imgLoadXVIO, lib, "IMG_LoadXV_IO")
	purego.RegisterLibFunc(&imgReadXPMFromArray, lib, "IMG_ReadXPMFromArray")
	purego.RegisterLibFunc(&imgReadXPMFromArrayToRGB888, lib, "IMG_ReadXPMFromArrayToRGB888")
	purego.RegisterLibFunc(&imgSaveAVIF, lib, "IMG_SaveAVIF")
	purego.RegisterLibFunc(&imgSaveAVIFIO, lib, "IMG_SaveAVIF_IO")
	purego.RegisterLibFunc(&imgSaveJPG, lib, "IMG_SaveJPG")
	purego.RegisterLibFunc(&imgSaveJPGIO, lib, "IMG_SaveJPG_IO")
	purego.RegisterLibFunc(&imgSavePNG, lib, "IMG_SavePNG")
	purego.RegisterLibFunc(&imgSavePNGIO, lib, "IMG_SavePNG_IO")
	purego.RegisterLibFunc(&imgVersion, lib, "IMG_Version")

	// Load functions available since 3.4.0
	versionMajor, versionMinor, _ := Version()
	if versionMajor >= 3 && versionMinor >= 4 {
		purego.RegisterLibFunc(&imgLoadGPUTexture, lib, "IMG_LoadGPUTexture")
		purego.RegisterLibFunc(&imgLoadGPUTextureIO, lib, "IMG_LoadGPUTexture_IO")
		purego.RegisterLibFunc(&imgLoadGPUTextureTypedIO, lib, "IMG_LoadGPUTextureTyped_IO")
		purego.RegisterLibFunc(&imgGetClipboardImage, lib, "IMG_GetClipboardImage")
		purego.RegisterLibFunc(&imgIsANI, lib, "IMG_isANI")
		purego.RegisterLibFunc(&imgSave, lib, "IMG_Save")
		purego.RegisterLibFunc(&imgSaveTypedIO, lib, "IMG_SaveTyped_IO")
		purego.RegisterLibFunc(&imgSaveBMP, lib, "IMG_SaveBMP")
		purego.RegisterLibFunc(&imgSaveBMPIO, lib, "IMG_SaveBMP_IO")
		purego.RegisterLibFunc(&imgSaveCUR, lib, "IMG_SaveCUR")
		purego.RegisterLibFunc(&imgSaveCURIO, lib, "IMG_SaveCUR_IO")
		purego.RegisterLibFunc(&imgSaveGIF, lib, "IMG_SaveGIF")
		purego.RegisterLibFunc(&imgSaveGIFIO, lib, "IMG_SaveGIF_IO")
		purego.RegisterLibFunc(&imgSaveICO, lib, "IMG_SaveICO")
		purego.RegisterLibFunc(&imgSaveICOIO, lib, "IMG_SaveICO_IO")
		purego.RegisterLibFunc(&imgSaveTGA, lib, "IMG_SaveTGA")
		purego.RegisterLibFunc(&imgSaveTGAIO, lib, "IMG_SaveTGA_IO")
		purego.RegisterLibFunc(&imgSaveWEBP, lib, "IMG_SaveWEBP")
		purego.RegisterLibFunc(&imgSaveWEBPIO, lib, "IMG_SaveWEBP_IO")
		purego.RegisterLibFunc(&imgLoadANIAnimationIO, lib, "IMG_LoadANIAnimation_IO")
		purego.RegisterLibFunc(&imgLoadAPNGAnimationIO, lib, "IMG_LoadAPNGAnimation_IO")
		purego.RegisterLibFunc(&imgLoadAVIFAnimationIO, lib, "IMG_LoadAVIFAnimation_IO")
		purego.RegisterLibFunc(&imgSaveAnimation, lib, "IMG_SaveAnimation")
		purego.RegisterLibFunc(&imgSaveAnimationTypedIO, lib, "IMG_SaveAnimationTyped_IO")
		purego.RegisterLibFunc(&imgSaveANIAnimationIO, lib, "IMG_SaveANIAnimation_IO")
		purego.RegisterLibFunc(&imgSaveAPNGAnimationIO, lib, "IMG_SaveAPNGAnimation_IO")
		purego.RegisterLibFunc(&imgSaveAVIFAnimationIO, lib, "IMG_SaveAVIFAnimation_IO")
		purego.RegisterLibFunc(&imgSaveGIFAnimationIO, lib, "IMG_SaveGIFAnimation_IO")
		purego.RegisterLibFunc(&imgSaveWEBPAnimationIO, lib, "IMG_SaveWEBPAnimation_IO")
		purego.RegisterLibFunc(&imgCreateAnimatedCursor, lib, "IMG_CreateAnimatedCursor")
		purego.RegisterLibFunc(&imgCreateAnimationEncoder, lib, "IMG_CreateAnimationEncoder")
		purego.RegisterLibFunc(&imgCreateAnimationEncoderIO, lib, "IMG_CreateAnimationEncoder_IO")
		purego.RegisterLibFunc(&imgCreateAnimationEncoderWithProperties, lib, "IMG_CreateAnimationEncoderWithProperties")
		purego.RegisterLibFunc(&imgAddAnimationEncoderFrame, lib, "IMG_AddAnimationEncoderFrame")
		purego.RegisterLibFunc(&imgCloseAnimationEncoder, lib, "IMG_CloseAnimationEncoder")
		purego.RegisterLibFunc(&imgCreateAnimationDecoder, lib, "IMG_CreateAnimationDecoder")
		purego.RegisterLibFunc(&imgCreateAnimationDecoderIO, lib, "IMG_CreateAnimationDecoder_IO")
		purego.RegisterLibFunc(&imgCreateAnimationDecoderWithProperties, lib, "IMG_CreateAnimationDecoderWithProperties")
		purego.RegisterLibFunc(&imgGetAnimationDecoderProperties, lib, "IMG_GetAnimationDecoderProperties")
		purego.RegisterLibFunc(&imgGetAnimationDecoderFrame, lib, "IMG_GetAnimationDecoderFrame")
		purego.RegisterLibFunc(&imgGetAnimationDecoderStatus, lib, "IMG_GetAnimationDecoderStatus")
		purego.RegisterLibFunc(&imgResetAnimationDecoder, lib, "IMG_ResetAnimationDecoder")
		purego.RegisterLibFunc(&imgCloseAnimationDecoder, lib, "IMG_CloseAnimationDecoder")
	}
}
