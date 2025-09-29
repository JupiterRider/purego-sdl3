package sdl

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

const SoftwareRenderer = "software"

const (
	PropRendererCreateNameString                           = "SDL.renderer.create.name"
	PropRendererCreateWindowPointer                        = "SDL.renderer.create.window"
	PropRendererCreateSurfacePointer                       = "SDL.renderer.create.surface"
	PropRendererCreateOutputColorspaceNumber               = "SDL.renderer.create.output_colorspace"
	PropRendererCreatePresentVSyncNumber                   = "SDL.renderer.create.present_vsync"
	PropRendererCreateVulkanInstancePointer                = "SDL.renderer.create.vulkan.instance"
	PropRendererCreateVulkanSurfaceNumber                  = "SDL.renderer.create.vulkan.surface"
	PropRendererCreateVulkanPhysicalDevicePointer          = "SDL.renderer.create.vulkan.physical_device"
	PropRendererCreateVulkanDevicePointer                  = "SDL.renderer.create.vulkan.device"
	PropRendererCreateVulkanGraphicsQueueFamilyIndexNumber = "SDL.renderer.create.vulkan.graphics_queue_family_index"
	PropRendererCreateVulkanPresentQueueFamilyIndexNumber  = "SDL.renderer.create.vulkan.present_queue_family_index"

	PropRendererNameString                           = "SDL.renderer.name"
	PropRendererWindowPointer                        = "SDL.renderer.window"
	PropRendererSurfacePointer                       = "SDL.renderer.surface"
	PropRendererVSyncNumber                          = "SDL.renderer.vsync"
	PropRendererMaxTextureSizeNumber                 = "SDL.renderer.max_texture_size"
	PropRendererTextureFormatsPointer                = "SDL.renderer.texture_formats"
	PropRendererOutputColorspaceNumber               = "SDL.renderer.output_colorspace"
	PropRendererHDREnabledBoolean                    = "SDL.renderer.HDR_enabled"
	PropRendererSDRWhitePointFloat                   = "SDL.renderer.SDR_white_point"
	PropRendererHDRHeadroomFloat                     = "SDL.renderer.HDR_headroom"
	PropRendererD3D9DevicePointer                    = "SDL.renderer.d3d9.device"
	PropRendererD3D11DevicePointer                   = "SDL.renderer.d3d11.device"
	PropRendererD3D11SwapchainPointer                = "SDL.renderer.d3d11.swap_chain"
	PropRendererD3D12DevicePointer                   = "SDL.renderer.d3d12.device"
	PropRendererD3D12SwapchainPointer                = "SDL.renderer.d3d12.swap_chain"
	PropRendererD3D12CommandQueuePointer             = "SDL.renderer.d3d12.command_queue"
	PropRendererVulkanInstancePointer                = "SDL.renderer.vulkan.instance"
	PropRendererVulkanSurfaceNumber                  = "SDL.renderer.vulkan.surface"
	PropRendererVulkanPhysicalDevicePointer          = "SDL.renderer.vulkan.physical_device"
	PropRendererVulkanDevicePointer                  = "SDL.renderer.vulkan.device"
	PropRendererVulkanGraphicsQueueFamilyIndexNumber = "SDL.renderer.vulkan.graphics_queue_family_index"
	PropRendererVulkanPresentQueueFamilyIndexNumber  = "SDL.renderer.vulkan.present_queue_family_index"
	PropRendererVulkanSwapchainImageCountNumber      = "SDL.renderer.vulkan.swapchain_image_count"
	PropRendererGPUDevicePointer                     = "SDL.renderer.gpu.device"

	PropTextureCreateColorspaceNumber         = "SDL.texture.create.colorspace"
	PropTextureCreateFormatNumber             = "SDL.texture.create.format"
	PropTextureCreateAccessNumber             = "SDL.texture.create.access"
	PropTextureCreateWidthNumber              = "SDL.texture.create.width"
	PropTextureCreateHeightNumber             = "SDL.texture.create.height"
	PropTextureCreateSDRWhitePointFloat       = "SDL.texture.create.SDR_white_point"
	PropTextureCreateHDRHeadroomFloat         = "SDL.texture.create.HDR_headroom"
	PropTextureCreateD3D11TexturePointer      = "SDL.texture.create.d3d11.texture"
	PropTextureCreateD3D11TextureUPointer     = "SDL.texture.create.d3d11.texture_u"
	PropTextureCreateD3D11TextureVPointer     = "SDL.texture.create.d3d11.texture_v"
	PropTextureCreateD3D12TexturePointer      = "SDL.texture.create.d3d12.texture"
	PropTextureCreateD3D12TextureUPointer     = "SDL.texture.create.d3d12.texture_u"
	PropTextureCreateD3D12TextureVPointer     = "SDL.texture.create.d3d12.texture_v"
	PropTextureCreateMetalPixelbufferPointer  = "SDL.texture.create.metal.pixelbuffer"
	PropTextureCreateOpenGLTextureNumber      = "SDL.texture.create.opengl.texture"
	PropTextureCreateOpenGLTextureUvNumber    = "SDL.texture.create.opengl.texture_uv"
	PropTextureCreateOpenGLTextureUNumber     = "SDL.texture.create.opengl.texture_u"
	PropTextureCreateOpenGLTextureVNumber     = "SDL.texture.create.opengl.texture_v"
	PropTextureCreateOpenGLES2TextureNumber   = "SDL.texture.create.opengles2.texture"
	PropTextureCreateOpenGLES2TextureUvNumber = "SDL.texture.create.opengles2.texture_uv"
	PropTextureCreateOpenGLES2TextureUNumber  = "SDL.texture.create.opengles2.texture_u"
	PropTextureCreateOpenGLES2TextureVNumber  = "SDL.texture.create.opengles2.texture_v"
	PropTextureCreateVulkanTextureNumber      = "SDL.texture.create.vulkan.texture"

	PropTextureColorspaceNumber             = "SDL.texture.colorspace"
	PropTextureFormatNumber                 = "SDL.texture.format"
	PropTextureAccessNumber                 = "SDL.texture.access"
	PropTextureWidthNumber                  = "SDL.texture.width"
	PropTextureHeightNumber                 = "SDL.texture.height"
	PropTextureSDRWhitePointFloat           = "SDL.texture.SDR_white_point"
	PropTextureHDRHeadroomFloat             = "SDL.texture.HDR_headroom"
	PropTextureD3D11TexturePointer          = "SDL.texture.d3d11.texture"
	PropTextureD3D11TextureUPointer         = "SDL.texture.d3d11.texture_u"
	PropTextureD3D11TextureVPointer         = "SDL.texture.d3d11.texture_v"
	PropTextureD3D12TexturePointer          = "SDL.texture.d3d12.texture"
	PropTextureD3D12TextureUPointer         = "SDL.texture.d3d12.texture_u"
	PropTextureD3D12TextureVPointer         = "SDL.texture.d3d12.texture_v"
	PropTextureOpenGLTextureNumber          = "SDL.texture.opengl.texture"
	PropTextureOpenGLTextureUvNumber        = "SDL.texture.opengl.texture_uv"
	PropTextureOpenGLTextureUNumber         = "SDL.texture.opengl.texture_u"
	PropTextureOpenGLTextureVNumber         = "SDL.texture.opengl.texture_v"
	PropTextureOpenGLTextureTargetNumber    = "SDL.texture.opengl.target"
	PropTextureOpenGLTexWFloat              = "SDL.texture.opengl.tex_w"
	PropTextureOpenGLTexHFloat              = "SDL.texture.opengl.tex_h"
	PropTextureOpenGLES2TextureNumber       = "SDL.texture.opengles2.texture"
	PropTextureOpenGLES2TextureUvNumber     = "SDL.texture.opengles2.texture_uv"
	PropTextureOpenGLES2TextureUNumber      = "SDL.texture.opengles2.texture_u"
	PropTextureOpenGLES2TextureVNumber      = "SDL.texture.opengles2.texture_v"
	PropTextureOpenGLES2TextureTargetNumber = "SDL.texture.opengles2.target"
	PropTextureVulkanTextureNumber          = "SDL.texture.vulkan.texture"
)

