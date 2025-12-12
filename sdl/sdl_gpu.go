package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

// [GPUSwapchainComposition] is a structure specifying the texture format and colorspace of the swapchain textures.
//
// [GPUSwapchainComposition]: https://wiki.libsdl.org/SDL3/SDL_GPUSwapchainComposition
type GPUSwapchainComposition uint32

const (
	GPUSwapchainCompositionSDR               GPUSwapchainComposition = iota // B8G8R8A8 or R8G8B8A8 swapchain. Pixel values are in sRGB encoding
	GPUSwapchainCompositionSDRLinear                                        // B8G8R8A8_SRGB or R8G8B8A8_SRGB swapchain. Pixel values are stored in memory in sRGB encoding but accessed in shaders in "linear sRGB" encoding which is sRGB but with a linear transfer function.
	GPUSwapchainCompositionHDRExtendedLinear                                // R16G16B16A16_FLOAT swapchain. Pixel values are in extended linear sRGB encoding and permits values outside of the [0, 1] range.
	GPUSwapchainCompositionHDR10ST2084                                      // A2R10G10B10 or A2B10G10R10 swapchain. Pixel values are in BT.2020 ST2084 (PQ) encoding.
)

// [GPUPresentMode] is a structure specifying the timing that will be used to present swapchain textures to the OS.
//
// [GPUPresentMode]: https://wiki.libsdl.org/SDL3/SDL_GPUPresentMode
type GPUPresentMode uint32

const (
	GPUPresentModeVSync     GPUPresentMode = iota // Waits for vblank before presenting. No tearing is possible. If there is a pending image to present, the new image is enqueued for presentation. Disallows tearing at the cost of visual latency.
	GPUPresentModeImmediate                       // Immediately presents. Lowest latency option, but tearing may occur.
	GPUPresentModeMailbox                         // Waits for vblank before presenting. No tearing is possible. If there is a pending image to present, the pending image is replaced by the new image. Similar to VSYNC, but with reduced visual latency.
)

// [GPUSamplerAddressMode] is a structure specifying behavior of texture sampling when the coordinates exceed the 0-1 range.
//
// [GPUSamplerAddressMode]: https://wiki.libsdl.org/SDL3/SDL_GPUSamplerAddressMode
type GPUSamplerAddressMode uint32

const (
	GPUSamplerAddressModeRepeat         GPUSamplerAddressMode = iota // Specifies that the coordinates will wrap around.
	GPUSamplerAddressModeMirroredRepeat                              // Specifies that the coordinates will wrap around mirrored.
	GPUSamplerAddressModeClampToEdge                                 // Specifies that the coordinates will clamp to the 0-1 range.
)

// [GPUSamplerMipmapMode] is a structure specifying a mipmap mode used by a sampler.
//
// [GPUSamplerMipmapMode]: https://wiki.libsdl.org/SDL3/SDL_GPUSamplerMipmapMode
type GPUSamplerMipmapMode uint32

const (
	GPUSamplerMipmapModeNearest GPUSamplerMipmapMode = iota // Point filtering.
	GPUSamplerMipmapModeLinear                              // Linear filtering.
)

// [GPUFilter] is a structure specifying a filter operation used by a sampler.
//
// [GPUFilter]: https://wiki.libsdl.org/SDL3/SDL_GPUFilter
type GPUFilter uint32

const (
	GPUFilterNearest GPUFilter = iota // Point filtering.
	GPUFilterLinear                   // Linear filtering.
)

// [GPUBlendFactor] is a structure specifying a blending factor to be used when pixels in a render target are blended with existing pixels in the texture.
//
// [GPUBlendFactor]: https://wiki.libsdl.org/SDL3/SDL_GPUBlendFactor
type GPUBlendFactor uint32

const (
	GPUBlendFactorInvalid               GPUBlendFactor = iota
	GPUBlendFactorZero                                 // 0
	GPUBlendFactorOne                                  // 1
	GPUBlendFactorSrcColor                             // source color
	GPUBlendFactorOneMinusSrcColor                     // 1 - source color
	GPUBlendFactorDstColor                             // destination color
	GPUBlendFactorOneMinusDstColor                     // 1 - destination color
	GPUBlendFactorSrcAlpha                             // source alpha
	GPUBlendFactorOneMinusSrcAlpha                     // 1 - source alpha
	GPUBlendFactorDstAlpha                             // destination alpha
	GPUBlendFactorOneMinusDstAlpha                     // 1 - destination alpha
	GPUBlendFactorConstantColor                        // blend constant
	GPUBlendFactorOneMinusConstantColor                // 1 - blend constant
	GPUBlendFactorSrcAlphaSaturate                     // min(source alpha, 1 - destination alpha)
)

// [GPUColorComponentFlags] is a structure specifying which color components are written in a graphics pipeline.
//
// [GPUColorComponentFlags]: https://wiki.libsdl.org/SDL3/SDL_GPUColorComponentFlags
type GPUColorComponentFlags uint8

const (
	GPUColorComponentR GPUColorComponentFlags = 1 << 0 // The red component
	GPUColorComponentG GPUColorComponentFlags = 1 << 1 // The green component
	GPUColorComponentB GPUColorComponentFlags = 1 << 2 // The blue component
	GPUColorComponentA GPUColorComponentFlags = 1 << 3 // The alpha component
)

// [GPUBlendOp] is a structure specifying the operator to be used when pixels in a render target are blended with existing pixels in the texture.
//
// [GPUBlendOp]: https://wiki.libsdl.org/SDL3/SDL_GPUBlendOp
type GPUBlendOp uint32

const (
	GPUBlendOpInvalid         GPUBlendOp = iota
	GPUBlendOpAdd                        // (source * source_factor) + (destination * destination_factor)
	GPUBlendOpSubtract                   // (source * source_factor) - (destination * destination_factor)
	GPUBlendOpReverseSubtract            // (destination * destination_factor) - (source * source_factor)
	GPUBlendOpMin                        // min(source, destination)
	GPUBlendOpMax                        // max(source, destination)
)

// [GPUStencilOp] is a structure specifying what happens to a stored stencil value if stencil tests fail or pass.
//
// [GPUStencilOp]: https://wiki.libsdl.org/SDL3/SDL_GPUStencilOp
type GPUStencilOp uint32

const (
	GPUStencilOpInvalid           GPUStencilOp = iota
	GPUStencilOpKeep                           // Keeps the current value.
	GPUStencilOpZero                           // Sets the value to 0.
	GPUStencilOpReplace                        // Sets the value to reference.
	GPUStencilOpIncrementAndClamp              // Increments the current value and clamps to the maximum value.
	GPUStencilOpDecrementAndClamp              // Decrements the current value and clamps to 0.
	GPUStencilOpInvert                         // Bitwise-inverts the current value.
	GPUStencilOpIncrementAndWrap               // Increments the current value and wraps back to 0.
	GPUStencilOpDecrementAndWrap               // Decrements the current value and wraps to the maximum value.
)

// [GPUCompareOp] is a structure specifying a comparison operator for depth, stencil and sampler operations.
//
// [GPUCompareOp]: https://wiki.libsdl.org/SDL3/SDL_GPUCompareOp
type GPUCompareOp uint32

const (
	GPUCompareOpInvalid        GPUCompareOp = iota
	GPUCompareOpNever                       // The comparison always evaluates false.
	GPUCompareOpLess                        // The comparison evaluates reference < test.
	GPUCompareOpEqual                       // The comparison evaluates reference == test.
	GPUCompareOpLessOrEqual                 // The comparison evaluates reference <= test.
	GPUCompareOpGreater                     // The comparison evaluates reference > test.
	GPUCompareOpNotEqual                    // The comparison evaluates reference != test.
	GPUCompareOpGreaterOrEqual              // The comparison evaluates reference >= test.
	GPUCompareOpAlways                      // The comparison always evaluates true.
)

// [GPUFrontFace] is a structure specifying the vertex winding that will cause a triangle to be determined to be front-facing.
//
// [GPUFrontFace]: https://wiki.libsdl.org/SDL3/SDL_GPUFrontFace
type GPUFrontFace uint32

const (
	GPUFrontFaceCounterClockwise GPUFrontFace = iota // A triangle with counter-clockwise vertex winding will be considered front-facing.
	GPUFrontFaceClockwise                            // A triangle with clockwise vertex winding will be considered front-facing.
)

// [GPUCullMode] is a structure specifying the facing direction in which triangle faces will be culled.
//
// [GPUCullMode]: https://wiki.libsdl.org/SDL3/SDL_GPUCullMode
type GPUCullMode uint32

const (
	GPUCullModeNone  GPUCullMode = iota // No triangles are culled.
	GPUCullModeFront                    // Front-facing triangles are culled.
	GPUCullModeBack                     // Back-facing triangles are culled.
)

