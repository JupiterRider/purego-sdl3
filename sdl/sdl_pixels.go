package sdl

import "unsafe"

const (
	AlphaOpaque           = 255
	AlphaOpaqueFloat      = 1.0
	AlphaTransparent      = 0
	AlphaTransparentFloat = 0.0
)

// [Colorspace] defines the colorspace definitions.
//
// [Colorspace]: https://wiki.libsdl.org/SDL3/SDL_Colorspace
type Colorspace uint32

const (
	ColorspaceUnknown Colorspace = 0
	// sRGB is a gamma corrected colorspace, and the default colorspace for SDL rendering and 8-bit RGB surfaces.
	//
	// Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G22_NONE_P709
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT709,
	//                        SDL_TRANSFER_CHARACTERISTICS_SRGB,
	//                        SDL_MATRIX_COEFFICIENTS_IDENTITY,
	//                        SDL_CHROMA_LOCATION_NONE),
	ColorspaceSRGB Colorspace = 0x120005a0
	// This is a linear colorspace and the default colorspace for floating point surfaces. On Windows this is the scRGB colorspace, and on Apple platforms this is kCGColorSpaceExtendedLinearSRGB for EDR content.
	//
	// Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G10_NONE_P709
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT709,
	//                        SDL_TRANSFER_CHARACTERISTICS_LINEAR,
	//                        SDL_MATRIX_COEFFICIENTS_IDENTITY,
	//                        SDL_CHROMA_LOCATION_NONE)
	ColorspaceSRGBLinear Colorspace = 0x12000500
	// HDR10 is a non-linear HDR colorspace and the default colorspace for 10-bit surfaces.
	//
	// Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G2084_NONE_P2020
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT2020,
	//                        SDL_TRANSFER_CHARACTERISTICS_PQ,
	//                        SDL_MATRIX_COEFFICIENTS_IDENTITY,
	//                        SDL_CHROMA_LOCATION_NONE)
	ColorspaceHDR10 Colorspace = 0x12002600
	// Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_NONE_P709_X601
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT709,
	//                        SDL_TRANSFER_CHARACTERISTICS_BT601,
	//                        SDL_MATRIX_COEFFICIENTS_BT601,
	//                        SDL_CHROMA_LOCATION_NONE)
	ColorspaceJPEG Colorspace = 0x220004c6
	// Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_LIMITED,
	//                        SDL_COLOR_PRIMARIES_BT601,
	//                        SDL_TRANSFER_CHARACTERISTICS_BT601,
	//                        SDL_MATRIX_COEFFICIENTS_BT601,
	//                        SDL_CHROMA_LOCATION_LEFT)
	ColorspaceBT601Limited Colorspace = 0x211018c6
	// Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT601,
	//                        SDL_TRANSFER_CHARACTERISTICS_BT601,
	//                        SDL_MATRIX_COEFFICIENTS_BT601,
	//                        SDL_CHROMA_LOCATION_LEFT)
	ColorspaceBT601Full Colorspace = 0x221018c6
	// Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_LIMITED,
	//                        SDL_COLOR_PRIMARIES_BT709,
	//                        SDL_TRANSFER_CHARACTERISTICS_BT709,
	//                        SDL_MATRIX_COEFFICIENTS_BT709,
	//                        SDL_CHROMA_LOCATION_LEFT)
	ColorspaceBT709Limited Colorspace = 0x21100421
	// Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT709,
	//                        SDL_TRANSFER_CHARACTERISTICS_BT709,
	//                        SDL_MATRIX_COEFFICIENTS_BT709,
	//                        SDL_CHROMA_LOCATION_LEFT)
	ColorspaceBT709Full Colorspace = 0x22100421
	// Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P2020
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_LIMITED,
	//                        SDL_COLOR_PRIMARIES_BT2020,
	//                        SDL_TRANSFER_CHARACTERISTICS_PQ,
	//                        SDL_MATRIX_COEFFICIENTS_BT2020_NCL,
	//                        SDL_CHROMA_LOCATION_LEFT)
	ColorspaceBT2020Limited Colorspace = 0x21102609
	//Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_LEFT_P2020
	//  SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	//                        SDL_COLOR_RANGE_FULL,
	//                        SDL_COLOR_PRIMARIES_BT2020,
	//                        SDL_TRANSFER_CHARACTERISTICS_PQ,
	//                        SDL_MATRIX_COEFFICIENTS_BT2020_NCL,
	//                        SDL_CHROMA_LOCATION_LEFT)
	ColorspaceBT2020Full Colorspace = 0x22102609
	// The default colorspace for RGB surfaces if no colorspace is specified.
	ColorspaceRGBDefault Colorspace = ColorspaceSRGB
	// The default colorspace for YUV surfaces if no colorspace is specified.
	ColorspaceYUVDefault Colorspace = ColorspaceBT601Limited
)