const (
	RendererVSyncDisabled = 0
	RendererVSyncAdaptive = -1
)

const (
	DebugTextFontCharacterSize = 8
)

// [RendererLogicalPresentation] describes how the logical size is mapped to the output.
//
// [RendererLogicalPresentation]: https://wiki.libsdl.org/SDL3/SDL_RendererLogicalPresentation
type RendererLogicalPresentation uint32

const (
	LogicalPresentationDisabled     RendererLogicalPresentation = iota // There is no logical size in effect
	LogicalPresentationStretch                                         // The rendered content is stretched to the output resolution
	LogicalPresentationLetterbox                                       // The rendered content is fit to the largest dimension and the other dimension is letterboxed with the clear color
	LogicalPresentationOverscan                                        // The rendered content is fit to the smallest dimension and the other dimension extends beyond the output bounds
	LogicalPresentationIntegerScale                                    // The rendered content is scaled up by integer multiples to fit the output resolution
)

// [TextureAccess] the access pattern allowed for a texture.
//
// [TextureAccess]: https://wiki.libsdl.org/SDL3/SDL_TextureAccess
type TextureAccess uint32

const (
	TextureAccessStatic    TextureAccess = iota // Changes rarely, not lockable
	TextureAccessStreaming                      // Changes frequently, lockable
	TextureAccessTarget                         // Texture can be used as a render target
)

// [Renderer] is a structure representing rendering state
//
// [Renderer]: https://wiki.libsdl.org/SDL3/SDL_Renderer
type Renderer struct{}

// [Texture] is a efficient driver-specific representation of pixel data.
//
// [Texture]: https://wiki.libsdl.org/SDL3/SDL_Texture
type Texture struct {
	Format   PixelFormat // The format of the texture, read-only
	W        int32       // The width of the texture, read-only.
	H        int32       // The height of the texture, read-only.
	Refcount int32       // Application reference count, used when freeing texture
}

// [Vertex] is a vertex structure.
//
// [Vertex]: https://wiki.libsdl.org/SDL3/SDL_Vertex
type Vertex struct {
	Position FPoint // Vertex position, in SDL_Renderer coordinates
	Color    FColor // Vertex color
	TexCoord FPoint // Normalized texture coordinates, if needed
}

