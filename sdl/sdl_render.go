package sdl

type Renderer struct{}

// Texture is a efficient driver-specific representation of pixel data.
type Texture struct {
	Format   PixelFormat
	W        int32
	H        int32
	Refcount int32
}

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(title string, width, height int32, flags WindowFlags, window **Window, renderer **Renderer) bool {
	return sdlCreateWindowAndRenderer(title, width, height, flags, window, renderer)
}

// SetRenderDrawColor sets the color used for drawing operations.
func SetRenderDrawColor(renderer *Renderer, r, g, b, a uint8) bool {
	return sdlSetRenderDrawColor(renderer, r, g, b, a)
}

// SetDrawColor sets the color used for drawing operations.
func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) bool {
	return SetRenderDrawColor(renderer, r, g, b, a)
}

// RenderPresent updates the screen with any rendering performed since the previous call.
func RenderPresent(renderer *Renderer) bool {
	return sdlRenderPresent(renderer)
}

// Present updates the screen with any rendering performed since the previous call.
func (renderer *Renderer) Present() bool {
	return RenderPresent(renderer)
}

// RenderClear clears the current rendering target with the drawing color.
func RenderClear(renderer *Renderer) bool {
	return sdlRenderClear(renderer)
}

// Clear clears the current rendering target with the drawing color.
func (renderer *Renderer) Clear() bool {
	return RenderClear(renderer)
}

// DestroyRenderer destroys the rendering context for a window and free all associated textures.
func DestroyRenderer(renderer *Renderer) {
	sdlDestroyRenderer(renderer)
}

// Destroy destroys the rendering context for a window and free all associated textures.
func (renderer *Renderer) Destroy() {
	DestroyRenderer(renderer)
}

// RenderRect draws a rectangle on the current rendering target at subpixel precision.
func RenderRect(renderer *Renderer, rect *FRect) bool {
	return sdlRenderRect(renderer, rect)
}

// RenderRect draws a rectangle on the current rendering target at subpixel precision.
func (renderer *Renderer) RenderRect(rect *FRect) bool {
	return RenderRect(renderer, rect)
}

// RenderFillRect fills a rectangle on the current rendering target with the drawing color at subpixel precision.
func RenderFillRect(renderer *Renderer, rect *FRect) bool {
	return sdlRenderFillRect(renderer, rect)
}

// RenderFillRect fills a rectangle on the current rendering target with the drawing color at subpixel precision.
func (renderer *Renderer) RenderFillRect(rect *FRect) bool {
	return RenderFillRect(renderer, rect)
}

// RenderDebugText draws debug text to a [Renderer].
func RenderDebugText(renderer *Renderer, x, y float32, str string) bool {
	return sdlRenderDebugText(renderer, x, y, str)
}

// RenderDebugText draws debug text to a [Renderer].
func (renderer *Renderer) RenderDebugText(x, y float32, str string) bool {
	return RenderDebugText(renderer, x, y, str)
}

// CreateTextureFromSurface creates a texture from an existing surface.
func CreateTextureFromSurface(renderer *Renderer, surface *Surface) *Texture {
	return sdlCreateTextureFromSurface(renderer, surface)
}

// CreateTextureFromSurface creates a texture from an existing surface.
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) *Texture {
	return CreateTextureFromSurface(renderer, surface)
}

// RenderTexture copies a portion of the texture to the current rendering target at subpixel precision.
func RenderTexture(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) bool {
	return sdlRenderTexture(renderer, texture, srcrect, dstrect)
}

// RenderTexture copies a portion of the texture to the current rendering target at subpixel precision.
func (renderer *Renderer) RenderTexture(texture *Texture, srcrect *FRect, dstrect *FRect) bool {
	return RenderTexture(renderer, texture, srcrect, dstrect)
}

// DestroyTexture destroys the specified texture.
func DestroyTexture(texture *Texture) {
	sdlDestroyTexture(texture)
}

// Destroy destroys the specified texture.
func (texture *Texture) Destroy() {
	DestroyTexture(texture)
}

// func AddVulkanRenderSemaphores(renderer *Renderer, wait_stage_mask uint32, wait_semaphore int64, signal_semaphore int64) bool {
//	return sdlAddVulkanRenderSemaphores(renderer, wait_stage_mask, wait_semaphore, signal_semaphore)
// }

