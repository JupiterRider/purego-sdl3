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

type RendererLogicalPresentation uint32

const (
	LogicalPresentationDisabled RendererLogicalPresentation = iota
	LogicalPresentationStretch
	LogicalPresentationLetterbox
	LogicalPresentationOverscan
	LogicalPresentationIntegerScale
)

type TextureAccess uint32

const (
	TextureAccessStatic TextureAccess = iota
	TextureAccessStreaming
	TextureAccessTarget
)

type Renderer struct{}

// Texture is a efficient driver-specific representation of pixel data.
type Texture struct {
	Format   PixelFormat
	W        int32
	H        int32
	Refcount int32
}

type Vertex struct {
	Position FPoint
	Color    FColor
	TexCoord FPoint
}

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(title string, width, height int32, flags WindowFlags, window **Window, renderer **Renderer) bool {
	return sdlCreateWindowAndRenderer(title, width, height, flags, window, renderer)
}

// SetRenderDrawColor sets the color used for drawing operations.
func SetRenderDrawColor(renderer *Renderer, r, g, b, a uint8) bool {
	ret, _, _ := purego.SyscallN(sdlSetRenderDrawColor, uintptr(unsafe.Pointer(renderer)), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	return byte(ret) != 0
}

// RenderPresent updates the screen with any rendering performed since the previous call.
func RenderPresent(renderer *Renderer) bool {
	ret, _, _ := purego.SyscallN(sdlRenderPresent, uintptr(unsafe.Pointer(renderer)))
	return byte(ret) != 0
}

// RenderClear clears the current rendering target with the drawing color.
func RenderClear(renderer *Renderer) bool {
	ret, _, _ := purego.SyscallN(sdlRenderClear, uintptr(unsafe.Pointer(renderer)))
	return byte(ret) != 0
}

// DestroyRenderer destroys the rendering context for a window and free all associated textures.
func DestroyRenderer(renderer *Renderer) {
	sdlDestroyRenderer(renderer)
}

// RenderRect draws a rectangle on the current rendering target at subpixel precision.
func RenderRect(renderer *Renderer, rect *FRect) bool {
	ret, _, _ := purego.SyscallN(sdlRenderRect, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect)))
	return byte(ret) != 0
}

// RenderFillRect fills a rectangle on the current rendering target with the drawing color at subpixel precision.
func RenderFillRect(renderer *Renderer, rect *FRect) bool {
	ret, _, _ := purego.SyscallN(sdlRenderFillRect, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect)))
	return byte(ret) != 0
}

// RenderDebugText draws debug text to a [Renderer].
func RenderDebugText(renderer *Renderer, x, y float32, str string) bool {
	return sdlRenderDebugText(renderer, x, y, str)
}

// CreateTextureFromSurface creates a texture from an existing surface.
func CreateTextureFromSurface(renderer *Renderer, surface *Surface) *Texture {
	return sdlCreateTextureFromSurface(renderer, surface)
}

// RenderTexture copies a portion of the texture to the current rendering target at subpixel precision.
func RenderTexture(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) bool {
	ret, _, _ := purego.SyscallN(sdlRenderTexture, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dstrect)))
	return byte(ret) != 0
}

// DestroyTexture destroys the specified texture.
func DestroyTexture(texture *Texture) {
	sdlDestroyTexture(texture)
}

func AddVulkanRenderSemaphores(renderer *Renderer, waitStageMask uint32, waitSemaphore int64, signalSemaphore int64) bool {
	return sdlAddVulkanRenderSemaphores(renderer, waitStageMask, waitSemaphore, signalSemaphore)
}

func ConvertEventToRenderCoordinates(renderer *Renderer, event *Event) bool {
	return sdlConvertEventToRenderCoordinates(renderer, event)
}

// CreateRenderer creates a 2D rendering context for a window.
// The name parameter can be one driver, a comma-separated list of drivers or "" to let SDL choose one.
func CreateRenderer(window *Window, name string) *Renderer {
	return sdlCreateRenderer(window, convert.ToBytePtrNullable(name))
}

func CreateRendererWithProperties(props PropertiesID) *Renderer {
	return sdlCreateRendererWithProperties(props)
}

func CreateSoftwareRenderer(surface *Surface) *Renderer {
	return sdlCreateSoftwareRenderer(surface)
}