// [CreateWindowAndRenderer] creates a window and default renderer.
//
// [CreateWindowAndRenderer]: https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer
func CreateWindowAndRenderer(title string, width, height int32, flags WindowFlags, window **Window, renderer **Renderer) bool {
	return sdlCreateWindowAndRenderer(title, width, height, flags, window, renderer)
}

// [SetRenderDrawColor] sets the color used for drawing operations.
//
// [SetRenderDrawColor]: https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor
func SetRenderDrawColor(renderer *Renderer, r, g, b, a uint8) bool {
	ret, _, _ := purego.SyscallN(sdlSetRenderDrawColor, uintptr(unsafe.Pointer(renderer)), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	return byte(ret) != 0
}

// [RenderPresent] updates the screen with any rendering performed since the previous call.
//
// [RenderPresent]: https://wiki.libsdl.org/SDL3/SDL_RenderPresent
func RenderPresent(renderer *Renderer) bool {
	ret, _, _ := purego.SyscallN(sdlRenderPresent, uintptr(unsafe.Pointer(renderer)))
	return byte(ret) != 0
}

// [RenderClear] clears the current rendering target with the drawing color.
//
// [RenderClear]: https://wiki.libsdl.org/SDL3/SDL_RenderClear
func RenderClear(renderer *Renderer) bool {
	ret, _, _ := purego.SyscallN(sdlRenderClear, uintptr(unsafe.Pointer(renderer)))
	return byte(ret) != 0
}

// [DestroyRenderer] destroys the rendering context for a window and free all associated textures.
//
// [DestroyRenderer]: https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer
func DestroyRenderer(renderer *Renderer) {
	sdlDestroyRenderer(renderer)
}

// [RenderRect] draws a rectangle on the current rendering target at subpixel precision.
//
// [RenderRect]: https://wiki.libsdl.org/SDL3/SDL_RenderRect
func RenderRect(renderer *Renderer, rect *FRect) bool {
	ret, _, _ := purego.SyscallN(sdlRenderRect, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect)))
	return byte(ret) != 0
}

// [RenderFillRect] fills a rectangle on the current rendering target with the drawing color at subpixel precision.
//
// [RenderFillRect]: https://wiki.libsdl.org/SDL3/SDL_RenderFillRect
func RenderFillRect(renderer *Renderer, rect *FRect) bool {
	ret, _, _ := purego.SyscallN(sdlRenderFillRect, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect)))
	return byte(ret) != 0
}

// [RenderDebugText] draws debug text to a [Renderer].
//
// [RenderDebugText]: https://wiki.libsdl.org/SDL3/SDL_RenderDebugText
func RenderDebugText(renderer *Renderer, x, y float32, str string) bool {
	return sdlRenderDebugText(renderer, x, y, str)
}

// [CreateTextureFromSurface] creates a texture from an existing surface.
//
// [CreateTextureFromSurface]: https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface
func CreateTextureFromSurface(renderer *Renderer, surface *Surface) *Texture {
	return sdlCreateTextureFromSurface(renderer, surface)
}

// [RenderTexture] copies a portion of the texture to the current rendering target at subpixel precision.
//
// [RenderTexture]: https://wiki.libsdl.org/SDL3/SDL_RenderTexture
func RenderTexture(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) bool {
	ret, _, _ := purego.SyscallN(sdlRenderTexture, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dstrect)))
	return byte(ret) != 0
}

// [DestroyTexture] destroys the specified texture.
//
// [DestroyTexture]: https://wiki.libsdl.org/SDL3/SDL_DestroyTexture
func DestroyTexture(texture *Texture) {
	sdlDestroyTexture(texture)
}

// [AddVulkanRenderSemaphores] adds a set of synchronization semaphores for the current frame.
//
// [AddVulkanRenderSemaphores]: https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores
func AddVulkanRenderSemaphores(renderer *Renderer, waitStageMask uint32, waitSemaphore int64, signalSemaphore int64) bool {
	return sdlAddVulkanRenderSemaphores(renderer, waitStageMask, waitSemaphore, signalSemaphore)
}

// [ConvertEventToRenderCoordinates] converts the coordinates in an event to render coordinates.
//
// [ConvertEventToRenderCoordinates]: https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates
func ConvertEventToRenderCoordinates(renderer *Renderer, event *Event) bool {
	return sdlConvertEventToRenderCoordinates(renderer, event)
}

// [CreateRenderer] creates a 2D rendering context for a window.
// The name parameter can be one driver, a comma-separated list of drivers or "" to let SDL choose one.
//
// [CreateRenderer]: https://wiki.libsdl.org/SDL3/SDL_CreateRenderer
func CreateRenderer(window *Window, name string) *Renderer {
	return sdlCreateRenderer(window, convert.ToBytePtrNullable(name))
}

// [CreateRendererWithProperties] creates a 2D rendering context for a window, with the specified properties.
//
// [CreateRendererWithProperties]: https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties
func CreateRendererWithProperties(props PropertiesID) *Renderer {
	return sdlCreateRendererWithProperties(props)
}