// func ConvertEventToRenderCoordinates(renderer *Renderer, event *Event) bool {
//	return sdlConvertEventToRenderCoordinates(renderer, event)
// }

// func CreateRenderer(window *Window, name string) *Renderer {
//	return sdlCreateRenderer(window, name)
// }

// func CreateRendererWithProperties(props PropertiesID) *Renderer {
//	return sdlCreateRendererWithProperties(props)
// }

// func CreateSoftwareRenderer(surface *Surface) *Renderer {
//	return sdlCreateSoftwareRenderer(surface)
// }

// func CreateTexture(renderer *Renderer, format PixelFormat, access TextureAccess, w int32, h int32) *Texture {
//	return sdlCreateTexture(renderer, format, access, w, h)
// }

// func CreateTextureWithProperties(renderer *Renderer, props PropertiesID) *Texture {
//	return sdlCreateTextureWithProperties(renderer, props)
// }

// func FlushRenderer(renderer *Renderer) bool {
//	return sdlFlushRenderer(renderer)
// }

// func GetCurrentRenderOutputSize(renderer *Renderer, w *int32, h *int32) bool {
//	return sdlGetCurrentRenderOutputSize(renderer, w, h)
// }

// func GetNumRenderDrivers() int32 {
//	return sdlGetNumRenderDrivers()
// }

// func GetRenderClipRect(renderer *Renderer, rect *Rect) bool {
//	return sdlGetRenderClipRect(renderer, rect)
// }

// func GetRenderColorScale(renderer *Renderer, scale *float32) bool {
//	return sdlGetRenderColorScale(renderer, scale)
// }

// func GetRenderDrawBlendMode(renderer *Renderer, blendMode *BlendMode) bool {
//	return sdlGetRenderDrawBlendMode(renderer, blendMode)
// }

// func GetRenderDrawColor(renderer *Renderer, r *uint8, g *uint8, b *uint8, a *uint8) bool {
//	return sdlGetRenderDrawColor(renderer, r, g, b, a)
// }

// func GetRenderDrawColorFloat(renderer *Renderer, r *float32, g *float32, b *float32, a *float32) bool {
//	return sdlGetRenderDrawColorFloat(renderer, r, g, b, a)
// }

// func GetRenderDriver(index int32) string {
//	return sdlGetRenderDriver(index)
// }

// func GetRenderer(window *Window) *Renderer {
//	return sdlGetRenderer(window)
// }

// func GetRendererFromTexture(texture *Texture) *Renderer {
//	return sdlGetRendererFromTexture(texture)
// }

// func GetRendererName(renderer *Renderer) string {
//	return sdlGetRendererName(renderer)
// }

// func GetRendererProperties(renderer *Renderer) PropertiesID {
//	return sdlGetRendererProperties(renderer)
// }

// func GetRenderLogicalPresentation(renderer *Renderer, w *int32, h *int32, mode *RendererLogicalPresentation) bool {
//	return sdlGetRenderLogicalPresentation(renderer, w, h, mode)
// }

// func GetRenderLogicalPresentationRect(renderer *Renderer, rect *FRect) bool {
//	return sdlGetRenderLogicalPresentationRect(renderer, rect)
// }

// func GetRenderMetalCommandEncoder(renderer *Renderer) unsafe.Pointer {
//	return sdlGetRenderMetalCommandEncoder(renderer)
// }

// func GetRenderMetalLayer(renderer *Renderer) unsafe.Pointer {
//	return sdlGetRenderMetalLayer(renderer)
// }

// func GetRenderOutputSize(renderer *Renderer, w *int32, h *int32) bool {
//	return sdlGetRenderOutputSize(renderer, w, h)
// }

// func GetRenderSafeArea(renderer *Renderer, rect *Rect) bool {
//	return sdlGetRenderSafeArea(renderer, rect)
// }

// func GetRenderScale(renderer *Renderer, scaleX *float32, scaleY *float32) bool {
//	return sdlGetRenderScale(renderer, scaleX, scaleY)
// }

// func GetRenderTarget(renderer *Renderer) *Texture {
//	return sdlGetRenderTarget(renderer)
// }

// func GetRenderViewport(renderer *Renderer, rect *Rect) bool {
//	return sdlGetRenderViewport(renderer, rect)
// }

// func GetRenderVSync(renderer *Renderer, vsync *int32) bool {
//	return sdlGetRenderVSync(renderer, vsync)
// }

