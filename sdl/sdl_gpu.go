package sdl

import "github.com/jupiterrider/purego-sdl3/internal/convert"

type GPUSwapchainComposition uint32

const (
	GPUSwapchainCompositionSDR GPUSwapchainComposition = iota
	GPUSwapchainCompositionSDRLinear
	GPUSwapchainCompositionHDRExtendedLinear
	GPUSwapchainCompositionHDR10ST2084
)

type GPUPresentMode uint32

const (
	GPUPresentModeVSync GPUPresentMode = iota
	GPUPresentModeImmediate
	GPUPresentModeMailbox
)

type GPUSamplerAddressMode uint32

const (
	GPUSamplerAddressModeRepeat GPUSamplerAddressMode = iota
	GPUSamplerAddressModeMirroredRepeat
	GPUSamplerAddressModeClampToEdge
)

type GPUSamplerMipmapMode uint32

const (
	GPUSamplerMipmapModeNearest GPUSamplerMipmapMode = iota
	GPUSamplerMipmapModeLinear
)

type GPUFilter uint32

const (
	GPUFilterNearest GPUFilter = iota
	GPUFilterLinear
)

type GPUBlendFactor uint32

const (
	GPUBlendFactorInvalid GPUBlendFactor = iota
	GPUBlendFactorZero
	GPUBlendFactorOne
	GPUBlendFactorSrcColor
	GPUBlendFactorOneMinusSrcColor
	GPUBlendFactorDstColor
	GPUBlendFactorOneMinusDstColor
	GPUBlendFactorSrcAlpha
	GPUBlendFactorOneMinusSrcAlpha
	GPUBlendFactorDstAlpha
	GPUBlendFactorOneMinusDstAlpha
	GPUBlendFactorConstantColor
	GPUBlendFactorOneMinusConstantColor
	GPUBlendFactorSrcAlphaSaturate
)

type GPUColorComponentFlags uint8

const (
	GPUColorComponentR GPUColorComponentFlags = 1 << 0
	GPUColorComponentG GPUColorComponentFlags = 1 << 1
	GPUColorComponentB GPUColorComponentFlags = 1 << 2
	GPUColorComponentA GPUColorComponentFlags = 1 << 3
)

type GPUBlendOp uint32

const (
	GPUBlendOpInvalid GPUBlendOp = iota
	GPUBlendOpAdd
	GPUBlendOpSubtract
	GPUBlendOpReverseSubtract
	GPUBlendOpMin
	GPUBlendOpMax
)

type GPUStencilOp uint32

const (
	GPUStencilOpInvalid GPUStencilOp = iota
	GPUStencilOpKeep
	GPUStencilOpZero
	GPUStencilOpReplace
	GPUStencilOpIncrementAndClamp
	GPUStencilOpDecrementAndClamp
	GPUStencilOpInvert
	GPUStencilOpIncrementAndWrap
	GPUStencilOpDecrementAndWrap
)

type GPUCompareOp uint32

const (
	GPUCompareOpInvalid GPUCompareOp = iota
	GPUCompareOpNever
	GPUCompareOpLess
	GPUCompareOpEqual
	GPUCompareOpLessOrEqual
	GPUCompareOpGreater
	GPUCompareOpNotEqual
	GPUCompareOpGreaterOrEqual
	GPUCompareOpAlways
)

type GPUFrontFace uint32

const (
	GPUFrontFaceCounterClockwise GPUFrontFace = iota
	GPUFrontFaceClockwise
)

type GPUCullMode uint32

const (
	GPUCullModeNone GPUCullMode = iota
	GPUCullModeFront
	GPUCullModeBack
)

type GPUFillMode uint32

const (
	GPUFillModeFill GPUFillMode = iota
	GPUFillModeLine
)

type GPUVertexInputRate uint32

const (
	GPUVertexInputRateVertex GPUVertexInputRate = iota
	GPUVertexInputRateInstance
)

type GPUVertexElementFormat uint32