// [GPUFillMode] is a structure specifying the fill mode of the graphics pipeline.
//
// [GPUFillMode]: https://wiki.libsdl.org/SDL3/SDL_GPUFillMode
type GPUFillMode uint32

const (
	GPUFillModeFill GPUFillMode = iota // Polygons will be rendered via rasterization.
	GPUFillModeLine                    // Polygon edges will be drawn as line segments.
)

// [GPUVertexInputRate] is a structure specifying the rate at which vertex attributes are pulled from buffers.
//
// [GPUVertexInputRate]: https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputRate
type GPUVertexInputRate uint32

const (
	GPUVertexInputRateVertex   GPUVertexInputRate = iota // Attribute addressing is a function of the vertex index.
	GPUVertexInputRateInstance                           // Attribute addressing is a function of the instance index.
)

// [GPUVertexElementFormat] is a structure specifying the format of a vertex attribute.
//
// [GPUVertexElementFormat]: https://wiki.libsdl.org/SDL3/SDL_GPUVertexElementFormat
type GPUVertexElementFormat uint32

const (
	GPUVertexElementFormatInvalid GPUVertexElementFormat = iota
	// 32-bit Signed Integers
	GPUVertexElementFormatInt
	GPUVertexElementFormatInt2
	GPUVertexElementFormatInt3
	GPUVertexElementFormatInt4
	// 32-bit Unsigned Integers
	GPUVertexElementFormatUint
	GPUVertexElementFormatUint2
	GPUVertexElementFormatUint3
	GPUVertexElementFormatUint4
	// 32-bit Floats
	GPUVertexElementFormatFloat
	GPUVertexElementFormatFloat2
	GPUVertexElementFormatFloat3
	GPUVertexElementFormatFloat4
	// 8-bit Signed Integers
	GPUVertexElementFormatByte2
	GPUVertexElementFormatByte4
	// 8-bit Unsigned Integers
	GPUVertexElementFormatUbyte2
	GPUVertexElementFormatUbyte4
	// 8-bit Signed Normalized
	GPUVertexElementFormatByte2Norm
	GPUVertexElementFormatByte4Norm
	// 8-bit Unsigned Normalized
	GPUVertexElementFormatUbyte2Norm
	GPUVertexElementFormatUbyte4Norm
	// 16-bit Signed Integers
	GPUVertexElementFormatShort2
	GPUVertexElementFormatShort4
	// 16-bit Unsigned Integers
	GPUVertexElementFormatUshort2
	GPUVertexElementFormatUshort4
	// 16-bit Signed Normalized
	GPUVertexElementFormatShort2Norm
	GPUVertexElementFormatShort4Norm
	// 16-bit Unsigned Normalized
	GPUVertexElementFormatUshort2Norm
	GPUVertexElementFormatUshort4Norm
	// 16-bit Floats
	GPUVertexElementFormatHalf2
	GPUVertexElementFormatHalf4
)

// [GPUShaderStage] is a structure specifying which stage a shader program corresponds to.
//
// [GPUShaderStage]: https://wiki.libsdl.org/SDL3/SDL_GPUShaderStage
type GPUShaderStage uint32

const (
	GPUShaderStageVertex GPUShaderStage = iota
	GPUShaderStageFragment
)

// [GPUTransferBufferUsage] is a structure specifying how a transfer buffer is intended to be used by the client.
//
// [GPUTransferBufferUsage]: https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferUsage
type GPUTransferBufferUsage uint32

const (
	GPUTransferBufferUsageUpload GPUTransferBufferUsage = iota
	GPUTransferBufferUsageDownload
)

// [GPUCubeMapFace] is a structure specifying the face of a cube map.
//
// [GPUCubeMapFace]: https://wiki.libsdl.org/SDL3/SDL_GPUCubeMapFace
type GPUCubeMapFace uint32

const (
	GPUCubeMapFacePositiveX GPUCubeMapFace = iota
	GPUCubeMapFaceNegativeX
	GPUCubeMapFacePositiveY
	GPUCubeMapFaceNegativeY
	GPUCubeMapFacePositiveZ
	GPUCubeMapFaceNegativeZ
)

// [GPUSampleCount] is a structure specifying the sample count of a texture.
//
// [GPUSampleCount]: https://wiki.libsdl.org/SDL3/SDL_GPUSampleCount
type GPUSampleCount uint32

const (
	GPUSampleCount1 GPUSampleCount = iota // No multisampling.
	GPUSampleCount2                       // MSAA 2x
	GPUSampleCount4                       // MSAA 4x
	GPUSampleCount8                       // MSAA 8x
)

// [GPUTextureType] is a structure specifying the type of a texture.
//
// [GPUTextureType]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureType
type GPUTextureType uint32

const (
	GPUTextureType2D        GPUTextureType = iota // The texture is a 2-dimensional image.
	GPUTextureType2DArray                         // The texture is a 2-dimensional array image.
	GPUTextureType3D                              // The texture is a 3-dimensional image.
	GPUTextureTypeCube                            // The texture is a cube image.
	GPUTextureTypeCubeArray                       // The texture is a cube array image.
)

// [GPUTextureFormat] is a structure specifying the pixel format of a texture.
//
// [GPUTextureFormat]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormat
type GPUTextureFormat uint32

const (
	GPUTextureFormatInvalid GPUTextureFormat = iota
	// Unsigned Normalized Float Color Formats
	GPUTextureFormatA8Unorm
	GPUTextureFormatR8Unorm
	GPUTextureFormatR8G8Unorm
	GPUTextureFormatR8G8B8A8Unorm
	GPUTextureFormatR16Unorm
	GPUTextureFormatR16G16Unorm
	GPUTextureFormatR16G16B16A16Unorm
	GPUTextureFormatR10G10B10A2Unorm
	GPUTextureFormatB5G6R5Unorm
	GPUTextureFormatB5G5R5A1Unorm
	GPUTextureFormatB4G4R4A4Unorm
	GPUTextureFormatB8G8R8A8Unorm
	// Compressed Unsigned Normalized Float Color Formats
	GPUTextureFormatBC1RGBAUnorm
	GPUTextureFormatBC2RGBAUnorm
	GPUTextureFormatBC3RGBAUnorm
	GPUTextureFormatBC4RUnorm
	GPUTextureFormatBC5RGUnorm
	GPUTextureFormatBC7RGBAUnorm
	// Compressed Signed Float Color Formats
	GPUTextureFormatBC6HRGBFloat
	// Compressed Unsigned Float Color Formats
	GPUTextureFormatBC6HRGBUfloat
	// Signed Normalized Float Color Formats
	GPUTextureFormatR8Snorm
	GPUTextureFormatR8G8Snorm
	GPUTextureFormatR8G8B8A8Snorm
	GPUTextureFormatR16Snorm
	GPUTextureFormatR16G16Snorm
	GPUTextureFormatR16G16B16A16Snorm
	// Signed Float Color Formats
	GPUTextureFormatR16Float
	GPUTextureFormatR16G16Float
	GPUTextureFormatR16G16B16A16Float
	GPUTextureFormatR32Float
	GPUTextureFormatR32G32Float
	GPUTextureFormatR32G32B32A32Float
	// Unsigned Float Color Formats
	GPUTextureFormatR11G11B10Ufloat
	// Unsigned Integer Color Formats
	GPUTextureFormatR8Uint
	GPUTextureFormatR8G8Uint
	GPUTextureFormatR8G8B8A8Uint
	GPUTextureFormatR16Uint
	GPUTextureFormatR16G16Uint
	GPUTextureFormatR16G16B16A16Uint
	GPUTextureFormatR32Uint
	GPUTextureFormatR32G32Uint
	GPUTextureFormatR32G32B32A32Uint
	// Signed Integer Color Formats
	GPUTextureFormatR8Int
	GPUTextureFormatR8G8Int
	GPUTextureFormatR8G8B8A8Int
	GPUTextureFormatR16Int
	GPUTextureFormatR16G16Int
	GPUTextureFormatR16G16B16A16Int
	GPUTextureFormatR32Int
	GPUTextureFormatR32G32Int
	GPUTextureFormatR32G32B32A32Int
	// SRGB Unsigned Normalized Color Formats
	GPUTextureFormatR8G8B8A8UnormSRGB
	GPUTextureFormatB8G8R8A8UnormSRGB
	// Compressed SRGB Unsigned Normalized Color Formats
	GPUTextureFormatBC1RGBAUnormSRGB
	GPUTextureFormatBC2RGBAUnormSRGB
	GPUTextureFormatBC3RGBAUnormSRGB
	GPUTextureFormatBC7RGBAUnormSRGB
	// Depth Formats
	GPUTextureFormatD16Unorm
	GPUTextureFormatD24Unorm
	GPUTextureFormatD32Float
	GPUTextureFormatD24UnormS8Uint
	GPUTextureFormatD32FloatS8Uint
	// Compressed ASTC Normalized Float Color Formats
	GPUTextureFormatAstc4x4Unorm
	GPUTextureFormatAstc5x4Unorm
	GPUTextureFormatAstc5x5Unorm
	GPUTextureFormatAstc6x5Unorm
	GPUTextureFormatAstc6x6Unorm
	GPUTextureFormatAstc8x5Unorm
	GPUTextureFormatAstc8x6Unorm
	GPUTextureFormatAstc8x8Unorm
	GPUTextureFormatAstc10x5Unorm
	GPUTextureFormatAstc10x6Unorm
	GPUTextureFormatAstc10x8Unorm
	GPUTextureFormatAstc10x10Unorm
	GPUTextureFormatAstc12x10Unorm
	GPUTextureFormatAstc12x12Unorm
	// Compressed SRGB ASTC Normalized Float Color Formats
	GPUTextureFormatAstc4x4UnormSRGB
	GPUTextureFormatAstc5x4UnormSRGB
	GPUTextureFormatAstc5x5UnormSRGB
	GPUTextureFormatAstc6x5UnormSRGB
	GPUTextureFormatAstc6x6UnormSRGB
	GPUTextureFormatAstc8x5UnormSRGB
	GPUTextureFormatAstc8x6UnormSRGB
	GPUTextureFormatAstc8x8UnormSRGB
	GPUTextureFormatAstc10x5UnormSRGB
	GPUTextureFormatAstc10x6UnormSRGB
	GPUTextureFormatAstc10x8UnormSRGB
	GPUTextureFormatAstc10x10UnormSRGB
	GPUTextureFormatAstc12x10UnormSRGB
	GPUTextureFormatAstc12x12UnormSRGB
	// Compressed ASTC Signed Float Color Formats
	GPUTextureFormatAstc4x4Float
	GPUTextureFormatAstc5x4Float
	GPUTextureFormatAstc5x5Float
	GPUTextureFormatAstc6x5Float
	GPUTextureFormatAstc6x6Float
	GPUTextureFormatAstc8x5Float
	GPUTextureFormatAstc8x6Float
	GPUTextureFormatAstc8x8Float
	GPUTextureFormatAstc10x5Float
	GPUTextureFormatAstc10x6Float
	GPUTextureFormatAstc10x8Float
	GPUTextureFormatAstc10x10Float
	GPUTextureFormatAstc12x10Float
	GPUTextureFormatAstc12x12Float
)