// func GetRenderWindow(renderer *Renderer) *Window {
//	return sdlGetRenderWindow(renderer)
// }

// func GetTextureAlphaMod(texture *Texture, alpha *uint8) bool {
//	return sdlGetTextureAlphaMod(texture, alpha)
// }

// func GetTextureAlphaModFloat(texture *Texture, alpha *float32) bool {
//	return sdlGetTextureAlphaModFloat(texture, alpha)
// }

// func GetTextureBlendMode(texture *Texture, blendMode *BlendMode) bool {
//	return sdlGetTextureBlendMode(texture, blendMode)
// }

// func GetTextureColorMod(texture *Texture, r *uint8, g *uint8, b *uint8) bool {
//	return sdlGetTextureColorMod(texture, r, g, b)
// }

// func GetTextureColorModFloat(texture *Texture, r *float32, g *float32, b *float32) bool {
//	return sdlGetTextureColorModFloat(texture, r, g, b)
// }

// func GetTextureProperties(texture *Texture) PropertiesID {
//	return sdlGetTextureProperties(texture)
// }

// func GetTextureScaleMode(texture *Texture, scaleMode *ScaleMode) bool {
//	return sdlGetTextureScaleMode(texture, scaleMode)
// }

// func GetTextureSize(texture *Texture, w *float32, h *float32) bool {
//	return sdlGetTextureSize(texture, w, h)
// }

// func LockTexture(texture *Texture, rect *Rect, pixels *unsafe.Pointer, pitch *int32) bool {
//	return sdlLockTexture(texture, rect, pixels, pitch)
// }

// func LockTextureToSurface(texture *Texture, rect *Rect, surface **Surface) bool {
//	return sdlLockTextureToSurface(texture, rect, surface)
// }

// func RenderClipEnabled(renderer *Renderer) bool {
//	return sdlRenderClipEnabled(renderer)
// }

// func RenderCoordinatesFromWindow(renderer *Renderer, window_x float32, window_y float32, x *float32, y *float32) bool {
//	return sdlRenderCoordinatesFromWindow(renderer, window_x, window_y, x, y)
// }

// func RenderCoordinatesToWindow(renderer *Renderer, x float32, y float32, window_x *float32, window_y *float32) bool {
//	return sdlRenderCoordinatesToWindow(renderer, x, y, window_x, window_y)
// }

// func RenderDebugTextFormat(renderer *Renderer, x float32, y float32, fmt string) bool {
//	return sdlRenderDebugTextFormat(renderer, x, y, fmt)
// }

// func RenderFillRects(renderer *Renderer, rects *FRect, count int32) bool {
//	return sdlRenderFillRects(renderer, rects, count)
// }

// func RenderGeometry(renderer *Renderer, texture *Texture, vertices *Vertex, num_vertices int32, indices *int32, num_indices int32) bool {
//	return sdlRenderGeometry(renderer, texture, vertices, num_vertices, indices, num_indices)
// }

// func RenderGeometryRaw(renderer *Renderer, texture *Texture, xy *float32, xy_stride int32, color *FColor, color_stride int32, uv *float32, uv_stride int32, num_vertices int32, indices unsafe.Pointer, num_indices int32, size_indices int32) bool {
//	return sdlRenderGeometryRaw(renderer, texture, xy, xy_stride, color, color_stride, uv, uv_stride, num_vertices, indices, num_indices, size_indices)
// }

// func RenderLine(renderer *Renderer, x1 float32, y1 float32, x2 float32, y2 float32) bool {
//	return sdlRenderLine(renderer, x1, y1, x2, y2)
// }

// func RenderLines(renderer *Renderer, points *FPoint, count int32) bool {
//	return sdlRenderLines(renderer, points, count)
// }

// func RenderPoint(renderer *Renderer, x float32, y float32) bool {
//	return sdlRenderPoint(renderer, x, y)
// }

// func RenderPoints(renderer *Renderer, points *FPoint, count int32) bool {
//	return sdlRenderPoints(renderer, points, count)
// }

// func RenderReadPixels(renderer *Renderer, rect *Rect) *Surface {
//	return sdlRenderReadPixels(renderer, rect)
// }

// func RenderRects(renderer *Renderer, rects *FRect, count int32) bool {
//	return sdlRenderRects(renderer, rects, count)
// }