const (
	GPUVertexElementFormatInvalid GPUVertexElementFormat = iota
	GPUVertexElementFormatInt
	GPUVertexElementFormatInt2
	GPUVertexElementFormatInt3
	GPUVertexElementFormatInt4
	GPUVertexElementFormatUint
	GPUVertexElementFormatUint2
	GPUVertexElementFormatUint3
	GPUVertexElementFormatUint4
	GPUVertexElementFormatFloat
	GPUVertexElementFormatFloat2
	GPUVertexElementFormatFloat3
	GPUVertexElementFormatFloat4
	GPUVertexElementFormatByte2
	GPUVertexElementFormatByte4
	GPUVertexElementFormatUbyte2
	GPUVertexElementFormatUbyte4
	GPUVertexElementFormatByte2Norm
	GPUVertexElementFormatByte4Norm
	GPUVertexElementFormatUbyte2Norm
	GPUVertexElementFormatUbyte4Norm
	GPUVertexElementFormatShort2
	GPUVertexElementFormatShort4
	GPUVertexElementFormatUshort2
	GPUVertexElementFormatUshort4
	GPUVertexElementFormatShort2Norm
	GPUVertexElementFormatShort4Norm
	GPUVertexElementFormatUshort2Norm
	GPUVertexElementFormatUshort4Norm
	GPUVertexElementFormatHalf2
	GPUVertexElementFormatHalf4
)

type GPUShaderStage uint32

const (
	GPUShaderStageVertex GPUShaderStage = iota
	GPUShaderStageFragment
)

type GPUTransferBufferUsage uint32

const (
	GPUTransferBufferUsageUpload GPUTransferBufferUsage = iota
	GPUTransferBufferUsageDownload
)

type GPUCubeMapFace uint32

const (
	GPUCubeMapFacePositiveX GPUCubeMapFace = iota
	GPUCubeMapFaceNegativeX
	GPUCubeMapFacePositiveY
	GPUCubeMapFaceNegativeY
	GPUCubeMapFacePositiveZ
	GPUCubeMapFaceNegativeZ
)

type GPUSampleCount uint32

const (
	GPUSampleCount1 GPUSampleCount = iota
	GPUSampleCount2
	GPUSampleCount4
	GPUSampleCount8
)

type GPUTextureType uint32

const (
	GPUTextureType2D GPUTextureType = iota
	GPUTextureType2DArray
	GPUTextureType3D
	GPUTextureTypeCube
	GPUTextureTypeCubeArray
)

type GPUTextureFormat uint32