// [CreateSoftwareRenderer] creates a 2D software rendering context for a surface.
//
// [CreateSoftwareRenderer]: https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer
func CreateSoftwareRenderer(surface *Surface) *Renderer {
	return sdlCreateSoftwareRenderer(surface)
}

// [CreateTexture] creates a texture for a rendering context.
//
// [CreateTexture]: https://wiki.libsdl.org/SDL3/SDL_CreateTexture
func CreateTexture(renderer *Renderer, format PixelFormat, access TextureAccess, w int32, h int32) *Texture {
	return sdlCreateTexture(renderer, format, access, w, h)
}

// [CreateTextureWithProperties] creates a texture for a rendering context with the specified properties.
//
// [CreateTextureWithProperties]: https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties
func CreateTextureWithProperties(renderer *Renderer, props PropertiesID) *Texture {
	return sdlCreateTextureWithProperties(renderer, props)
}

// [FlushRenderer] forces the rendering context to flush any pending commands and state.
//
// [FlushRenderer]: https://wiki.libsdl.org/SDL3/SDL_FlushRenderer
func FlushRenderer(renderer *Renderer) bool {
	ret, _, _ := purego.SyscallN(sdlFlushRenderer, uintptr(unsafe.Pointer(renderer)))
	return byte(ret) != 0
}

// [GetCurrentRenderOutputSize] gets the current output size in pixels of a rendering context.
//
// [GetCurrentRenderOutputSize]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize
func GetCurrentRenderOutputSize(renderer *Renderer, w *int32, h *int32) bool {
	return sdlGetCurrentRenderOutputSize(renderer, w, h)
}

// [GetNumRenderDrivers] gets the number of 2D rendering drivers available for the current display.
//
// [GetNumRenderDrivers]: https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers
func GetNumRenderDrivers() int32 {
	return sdlGetNumRenderDrivers()
}

// [GetRenderClipRect] gets the clip rectangle for the current target.
//
// [GetRenderClipRect]: https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect
func GetRenderClipRect(renderer *Renderer, rect *Rect) bool {
	return sdlGetRenderClipRect(renderer, rect)
}

// [GetRenderColorScale] gets the color scale used for render operations.
//
// [GetRenderColorScale]: https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale
func GetRenderColorScale(renderer *Renderer, scale *float32) bool {
	return sdlGetRenderColorScale(renderer, scale)
}

// [GetRenderDrawBlendMode] gets the blend mode used for drawing operations.
//
// [GetRenderDrawBlendMode]: https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode
func GetRenderDrawBlendMode(renderer *Renderer, blendMode *BlendMode) bool {
	return sdlGetRenderDrawBlendMode(renderer, blendMode)
}

// [GetRenderDrawColor] gets the color used for drawing operations (Rect, Line and Clear).
//
// [GetRenderDrawColor]: https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor
func GetRenderDrawColor(renderer *Renderer, r *uint8, g *uint8, b *uint8, a *uint8) bool {
	return sdlGetRenderDrawColor(renderer, r, g, b, a)
}

// [GetRenderDrawColorFloat] gets the color used for drawing operations (Rect, Line and Clear).
//
// [GetRenderDrawColorFloat]: https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat
func GetRenderDrawColorFloat(renderer *Renderer, r *float32, g *float32, b *float32, a *float32) bool {
	return sdlGetRenderDrawColorFloat(renderer, r, g, b, a)
}

// [GetRenderDriver] gets the name of a built in 2D rendering driver.
//
// [GetRenderDriver]: https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver
func GetRenderDriver(index int32) string {
	return sdlGetRenderDriver(index)
}

// [GetRenderer] gets the renderer associated with a window.
//
// [GetRenderer]: https://wiki.libsdl.org/SDL3/SDL_GetRenderer
func GetRenderer(window *Window) *Renderer {
	return sdlGetRenderer(window)
}

// [GetRendererFromTexture] gets the renderer that created an [Texture].
//
// [GetRendererFromTexture]: https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture
func GetRendererFromTexture(texture *Texture) *Renderer {
	return sdlGetRendererFromTexture(texture)
}

// [GetRendererName] returns the name of the selected renderer, or "" on failure.
//
// [GetRendererName]: https://wiki.libsdl.org/SDL3/SDL_GetRendererName
func GetRendererName(renderer *Renderer) string {
	return sdlGetRendererName(renderer)
}

// [GetRendererProperties] gets the properties associated with a renderer.
//
// [GetRendererProperties]: https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties
func GetRendererProperties(renderer *Renderer) PropertiesID {
	return sdlGetRendererProperties(renderer)
}

// [GetRenderLogicalPresentation] gets device independent resolution and presentation mode for rendering.
//
// [GetRenderLogicalPresentation]: https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation
func GetRenderLogicalPresentation(renderer *Renderer, w *int32, h *int32, mode *RendererLogicalPresentation) bool {
	return sdlGetRenderLogicalPresentation(renderer, w, h, mode)
}

// [GetRenderLogicalPresentationRect] gets the final presentation rectangle for rendering.
//
// [GetRenderLogicalPresentationRect]: https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect
func GetRenderLogicalPresentationRect(renderer *Renderer, rect *FRect) bool {
	return sdlGetRenderLogicalPresentationRect(renderer, rect)
}