// func RenderTexture9Grid(renderer *Renderer, texture *Texture, srcrect *FRect, left_width float32, right_width float32, top_height float32, bottom_height float32, scale float32, dstrect *FRect) bool {
//	return sdlRenderTexture9Grid(renderer, texture, srcrect, left_width, right_width, top_height, bottom_height, scale, dstrect)
// }

// func RenderTextureAffine(renderer *Renderer, texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) bool {
//	return sdlRenderTextureAffine(renderer, texture, srcrect, origin, right, down)
// }

// func RenderTextureRotated(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) bool {
//	return sdlRenderTextureRotated(renderer, texture, srcrect, dstrect, angle, center, flip)
// }

// func RenderTextureTiled(renderer *Renderer, texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) bool {
//	return sdlRenderTextureTiled(renderer, texture, srcrect, scale, dstrect)
// }

// func RenderViewportSet(renderer *Renderer) bool {
//	return sdlRenderViewportSet(renderer)
// }

// func SetRenderClipRect(renderer *Renderer, rect *Rect) bool {
//	return sdlSetRenderClipRect(renderer, rect)
// }

// func SetRenderColorScale(renderer *Renderer, scale float32) bool {
//	return sdlSetRenderColorScale(renderer, scale)
// }

// func SetRenderDrawBlendMode(renderer *Renderer, blendMode BlendMode) bool {
//	return sdlSetRenderDrawBlendMode(renderer, blendMode)
// }

// func SetRenderDrawColorFloat(renderer *Renderer, r float32, g float32, b float32, a float32) bool {
//	return sdlSetRenderDrawColorFloat(renderer, r, g, b, a)
// }

// func SetRenderLogicalPresentation(renderer *Renderer, w int32, h int32, mode RendererLogicalPresentation) bool {
//	return sdlSetRenderLogicalPresentation(renderer, w, h, mode)
// }

// func SetRenderScale(renderer *Renderer, scaleX float32, scaleY float32) bool {
//	return sdlSetRenderScale(renderer, scaleX, scaleY)
// }

// func SetRenderTarget(renderer *Renderer, texture *Texture) bool {
//	return sdlSetRenderTarget(renderer, texture)
// }

// func SetRenderViewport(renderer *Renderer, rect *Rect) bool {
//	return sdlSetRenderViewport(renderer, rect)
// }

// func SetRenderVSync(renderer *Renderer, vsync int32) bool {
//	return sdlSetRenderVSync(renderer, vsync)
// }

// func SetTextureAlphaMod(texture *Texture, alpha uint8) bool {
//	return sdlSetTextureAlphaMod(texture, alpha)
// }

// func SetTextureAlphaModFloat(texture *Texture, alpha float32) bool {
//	return sdlSetTextureAlphaModFloat(texture, alpha)
// }

// func SetTextureBlendMode(texture *Texture, blendMode BlendMode) bool {
//	return sdlSetTextureBlendMode(texture, blendMode)
// }

// func SetTextureColorMod(texture *Texture, r uint8, g uint8, b uint8) bool {
//	return sdlSetTextureColorMod(texture, r, g, b)
// }

// func SetTextureColorModFloat(texture *Texture, r float32, g float32, b float32) bool {
//	return sdlSetTextureColorModFloat(texture, r, g, b)
// }

// func SetTextureScaleMode(texture *Texture, scaleMode ScaleMode) bool {
//	return sdlSetTextureScaleMode(texture, scaleMode)
// }

// func UnlockTexture(texture *Texture)  {
//	sdlUnlockTexture(texture)
// }

// func UpdateNVTexture(texture *Texture, rect *Rect, Yplane *uint8, Ypitch int32, UVplane *uint8, UVpitch int32) bool {
//	return sdlUpdateNVTexture(texture, rect, Yplane, Ypitch, UVplane, UVpitch)
// }

// func UpdateTexture(texture *Texture, rect *Rect, pixels unsafe.Pointer, pitch int32) bool {
//	return sdlUpdateTexture(texture, rect, pixels, pitch)
// }

// func UpdateYUVTexture(texture *Texture, rect *Rect, Yplane *uint8, Ypitch int32, Uplane *uint8, Upitch int32, Vplane *uint8, Vpitch int32) bool {
//	return sdlUpdateYUVTexture(texture, rect, Yplane, Ypitch, Uplane, Upitch, Vplane, Vpitch)
// }