// [ChromaLocation] defines the colorspace chroma sample location.
//
// [ChromaLocation]: https://wiki.libsdl.org/SDL3/SDL_ChromaLocation
type ChromaLocation uint32

const (
	ChromaLocationNone   ChromaLocation = iota // RGB, no chroma sampling.
	ChromaLocationLeft                         // In MPEG-2, MPEG-4, and AVC, Cb and Cr are taken on midpoint of the left-edge of the 2x2 square. In other words, they have the same horizontal location as the top-left pixel, but is shifted one-half pixel down vertically.
	ChromaLocationCenter                       // In JPEG/JFIF, H.261, and MPEG-1, Cb and Cr are taken at the center of the 2x2 square. In other words, they are offset one-half pixel to the right and one-half pixel down compared to the top-left pixel.
	ChromaLocationTopLeft
)

// [MatrixCoefficients] defines the colorspace matrix coefficients.
//
// [MatrixCoefficients]: https://wiki.libsdl.org/SDL3/SDL_MatrixCoefficients
type MatrixCoefficients uint32

const (
	MatrixCoefficientsIdentity         MatrixCoefficients = 0
	MatrixCoefficientsBT709            MatrixCoefficients = 1 // ITU-R BT.709-6.
	MatrixCoefficientsUnspecified      MatrixCoefficients = 2
	MatrixCoefficientsFCC              MatrixCoefficients = 4 // US FCC Title 47.
	MatrixCoefficientsBT470BG          MatrixCoefficients = 5 // ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625, functionally the same as SDL_MATRIX_COEFFICIENTS_BT601.
	MatrixCoefficientsBT601            MatrixCoefficients = 6 // ITU-R BT.601-7 525.
	MatrixCoefficientsSMPTE240         MatrixCoefficients = 7 // SMPTE 240M.
	MatrixCoefficientsYCGCO            MatrixCoefficients = 8
	MatrixCoefficientsBT2020NCL        MatrixCoefficients = 9  // ITU-R BT.2020-2 non-constant luminance.
	MatrixCoefficientsBT2020CL         MatrixCoefficients = 10 // ITU-R BT.2020-2 constant luminance.
	MatrixCoefficientsSMPTE2085        MatrixCoefficients = 11 // SMPTE ST 2085.
	MatrixCoefficientsChromaDerivedNCL MatrixCoefficients = 12
	MatrixCoefficientsChromaDerivedCL  MatrixCoefficients = 13
	MatrixCoefficientsICTCP            MatrixCoefficients = 14 // ITU-R BT.2100-0 ICTCP.
	MatrixCoefficientsCustom           MatrixCoefficients = 31
)

// [TransferCharacteristics] defines the colorspace transfer characteristics.
//
// [TransferCharacteristics]: https://wiki.libsdl.org/SDL3/SDL_TransferCharacteristics
type TransferCharacteristics uint32