// [GPUIndexElementSize] is a structure specifying the size of elements in an index buffer.
//
// [GPUIndexElementSize]: https://wiki.libsdl.org/SDL3/SDL_GPUIndexElementSize
type GPUIndexElementSize uint32

const (
	GPUIndexElementSize16Bit GPUIndexElementSize = iota // The index elements are 16-bit.
	GPUIndexElementSize32Bit                            // The index elements are 32-bit.
)

// [GPUStoreOp] is a structure specifying how the contents of a texture attached to a render pass are treated at the end of the render pass.
//
// [GPUStoreOp]: https://wiki.libsdl.org/SDL3/SDL_GPUStoreOp
type GPUStoreOp uint32

const (
	GPUStoreOpStore           GPUStoreOp = iota // Stores the results of the render pass in the texture. Not recommended for multisample textures as it requires significant memory bandwidth.
	GPUStoreOpDontCare                          // The driver will do whatever it wants with the texture memory. This is often a good option for depth/stencil textures.
	GPUStoreOpResolve                           // Resolves a multisample texture into resolve_texture, which must have a sample count of 1. Then the driver may discard the multisample texture memory. This is the most performant method of resolving a multisample target.
	GPUStoreOpResolveAndStore                   // Resolves a multisample texture into the resolve_texture, which must have a sample count of 1. Then the driver stores the multisample texture's contents. Not recommended as it requires significant memory bandwidth.
)

// [GPULoadOp] is a structure specifying how the contents of a texture attached to a render pass are treated at the beginning of the render pass.
//
// [GPULoadOp]: https://wiki.libsdl.org/SDL3/SDL_GPULoadOp
type GPULoadOp uint32

const (
	GPULoadOpLoad     GPULoadOp = iota // Loads the data currently in the texture. Not recommended for multisample textures as it requires significant memory bandwidth.
	GPULoadOpClear                     // Clears the texture to a single color.
	GPULoadOpDontCare                  // The driver will do whatever it wants with the texture memory. This is a good option if you know that every single pixel will be touched in the render pass.
)

// [GPUPrimitiveType] is a structure specifying the primitive topology of a graphics pipeline.
//
// [GPUPrimitiveType]: https://wiki.libsdl.org/SDL3/SDL_GPUPrimitiveType
type GPUPrimitiveType uint32

const (
	GPUPrimitiveTypeTrianglelist  GPUPrimitiveType = iota // A series of separate triangles.
	GPUPrimitiveTypeTrianglestrip                         // A series of connected triangles.
	GPUPrimitiveTypeLinelist                              // A series of separate lines.
	GPUPrimitiveTypeLinestrip                             // A series of connected lines.
	GPUPrimitiveTypePointlist                             // A series of separate points.
)

// [GPUShaderFormat] is a structure specifying the format of shader code.
//
// [GPUShaderFormat]: https://wiki.libsdl.org/SDL3/SDL_GPUShaderFormat
type GPUShaderFormat uint32

const (
	GPUShaderFormatInvalid  GPUShaderFormat = 0
	GPUShaderFormatPrivate  GPUShaderFormat = 1 << 0 // Shaders for NDA'd platforms.
	GPUShaderFormatSPIRV    GPUShaderFormat = 1 << 1 // SPIR-V shaders for Vulkan.
	GPUShaderFormatDXBC     GPUShaderFormat = 1 << 2 // DXBC SM5_1 shaders for D3D12.
	GPUShaderFormatDXIL     GPUShaderFormat = 1 << 3 // DXIL SM6_0 shaders for D3D12.
	GPUShaderFormatMSL      GPUShaderFormat = 1 << 4 // MSL shaders for Metal.
	GPUShaderFormatMetallib GPUShaderFormat = 1 << 5 // Precompiled metallib shaders for Metal.
)

// [GPUBufferUsageFlags] is a structure specifying how a buffer is intended to be used by the client.
//
// [GPUBufferUsageFlags]: https://wiki.libsdl.org/SDL3/SDL_GPUBufferUsageFlags
type GPUBufferUsageFlags uint32

const (
	GPUBufferUsageVertex              GPUBufferUsageFlags = 1 << 0 // Buffer is a vertex buffer.
	GPUBufferUsageIndex               GPUBufferUsageFlags = 1 << 1 // Buffer is an index buffer.
	GPUBufferUsageIndirect            GPUBufferUsageFlags = 1 << 2 // Buffer is an indirect buffer.
	GPUBufferUsageGraphicsStorageRead GPUBufferUsageFlags = 1 << 3 // Buffer supports storage reads in graphics stages.
	GPUBufferUsageComputeStorageRead  GPUBufferUsageFlags = 1 << 4 // Buffer supports storage reads in the compute stage.
	GPUBufferUsageComputeStorageWrite GPUBufferUsageFlags = 1 << 5 // Buffer supports storage writes in the compute stage.
)

// [GPUDevice] is an opaque handle representing the SDL_GPU context.
//
// [GPUDevice]: https://wiki.libsdl.org/SDL3/SDL_GPUDevice
type GPUDevice struct{}

// [GPUTexture] is an opaque handle representing a texture.
//
// [GPUTexture]: https://wiki.libsdl.org/SDL3/SDL_GPUTexture
type GPUTexture struct{}

// [GPUCommandBuffer] is an opaque handle representing a command buffer.
//
// [GPUCommandBuffer]: https://wiki.libsdl.org/SDL3/SDL_GPUCommandBuffer
type GPUCommandBuffer struct{}

// [GPURenderPass] is an opaque handle representing a render pass.
//
// [GPURenderPass]: https://wiki.libsdl.org/SDL3/SDL_GPURenderPass
type GPURenderPass struct{}

// [GPUColorTargetInfo] is a structure specifying the parameters of a color target used by a render pass.
//
// [GPUColorTargetInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetInfo
type GPUColorTargetInfo struct {
	Texture             *GPUTexture // The texture that will be used as a color target by a render pass.
	MipLevel            uint32      // The mip level to use as a color target.
	LayerOrDepthPlane   uint32      // The layer index or depth plane to use as a color target. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures.
	ClearColor          FColor      // The color to clear the color target to at the start of the render pass. Ignored if GPULoadOpClear is not used.
	LoadOp              GPULoadOp   // What is done with the contents of the color target at the beginning of the render pass.
	StoreOp             GPUStoreOp  // What is done with the results of the render pass.
	ResolveTexture      *GPUTexture // The texture that will receive the results of a multisample resolve operation. Ignored if a RESOLVE* store_op is not used.
	ResolveMipLevel     uint32      // The mip level of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used.
	ResolveLayer        uint32      // The layer index of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used.
	Cycle               bool        // true cycles the texture if the texture is bound and load_op is not LOAD
	CycleResolveTexture bool        // true cycles the resolve texture if the resolve texture is bound. Ignored if a RESOLVE* store_op is not used.
	_                   uint8       // padding1
	_                   uint8       // padding2
}