// [GetRenderMetalCommandEncoder] gets the Metal command encoder for the current frame.
//
// [GetRenderMetalCommandEncoder]: https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder
func GetRenderMetalCommandEncoder(renderer *Renderer) unsafe.Pointer {
	return sdlGetRenderMetalCommandEncoder(renderer)
}

// [GetRenderMetalLayer] gets the CAMetalLayer associated with the given Metal renderer.
//
// [GetRenderMetalLayer]: https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer
func GetRenderMetalLayer(renderer *Renderer) unsafe.Pointer {
	return sdlGetRenderMetalLayer(renderer)
}

// [GetRenderOutputSize] gets the output size in pixels of a rendering context.
//
// [GetRenderOutputSize]: https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize
func GetRenderOutputSize(renderer *Renderer, w *int32, h *int32) bool {
	return sdlGetRenderOutputSize(renderer, w, h)
}

// [GetRenderSafeArea] gets the safe area for rendering within the current viewport.
//
// [GetRenderSafeArea]: https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea
func GetRenderSafeArea(renderer *Renderer, rect *Rect) bool {
	return sdlGetRenderSafeArea(renderer, rect)
}

// [GetRenderScale] gets the drawing scale for the current target.
//
// [GetRenderScale]: https://wiki.libsdl.org/SDL3/SDL_GetRenderScale
func GetRenderScale(renderer *Renderer, scaleX *float32, scaleY *float32) bool {
	return sdlGetRenderScale(renderer, scaleX, scaleY)
}

// [GetRenderTarget] gets the current render target.
//
// [GetRenderTarget]: https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget
func GetRenderTarget(renderer *Renderer) *Texture {
	return sdlGetRenderTarget(renderer)
}

// [GetRenderViewport] gets the drawing area for the current target.
//
// [GetRenderViewport]: https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport
func GetRenderViewport(renderer *Renderer, rect *Rect) bool {
	return sdlGetRenderViewport(renderer, rect)
}

// [GetRenderVSync] gets VSync of the given renderer.
//
// [GetRenderVSync]: https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync
func GetRenderVSync(renderer *Renderer, vsync *int32) bool {
	return sdlGetRenderVSync(renderer, vsync)
}

// [GetRenderWindow] gets the window associated with a renderer.
//
// [GetRenderWindow]: https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow
func GetRenderWindow(renderer *Renderer) *Window {
	return sdlGetRenderWindow(renderer)
}

// [GetTextureAlphaMod] gets the additional alpha value multiplied into render copy operations.
//
// [GetTextureAlphaMod]: https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod
func GetTextureAlphaMod(texture *Texture, alpha *uint8) bool {
	return sdlGetTextureAlphaMod(texture, alpha)
}

// [GetTextureAlphaModFloat] gets the additional alpha value multiplied into render copy operations.
//
// [GetTextureAlphaModFloat]: https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat
func GetTextureAlphaModFloat(texture *Texture, alpha *float32) bool {
	return sdlGetTextureAlphaModFloat(texture, alpha)
}

// [GetTextureBlendMode] gets the blend mode used for texture copy operations.
//
// [GetTextureBlendMode]: https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode
func GetTextureBlendMode(texture *Texture, blendMode *BlendMode) bool {
	return sdlGetTextureBlendMode(texture, blendMode)
}

// [GetTextureColorMod] gets the additional color value multiplied into render copy operations.
//
// [GetTextureColorMod]: https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod
func GetTextureColorMod(texture *Texture, r *uint8, g *uint8, b *uint8) bool {
	return sdlGetTextureColorMod(texture, r, g, b)
}

// [GetTextureColorModFloat] gets the additional color value multiplied into render copy operations.
//
// [GetTextureColorModFloat]: https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat
func GetTextureColorModFloat(texture *Texture, r *float32, g *float32, b *float32) bool {
	return sdlGetTextureColorModFloat(texture, r, g, b)
}

// [GetTextureProperties] gets the properties associated with a texture.
//
// [GetTextureProperties]: https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties
func GetTextureProperties(texture *Texture) PropertiesID {
	return sdlGetTextureProperties(texture)
}

// [GetTextureScaleMode] gets the scale mode used for texture scale operations.
//
// [GetTextureScaleMode]: https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode
func GetTextureScaleMode(texture *Texture, scaleMode *ScaleMode) bool {
	return sdlGetTextureScaleMode(texture, scaleMode)
}

// [GetTextureSize] gets the size of a texture, as floating point values.
//
// [GetTextureSize]: https://wiki.libsdl.org/SDL3/SDL_GetTextureSize
func GetTextureSize(texture *Texture, w *float32, h *float32) bool {
	return sdlGetTextureSize(texture, w, h)
}

// [LockTexture] locks a portion of the texture for write-only pixel access.
//
// [LockTexture]: https://wiki.libsdl.org/SDL3/SDL_LockTexture
func LockTexture(texture *Texture, rect *Rect, pixels *unsafe.Pointer, pitch *int32) bool {
	return sdlLockTexture(texture, rect, pixels, pitch)
}