const (
	GPUTextureFormatInvalid GPUTextureFormat = iota
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
	GPUTextureFormatBC1RGBAUnorm
	GPUTextureFormatBC2RGBAUnorm
	GPUTextureFormatBC3RGBAUnorm
	GPUTextureFormatBC4RUnorm
	GPUTextureFormatBC5RGUnorm
	GPUTextureFormatBC7RGBAUnorm
	GPUTextureFormatBC6HRGBFloat
	GPUTextureFormatBC6HRGBUfloat
	GPUTextureFormatR8Snorm
	GPUTextureFormatR8G8Snorm
	GPUTextureFormatR8G8B8A8Snorm
	GPUTextureFormatR16Snorm
	GPUTextureFormatR16G16Snorm
	GPUTextureFormatR16G16B16A16Snorm
	GPUTextureFormatR16Float
	GPUTextureFormatR16G16Float
	GPUTextureFormatR16G16B16A16Float
	GPUTextureFormatR32Float
	GPUTextureFormatR32G32Float
	GPUTextureFormatR32G32B32A32Float
	GPUTextureFormatR11G11B10Ufloat
	GPUTextureFormatR8Uint
	GPUTextureFormatR8G8Uint
	GPUTextureFormatR8G8B8A8Uint
	GPUTextureFormatR16Uint
	GPUTextureFormatR16G16Uint
	GPUTextureFormatR16G16B16A16Uint
	GPUTextureFormatR32Uint
	GPUTextureFormatR32G32Uint
	GPUTextureFormatR32G32B32A32Uint
	GPUTextureFormatR8Int
	GPUTextureFormatR8G8Int
	GPUTextureFormatR8G8B8A8Int
	GPUTextureFormatR16Int
	GPUTextureFormatR16G16Int
	GPUTextureFormatR16G16B16A16Int
	GPUTextureFormatR32Int
	GPUTextureFormatR32G32Int
	GPUTextureFormatR32G32B32A32Int
	GPUTextureFormatR8G8B8A8UnormSRGB
	GPUTextureFormatB8G8R8A8UnormSRGB
	GPUTextureFormatBC1RGBAUnormSRGB
	GPUTextureFormatBC2RGBAUnormSRGB
	GPUTextureFormatBC3RGBAUnormSRGB
	GPUTextureFormatBC7RGBAUnormSRGB
	GPUTextureFormatD16Unorm
	GPUTextureFormatD24Unorm
	GPUTextureFormatD32Float
	GPUTextureFormatD24UnormS8Uint
	GPUTextureFormatD32FloatS8Uint
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

type GPUIndexElementSize uint32

const (
	GPUIndexElementSize16Bit GPUIndexElementSize = iota
	GPUIndexElementSize32Bit
)

type GPUStoreOp uint32

const (
	GPUStoreOpStore GPUStoreOp = iota
	GPUStoreOpDontCare
	GPUStoreOpResolve
	GPUStoreOpResolveAndStore
)

type GPULoadOp uint32

const (
	GPULoadOpLoad GPULoadOp = iota
	GPULoadOpClear
	GPULoadOpDontCare
)

type GPUPrimitiveType uint32

const (
	GPUPrimitiveTypeTrianglelist GPUPrimitiveType = iota
	GPUPrimitiveTypeTrianglestrip
	GPUPrimitiveTypeLinelist
	GPUPrimitiveTypeLinestrip
	GPUPrimitiveTypePointlist
)

type GPUShaderFormat uint32

const (
	GPUShaderFormatInvalid  GPUShaderFormat = 0
	GPUShaderFormatPrivate  GPUShaderFormat = 1 << 0
	GPUShaderFormatSPIRV    GPUShaderFormat = 1 << 1
	GPUShaderFormatDXBC     GPUShaderFormat = 1 << 2
	GPUShaderFormatDXIL     GPUShaderFormat = 1 << 3
	GPUShaderFormatMSL      GPUShaderFormat = 1 << 4
	GPUShaderFormatMetallib GPUShaderFormat = 1 << 5
)

type GPUDevice struct{}

type GPUTexture struct{}

type GPUCommandBuffer struct{}

type GPURenderPass struct{}

type GPUColorTargetInfo struct {
	Texture             *GPUTexture
	MipLevel            uint32
	LayerOrDepthPlane   uint32
	ClearColor          FColor
	LoadOp              GPULoadOp
	StoreOp             GPUStoreOp
	ResolveTexture      *GPUTexture
	ResolveMipLevel     uint32
	ResolveLayer        uint32
	Cycle               bool
	CycleResolveTexture bool
	Padding1            uint8
	Padding2            uint8
}

type GPUDepthStencilTargetInfo struct {
	Texture        *GPUTexture
	ClearDepth     float32
	LoadOp         GPULoadOp
	StoreOp        GPUStoreOp
	StencilLoadOp  GPULoadOp
	StencilStoreOp GPUStoreOp
	Cycle          bool
	ClearStencil   uint8
	Padding1       uint8
	Padding2       uint8
}

type GPUShader struct{}

type GPUVertexBufferDescription struct {
	Slot             uint32
	Pitch            uint32
	InputRate        GPUVertexInputRate
	InstanceStepRate uint32
}

type GPUVertexAttribute struct {
	Location   uint32
	BufferSlot uint32
	Format     GPUVertexElementFormat
	Offset     uint32
}

type GPUVertexInputState struct {
	VertexBufferDescriptions *GPUVertexBufferDescription
	NumVertexBuffers         uint32
	VertexAttributes         *GPUVertexAttribute
	NumVertexAttributes      uint32
}

type GPUStencilOpState struct {
	FailOp      GPUStencilOp
	PassOp      GPUStencilOp
	DepthFailOp GPUStencilOp
	CompareOp   GPUCompareOp
}

type GPUColorTargetBlendState struct {
	SrcColorBlendFactor  GPUBlendFactor
	DstColorBlendFactor  GPUBlendFactor
	ColorBlendOp         GPUBlendOp
	SrcAlphaBlendFactor  GPUBlendFactor
	DstAlphaBlendFactor  GPUBlendFactor
	AlphaBlendOp         GPUBlendOp
	ColorWriteMask       GPUColorComponentFlags
	EnableBlend          bool
	EnableColorWriteMask bool
	Padding1             uint8
	Padding2             uint8
}

type GPUShaderCreateInfo struct {
	CodeSize           uint64
	Code               *uint8
	EntryPoint         *byte // must be null-terminated string. use convert.ToBytePtr().
	Format             GPUShaderFormat
	Stage              GPUShaderStage
	NumSamplers        uint32
	NumStorageTextures uint32
	NumStorageBuffers  uint32
	NumUniformBuffers  uint32
	Props              PropertiesID
}

type GPURasterizerState struct {
	FillMode                GPUFillMode
	CullMode                GPUCullMode
	FrontFace               GPUFrontFace
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	EnableDepthBias         bool
	EnableDepthClip         bool
	Padding1                uint8
	Padding2                uint8
}

type GPUMultisampleState struct {
	SampleCount GPUSampleCount
	SampleMask  uint32
	EnableMask  bool
	Padding1    uint8
	Padding2    uint8
	Padding3    uint8
}

type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp
	BackStencilState  GPUStencilOpState
	FrontStencilState GPUStencilOpState
	CompareMask       uint8
	WriteMask         uint8
	EnableDepthTest   bool
	EnableDepthWrite  bool
	EnableStencilTest bool
	Padding1          uint8
	Padding2          uint8
	Padding3          uint8
}