func CreateTexture(renderer *Renderer, format PixelFormat, access TextureAccess, w int32, h int32) *Texture {
	return sdlCreateTexture(renderer, format, access, w, h)
}

func CreateTextureWithProperties(renderer *Renderer, props PropertiesID) *Texture {
	return sdlCreateTextureWithProperties(renderer, props)
}

func FlushRenderer(renderer *Renderer) bool {
	ret, _, _ := purego.SyscallN(sdlFlushRenderer, uintptr(unsafe.Pointer(renderer)))
	return byte(ret) != 0
}

func GetCurrentRenderOutputSize(renderer *Renderer, w *int32, h *int32) bool {
	return sdlGetCurrentRenderOutputSize(renderer, w, h)
}

func GetNumRenderDrivers() int32 {
	return sdlGetNumRenderDrivers()
}

func GetRenderClipRect(renderer *Renderer, rect *Rect) bool {
	return sdlGetRenderClipRect(renderer, rect)
}

func GetRenderColorScale(renderer *Renderer, scale *float32) bool {
	return sdlGetRenderColorScale(renderer, scale)
}

func GetRenderDrawBlendMode(renderer *Renderer, blendMode *BlendMode) bool {
	return sdlGetRenderDrawBlendMode(renderer, blendMode)
}

func GetRenderDrawColor(renderer *Renderer, r *uint8, g *uint8, b *uint8, a *uint8) bool {
	return sdlGetRenderDrawColor(renderer, r, g, b, a)
}

func GetRenderDrawColorFloat(renderer *Renderer, r *float32, g *float32, b *float32, a *float32) bool {
	return sdlGetRenderDrawColorFloat(renderer, r, g, b, a)
}

func GetRenderDriver(index int32) string {
	return sdlGetRenderDriver(index)
}

func GetRenderer(window *Window) *Renderer {
	return sdlGetRenderer(window)
}

func GetRendererFromTexture(texture *Texture) *Renderer {
	return sdlGetRendererFromTexture(texture)
}

// GetRendererName returns the name of the selected renderer, or "" on failure.
func GetRendererName(renderer *Renderer) string {
	return sdlGetRendererName(renderer)
}

func GetRendererProperties(renderer *Renderer) PropertiesID {
	return sdlGetRendererProperties(renderer)
}

func GetRenderLogicalPresentation(renderer *Renderer, w *int32, h *int32, mode *RendererLogicalPresentation) bool {
	return sdlGetRenderLogicalPresentation(renderer, w, h, mode)
}

func GetRenderLogicalPresentationRect(renderer *Renderer, rect *FRect) bool {
	return sdlGetRenderLogicalPresentationRect(renderer, rect)
}

func GetRenderMetalCommandEncoder(renderer *Renderer) unsafe.Pointer {
	return sdlGetRenderMetalCommandEncoder(renderer)
}

func GetRenderMetalLayer(renderer *Renderer) unsafe.Pointer {
	return sdlGetRenderMetalLayer(renderer)
}

func GetRenderOutputSize(renderer *Renderer, w *int32, h *int32) bool {
	return sdlGetRenderOutputSize(renderer, w, h)
}

func GetRenderSafeArea(renderer *Renderer, rect *Rect) bool {
	return sdlGetRenderSafeArea(renderer, rect)
}

func GetRenderScale(renderer *Renderer, scaleX *float32, scaleY *float32) bool {
	return sdlGetRenderScale(renderer, scaleX, scaleY)
}

func GetRenderTarget(renderer *Renderer) *Texture {
	return sdlGetRenderTarget(renderer)
}

func GetRenderViewport(renderer *Renderer, rect *Rect) bool {
	return sdlGetRenderViewport(renderer, rect)
}

func GetRenderVSync(renderer *Renderer, vsync *int32) bool {
	return sdlGetRenderVSync(renderer, vsync)
}

func GetRenderWindow(renderer *Renderer) *Window {
	return sdlGetRenderWindow(renderer)
}

func GetTextureAlphaMod(texture *Texture, alpha *uint8) bool {
	return sdlGetTextureAlphaMod(texture, alpha)
}

func GetTextureAlphaModFloat(texture *Texture, alpha *float32) bool {
	return sdlGetTextureAlphaModFloat(texture, alpha)
}

func GetTextureBlendMode(texture *Texture, blendMode *BlendMode) bool {
	return sdlGetTextureBlendMode(texture, blendMode)
}