// [LockTextureToSurface] locks a portion of the texture for write-only pixel access, and expose it as a SDL surface.
//
// [LockTextureToSurface]: https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface
func LockTextureToSurface(texture *Texture, rect *Rect, surface **Surface) bool {
	return sdlLockTextureToSurface(texture, rect, surface)
}

// [RenderClipEnabled] gets whether clipping is enabled on the given render target.
//
// [RenderClipEnabled]: https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled
func RenderClipEnabled(renderer *Renderer) bool {
	return sdlRenderClipEnabled(renderer)
}

// [RenderCoordinatesFromWindow] gets a point in render coordinates when given a point in window coordinates.
//
// [RenderCoordinatesFromWindow]: https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow
func RenderCoordinatesFromWindow(renderer *Renderer, windowX float32, windowY float32, x *float32, y *float32) bool {
	return sdlRenderCoordinatesFromWindow(renderer, windowX, windowY, x, y)
}

// [RenderCoordinatesToWindow] gets a point in window coordinates when given a point in render coordinates.
//
// [RenderCoordinatesToWindow]: https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow
func RenderCoordinatesToWindow(renderer *Renderer, x float32, y float32, windowX *float32, windowY *float32) bool {
	return sdlRenderCoordinatesToWindow(renderer, x, y, windowX, windowY)
}

// [RenderDebugTextFormat] draws debug text to an [Renderer].
//
// [RenderDebugTextFormat]: https://wiki.libsdl.org/SDL3/SDL_RenderDebugTextFormat
func RenderDebugTextFormat(renderer *Renderer, x float32, y float32, format string, a ...any) bool {
	return sdlRenderDebugTextFormat(renderer, x, y, fmt.Sprintf(format, a...))
}

// [RenderFillRects] fills some number of rectangles on the current rendering target with the drawing color at subpixel precision.
//
// [RenderFillRects]: https://wiki.libsdl.org/SDL3/SDL_RenderFillRects
func RenderFillRects(renderer *Renderer, rects []FRect) bool {
	count := len(rects)
	var rectsPtr *FRect
	if count > 0 {
		rectsPtr = &rects[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderFillRects, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rectsPtr)), uintptr(count))
	return byte(ret) != 0
}

// [RenderGeometry] renders a list of triangles, optionally using a texture and indices into the vertex array Color and alpha
// modulation is done per vertex ([SetTextureColorMod] and [SetTextureAlphaMod] are ignored).
//
// [RenderGeometry]: https://wiki.libsdl.org/SDL3/SDL_RenderGeometry
func RenderGeometry(renderer *Renderer, texture *Texture, vertices []Vertex, indices []int32) bool {
	numVertices := len(vertices)
	var verticesPtr *Vertex
	if numVertices > 0 {
		verticesPtr = &vertices[0]
	}

	numIndices := len(indices)
	var indicesPtr *int32
	if numIndices > 0 {
		indicesPtr = &indices[0]
	}

	ret, _, _ := purego.SyscallN(sdlRenderGeometry,
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(verticesPtr)),
		uintptr(numVertices),
		uintptr(unsafe.Pointer(indicesPtr)),
		uintptr(numIndices))

	return byte(ret) != 0
}

// [RenderGeometryRaw] renders a list of triangles, optionally using a texture and indices into the vertex arrays Color and alpha
// modulation is done per vertex ([SetTextureColorMod] and [SetTextureAlphaMod] are ignored).
//
// [RenderGeometryRaw]: https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw
func RenderGeometryRaw(renderer *Renderer, texture *Texture, xy []FPoint, color []FColor, uv []FPoint, indices []int32) bool {
	var xyPtr, uvPtr *FPoint
	numXY := len(xy)
	if numXY > 0 {
		xyPtr = &xy[0]
	}
	if len(uv) > 0 {
		uvPtr = &uv[0]
	}

	var colorPtr *FColor
	if len(color) > 0 {
		colorPtr = &color[0]
	}

	var indicesPtr *int32
	numIndices := len(indices)
	if numIndices > 0 {
		indicesPtr = &indices[0]
	}

	ret, _, _ := purego.SyscallN(sdlRenderGeometryRaw,
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(xyPtr)),
		uintptr(8), // xyStride -> unsafe.Sizeof(sdl.FPoint{}) == 8
		uintptr(unsafe.Pointer(colorPtr)),
		uintptr(16), // colorStride -> unsafe.Sizeof(sdl.FPoint{}) == 16
		uintptr(unsafe.Pointer(uvPtr)),
		uintptr(8), // uvStride -> unsafe.Sizeof(sdl.FPoint{}) == 8
		uintptr(numXY),
		uintptr(unsafe.Pointer(indicesPtr)),
		uintptr(numIndices),
		uintptr(4)) // sizeIndices -> unsafe.Sizeof(int32(0)) == 4

	return byte(ret) != 0
}

// [RenderLine] draws a line on the current rendering target at subpixel precision.
//
// [RenderLine]: https://wiki.libsdl.org/SDL3/SDL_RenderLine
func RenderLine(renderer *Renderer, x1 float32, y1 float32, x2 float32, y2 float32) bool {
	return sdlRenderLine(renderer, x1, y1, x2, y2)
}