type GPUColorTargetDescription struct {
	Format     GPUTextureFormat
	BlendState GPUColorTargetBlendState
}

type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions *GPUColorTargetDescription
	NumColorTargets         uint32
	DepthStencilFormat      GPUTextureFormat
	HasDepthStencilTarget   bool
	Padding1                uint8
	Padding2                uint8
	Padding3                uint8
}

type GPUGraphicsPipelineCreateInfo struct {
	VertextShader     *GPUShader
	FragmentShader    *GPUShader
	VertexInputState  GPUVertexInputState
	PrimitiveType     GPUPrimitiveType
	RasterizerState   GPURasterizerState
	MultisampleState  GPUMultisampleState
	DepthStencilState GPUDepthStencilState
	TargetInfo        GPUGraphicsPipelineTargetInfo
}

type GPUGraphicsPipeline struct{}

type GPUViewport struct {
	X, Y, W, H, MinDepth, MaxDepth float32
}

func AcquireGPUCommandBuffer(device *GPUDevice) *GPUCommandBuffer {
	return sdlAcquireGPUCommandBuffer(device)
}

func AcquireGPUSwapchainTexture(commandBuffer *GPUCommandBuffer, window *Window, swapchainTexture **GPUTexture, swapchainTextureWidth *uint32, swapchainTextureHeight *uint32) bool {
	return sdlAcquireGPUSwapchainTexture(commandBuffer, window, swapchainTexture, swapchainTextureWidth, swapchainTextureHeight)
}