func GetTextureColorMod(texture *Texture, r *uint8, g *uint8, b *uint8) bool {
	return sdlGetTextureColorMod(texture, r, g, b)
}

func GetTextureColorModFloat(texture *Texture, r *float32, g *float32, b *float32) bool {
	return sdlGetTextureColorModFloat(texture, r, g, b)
}

func GetTextureProperties(texture *Texture) PropertiesID {
	return sdlGetTextureProperties(texture)
}

func GetTextureScaleMode(texture *Texture, scaleMode *ScaleMode) bool {
	return sdlGetTextureScaleMode(texture, scaleMode)
}

func GetTextureSize(texture *Texture, w *float32, h *float32) bool {
	return sdlGetTextureSize(texture, w, h)
}

func LockTexture(texture *Texture, rect *Rect, pixels *unsafe.Pointer, pitch *int32) bool {
	return sdlLockTexture(texture, rect, pixels, pitch)
}

func LockTextureToSurface(texture *Texture, rect *Rect, surface **Surface) bool {
	return sdlLockTextureToSurface(texture, rect, surface)
}

func RenderClipEnabled(renderer *Renderer) bool {
	return sdlRenderClipEnabled(renderer)
}

func RenderCoordinatesFromWindow(renderer *Renderer, windowX float32, windowY float32, x *float32, y *float32) bool {
	return sdlRenderCoordinatesFromWindow(renderer, windowX, windowY, x, y)
}

func RenderCoordinatesToWindow(renderer *Renderer, x float32, y float32, windowX *float32, windowY *float32) bool {
	return sdlRenderCoordinatesToWindow(renderer, x, y, windowX, windowY)
}

func RenderDebugTextFormat(renderer *Renderer, x float32, y float32, format string, a ...any) bool {
	return sdlRenderDebugTextFormat(renderer, x, y, fmt.Sprintf(format, a...))
}

func RenderFillRects(renderer *Renderer, rects []FRect) bool {
	count := len(rects)
	var rectsPtr *FRect
	if count > 0 {
		rectsPtr = &rects[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderFillRects, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rectsPtr)), uintptr(count))
	return byte(ret) != 0
}

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

func RenderGeometryRaw(renderer *Renderer, texture *Texture, xy []FPoint, color []FColor, uv []FPoint, indices []int32) bool {
	ret, _, _ := purego.SyscallN(sdlRenderGeometryRaw,
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(&xy[0])),
		uintptr(8), // xyStride -> unsafe.Sizeof(sdl.FPoint{}) == 8
		uintptr(unsafe.Pointer(&color[0])),
		uintptr(16), // colorStride -> unsafe.Sizeof(sdl.FPoint{}) == 16
		uintptr(unsafe.Pointer(&uv[0])),
		uintptr(8), // uvStride -> unsafe.Sizeof(sdl.FPoint{}) == 8
		uintptr(len(xy)),
		uintptr(unsafe.Pointer(&indices[0])),
		uintptr(len(indices)),
		uintptr(4)) // sizeIndices -> unsafe.Sizeof(int32(0)) == 4

	return byte(ret) != 0
}

func RenderLine(renderer *Renderer, x1 float32, y1 float32, x2 float32, y2 float32) bool {
	return sdlRenderLine(renderer, x1, y1, x2, y2)
}

func RenderLines(renderer *Renderer, points []FPoint) bool {
	count := len(points)
	var pointsPtr *FPoint
	if count > 0 {
		pointsPtr = &points[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderLines, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(pointsPtr)), uintptr(count))
	return byte(ret) != 0
}

func RenderPoint(renderer *Renderer, x float32, y float32) bool {
	return sdlRenderPoint(renderer, x, y)
}

func RenderPoints(renderer *Renderer, points []FPoint) bool {
	count := len(points)
	var pointsPtr *FPoint
	if count > 0 {
		pointsPtr = &points[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderPoints, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(pointsPtr)), uintptr(count))
	return byte(ret) != 0
}

func RenderReadPixels(renderer *Renderer, rect *Rect) *Surface {
	return sdlRenderReadPixels(renderer, rect)
}

func RenderRects(renderer *Renderer, rects []FRect) bool {
	count := len(rects)
	var rectsPtr *FRect
	if count > 0 {
		rectsPtr = &rects[0]
	}
	ret, _, _ := purego.SyscallN(sdlRenderRects, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rectsPtr)), uintptr(count))
	return byte(ret) != 0
}

