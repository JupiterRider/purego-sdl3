package sdl

// [BlendOperation] is a structure specifying the blend operation used when combining source and destination pixel components.
//
// [BlendOperation]: https://wiki.libsdl.org/SDL3/SDL_BlendOperation
type BlendOperation uint32

const (
	BlendOperationAdd         BlendOperation = iota + 1 // dst + src: supported by all renderers.
	BlendOperationSubtract                              // src - dst : supported by D3D, OpenGL, OpenGLES, and Vulkan.
	BlendOperationRevSubtract                           // dst - src : supported by D3D, OpenGL, OpenGLES, and Vulkan.
	BlendOperationMinimum                               // min(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan.
	BlendOperationMaximum                               // max(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan.
)

// [BlendFactor] is a structure specifying the normalized factor used to multiply pixel components.
//
// [BlendFactor]: https://wiki.libsdl.org/SDL3/SDL_BlendFactor
type BlendFactor uint32

const (
	BlendFactorZero             BlendFactor = iota + 1 // 0, 0, 0, 0.
	BlendFactorOne                                     // 1, 1, 1, 1.
	BlendFactorSrcColor                                // srcR, srcG, srcB, srcA.
	BlendFactorOneMinusSrcColor                        // 1-srcR, 1-srcG, 1-srcB, 1-srcA.
	BlendFactorSrcAlpha                                // srcA, srcA, srcA, srcA.
	BlendFactorOneMinusSrcAlpha                        // 1-srcA, 1-srcA, 1-srcA, 1-srcA.
	BlendFactorDstColor                                // dstR, dstG, dstB, dstA.
	BlendFactorOneMinusDstColor                        // 1-dstR, 1-dstG, 1-dstB, 1-dstA.
	BlendFactorDstAlpha                                // dstA, dstA, dstA, dstA.
	BlendFactorOneMinusDstAlpha                        // 1-dstA, 1-dstA, 1-dstA, 1-dstA.
)

// [BlendMode] is a set of blend modes used in drawing operations.
//
// [BlendMode]: https://wiki.libsdl.org/SDL3/SDL_BlendMode
type BlendMode uint32

const (
	BlendModeNone               = 0x00000000 // No blending: dstRGBA = srcRGBA.
	BlendModeBlend              = 0x00000001 // Alpha blending: dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)), dstA = srcA + (dstA * (1-srcA)).
	BlendModeBlendPremultiplied = 0x00000010 // Pre-multiplied alpha blending: dstRGBA = srcRGBA + (dstRGBA * (1-srcA)).
	BlendModeAdd                = 0x00000002 // Additive blending: dstRGB = (srcRGB * srcA) + dstRGB, dstA = dstA.
	BlendModeAddPremultiplied   = 0x00000020 // Pre-multiplied additive blending: dstRGB = srcRGB + dstRGB, dstA = dstA.
	BlendModeMod                = 0x00000004 // Color modulate: dstRGB = srcRGB * dstRGB, dstA = dstA.
	BlendModeMul                = 0x00000008 // Color multiply: dstRGB = (srcRGB * dstRGB) + (dstRGB * (1-srcA)), dstA = dstA.
	BlendModeInvalid            = 0x7FFFFFFF
)

// func ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
//	return sdlComposeCustomBlendMode(srcColorFactor, dstColorFactor, colorOperation, srcAlphaFactor, dstAlphaFactor, alphaOperation)
// }