// func BeginGPUComputePass(command_buffer *GPUCommandBuffer, storage_texture_bindings *GPUStorageTextureReadWriteBinding, num_storage_texture_bindings uint32, storage_buffer_bindings *GPUStorageBufferReadWriteBinding, num_storage_buffer_bindings uint32) *GPUComputePass {
//	return sdlBeginGPUComputePass(command_buffer, storage_texture_bindings, num_storage_texture_bindings, storage_buffer_bindings, num_storage_buffer_bindings)
// }

// func BeginGPUCopyPass(command_buffer *GPUCommandBuffer) *GPUCopyPass {
//	return sdlBeginGPUCopyPass(command_buffer)
// }

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

// func BindGPUFragmentSamplers(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUFragmentSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// func BindGPUFragmentStorageBuffers(render_pass *GPURenderPass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32)  {
//	sdlBindGPUFragmentStorageBuffers(render_pass, first_slot, storage_buffers, num_bindings)
// }

// func BindGPUFragmentStorageTextures(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUFragmentStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
// }

func BindGPUGraphicsPipeline(renderPass *GPURenderPass, graphicsPipeline *GPUGraphicsPipeline) {
	sdlBindGPUGraphicsPipeline(renderPass, graphicsPipeline)
}

// func BindGPUIndexBuffer(render_pass *GPURenderPass, binding *GPUBufferBinding, index_element_size GPUIndexElementSize)  {
//	sdlBindGPUIndexBuffer(render_pass, binding, index_element_size)
// }

// func BindGPUVertexBuffers(render_pass *GPURenderPass, first_slot uint32, bindings *GPUBufferBinding, num_bindings uint32)  {
//	sdlBindGPUVertexBuffers(render_pass, first_slot, bindings, num_bindings)
// }

// func BindGPUVertexSamplers(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUVertexSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// func BindGPUVertexStorageBuffers(render_pass *GPURenderPass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32)  {
//	sdlBindGPUVertexStorageBuffers(render_pass, first_slot, storage_buffers, num_bindings)
// }

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

func ClaimWindowForGPUDevice(device *GPUDevice, window *Window) bool {
	return sdlClaimWindowForGPUDevice(device, window)
}

// func CopyGPUBufferToBuffer(copy_pass *GPUCopyPass, source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool)  {
//	sdlCopyGPUBufferToBuffer(copy_pass, source, destination, size, cycle)
// }

// func CopyGPUTextureToTexture(copy_pass *GPUCopyPass, source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool)  {
//	sdlCopyGPUTextureToTexture(copy_pass, source, destination, w, h, d, cycle)
// }

// func CreateGPUBuffer(device *GPUDevice, createinfo *GPUBufferCreateInfo) *GPUBuffer {
//	return sdlCreateGPUBuffer(device, createinfo)
// }

// func CreateGPUComputePipeline(device *GPUDevice, createinfo *GPUComputePipelineCreateInfo) *GPUComputePipeline {
//	return sdlCreateGPUComputePipeline(device, createinfo)
// }

func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) *GPUDevice {
	return sdlCreateGPUDevice(formatFlags, debugMode, convert.ToBytePtrNullable(name))
}

// func CreateGPUDeviceWithProperties(props PropertiesID) *GPUDevice {
//	return sdlCreateGPUDeviceWithProperties(props)
// }

func CreateGPUGraphicsPipeline(device *GPUDevice, createInfo *GPUGraphicsPipelineCreateInfo) *GPUGraphicsPipeline {
	return sdlCreateGPUGraphicsPipeline(device, createInfo)
}

// func CreateGPUSampler(device *GPUDevice, createinfo *GPUSamplerCreateInfo) *GPUSampler {
//	return sdlCreateGPUSampler(device, createinfo)
// }

func CreateGPUShader(device *GPUDevice, createInfo *GPUShaderCreateInfo) *GPUShader {
	return sdlCreateGPUShader(device, createInfo)
}

// func CreateGPUTexture(device *GPUDevice, createinfo *GPUTextureCreateInfo) *GPUTexture {
//	return sdlCreateGPUTexture(device, createinfo)
// }

