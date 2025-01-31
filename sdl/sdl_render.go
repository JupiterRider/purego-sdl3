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

// SetRenderDrawColor sets the color used for drawing operations.
func (renderer *Renderer) SetRenderDrawColor(r, g, b, a uint8) bool {
	return SetRenderDrawColor(renderer, r, g, b, a)
}

// RenderPresent updates the screen with any rendering performed since the previous call.
func RenderPresent(renderer *Renderer) bool {
	return sdlRenderPresent(renderer)
}

// RenderPresent updates the screen with any rendering performed since the previous call.
func (renderer *Renderer) RenderPresent() bool {
	return RenderPresent(renderer)
}

// RenderClear clears the current rendering target with the drawing color.
func RenderClear(renderer *Renderer) bool {
	return sdlRenderClear(renderer)
}

// RenderClear clears the current rendering target with the drawing color.
func (renderer *Renderer) RenderClear() bool {
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
func RenderTexture(renderer *Renderer, texture *Texture, srcrect, dstrect *FRect) bool {
	return sdlRenderTexture(renderer, texture, srcrect, dstrect)
}

// RenderTexture copies a portion of the texture to the current rendering target at subpixel precision.
func (renderer *Renderer) RenderTexture(texture *Texture, srcrect, dstrect *FRect) bool {
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