const (
	TransferCharacteristicsUnknown      TransferCharacteristics = 0
	TransferCharacteristicsBT709        TransferCharacteristics = 1 // Rec. ITU-R BT.709-6 / ITU-R BT1361.
	TransferCharacteristicsUnspecified  TransferCharacteristics = 2
	TransferCharacteristicsGamma22      TransferCharacteristics = 4 // ITU-R BT.470-6 System M / ITU-R BT1700 625 PAL & SECAM.
	TransferCharacteristicsGamma28      TransferCharacteristics = 5 // ITU-R BT.470-6 System B, G.
	TransferCharacteristicsBT601        TransferCharacteristics = 6 // SMPTE ST 170M / ITU-R BT.601-7 525 or 625.
	TransferCharacteristicsSMPTE240     TransferCharacteristics = 7 // SMPTE ST 240M.
	TransferCharacteristicsLinear       TransferCharacteristics = 8
	TransferCharacteristicsLog100       TransferCharacteristics = 9
	TransferCharacteristicsLog100Sqrt10 TransferCharacteristics = 10
	TransferCharacteristicsIEC61966     TransferCharacteristics = 11 // IEC 61966-2-4.
	TransferCharacteristicsBT1361       TransferCharacteristics = 12 // ITU-R BT1361 Extended Colour Gamut.
	TransferCharacteristicsSRGB         TransferCharacteristics = 13 // IEC 61966-2-1 (sRGB or sYCC).
	TransferCharacteristicsBT202010Bit  TransferCharacteristics = 14 // ITU-R BT2020 for 10-bit system.
	TransferCharacteristicsBT202012Bit  TransferCharacteristics = 15 // ITU-R BT2020 for 12-bit system.
	TransferCharacteristicsPQ           TransferCharacteristics = 16 // SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems.
	TransferCharacteristicsSMPTE428     TransferCharacteristics = 17 // SMPTE ST 428-1.
	TransferCharacteristicsHLG          TransferCharacteristics = 18 // ARIB STD-B67, known as "hybrid log-gamma" (HLG).
	TransferCharacteristicsCustom       TransferCharacteristics = 31
)

// [ColorPrimaries] defines the colorspace color primaries, as described by https://www.itu.int/rec/T-REC-H.273-201612-S/en.
//
// [ColorPrimaries]: https://wiki.libsdl.org/SDL3/SDL_ColorPrimaries
type ColorPrimaries uint32

const (
	ColorPrimariesUnknown     ColorPrimaries = 0
	ColorPrimariesBT709       ColorPrimaries = 1 // ITU-R BT.709-6.
	ColorPrimariesUnspecified ColorPrimaries = 2
	ColorPrimariesBT470M      ColorPrimaries = 4  // ITU-R BT.470-6 System M.
	ColorPrimariesBT470BG     ColorPrimaries = 5  // ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625.
	ColorPrimariesBT601       ColorPrimaries = 6  // ITU-R BT.601-7 525, SMPTE 170M.
	ColorPrimariesSMPTE240    ColorPrimaries = 7  // SMPTE 240M, functionally the same as SDL_COLOR_PRIMARIES_BT601.
	ColorPrimariesGenericFilm ColorPrimaries = 8  // Generic film (color filters using Illuminant C).
	ColorPrimariesBT2020      ColorPrimaries = 9  // ITU-R BT.2020-2 / ITU-R BT.2100-0.
	ColorPrimariesXYZ         ColorPrimaries = 10 // SMPTE ST 428-1.
	ColorPrimariesSMPTE431    ColorPrimaries = 11 // SMPTE RP 431-2.
	ColorPrimariesSMPTE432    ColorPrimaries = 12 // SMPTE EG 432-1 / DCI P3.
	ColorPrimariesEBU3213     ColorPrimaries = 22 // EBU Tech. 3213-E.
	ColorPrimariesCustom      ColorPrimaries = 31
)

// [ColorRange] defines the colorspace color range, as described by https://www.itu.int/rec/R-REC-BT.2100-2-201807-I/en.
//
// [ColorRange]: https://wiki.libsdl.org/SDL3/SDL_ColorRange
type ColorRange uint32

const (
	ColorRangeUnknown ColorRange = iota
	ColorRangeLimited            // Narrow range, e.g. 16-235 for 8-bit RGB and luma, and 16-240 for 8-bit chroma.
	ColorRangeFull               // Full range, e.g. 0-255 for 8-bit RGB and luma, and 1-255 for 8-bit chroma.
)

// [ColorType] defines the colorspace color type.
//
// [ColorType]: https://wiki.libsdl.org/SDL3/SDL_ColorType
type ColorType uint32

const (
	ColorTypeUnknown ColorType = iota
	ColorTypeRGB
	ColorTypeYCBCR
)

// [PixelType] defines the pixel type.
//
// [PixelType]: https://wiki.libsdl.org/SDL3/SDL_PixelType
type PixelType uint32