// [GPUDepthStencilTargetInfo] is a structure specifying the parameters of a depth-stencil target used by a render pass.
//
// [GPUDepthStencilTargetInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUDepthStencilTargetInfo
type GPUDepthStencilTargetInfo struct {
	Texture        *GPUTexture // The texture that will be used as the depth stencil target by the render pass.
	ClearDepth     float32     // The value to clear the depth component to at the beginning of the render pass. Ignored if GPU_LOADOP_CLEAR is not used.
	LoadOp         GPULoadOp   // What is done with the depth contents at the beginning of the render pass.
	StoreOp        GPUStoreOp  // What is done with the depth results of the render pass.
	StencilLoadOp  GPULoadOp   // What is done with the stencil contents at the beginning of the render pass.
	StencilStoreOp GPUStoreOp  // What is done with the stencil results of the render pass.
	Cycle          bool        // true cycles the texture if the texture is bound and any load ops are not LOAD
	ClearStencil   uint8       // The value to clear the stencil component to at the beginning of the render pass. Ignored if GPU_LOADOP_CLEAR is not used.
	MipLevel       uint8       // The mip level to use as the depth stencil target.
	Layer          uint8       // The layer index to use as the depth stencil target.
}

// [GPUShader] is an opaque handle representing a compiled shader object.
//
// [GPUShader]: https://wiki.libsdl.org/SDL3/SDL_GPUShader
type GPUShader struct{}

// [GPUVertexBufferDescription] is a structure specifying the parameters of vertex buffers used in a graphics pipeline.
//
// [GPUVertexBufferDescription]: https://wiki.libsdl.org/SDL3/SDL_GPUVertexBufferDescription
type GPUVertexBufferDescription struct {
	Slot             uint32             // The binding slot of the vertex buffer.
	Pitch            uint32             // The size of a single element + the offset between elements.
	InputRate        GPUVertexInputRate // Whether attribute addressing is a function of the vertex index or instance index.
	InstanceStepRate uint32             // Reserved for future use. Must be set to 0.
}

// [GPUVertexAttribute] is a structure specifying a vertex attribute.
//
// [GPUVertexAttribute]: https://wiki.libsdl.org/SDL3/SDL_GPUVertexAttribute
type GPUVertexAttribute struct {
	Location   uint32                 // The shader input location index.
	BufferSlot uint32                 // The binding slot of the associated vertex buffer.
	Format     GPUVertexElementFormat // The size and type of the attribute data.
	Offset     uint32                 // The byte offset of this attribute relative to the start of the vertex element.
}

// [GPUVertexInputState] is a structure specifying the parameters of a graphics pipeline vertex input state.
//
// [GPUVertexInputState]: https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputState
type GPUVertexInputState struct {
	VertexBufferDescriptions *GPUVertexBufferDescription // A pointer to an array of vertex buffer descriptions.
	NumVertexBuffers         uint32                      // The number of vertex buffer descriptions in the above array.
	VertexAttributes         *GPUVertexAttribute         // A pointer to an array of vertex attribute descriptions.
	NumVertexAttributes      uint32                      // The number of vertex attribute descriptions in the above array.
}

// [GPUStencilOpState] is a structure specifying the stencil operation state of a graphics pipeline.
//
// [GPUStencilOpState]: https://wiki.libsdl.org/SDL3/SDL_GPUStencilOpState
type GPUStencilOpState struct {
	FailOp      GPUStencilOp // The action performed on samples that fail the stencil test.
	PassOp      GPUStencilOp // The action performed on samples that pass the depth and stencil tests.
	DepthFailOp GPUStencilOp // The action performed on samples that pass the stencil test and fail the depth test.
	CompareOp   GPUCompareOp // The comparison operator used in the stencil test.
}

// [GPUColorTargetBlendState] is a structure specifying the blend state of a color target.
//
// [GPUColorTargetBlendState]: https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetBlendState
type GPUColorTargetBlendState struct {
	SrcColorBlendFactor  GPUBlendFactor         // The value to be multiplied by the source RGB value.
	DstColorBlendFactor  GPUBlendFactor         // The value to be multiplied by the destination RGB value.
	ColorBlendOp         GPUBlendOp             // The blend operation for the RGB components.
	SrcAlphaBlendFactor  GPUBlendFactor         // The value to be multiplied by the source alpha.
	DstAlphaBlendFactor  GPUBlendFactor         // The value to be multiplied by the destination alpha.
	AlphaBlendOp         GPUBlendOp             // The blend operation for the alpha component.
	ColorWriteMask       GPUColorComponentFlags // A bitmask specifying which of the RGBA components are enabled for writing. Writes to all channels if enable_color_write_mask is false.
	EnableBlend          bool                   // Whether blending is enabled for the color target.
	EnableColorWriteMask bool                   // Whether the color write mask is enabled.
	_                    uint8                  // padding1
	_                    uint8                  // padding2
}

// [GPUShaderCreateInfo] is a structure specifying code and metadata for creating a shader object.
//
// [GPUShaderCreateInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUShaderCreateInfo
type GPUShaderCreateInfo struct {
	CodeSize uint64 // The size in bytes of the code pointed to.
	Code     *uint8 // A pointer to shader code.
	// entryPoint must be a null-terminated C-style string.
	// Use the methods GPUShaderCreateInfo.EntryPoint and GPUShaderCreateInfo.SetEntryPoint instead.
	entryPoint         *byte
	Format             GPUShaderFormat // The format of the shader code.
	Stage              GPUShaderStage  // The stage the shader program corresponds to.
	NumSamplers        uint32          // The number of samplers defined in the shader.
	NumStorageTextures uint32          // The number of storage textures defined in the shader.
	NumStorageBuffers  uint32          // The number of storage buffers defined in the shader.
	NumUniformBuffers  uint32          // The number of uniform buffers defined in the shader.
	Props              PropertiesID    // A properties ID for extensions. Should be 0 if no extensions are needed.
}

func (createInfo *GPUShaderCreateInfo) EntryPoint() string {
	return convert.ToString(createInfo.entryPoint)
}

// SetEntryPoint specifies the entry point function name for the shader.
func (createInfo *GPUShaderCreateInfo) SetEntryPoint(entryPoint string) {
	createInfo.entryPoint = convert.ToBytePtr(entryPoint)
}

// [GPURasterizerState] is a structure specifying the parameters of the graphics pipeline rasterizer state.
//
// [GPURasterizerState]: https://wiki.libsdl.org/SDL3/SDL_GPURasterizerState
type GPURasterizerState struct {
	FillMode                GPUFillMode  // Whether polygons will be filled in or drawn as lines.
	CullMode                GPUCullMode  // The facing direction in which triangles will be culled.
	FrontFace               GPUFrontFace // The vertex winding that will cause a triangle to be determined as front-facing.
	DepthBiasConstantFactor float32      // A scalar factor controlling the depth value added to each fragment.
	DepthBiasClamp          float32      // The maximum depth bias of a fragment.
	DepthBiasSlopeFactor    float32      // A scalar factor applied to a fragment's slope in depth calculations.
	EnableDepthBias         bool         // true to bias fragment depth values.
	EnableDepthClip         bool         // true to enable depth clip, false to enable depth clamp.
	_                       uint8        // Padding1
	_                       uint8        // Padding2
}

// [GPUMultisampleState] is a structure specifying the parameters of the graphics pipeline multisample state.
//
// [GPUMultisampleState]: https://wiki.libsdl.org/SDL3/SDL_GPUMultisampleState
type GPUMultisampleState struct {
	SampleCount           GPUSampleCount // The number of samples to be used in rasterization.
	SampleMask            uint32         // Reserved for future use. Must be set to 0.
	EnableMask            bool           // Reserved for future use. Must be set to false.
	EnableAlphaToCoverage bool           // true enables the alpha-to-coverage feature.
	_                     uint8          // Padding2
	_                     uint8          // Padding3
}

// [GPUDepthStencilState] is a structure specifying the parameters of the graphics pipeline depth stencil state.
//
// [GPUDepthStencilState]: https://wiki.libsdl.org/SDL3/SDL_GPUDepthStencilState
type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp      // The comparison operator used for depth testing.
	BackStencilState  GPUStencilOpState // The stencil op state for back-facing triangles.
	FrontStencilState GPUStencilOpState // The stencil op state for front-facing triangles.
	CompareMask       uint8             // Selects the bits of the stencil values participating in the stencil test.
	WriteMask         uint8             // Selects the bits of the stencil values updated by the stencil test.
	EnableDepthTest   bool              // true enables the depth test.
	EnableDepthWrite  bool              // true enables depth writes. Depth writes are always disabled when enable_depth_test is false.
	EnableStencilTest bool              // true enables the stencil test.
	_                 uint8             // Padding1
	_                 uint8             // Padding2
	_                 uint8             // Padding3
}