// [RenderLines] draws a series of connected lines on the current rendering target at subpixel precision.
//
// [RenderLines]: https://wiki.libsdl.org/SDL3/SDL_RenderLines
func RenderLines(renderer *Renderer, points []FPoint) bool {
	count := len(points)
	var pointsPtr *FPoint
	if count > 0 {
		pointsPtr = &points[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderLines, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(pointsPtr)), uintptr(count))
	return byte(ret) != 0
}

// [RenderPoint] draws a point on the current rendering target at subpixel precision.
//
// [RenderPoint]: https://wiki.libsdl.org/SDL3/SDL_RenderPoint
func RenderPoint(renderer *Renderer, x float32, y float32) bool {
	return sdlRenderPoint(renderer, x, y)
}

// [RenderPoints] draws multiple points on the current rendering target at subpixel precision.
//
// [RenderPoints]: https://wiki.libsdl.org/SDL3/SDL_RenderPoints
func RenderPoints(renderer *Renderer, points []FPoint) bool {
	count := len(points)
	var pointsPtr *FPoint
	if count > 0 {
		pointsPtr = &points[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderPoints, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(pointsPtr)), uintptr(count))
	return byte(ret) != 0
}

// [RenderReadPixels] reads pixels from the current rendering target.
//
// [RenderReadPixels]: https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels
func RenderReadPixels(renderer *Renderer, rect *Rect) *Surface {
	return sdlRenderReadPixels(renderer, rect)
}

// [RenderRects] draws some number of rectangles on the current rendering target at subpixel precision.
//
// [RenderRects]: https://wiki.libsdl.org/SDL3/SDL_RenderRects
func RenderRects(renderer *Renderer, rects []FRect) bool {
	count := len(rects)
	var rectsPtr *FRect
	if count > 0 {
		rectsPtr = &rects[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderRects, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rectsPtr)), uintptr(count))
	return byte(ret) != 0
}

// [RenderTexture9Grid] performs a scaled copy using the 9-grid algorithm to the current rendering target at subpixel precision.
//
// [RenderTexture9Grid]: https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid
func RenderTexture9Grid(renderer *Renderer, texture *Texture, srcrect *FRect, leftWidth, rightWidth, topHeight, bottomHeight, scale float32, dstrect *FRect) bool {
	return sdlRenderTexture9Grid(renderer, texture, srcrect, leftWidth, rightWidth, topHeight, bottomHeight, scale, dstrect)
}

// [RenderTextureAffine] copies a portion of the source texture to the current rendering target, with affine transform, at subpixel precision.
//
// [RenderTextureAffine]: https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine
func RenderTextureAffine(renderer *Renderer, texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) bool {
	ret, _, _ := purego.SyscallN(
		sdlRenderTextureAffine,
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(origin)),
		uintptr(unsafe.Pointer(right)),
		uintptr(unsafe.Pointer(down)))

	return byte(ret) != 0
}

// [RenderTextureRotated] copies a portion of the source texture to the current rendering target, with rotation and flipping, at subpixel precision.
//
// [RenderTextureRotated]: https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated
func RenderTextureRotated(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) bool {
	return sdlRenderTextureRotated(renderer, texture, srcrect, dstrect, angle, center, flip)
}

// [RenderTextureTiled] tiles a portion of the texture to the current rendering target at subpixel precision.
//
// [RenderTextureTiled]: https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled
func RenderTextureTiled(renderer *Renderer, texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) bool {
	return sdlRenderTextureTiled(renderer, texture, srcrect, scale, dstrect)
}

// [RenderViewportSet] returns whether an explicit rectangle was set as the viewport.
//
// [RenderViewportSet]: https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet
func RenderViewportSet(renderer *Renderer) bool {
	return sdlRenderViewportSet(renderer)
}

// [SetRenderClipRect] sets the clip rectangle for rendering on the specified target.
//
// [SetRenderClipRect]: https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect
func SetRenderClipRect(renderer *Renderer, rect *Rect) bool {
	return sdlSetRenderClipRect(renderer, rect)
}

// [SetRenderColorScale] sets the color scale used for render operations.
//
// [SetRenderColorScale]: https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale
func SetRenderColorScale(renderer *Renderer, scale float32) bool {
	return sdlSetRenderColorScale(renderer, scale)
}

// [SetRenderDrawBlendMode] sets the blend mode used for drawing operations (Fill and Line).
//
// [SetRenderDrawBlendMode]: https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode
func SetRenderDrawBlendMode(renderer *Renderer, blendMode BlendMode) bool {
	return sdlSetRenderDrawBlendMode(renderer, blendMode)
}

// [SetRenderDrawColorFloat] sets the color used for drawing operations (Rect, Line and Clear).
//
// [SetRenderDrawColorFloat]: https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat
func SetRenderDrawColorFloat(renderer *Renderer, r float32, g float32, b float32, a float32) bool {
	return sdlSetRenderDrawColorFloat(renderer, r, g, b, a)
}

// [SetRenderLogicalPresentation] sets a device-independent resolution and presentation mode for rendering.
//
// [SetRenderLogicalPresentation]: https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation
func SetRenderLogicalPresentation(renderer *Renderer, w int32, h int32, mode RendererLogicalPresentation) bool {
	return sdlSetRenderLogicalPresentation(renderer, w, h, mode)
}