const (
	PixelTypeUnknown PixelType = iota
	PixelTypeIndex1
	PixelTypeIndex4
	PixelTypeIndex8
	PixelTypePacked8
	PixelTypePacked16
	PixelTypePacked32
	PixelTypeArrayU8
	PixelTypeArrayU16
	PixelTypeArrayU32
	PixelTypeArrayF16
	PixelTypeArrayF32
	PixelTypeIndex2
)

// [BitmapOrder] defines the bitmap pixel order, high bit -> low bit.
//
// [BitmapOrder]: https://wiki.libsdl.org/SDL3/SDL_BitmapOrder
type BitmapOrder uint32

const (
	BitmapOrderNone BitmapOrder = iota
	BitmapOrder4321
	BitmapOrder1234
)

// [PackedOrder] defines the packed component order, high bit -> low bit.
//
// [PackedOrder]: https://wiki.libsdl.org/SDL3/SDL_PackedOrder
type PackedOrder uint32

const (
	PackedOrderNone PackedOrder = iota
	PackedOrderXRGB
	PackedOrderRGBX
	PackedOrderARGB
	PackedOrderRGBA
	PackedOrderXBGR
	PackedOrderBGRX
	PackedOrderABGR
	PackedOrderBGRA
)

// [ArrayOrder] defines the array component order, low byte -> high byte.
//
// [ArrayOrder]: https://wiki.libsdl.org/SDL3/SDL_ArrayOrder
type ArrayOrder uint32

const (
	ArrayOrderNone ArrayOrder = iota
	ArrayOrderRGB
	ArrayOrderRGBA
	ArrayOrderARGB
	ArrayOrderBGR
	ArrayOrderBGRA
	ArrayOrderABGR
)

// [PackedLayout] defines the packed component layout.
//
// [PackedLayout]: https://wiki.libsdl.org/SDL3/SDL_PackedLayout
type PackedLayout uint32

const (
	PackedLayoutNone PackedLayout = iota
	PackedLayout332
	PackedLayout4444
	PackedLayout1555
	PackedLayout5551
	PackedLayout565
	PackedLayout8888
	PackedLayout2101010
	PackedLayout1010102
)

// [PixelFormat] defines the pixel format.
//
// [PixelFormat]: https://wiki.libsdl.org/SDL3/SDL_PixelFormat
type PixelFormat uint32