// func CreateGPUTransferBuffer(device *GPUDevice, createinfo *GPUTransferBufferCreateInfo) *GPUTransferBuffer {
//	return sdlCreateGPUTransferBuffer(device, createinfo)
// }

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

// func DrawGPUIndexedPrimitives(render_pass *GPURenderPass, num_indices uint32, num_instances uint32, first_index uint32, vertex_offset int32, first_instance uint32)  {
//	sdlDrawGPUIndexedPrimitives(render_pass, num_indices, num_instances, first_index, vertex_offset, first_instance)
// }

// func DrawGPUIndexedPrimitivesIndirect(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32)  {
//	sdlDrawGPUIndexedPrimitivesIndirect(render_pass, buffer, offset, draw_count)
// }

func DrawGPUPrimitives(renderPass *GPURenderPass, numVertices uint32, numInstances uint32, firstVertex uint32, firstInstance uint32) {
	sdlDrawGPUPrimitives(renderPass, numVertices, numInstances, firstVertex, firstInstance)
}

// func DrawGPUPrimitivesIndirect(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32)  {
//	sdlDrawGPUPrimitivesIndirect(render_pass, buffer, offset, draw_count)
// }

// func EndGPUComputePass(compute_pass *GPUComputePass)  {
//	sdlEndGPUComputePass(compute_pass)
// }

// func EndGPUCopyPass(copy_pass *GPUCopyPass)  {
//	sdlEndGPUCopyPass(copy_pass)
// }

func EndGPURenderPass(renderPass *GPURenderPass) {
	sdlEndGPURenderPass(renderPass)
}

// func GenerateMipmapsForGPUTexture(command_buffer *GPUCommandBuffer, texture *GPUTexture)  {
//	sdlGenerateMipmapsForGPUTexture(command_buffer, texture)
// }

// func GetGPUDeviceDriver(device *GPUDevice) string {
//	return sdlGetGPUDeviceDriver(device)
// }

func GetGPUDriver(index int32) string {
	return sdlGetGPUDriver(index)
}

func GetGPUShaderFormats(device *GPUDevice) GPUShaderFormat {
	return sdlGetGPUShaderFormats(device)
}

func GetGPUSwapchainTextureFormat(device *GPUDevice, window *Window) GPUTextureFormat {
	return sdlGetGPUSwapchainTextureFormat(device, window)
}

// func GetNumGPUDrivers() int32 {
//	return sdlGetNumGPUDrivers()
// }

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

// func MapGPUTransferBuffer(device *GPUDevice, transfer_buffer *GPUTransferBuffer, cycle bool) unsafe.Pointer {
//	return sdlMapGPUTransferBuffer(device, transfer_buffer, cycle)
// }

// func PopGPUDebugGroup(command_buffer *GPUCommandBuffer)  {
//	sdlPopGPUDebugGroup(command_buffer)
// }

// func PushGPUComputeUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUComputeUniformData(command_buffer, slot_index, data, length)
// }

// func PushGPUDebugGroup(command_buffer *GPUCommandBuffer, name string)  {
//	sdlPushGPUDebugGroup(command_buffer, name)
// }

// func PushGPUFragmentUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUFragmentUniformData(command_buffer, slot_index, data, length)
// }

// func PushGPUVertexUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUVertexUniformData(command_buffer, slot_index, data, length)
// }

// func QueryGPUFence(device *GPUDevice, fence *GPUFence) bool {
//	return sdlQueryGPUFence(device, fence)
// }

// func ReleaseGPUBuffer(device *GPUDevice, buffer *GPUBuffer)  {
//	sdlReleaseGPUBuffer(device, buffer)
// }

// func ReleaseGPUComputePipeline(device *GPUDevice, compute_pipeline *GPUComputePipeline)  {
//	sdlReleaseGPUComputePipeline(device, compute_pipeline)
// }

// func ReleaseGPUFence(device *GPUDevice, fence *GPUFence)  {
//	sdlReleaseGPUFence(device, fence)
// }