// [GPUColorTargetDescription] is a structure specifying the parameters of color targets used in a graphics pipeline.
//
// [GPUColorTargetDescription]: https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetDescription
type GPUColorTargetDescription struct {
	Format     GPUTextureFormat         // The pixel format of the texture to be used as a color target.
	BlendState GPUColorTargetBlendState // The blend state to be used for the color target.
}

// [GPUGraphicsPipelineTargetInfo] is a structure specifying the descriptions of render targets used in a graphics pipeline.
//
// [GPUGraphicsPipelineTargetInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineTargetInfo
type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions *GPUColorTargetDescription // A pointer to an array of color target descriptions.
	NumColorTargets         uint32                     // The number of color target descriptions in the above array.
	DepthStencilFormat      GPUTextureFormat           // The pixel format of the depth-stencil target. Ignored if HasDepthStencilTarget is false.
	HasDepthStencilTarget   bool                       // true specifies that the pipeline uses a depth-stencil target.
	_                       uint8                      // padding1
	_                       uint8                      // padding2
	_                       uint8                      // padding3
}

// [GPUGraphicsPipelineCreateInfo] is a structure specifying the parameters of a graphics pipeline state.
//
// [GPUGraphicsPipelineCreateInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineCreateInfo
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader      *GPUShader                    // The vertex shader used by the graphics pipeline.
	FragmentShader    *GPUShader                    // The fragment shader used by the graphics pipeline.
	VertexInputState  GPUVertexInputState           // The vertex layout of the graphics pipeline.
	PrimitiveType     GPUPrimitiveType              // The primitive topology of the graphics pipeline.
	RasterizerState   GPURasterizerState            // The rasterizer state of the graphics pipeline.
	MultisampleState  GPUMultisampleState           // The multisample state of the graphics pipeline.
	DepthStencilState GPUDepthStencilState          // The depth-stencil state of the graphics pipeline.
	TargetInfo        GPUGraphicsPipelineTargetInfo // Formats and blend modes for the render targets of the graphics pipeline.
	Props             PropertiesID                  // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// [GPUGraphicsPipeline] is an opaque handle representing a graphics pipeline.
//
// [GPUGraphicsPipeline]: https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipeline
type GPUGraphicsPipeline struct{}

// [GPUViewport] is a structure specifying a viewport.
//
// [GPUViewport]: https://wiki.libsdl.org/SDL3/SDL_GPUViewport
type GPUViewport struct {
	X        float32 // The left offset of the viewport.
	Y        float32 // The top offset of the viewport.
	W        float32 // The width of the viewport.
	H        float32 // The height of the viewport.
	MinDepth float32 // The minimum depth of the viewport.
	MaxDepth float32 // The maximum depth of the viewport.
}