func RenderTexture9Grid(renderer *Renderer, texture *Texture, srcrect *FRect, leftWidth, rightWidth, topHeight, bottomHeight, scale float32, dstrect *FRect) bool {
	return sdlRenderTexture9Grid(renderer, texture, srcrect, leftWidth, rightWidth, topHeight, bottomHeight, scale, dstrect)
}

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

func RenderTextureRotated(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) bool {
	return sdlRenderTextureRotated(renderer, texture, srcrect, dstrect, angle, center, flip)
}

func RenderTextureTiled(renderer *Renderer, texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) bool {
	return sdlRenderTextureTiled(renderer, texture, srcrect, scale, dstrect)
}

func RenderViewportSet(renderer *Renderer) bool {
	return sdlRenderViewportSet(renderer)
}

func SetRenderClipRect(renderer *Renderer, rect *Rect) bool {
	return sdlSetRenderClipRect(renderer, rect)
}

func SetRenderColorScale(renderer *Renderer, scale float32) bool {
	return sdlSetRenderColorScale(renderer, scale)
}

func SetRenderDrawBlendMode(renderer *Renderer, blendMode BlendMode) bool {
	return sdlSetRenderDrawBlendMode(renderer, blendMode)
}

func SetRenderDrawColorFloat(renderer *Renderer, r float32, g float32, b float32, a float32) bool {
	return sdlSetRenderDrawColorFloat(renderer, r, g, b, a)
}

func SetRenderLogicalPresentation(renderer *Renderer, w int32, h int32, mode RendererLogicalPresentation) bool {
	return sdlSetRenderLogicalPresentation(renderer, w, h, mode)
}

func SetRenderScale(renderer *Renderer, scaleX float32, scaleY float32) bool {
	return sdlSetRenderScale(renderer, scaleX, scaleY)
}

func SetRenderTarget(renderer *Renderer, texture *Texture) bool {
	return sdlSetRenderTarget(renderer, texture)
}

func SetRenderViewport(renderer *Renderer, rect *Rect) bool {
	return sdlSetRenderViewport(renderer, rect)
}

func SetRenderVSync(renderer *Renderer, vsync int32) bool {
	return sdlSetRenderVSync(renderer, vsync)
}

func SetTextureAlphaMod(texture *Texture, alpha uint8) bool {
	ret, _, _ := purego.SyscallN(sdlSetTextureAlphaMod, uintptr(unsafe.Pointer(texture)), uintptr(alpha))
	return byte(ret) != 0
}

func SetTextureAlphaModFloat(texture *Texture, alpha float32) bool {
	return sdlSetTextureAlphaModFloat(texture, alpha)
}

func SetTextureBlendMode(texture *Texture, blendMode BlendMode) bool {
	return sdlSetTextureBlendMode(texture, blendMode)
}

func SetTextureColorMod(texture *Texture, r uint8, g uint8, b uint8) bool {
	ret, _, _ := purego.SyscallN(sdlSetTextureColorMod, uintptr(unsafe.Pointer(texture)), uintptr(r), uintptr(g), uintptr(b))
	return byte(ret) != 0
}

func SetTextureColorModFloat(texture *Texture, r float32, g float32, b float32) bool {
	return sdlSetTextureColorModFloat(texture, r, g, b)
}

func SetTextureScaleMode(texture *Texture, scaleMode ScaleMode) bool {
	return sdlSetTextureScaleMode(texture, scaleMode)
}

func UnlockTexture(texture *Texture) {
	sdlUnlockTexture(texture)
}

func UpdateNVTexture(texture *Texture, rect *Rect, yPlane *uint8, yPitch int32, uvPlane *uint8, uvPitch int32) bool {
	ret, _, _ := purego.SyscallN(
		sdlUpdateNVTexture,
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(yPlane)), uintptr(yPitch),
		uintptr(unsafe.Pointer(uvPlane)), uintptr(uvPitch))

	return byte(ret) != 0
}

func UpdateTexture(texture *Texture, rect *Rect, pixels unsafe.Pointer, pitch int32) bool {
	ret, _, _ := purego.SyscallN(sdlUpdateTexture, uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(pixels), uintptr(pitch))
	return byte(ret) != 0
}

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