func ReleaseGPUGraphicsPipeline(device *GPUDevice, graphicsPipeline *GPUGraphicsPipeline) {
	sdlReleaseGPUGraphicsPipeline(device, graphicsPipeline)
}

// func ReleaseGPUSampler(device *GPUDevice, sampler *GPUSampler)  {
//	sdlReleaseGPUSampler(device, sampler)
// }

func ReleaseGPUShader(device *GPUDevice, shader *GPUShader) {
	sdlReleaseGPUShader(device, shader)
}

// func ReleaseGPUTexture(device *GPUDevice, texture *GPUTexture)  {
//	sdlReleaseGPUTexture(device, texture)
// }

// func ReleaseGPUTransferBuffer(device *GPUDevice, transfer_buffer *GPUTransferBuffer)  {
//	sdlReleaseGPUTransferBuffer(device, transfer_buffer)
// }

func ReleaseWindowFromGPUDevice(device *GPUDevice, window *Window) {
	sdlReleaseWindowFromGPUDevice(device, window)
}

// func SetGPUAllowedFramesInFlight(device *GPUDevice, allowed_frames_in_flight uint32) bool {
//	return sdlSetGPUAllowedFramesInFlight(device, allowed_frames_in_flight)
// }

// func SetGPUBlendConstants(render_pass *GPURenderPass, blend_constants FColor)  {
//	sdlSetGPUBlendConstants(render_pass, blend_constants)
// }

// func SetGPUBufferName(device *GPUDevice, buffer *GPUBuffer, text string)  {
//	sdlSetGPUBufferName(device, buffer, text)
// }

func SetGPUScissor(renderPass *GPURenderPass, scissor *Rect) {
	sdlSetGPUScissor(renderPass, scissor)
}

// func SetGPUStencilReference(render_pass *GPURenderPass, reference uint8)  {
//	sdlSetGPUStencilReference(render_pass, reference)
// }

func SetGPUSwapchainParameters(device *GPUDevice, window *Window, swapchainComposition GPUSwapchainComposition, presentMode GPUPresentMode) bool {
	return sdlSetGPUSwapchainParameters(device, window, swapchainComposition, presentMode)
}

// func SetGPUTextureName(device *GPUDevice, texture *GPUTexture, text string)  {
//	sdlSetGPUTextureName(device, texture, text)
// }

func SetGPUViewport(renderass *GPURenderPass, viewport *GPUViewport) {
	sdlSetGPUViewport(renderass, viewport)
}

func SubmitGPUCommandBuffer(commandBuffer *GPUCommandBuffer) bool {
	return sdlSubmitGPUCommandBuffer(commandBuffer)
}

// func SubmitGPUCommandBufferAndAcquireFence(command_buffer *GPUCommandBuffer) *GPUFence {
//	return sdlSubmitGPUCommandBufferAndAcquireFence(command_buffer)
// }

// func UnmapGPUTransferBuffer(device *GPUDevice, transfer_buffer *GPUTransferBuffer)  {
//	sdlUnmapGPUTransferBuffer(device, transfer_buffer)
// }

// func UploadToGPUBuffer(copy_pass *GPUCopyPass, source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool)  {
//	sdlUploadToGPUBuffer(copy_pass, source, destination, cycle)
// }

// func UploadToGPUTexture(copy_pass *GPUCopyPass, source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool)  {
//	sdlUploadToGPUTexture(copy_pass, source, destination, cycle)
// }

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

func WindowSupportsGPUPresentMode(device *GPUDevice, window *Window, presentMode GPUPresentMode) bool {
	return sdlWindowSupportsGPUPresentMode(device, window, presentMode)
}

// func WindowSupportsGPUSwapchainComposition(device *GPUDevice, window *Window, swapchain_composition GPUSwapchainComposition) bool {
//	return sdlWindowSupportsGPUSwapchainComposition(device, window, swapchain_composition)
// }