// [SetRenderScale] sets the drawing scale for rendering on the current target.
//
// [SetRenderScale]: https://wiki.libsdl.org/SDL3/SDL_SetRenderScale
func SetRenderScale(renderer *Renderer, scaleX float32, scaleY float32) bool {
	return sdlSetRenderScale(renderer, scaleX, scaleY)
}

// [SetRenderTarget] sets a texture as the current rendering target.
//
// [SetRenderTarget]: https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget
func SetRenderTarget(renderer *Renderer, texture *Texture) bool {
	return sdlSetRenderTarget(renderer, texture)
}

// [SetRenderViewport] sets the drawing area for rendering on the current target.
//
// [SetRenderViewport]: https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport
func SetRenderViewport(renderer *Renderer, rect *Rect) bool {
	return sdlSetRenderViewport(renderer, rect)
}

// [SetRenderVSync] sets VSync of the given renderer.
//
// [SetRenderVSync]: https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync
func SetRenderVSync(renderer *Renderer, vsync int32) bool {
	return sdlSetRenderVSync(renderer, vsync)
}

// [SetTextureAlphaMod] sets an additional alpha value multiplied into render copy operations.
//
// [SetTextureAlphaMod]: https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod
func SetTextureAlphaMod(texture *Texture, alpha uint8) bool {
	ret, _, _ := purego.SyscallN(sdlSetTextureAlphaMod, uintptr(unsafe.Pointer(texture)), uintptr(alpha))
	return byte(ret) != 0
}

// [SetTextureAlphaModFloat] sets an additional alpha value multiplied into render copy operations.
//
// [SetTextureAlphaModFloat]: https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat
func SetTextureAlphaModFloat(texture *Texture, alpha float32) bool {
	return sdlSetTextureAlphaModFloat(texture, alpha)
}

// [SetTextureBlendMode] sets the blend mode for a texture, used by [RenderTexture()].
//
// [SetTextureBlendMode]: https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode
func SetTextureBlendMode(texture *Texture, blendMode BlendMode) bool {
	return sdlSetTextureBlendMode(texture, blendMode)
}

// [SetTextureColorMod] sets an additional color value multiplied into render copy operations.
//
// [SetTextureColorMod]: https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod
func SetTextureColorMod(texture *Texture, r uint8, g uint8, b uint8) bool {
	ret, _, _ := purego.SyscallN(sdlSetTextureColorMod, uintptr(unsafe.Pointer(texture)), uintptr(r), uintptr(g), uintptr(b))
	return byte(ret) != 0
}

// [SetTextureColorModFloat] sets an additional color value multiplied into render copy operations.
//
// [SetTextureColorModFloat]: https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat
func SetTextureColorModFloat(texture *Texture, r float32, g float32, b float32) bool {
	return sdlSetTextureColorModFloat(texture, r, g, b)
}

// [SetTextureScaleMode] sets the scale mode used for texture scale operations.
//
// [SetTextureScaleMode]: https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode
func SetTextureScaleMode(texture *Texture, scaleMode ScaleMode) bool {
	return sdlSetTextureScaleMode(texture, scaleMode)
}

// [UnlockTexture] unlocks a texture, uploading the changes to video memory, if needed.
//
// [UnlockTexture]: https://wiki.libsdl.org/SDL3/SDL_UnlockTexture
func UnlockTexture(texture *Texture) {
	sdlUnlockTexture(texture)
}

// [UpdateNVTexture] updates a rectangle within a planar NV12 or NV21 texture with new pixels.
//
// [UpdateNVTexture]: https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture
func UpdateNVTexture(texture *Texture, rect *Rect, yPlane *uint8, yPitch int32, uvPlane *uint8, uvPitch int32) bool {
	ret, _, _ := purego.SyscallN(
		sdlUpdateNVTexture,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(yPlane)), uintptr(yPitch),
		uintptr(unsafe.Pointer(uvPlane)), uintptr(uvPitch))

	return byte(ret) != 0
}

// [UpdateTexture] updates the given texture rectangle with new pixel data.
//
// [UpdateTexture]: https://wiki.libsdl.org/SDL3/SDL_UpdateTexture
func UpdateTexture(texture *Texture, rect *Rect, pixels unsafe.Pointer, pitch int32) bool {
	ret, _, _ := purego.SyscallN(sdlUpdateTexture, uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(pixels), uintptr(pitch))
	return byte(ret) != 0
}

// [UpdateYUVTexture] updates a rectangle within a planar YV12 or IYUV texture with new pixel data.
//
// [UpdateYUVTexture]: https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture
func UpdateYUVTexture(texture *Texture, rect *Rect, yPlane *uint8, yPitch int32, uPlane *uint8, uPitch int32, vPlane *uint8, vPitch int32) bool {
	ret, _, _ := purego.SyscallN(
		sdlUpdateYUVTexture,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(yPlane)), uintptr(yPitch),
		uintptr(unsafe.Pointer(uPlane)), uintptr(uPitch),
		uintptr(unsafe.Pointer(vPlane)), uintptr(vPitch))

	return byte(ret) != 0
}