const (
	PixelFormatUnknown      PixelFormat = 0
	PixelFormatIndex1LSB    PixelFormat = 0x11100100
	PixelFormatIndex1MSB    PixelFormat = 0x11200100
	PixelFormatIndex2LSB    PixelFormat = 0x1C100200
	PixelFormatIndex2MSB    PixelFormat = 0x1C200200
	PixelFormatIndex4LSB    PixelFormat = 0x12100400
	PixelFormatIndex4MSB    PixelFormat = 0x12200400
	PixelFormatIndex8       PixelFormat = 0x13000801
	PixelFormatRGB332       PixelFormat = 0x14110801
	PixelFormatXRGB4444     PixelFormat = 0x15120C02
	PixelFormatXBGR4444     PixelFormat = 0x15520C02
	PixelFormatXRGB1555     PixelFormat = 0x15130F02
	PixelFormatXBGR1555     PixelFormat = 0x15530F02
	PixelFormatARGB4444     PixelFormat = 0x15321002
	PixelFormatRGBA4444     PixelFormat = 0x15421002
	PixelFormatABGR4444     PixelFormat = 0x15721002
	PixelFormatBGRA4444     PixelFormat = 0x15821002
	PixelFormatARGB1555     PixelFormat = 0x15331002
	PixelFormatRGBA5551     PixelFormat = 0x15441002
	PixelFormatABGR1555     PixelFormat = 0x15731002
	PixelFormatBGRA5551     PixelFormat = 0x15841002
	PixelFormatRGB565       PixelFormat = 0x15151002
	PixelFormatBGR565       PixelFormat = 0x15551002
	PixelFormatRGB24        PixelFormat = 0x17101803
	PixelFormatBGR24        PixelFormat = 0x17401803
	PixelFormatXRGB8888     PixelFormat = 0x16161804
	PixelFormatRGBX8888     PixelFormat = 0x16261804
	PixelFormatXBGR8888     PixelFormat = 0x16561804
	PixelFormatBGRX8888     PixelFormat = 0x16661804
	PixelFormatARGB8888     PixelFormat = 0x16362004
	PixelFormatRGBA8888     PixelFormat = 0x16462004
	PixelFormatABGR8888     PixelFormat = 0x16762004
	PixelFormatBGRA8888     PixelFormat = 0x16862004
	PixelFormatXRGB2101010  PixelFormat = 0x16172004
	PixelFormatXBGR2101010  PixelFormat = 0x16572004
	PixelFormatARGB2101010  PixelFormat = 0x16372004
	PixelFormatABGR2101010  PixelFormat = 0x16772004
	PixelFormatRGB48        PixelFormat = 0x18103006
	PixelFormatBGR48        PixelFormat = 0x18403006
	PixelFormatRGBA64       PixelFormat = 0x18204008
	PixelFormatARGB64       PixelFormat = 0x18304008
	PixelFormatBGRA64       PixelFormat = 0x18504008
	PixelFormatABGR64       PixelFormat = 0x18604008
	PixelFormatRGB48Float   PixelFormat = 0x1A103006
	PixelFormatBGR48Float   PixelFormat = 0x1A403006
	PixelFormatRGBA64Float  PixelFormat = 0x1A204008
	PixelFormatARGB64Float  PixelFormat = 0x1A304008
	PixelFormatBGRA64Float  PixelFormat = 0x1A504008
	PixelFormatABGR64Float  PixelFormat = 0x1A604008
	PixelFormatRGB96Float   PixelFormat = 0x1B10600C
	PixelFormatBGR96Float   PixelFormat = 0x1B40600C
	PixelFormatRGBA128Float PixelFormat = 0x1B208010
	PixelFormatARGB128Float PixelFormat = 0x1B308010
	PixelFormatBGRA128Float PixelFormat = 0x1B508010
	PixelFormatABGR128Float PixelFormat = 0x1B608010
	PixelFormatYV12         PixelFormat = 0x32315659
	PixelFormatIYUV         PixelFormat = 0x56555949
	PixelFormatYUY2         PixelFormat = 0x32595559
	PixelFormatUYVY         PixelFormat = 0x59565955
	PixelFormatYVYU         PixelFormat = 0x55595659
	PixelFormatNV12         PixelFormat = 0x3231564E
	PixelFormatNV21         PixelFormat = 0x3132564E
	PixelFormatP010         PixelFormat = 0x30313050
	PixelFormatExternalOES  PixelFormat = 0x2053454F
	PixelFormatRGBA32       PixelFormat = PixelFormatABGR8888
	PixelFormatARGB32       PixelFormat = PixelFormatBGRA8888
	PixelFormatBGRA32       PixelFormat = PixelFormatARGB8888
	PixelFormatABGR32       PixelFormat = PixelFormatRGBA8888
	PixelFormatRGBX32       PixelFormat = PixelFormatXBGR8888
	PixelFormatXRGB32       PixelFormat = PixelFormatBGRX8888
	PixelFormatBGRX32       PixelFormat = PixelFormatXRGB8888
	PixelFormatXBGR32       PixelFormat = PixelFormatRGBX8888
)

// [Color] is a structure that represents a color as RGBA components.
//
// [Color]: https://wiki.libsdl.org/SDL3/SDL_Color
type Color struct {
	R, G, B, A uint8
}

// [FColor] is a structure specifying the bits of this structure can be directly reinterpreted as a float-packed color which uses the [PIXELFORMAT_RGBA128_FLOAT] formas.
//
// [FColor]: https://wiki.libsdl.org/SDL3/SDL_FColor
type FColor struct {
	R, G, B, A float32
}

// [Palette] is a set of indexed colors representing a palette.
//
// [Palette]: https://wiki.libsdl.org/SDL3/SDL_Palette
type Palette struct {
	ncolors int32  // Number of elements in `colors`.
	colors  *Color // Number of elements in `colors`.
	version uint32 // Internal use only, do not touch.
	refcount int32 // Internal use only, do not touch.
}