// [GPUBufferCreateInfo] is a structure specifying the parameters of a buffer.
//
// [GPUBufferCreateInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUBufferCreateInfo
type GPUBufferCreateInfo struct {
	Usage GPUBufferUsageFlags // How the buffer is intended to be used by the client.
	Size  uint32              // The size in bytes of the buffer.
	Props PropertiesID        // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// [GPUBuffer] is an opaque handle representing a buffer.
//
// [GPUBuffer]: https://wiki.libsdl.org/SDL3/SDL_GPUBuffer
type GPUBuffer struct{}

// [GPUTransferBufferCreateInfo] is a structure specifying the parameters of a transfer buffer.
//
// [GPUTransferBufferCreateInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferCreateInfo
type GPUTransferBufferCreateInfo struct {
	Usage GPUTransferBufferUsage // How the transfer buffer is intended to be used by the client.
	Size  uint32                 // The size in bytes of the transfer buffer.
	Props PropertiesID           // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// [GPUTransferBuffer] is an opaque handle representing a transfer buffer.
//
// [GPUTransferBuffer]: https://wiki.libsdl.org/SDL3/SDL_GPUTransferBuffer
type GPUTransferBuffer struct{}

// [GPUCopyPass] is an opaque handle representing a copy pass.
//
// [GPUCopyPass]: https://wiki.libsdl.org/SDL3/SDL_GPUCopyPass
type GPUCopyPass struct{}

// [GPUTransferBufferLocation] is a structure specifying a location in a transfer buffer.
//
// [GPUTransferBufferLocation]: https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferLocation
type GPUTransferBufferLocation struct {
	TransferBuffer *GPUTransferBuffer // The transfer buffer used in the transfer operation.
	Offset         uint32             // The starting byte of the buffer data in the transfer buffer.
}

// [GPUBufferRegion] is a structure specifying a region of a buffer.
//
// [GPUBufferRegion]: https://wiki.libsdl.org/SDL3/SDL_GPUBufferRegion
type GPUBufferRegion struct {
	Buffer *GPUBuffer // The buffer.
	Offset uint32     // The starting byte within the buffer.
	Size   uint32     // The size in bytes of the region.
}

// [GPUBufferBinding] is a structure specifying parameters in a buffer binding call.
//
// [GPUBufferBinding]: https://wiki.libsdl.org/SDL3/SDL_GPUBufferBinding
type GPUBufferBinding struct {
	Buffer *GPUBuffer // The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_VERTEX for SDL_BindGPUVertexBuffers, or SDL_GPU_BUFFERUSAGE_INDEX for SDL_BindGPUIndexBuffer.
	Offset uint32     // The starting byte of the data to bind in the buffer.
}

// [GPUTextureUsageFlags] is a structure specifing how a texture is intended to be used by the client.
//
// [GPUTextureUsageFlags]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureUsageFlags
type GPUTextureUsageFlags uint32

const (
	GPUTextureUsageSampler                             GPUTextureUsageFlags = 1 << 0 // Texture supports sampling.
	GPUTextureUsageColorTarget                         GPUTextureUsageFlags = 1 << 1 // Texture is a color render target.
	GPUTextureUsageDepthStencilTarget                  GPUTextureUsageFlags = 1 << 2 // Texture is a depth stencil target.
	GPUTextureUsageGraphicsStorageRead                 GPUTextureUsageFlags = 1 << 3 // Texture supports storage reads in graphics stages.
	GPUTextureUsageComputeStorageRead                  GPUTextureUsageFlags = 1 << 4 // Texture supports storage reads in the compute stage.
	GPUTextureUsageComputeStorageWrite                 GPUTextureUsageFlags = 1 << 5 // Texture supports storage writes in the compute stage.
	GPUTextureUsageComputeStorageSimultaneousReadWrite GPUTextureUsageFlags = 1 << 6 // Texture supports reads and writes in the same compute shader. This is NOT equivalent to READ | WRITE.
)

// [GPUTextureCreateInfo] is a structure specifying the parameters of a texture.
//
// [GPUTextureCreateInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureCreateInfo
type GPUTextureCreateInfo struct {
	Type              GPUTextureType       // The base dimensionality of the texture.
	Format            GPUTextureFormat     // The pixel format of the texture.
	Usage             GPUTextureUsageFlags // How the texture is intended to be used by the client.
	Width             uint32               // The width of the texture.
	Height            uint32               // The height of the texture.
	LayerCountOrDepth uint32               // The layer count or depth of the texture. This value is treated as a layer count on 2D array textures, and as a depth value on 3D textures.
	NumLevels         uint32               // The number of mip levels in the texture.
	SampleCount       GPUSampleCount       // The number of samples per texel. Only applies if the texture is used as a render target.
	Props             PropertiesID         // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// [GPUTextureTransferInfo] is a structure specifying parameters related to transferring data to or from a texture.
//
// [GPUTextureTransferInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureTransferInfo
type GPUTextureTransferInfo struct {
	TransferBuffer *GPUTransferBuffer // The transfer buffer used in the transfer operation.
	Offset         uint32             // The starting byte of the image data in the transfer buffer.
	PixelsPerRow   uint32             // The number of pixels from one row to the next.
	RowsPerLayer   uint32             // The number of rows from one layer/depth-slice to the next.
}

// [GPUTextureRegion] is a structure specifying a region of a texture.
//
// [GPUTextureRegion]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureRegion
type GPUTextureRegion struct {
	Texture  *GPUTexture // The texture used in the copy operation.
	MipLevel uint32      // The mip level index to transfer.
	Layer    uint32      // The layer index to transfer.
	X        uint32      // The left offset of the region.
	Y        uint32      // The top offset of the region.
	Z        uint32      // The front offset of the region.
	W        uint32      // The width of the region.
	H        uint32      // The height of the region.
	D        uint32      // The depth of the region.
}

// [GPUSamplerCreateInfo] is a structure specifying the parameters of a sampler.
//
// [GPUSamplerCreateInfo]: https://wiki.libsdl.org/SDL3/SDL_GPUSamplerCreateInfo
type GPUSamplerCreateInfo struct {
	MinFilter        GPUFilter             // The minification filter to apply to lookups.
	MagFilter        GPUFilter             // The magnification filter to apply to lookups.
	MipmapMode       GPUSamplerMipmapMode  // The mipmap filter to apply to lookups.
	AddressModeU     GPUSamplerAddressMode // The addressing mode for U coordinates outside [0, 1).
	AddressModeV     GPUSamplerAddressMode // The addressing mode for V coordinates outside [0, 1).
	AddressModeW     GPUSamplerAddressMode // The addressing mode for W coordinates outside [0, 1).
	MipLodBias       float32               // The bias to be added to mipmap LOD calculation.
	MaxAnisotropy    float32               // The anisotropy value clamp used by the sampler. If enable_anisotropy is false, this is ignored.
	CompareOp        GPUCompareOp          // The comparison operator to apply to fetched data before filtering.
	MinLod           float32               // Clamps the minimum of the computed LOD value.
	MaxLod           float32               // Clamps the maximum of the computed LOD value.
	EnableAnisotropy bool                  // true to enable anisotropic filtering.
	EnableCompare    bool                  // true to enable comparison against a reference value during lookups.
	_                uint8                 // Padding1
	_                uint8                 // Padding2
	Props            PropertiesID          // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// [GPUSampler] is an opaque handle representing a sampler.
//
// [GPUSampler]: https://wiki.libsdl.org/SDL3/SDL_GPUSampler
type GPUSampler struct{}

// [GPUTextureSamplerBinding] is a structure specifying parameters in a sampler binding call.
//
// [GPUTextureSamplerBinding]: https://wiki.libsdl.org/SDL3/SDL_GPUTextureSamplerBinding
type GPUTextureSamplerBinding struct {
	Texture *GPUTexture // The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
	Sampler *GPUSampler // The sampler to bind.
}

// [AcquireGPUCommandBuffer] acquires a command buffer.
//
// [AcquireGPUCommandBuffer]: https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer
func AcquireGPUCommandBuffer(device *GPUDevice) *GPUCommandBuffer {
	return sdlAcquireGPUCommandBuffer(device)
}

// [AcquireGPUSwapchainTexture] acquires a texture to use in presentation.
//
// [AcquireGPUSwapchainTexture]: https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture
func AcquireGPUSwapchainTexture(commandBuffer *GPUCommandBuffer, window *Window, swapchainTexture **GPUTexture, swapchainTextureWidth *uint32, swapchainTextureHeight *uint32) bool {
	return sdlAcquireGPUSwapchainTexture(commandBuffer, window, swapchainTexture, swapchainTextureWidth, swapchainTextureHeight)
}

// func BeginGPUComputePass(command_buffer *GPUCommandBuffer, storage_texture_bindings *GPUStorageTextureReadWriteBinding, num_storage_texture_bindings uint32, storage_buffer_bindings *GPUStorageBufferReadWriteBinding, num_storage_buffer_bindings uint32) *GPUComputePass {
//	return sdlBeginGPUComputePass(command_buffer, storage_texture_bindings, num_storage_texture_bindings, storage_buffer_bindings, num_storage_buffer_bindings)
// }

// [BeginGPUCopyPass] begins a copy pass on a command buffer.
//
// [BeginGPUCopyPass]: https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass
func BeginGPUCopyPass(commandBuffer *GPUCommandBuffer) *GPUCopyPass {
	return sdlBeginGPUCopyPass(commandBuffer)
}

// [BeginGPURenderPass] begins a render pass on a command buffer.
//
// [BeginGPURenderPass]: https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass
func BeginGPURenderPass(commandBuffer *GPUCommandBuffer, colorTargetInfos []GPUColorTargetInfo, depthStencilTargetInfo *GPUDepthStencilTargetInfo) *GPURenderPass {
	numColorTargets := uint32(len(colorTargetInfos))
	var colorTargetInfosPtr *GPUColorTargetInfo
	if numColorTargets > 0 {
		colorTargetInfosPtr = &colorTargetInfos[0]
	}
	return sdlBeginGPURenderPass(commandBuffer, colorTargetInfosPtr, numColorTargets, depthStencilTargetInfo)
}

// func BindGPUComputePipeline(compute_pass *GPUComputePass, compute_pipeline *GPUComputePipeline)  {
//	sdlBindGPUComputePipeline(compute_pass, compute_pipeline)
// }

// func BindGPUComputeSamplers(compute_pass *GPUComputePass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUComputeSamplers(compute_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// func BindGPUComputeStorageBuffers(compute_pass *GPUComputePass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32)  {
//	sdlBindGPUComputeStorageBuffers(compute_pass, first_slot, storage_buffers, num_bindings)
// }

// func BindGPUComputeStorageTextures(compute_pass *GPUComputePass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUComputeStorageTextures(compute_pass, first_slot, storage_textures, num_bindings)
// }

// [BindGPUFragmentSamplers] binds texture-sampler pairs for use on the fragment shader.
//
// [BindGPUFragmentSamplers]: https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers
func BindGPUFragmentSamplers(renderPass *GPURenderPass, firstSlot uint32, textureSamplerBindings *GPUTextureSamplerBinding, numBindings uint32) {
	sdlBindGPUFragmentSamplers(renderPass, firstSlot, textureSamplerBindings, numBindings)
}

// [BindGPUFragmentStorageBuffers] binds storage buffers for use on the fragment shader.
//
// [BindGPUFragmentStorageBuffers]: https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers
func BindGPUFragmentStorageBuffers(renderPass *GPURenderPass, firstSlot uint32, storageBuffers **GPUBuffer, numBindings uint32) {
	sdlBindGPUFragmentStorageBuffers(renderPass, firstSlot, storageBuffers, numBindings)
}

// func BindGPUFragmentStorageTextures(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUFragmentStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
// }

// [BindGPUGraphicsPipeline] binds a graphics pipeline on a render pass to be used in rendering.
//
// [BindGPUGraphicsPipeline]: https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline
func BindGPUGraphicsPipeline(renderPass *GPURenderPass, graphicsPipeline *GPUGraphicsPipeline) {
	sdlBindGPUGraphicsPipeline(renderPass, graphicsPipeline)
}

// [BindGPUIndexBuffer] binds an index buffer on a command buffer for use with subsequent draw calls.
//
// [BindGPUIndexBuffer]: https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer
func BindGPUIndexBuffer(renderPass *GPURenderPass, binding *GPUBufferBinding, indexElementSize GPUIndexElementSize) {
	sdlBindGPUIndexBuffer(renderPass, binding, indexElementSize)
}

// [BindGPUVertexBuffers] binds vertex buffers on a command buffer for use with subsequent draw calls.
//
// [BindGPUVertexBuffers]: https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers
func BindGPUVertexBuffers(renderPass *GPURenderPass, firstSlot uint32, bindings *GPUBufferBinding, numBindings uint32) {
	sdlBindGPUVertexBuffers(renderPass, firstSlot, bindings, numBindings)
}

// func BindGPUVertexSamplers(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUVertexSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// [BindGPUVertexStorageBuffers] binds storage buffers for use on the vertex shader.
//
// [BindGPUVertexStorageBuffers]: https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers
func BindGPUVertexStorageBuffers(renderPass *GPURenderPass, firstSlot uint32, storageBuffers **GPUBuffer, numBindings uint32) {
	sdlBindGPUVertexStorageBuffers(renderPass, firstSlot, storageBuffers, numBindings)
}

// func BindGPUVertexStorageTextures(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUVertexStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
// }

// func BlitGPUTexture(command_buffer *GPUCommandBuffer, info *GPUBlitInfo)  {
//	sdlBlitGPUTexture(command_buffer, info)
// }

// func CalculateGPUTextureFormatSize(format GPUTextureFormat, width uint32, height uint32, depth_or_layer_count uint32) uint32 {
//	return sdlCalculateGPUTextureFormatSize(format, width, height, depth_or_layer_count)
// }

// func CancelGPUCommandBuffer(command_buffer *GPUCommandBuffer) bool {
//	return sdlCancelGPUCommandBuffer(command_buffer)
// }

// [ClaimWindowForGPUDevice] claims a window, creating a swapchain structure for it.
//
// [ClaimWindowForGPUDevice]: https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice
func ClaimWindowForGPUDevice(device *GPUDevice, window *Window) bool {
	return sdlClaimWindowForGPUDevice(device, window)
}

// func CopyGPUBufferToBuffer(copy_pass *GPUCopyPass, source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool)  {
//	sdlCopyGPUBufferToBuffer(copy_pass, source, destination, size, cycle)
// }

// func CopyGPUTextureToTexture(copy_pass *GPUCopyPass, source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool)  {
//	sdlCopyGPUTextureToTexture(copy_pass, source, destination, w, h, d, cycle)
// }

// [CreateGPUBuffer] creates a buffer object to be used in graphics or compute workflows.
//
// [CreateGPUBuffer]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer
func CreateGPUBuffer(device *GPUDevice, createinfo *GPUBufferCreateInfo) *GPUBuffer {
	return sdlCreateGPUBuffer(device, createinfo)
}

// func CreateGPUComputePipeline(device *GPUDevice, createinfo *GPUComputePipelineCreateInfo) *GPUComputePipeline {
//	return sdlCreateGPUComputePipeline(device, createinfo)
// }

// [CreateGPUDevice] creates a GPU context.
//
// [CreateGPUDevice]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice
func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) *GPUDevice {
	return sdlCreateGPUDevice(formatFlags, debugMode, convert.ToBytePtrNullable(name))
}

// func CreateGPUDeviceWithProperties(props PropertiesID) *GPUDevice {
//	return sdlCreateGPUDeviceWithProperties(props)
// }

// [CreateGPUGraphicsPipeline] creates a pipeline object to be used in a graphics workflow.
//
// [CreateGPUGraphicsPipeline]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline
func CreateGPUGraphicsPipeline(device *GPUDevice, createInfo *GPUGraphicsPipelineCreateInfo) *GPUGraphicsPipeline {
	return sdlCreateGPUGraphicsPipeline(device, createInfo)
}

// [CreateGPUSampler] creates a sampler object to be used when binding textures in a graphics workflow.
//
// [CreateGPUSampler]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler
func CreateGPUSampler(device *GPUDevice, createinfo *GPUSamplerCreateInfo) *GPUSampler {
	return sdlCreateGPUSampler(device, createinfo)
}

// [CreateGPUShader] creates a shader to be used when creating a graphics pipeline.
//
// [CreateGPUShader]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader
func CreateGPUShader(device *GPUDevice, createInfo *GPUShaderCreateInfo) *GPUShader {
	return sdlCreateGPUShader(device, createInfo)
}

// [CreateGPUTexture] creates a texture object to be used in graphics or compute workflows.
//
// [CreateGPUTexture]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture
func CreateGPUTexture(device *GPUDevice, createinfo *GPUTextureCreateInfo) *GPUTexture {
	return sdlCreateGPUTexture(device, createinfo)
}

// [CreateGPUTransferBuffer] creates a transfer buffer to be used when uploading to or downloading from graphics resources.
//
// [CreateGPUTransferBuffer]: https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer
func CreateGPUTransferBuffer(device *GPUDevice, createinfo *GPUTransferBufferCreateInfo) *GPUTransferBuffer {
	return sdlCreateGPUTransferBuffer(device, createinfo)
}

// [DestroyGPUDevice] destroys a GPU context previously returned by CreateGPUDevice.
//
// [DestroyGPUDevice]: https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice
func DestroyGPUDevice(device *GPUDevice) {
	sdlDestroyGPUDevice(device)
}

// func DispatchGPUCompute(compute_pass *GPUComputePass, groupcount_x uint32, groupcount_y uint32, groupcount_z uint32)  {
//	sdlDispatchGPUCompute(compute_pass, groupcount_x, groupcount_y, groupcount_z)
// }

// func DispatchGPUComputeIndirect(compute_pass *GPUComputePass, buffer *GPUBuffer, offset uint32)  {
//	sdlDispatchGPUComputeIndirect(compute_pass, buffer, offset)
// }

// func DownloadFromGPUBuffer(copy_pass *GPUCopyPass, source *GPUBufferRegion, destination *GPUTransferBufferLocation)  {
//	sdlDownloadFromGPUBuffer(copy_pass, source, destination)
// }

// func DownloadFromGPUTexture(copy_pass *GPUCopyPass, source *GPUTextureRegion, destination *GPUTextureTransferInfo)  {
//	sdlDownloadFromGPUTexture(copy_pass, source, destination)
// }

// [DrawGPUIndexedPrimitives] draws data using bound graphics state with an index buffer and instancing enabled.
//
// [DrawGPUIndexedPrimitives]: https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives
func DrawGPUIndexedPrimitives(renderPass *GPURenderPass, numIndices uint32, numInstances uint32, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	sdlDrawGPUIndexedPrimitives(renderPass, numIndices, numInstances, firstIndex, vertexOffset, firstInstance)
}

// func DrawGPUIndexedPrimitivesIndirect(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32)  {
//	sdlDrawGPUIndexedPrimitivesIndirect(render_pass, buffer, offset, draw_count)
// }

// [DrawGPUPrimitives] draws data using bound graphics state.
//
// [DrawGPUPrimitives]: https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives
func DrawGPUPrimitives(renderPass *GPURenderPass, numVertices uint32, numInstances uint32, firstVertex uint32, firstInstance uint32) {
	sdlDrawGPUPrimitives(renderPass, numVertices, numInstances, firstVertex, firstInstance)
}

// func DrawGPUPrimitivesIndirect(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32)  {
//	sdlDrawGPUPrimitivesIndirect(render_pass, buffer, offset, draw_count)
// }

// func EndGPUComputePass(compute_pass *GPUComputePass)  {
//	sdlEndGPUComputePass(compute_pass)
// }

// [EndGPUCopyPass] ends the current copy pass.
//
// [EndGPUCopyPass]: https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass
func EndGPUCopyPass(copyPass *GPUCopyPass) {
	sdlEndGPUCopyPass(copyPass)
}

// [EndGPURenderPass] ends the given render pass.
//
// [EndGPURenderPass]: https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass
func EndGPURenderPass(renderPass *GPURenderPass) {
	sdlEndGPURenderPass(renderPass)
}

// func GenerateMipmapsForGPUTexture(command_buffer *GPUCommandBuffer, texture *GPUTexture)  {
//	sdlGenerateMipmapsForGPUTexture(command_buffer, texture)
// }

// [GetGPUDeviceDriver] returns the name of the backend used to create this GPU context.
//
// [GetGPUDeviceDriver]: https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver
func GetGPUDeviceDriver(device *GPUDevice) string {
	return sdlGetGPUDeviceDriver(device)
}

// [GetGPUDriver] gets the name of a built in GPU driver.
//
// [GetGPUDriver]: https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver
func GetGPUDriver(index int32) string {
	return sdlGetGPUDriver(index)
}

// [GetGPUShaderFormats] returns the supported shader formats for this GPU context.
//
// [GetGPUShaderFormats]: https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats
func GetGPUShaderFormats(device *GPUDevice) GPUShaderFormat {
	return sdlGetGPUShaderFormats(device)
}

// [GetGPUSwapchainTextureFormat] obtains the texture format of the swapchain for the given window.
//
// [GetGPUSwapchainTextureFormat]: https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat
func GetGPUSwapchainTextureFormat(device *GPUDevice, window *Window) GPUTextureFormat {
	return sdlGetGPUSwapchainTextureFormat(device, window)
}

// [GetNumGPUDrivers] gets the number of GPU drivers compiled into SDL.
//
// [GetNumGPUDrivers]: https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers
func GetNumGPUDrivers() int32 {
	return sdlGetNumGPUDrivers()
}

// func GPUSupportsProperties(props PropertiesID) bool {
//	return sdlGPUSupportsProperties(props)
// }

// func GPUSupportsShaderFormats(format_flags GPUShaderFormat, name string) bool {
//	return sdlGPUSupportsShaderFormats(format_flags, name)
// }

// func GPUTextureFormatTexelBlockSize(format GPUTextureFormat) uint32 {
//	return sdlGPUTextureFormatTexelBlockSize(format)
// }

// func GPUTextureSupportsFormat(device *GPUDevice, format GPUTextureFormat, type GPUTextureType, usage GPUTextureUsageFlags) bool {
//	return sdlGPUTextureSupportsFormat(device, format, type, usage)
// }

// func GPUTextureSupportsSampleCount(device *GPUDevice, format GPUTextureFormat, sample_count GPUSampleCount) bool {
//	return sdlGPUTextureSupportsSampleCount(device, format, sample_count)
// }

// func InsertGPUDebugLabel(command_buffer *GPUCommandBuffer, text string)  {
//	sdlInsertGPUDebugLabel(command_buffer, text)
// }

// [MapGPUTransferBuffer] maps a transfer buffer into application address space.
//
// [MapGPUTransferBuffer]: https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer
func MapGPUTransferBuffer(device *GPUDevice, transferBuffer *GPUTransferBuffer, cycle bool) unsafe.Pointer {
	return sdlMapGPUTransferBuffer(device, transferBuffer, cycle)
}

// func PopGPUDebugGroup(command_buffer *GPUCommandBuffer)  {
//	sdlPopGPUDebugGroup(command_buffer)
// }

// func PushGPUComputeUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUComputeUniformData(command_buffer, slot_index, data, length)
// }

// func PushGPUDebugGroup(command_buffer *GPUCommandBuffer, name string)  {
//	sdlPushGPUDebugGroup(command_buffer, name)
// }

// [PushGPUFragmentUniformData] pushes data to a fragment uniform slot on the command buffer.
//
// [PushGPUFragmentUniformData]: https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData
func PushGPUFragmentUniformData(commandBuffer *GPUCommandBuffer, slotIndex uint32, data unsafe.Pointer, length uint32) {
	sdlPushGPUFragmentUniformData(commandBuffer, slotIndex, data, length)
}

// [PushGPUVertexUniformData] pushes data to a vertex uniform slot on the command buffer.
//
// [PushGPUVertexUniformData]: https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData
func PushGPUVertexUniformData(commandBuffer *GPUCommandBuffer, slotIndex uint32, data unsafe.Pointer, length uint32) {
	sdlPushGPUVertexUniformData(commandBuffer, slotIndex, data, length)
}

// func QueryGPUFence(device *GPUDevice, fence *GPUFence) bool {
//	return sdlQueryGPUFence(device, fence)
// }

// [ReleaseGPUBuffer] frees the given buffer as soon as it is safe to do so.
//
// [ReleaseGPUBuffer]: https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer
func ReleaseGPUBuffer(device *GPUDevice, buffer *GPUBuffer) {
	sdlReleaseGPUBuffer(device, buffer)
}

// func ReleaseGPUComputePipeline(device *GPUDevice, compute_pipeline *GPUComputePipeline)  {
//	sdlReleaseGPUComputePipeline(device, compute_pipeline)
// }

// func ReleaseGPUFence(device *GPUDevice, fence *GPUFence)  {
//	sdlReleaseGPUFence(device, fence)
// }

// [ReleaseGPUGraphicsPipeline] frees the given graphics pipeline as soon as it is safe to do so.
//
// [ReleaseGPUGraphicsPipeline]: https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline
func ReleaseGPUGraphicsPipeline(device *GPUDevice, graphicsPipeline *GPUGraphicsPipeline) {
	sdlReleaseGPUGraphicsPipeline(device, graphicsPipeline)
}

// [ReleaseGPUSampler] frees the given sampler as soon as it is safe to do so.
//
// [ReleaseGPUSampler]: https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler
func ReleaseGPUSampler(device *GPUDevice, sampler *GPUSampler) {
	sdlReleaseGPUSampler(device, sampler)
}

// [ReleaseGPUShader] frees the given shader as soon as it is safe to do so.
//
// [ReleaseGPUShader]: https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader
func ReleaseGPUShader(device *GPUDevice, shader *GPUShader) {
	sdlReleaseGPUShader(device, shader)
}

// [ReleaseGPUTexture] Frees the given texture as soon as it is safe to do so.
//
// [ReleaseGPUTexture]: https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture
func ReleaseGPUTexture(device *GPUDevice, texture *GPUTexture) {
	sdlReleaseGPUTexture(device, texture)
}

// [ReleaseGPUTransferBuffer] frees the given transfer buffer as soon as it is safe to do so.
//
// [ReleaseGPUTransferBuffer]: https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer
func ReleaseGPUTransferBuffer(device *GPUDevice, transferBuffer *GPUTransferBuffer) {
	sdlReleaseGPUTransferBuffer(device, transferBuffer)
}

// [ReleaseWindowFromGPUDevice] unclaims a window, destroying its swapchain structure.
//
// [ReleaseWindowFromGPUDevice]: https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice
func ReleaseWindowFromGPUDevice(device *GPUDevice, window *Window) {
	sdlReleaseWindowFromGPUDevice(device, window)
}

// func SetGPUAllowedFramesInFlight(device *GPUDevice, allowed_frames_in_flight uint32) bool {
//	return sdlSetGPUAllowedFramesInFlight(device, allowed_frames_in_flight)
// }

// func SetGPUBlendConstants(render_pass *GPURenderPass, blend_constants FColor)  {
//	sdlSetGPUBlendConstants(render_pass, blend_constants)
// }

// [SetGPUBufferName] sets an arbitrary string constant to label a buffer.
//
// [SetGPUBufferName]: https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName
func SetGPUBufferName(device *GPUDevice, buffer *GPUBuffer, text string) {
	sdlSetGPUBufferName(device, buffer, text)
}

// [SetGPUScissor] sets the current scissor state on a command buffer.
//
// [SetGPUScissor]: https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor
func SetGPUScissor(renderPass *GPURenderPass, scissor *Rect) {
	sdlSetGPUScissor(renderPass, scissor)
}

// func SetGPUStencilReference(render_pass *GPURenderPass, reference uint8)  {
//	sdlSetGPUStencilReference(render_pass, reference)
// }

// [SetGPUSwapchainParameters] changes the swapchain parameters for the given claimed window.
//
// [SetGPUSwapchainParameters]: https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters
func SetGPUSwapchainParameters(device *GPUDevice, window *Window, swapchainComposition GPUSwapchainComposition, presentMode GPUPresentMode) bool {
	return sdlSetGPUSwapchainParameters(device, window, swapchainComposition, presentMode)
}

// func SetGPUTextureName(device *GPUDevice, texture *GPUTexture, text string)  {
//	sdlSetGPUTextureName(device, texture, text)
// }

// [SetGPUViewport] sets the current viewport state on a command buffer.
//
// [SetGPUViewport]: https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport
func SetGPUViewport(renderass *GPURenderPass, viewport *GPUViewport) {
	sdlSetGPUViewport(renderass, viewport)
}

// [SubmitGPUCommandBuffer] submits a command buffer so its commands can be processed on the GPU.
//
// [SubmitGPUCommandBuffer]: https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer
func SubmitGPUCommandBuffer(commandBuffer *GPUCommandBuffer) bool {
	return sdlSubmitGPUCommandBuffer(commandBuffer)
}

// func SubmitGPUCommandBufferAndAcquireFence(command_buffer *GPUCommandBuffer) *GPUFence {
//	return sdlSubmitGPUCommandBufferAndAcquireFence(command_buffer)
// }

// [UnmapGPUTransferBuffer] unmaps a previously mapped transfer buffer.
//
// [UnmapGPUTransferBuffer]: https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer
func UnmapGPUTransferBuffer(device *GPUDevice, transferBuffer *GPUTransferBuffer) {
	sdlUnmapGPUTransferBuffer(device, transferBuffer)
}

// [UploadToGPUBuffer] uploads data from a transfer buffer to a buffer.
//
// [UploadToGPUBuffer]: https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer
func UploadToGPUBuffer(copyPass *GPUCopyPass, source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
	sdlUploadToGPUBuffer(copyPass, source, destination, cycle)
}

// [UploadToGPUTexture] uploads data from a transfer buffer to a texture.
//
// [UploadToGPUTexture]: https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture
func UploadToGPUTexture(copyPass *GPUCopyPass, source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
	sdlUploadToGPUTexture(copyPass, source, destination, cycle)
}

// [WaitAndAcquireGPUSwapchainTexture] blocks the thread until a swapchain texture is available to be acquired, and then acquires it.
//
// [WaitAndAcquireGPUSwapchainTexture]: https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture
func WaitAndAcquireGPUSwapchainTexture(commandBuffer *GPUCommandBuffer, window *Window, swapchainTexture **GPUTexture, swapchainTextureWidth *uint32, swapchainTextureHeight *uint32) bool {
	return sdlWaitAndAcquireGPUSwapchainTexture(commandBuffer, window, swapchainTexture, swapchainTextureWidth, swapchainTextureHeight)
}

// func WaitForGPUFences(device *GPUDevice, wait_all bool, fences **GPUFence, num_fences uint32) bool {
//	return sdlWaitForGPUFences(device, wait_all, fences, num_fences)
// }

// func WaitForGPUIdle(device *GPUDevice) bool {
//	return sdlWaitForGPUIdle(device)
// }

// func WaitForGPUSwapchain(device *GPUDevice, window *Window) bool {
//	return sdlWaitForGPUSwapchain(device, window)
// }

// [WindowSupportsGPUPresentMode] determines whether a presentation mode is supported by the window.
//
// [WindowSupportsGPUPresentMode]: https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode
func WindowSupportsGPUPresentMode(device *GPUDevice, window *Window, presentMode GPUPresentMode) bool {
	return sdlWindowSupportsGPUPresentMode(device, window, presentMode)
}

// func WindowSupportsGPUSwapchainComposition(device *GPUDevice, window *Window, swapchain_composition GPUSwapchainComposition) bool {
//	return sdlWindowSupportsGPUSwapchainComposition(device, window, swapchain_composition)
// }