// [PixelFormatDetails] defines the details about the format of a pixel.
//
// [PixelFormatDetails]: https://wiki.libsdl.org/SDL3/SDL_PixelFormatDetails
type PixelFormatDetails struct {
	Format                         PixelFormat
	BitsPerPixel                   uint8
	BytesPerPixel                  uint8
	Padding                        [2]uint8
	Rmask, Gmask, Bmask, Amask     uint32
	Rbits, Gbits, Bbits, Abits     uint8
	Rshift, Gshift, Bshift, Ashift uint8
}

// Colors gets available colors in the palette.
func (p *Palette) Colors() []Color {
	if p == nil || p.colors == nil {
		return nil
	}
	return unsafe.Slice(p.colors, p.ncolors)
}

// [DefinePixelFourCC] defines custom FourCC pixel formats.
//
// [DefinePixelFourCC]: https://wiki.libsdl.org/SDL3/SDL_DEFINE_PIXELFOURCC
func DefinePixelFourCC(a, b, c, d byte) PixelFormat {
	return PixelFormat(FourCC(a, b, c, d))
}

// [DefinePixelFormat] defines custom non-FourCC pixel formats.
//
// [DefinePixelFormat]: https://wiki.libsdl.org/SDL3/SDL_DEFINE_PIXELFORMAT
func DefinePixelFormat[T BitmapOrder | PackedOrder | ArrayOrder](pixelType PixelType, order T, layout PackedLayout, bitsPerPixel, bytesPerPixel byte) PixelFormat {
	return 1<<28 | PixelFormat(pixelType)<<24 | PixelFormat(order)<<20 | PixelFormat(layout)<<16 | PixelFormat(bitsPerPixel)<<8 | PixelFormat(bytesPerPixel)
}

// [CreatePalette] creates a palette structure with the specified number of color entries.
//
// [CreatePalette]: https://wiki.libsdl.org/SDL3/SDL_CreatePalette
func CreatePalette(ncolors int32) *Palette {
	return sdlCreatePalette(ncolors)
}

// [DestroyPalette] frees a palette created with [CreatePalette].
//
// [DestroyPalette]: https://wiki.libsdl.org/SDL3/SDL_DestroyPalette
func DestroyPalette(palette *Palette) {
	sdlDestroyPalette(palette)
}

// func GetMasksForPixelFormat(format PixelFormat, bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) bool {
//	return sdlGetMasksForPixelFormat(format, bpp, Rmask, Gmask, Bmask, Amask)
// }

// [GetPixelFormatDetails] creates an [PixelFormatDetails] structure corresponding to a pixel format.
//
// [GetPixelFormatDetails]: https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails
func GetPixelFormatDetails(format PixelFormat) *PixelFormatDetails {
	return sdlGetPixelFormatDetails(format)
}

// func GetPixelFormatForMasks(bpp int32, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) PixelFormat {
//	return sdlGetPixelFormatForMasks(bpp, Rmask, Gmask, Bmask, Amask)
// }

// func GetPixelFormatName(format PixelFormat) string {
//	return sdlGetPixelFormatName(format)
// }

// func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8)  {
//	sdlGetRGB(pixel, format, palette, r, g, b)
// }

// func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8, a *uint8)  {
//	sdlGetRGBA(pixel, format, palette, r, g, b, a)
// }

// [MapRGB] maps an RGB triple to an opaque pixel value for a given pixel format.
//
// [MapRGB]: https://wiki.libsdl.org/SDL3/SDL_MapRGB
func MapRGB(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8) uint32 {
	return sdlMapRGB(format, palette, r, g, b)
}

// func MapRGBA(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8, a uint8) uint32 {
//	return sdlMapRGBA(format, palette, r, g, b, a)
// }

// [SetPaletteColors] sets a range of colors in a palette.
//
// [SetPaletteColors]: https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors
func SetPaletteColors(palette *Palette, firstcolor int32, Colors ...Color) bool {
	ncolors := len(Colors)
	var ptrColors *Color
	if ncolors > 0 {
		ptrColors = &Colors[0]
	}
	return sdlSetPaletteColors(palette, ptrColors, firstcolor, int32(ncolors))
}
