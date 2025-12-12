package ttf

import (
	"github.com/jupiterrider/purego-sdl3/internal/convert"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

const (
	PropFontCreateFilenameString           = "SDL_ttf.font.create.filename"
	PropFontCreateIOStreamPointer          = "SDL_ttf.font.create.iostream"
	PropFontCreateIOStreamOffsetNumber     = "SDL_ttf.font.create.iostream.offset"
	PropFontCreateIOStreamAutocloseBoolean = "SDL_ttf.font.create.iostream.autoclose"
	PropFontCreateSizeFloat                = "SDL_ttf.font.create.size"
	PropFontCreateFaceNumber               = "SDL_ttf.font.create.face"
	PropFontCreateHorizontalDpiNumber      = "SDL_ttf.font.create.hdpi"
	PropFontCreateVerticalDpiNumber        = "SDL_ttf.font.create.vdpi"
	PropFontCreateExistingFont             = "SDL_ttf.font.create.existing_font"

	PropFontOutlineLineCapNumber    = "SDL_ttf.font.outline.line_cap"
	PropFontOutlineLineJoinNumber   = "SDL_ttf.font.outline.line_join"
	PropFontOutlineMiterLimitNumber = "SDL_ttf.font.outline.miter_limit"

	PropRendererTextEngineRenderer         = "SDL_ttf.renderer_text_engine.create.renderer"
	PropRendererTextEngineAtlasTextureSize = "SDL_ttf.renderer_text_engine.create.atlas_texture_size"

	PropGPUTextEngineDevice           = "SDL_ttf.gpu_text_engine.create.device"
	PropGPUTextEngineAtlasTextureSize = "SDL_ttf.gpu_text_engine.create.atlas_texture_size"
)

// [FontStyleFlags] defines the font style flags for [Font].
//
// [FontStyleFlags]: https://wiki.libsdl.org/SDL3_ttf/TTF_FontStyleFlags
type FontStyleFlags uint32

const (
	StyleNormal        = 0x00 // No special style.
	StyleBold          = 0x01 // Bold style.
	StyleItalic        = 0x02 // Italic style.
	StyleUnderline     = 0x04 // Underlined text.
	StyleStrikethrough = 0x08 // Strikethrough text.
)

// [SubStringFlags] defines the flags for [SubString].
//
// [SubStringFlags]: https://wiki.libsdl.org/SDL3_ttf/TTF_SubStringFlags
type SubStringFlags uint32

const (
	SubStringDirectionMask = 0x000000FF // The mask for the flow direction for this substring.
	SubStringTextStart     = 0x00000100 // This substring contains the beginning of the text.
	SubStringLineStart     = 0x00000200 // This substring contains the beginning of line `line_index`.
	SubStringLineEnd       = 0x00000400 // This substring contains the end of line `line_index`.
	SubStringTextEnd       = 0x00000800 // This substring contains the end of the text.
)

// [Direction] defines the direction flags.
//
// [Direction]: https://wiki.libsdl.org/SDL3_ttf/TTF_Direction
type Direction uint32

const (
	DirectionInvalid Direction = 0
	DirectionLtr     Direction = iota + 3 // Left to Right.
	DirectionRtl                          // Right to Left.
	DirectionTtb                          // Top to Bottom.
	DirectionBtt                          // Bottom to Top.
)

// [HintingFlags] defines the hinting flags for TTF (TrueType Fonts).
//
// [HintingFlags]: https://wiki.libsdl.org/SDL3_ttf/TTF_HintingFlags
type HintingFlags uint32

const (
	HintingNormal        HintingFlags = iota // Normal hinting applies standard grid-fitting.
	HintingLight                             // Light hinting applies subtle adjustments to improve rendering.
	HintingMono                              // Monochrome hinting adjusts the font for better rendering at lower resolutions.
	HintingNone                              // No hinting, the font is rendered without any grid-fitting.
	HintingLightSubpixel                     // Light hinting with subpixel rendering for more precise font edges.
)

// [HorizontalAlignment] is a structure specifying the horizontal alignment used when rendering wrapped texts.
//
// [HorizontalAlignment]: https://wiki.libsdl.org/SDL3_ttf/TTF_HorizontalAlignment
type HorizontalAlignment int32

const (
	HorizontalAlignInvalid HorizontalAlignment = iota - 1
	HorizontalAlignLeft
	HorizontalAlignCenter
	HorizontalAlignRight
)

// [ImageType] is a structure specifying the type of data in a glyph images.
//
// [ImageType]: https://wiki.libsdl.org/SDL3_ttf/TTF_ImageType
type ImageType uint32

const (
	ImageInvalid ImageType = iota
	ImageAlpha             // The color channels are white.
	ImageColor             // The color channels have image data.
	ImageSDF               // The alpha channel has signed distance field information.
)

// [GPUTextEngineWinding] is a structure specifying the winding order of the vertices returned by [GetGPUTextDrawData].
//
// [GPUTextEngineWinding]: https://wiki.libsdl.org/SDL3_ttf/TTF_GPUTextEngineWinding
type GPUTextEngineWinding int32

const (
	GPUTextEngineWindingInvalid GPUTextEngineWinding = iota - 1
	GPUTextEngineWindingClockwise
	GPUTextEngineWindingCounterClockwise
)

// [Font] is a structure specifying the internal structure containing font information.
//
// [Font]: https://wiki.libsdl.org/SDL3_ttf/TTF_Font
type Font struct{}

// [TextData] defines the internal data for [Text].
//
// [TextData]: https://wiki.libsdl.org/SDL3_ttf/TTF_TextData
type TextData struct{}

// [TextEngine] is a text engine used to create text objects.
//
// [TextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_TextEngine
type TextEngine struct{}

// [Text] defines the text created with [CreateText].
//
// [Text]: https://wiki.libsdl.org/SDL3_ttf/TTF_Text
type Text struct {
	text     *byte     // A copy of the UTF-8 string that this text object represents, useful for layout, debugging and retrieving substring text. This is updated when the text object is modified and will be freed automatically when the object is destroyed.
	NumLines int32     // The number of lines in the text, 0 if it's empty.
	Refcount int32     // Application reference count, used when freeing surface.
	Internal *TextData // Private.
}

// Text gets a copy of the UTF-8 string that this text object represents, useful for layout, debugging and retrieving substring text. This is updated when the text object is modified and will be freed automatically when the object is destroyed.
func (t *Text) Text() string {
	return convert.ToString(t.text)
}

// [SubString] is a structure specifying the representation of a substring within texts.
//
// [SubString]: https://wiki.libsdl.org/SDL3_ttf/TTF_SubString
type SubString struct {
	Flags        SubStringFlags // The flags for this substring.
	Offset       int32          // The byte offset from the beginning of the text.
	Length       int32          // The byte length starting at the offset.
	LineIndex    int32          // The index of the line that contains this substring.
	ClusterIndex int32          // The internal cluster index, used for quickly iterating.
	Rect         sdl.Rect       // The rectangle, relative to the top left of the text, containing the substring.
}

// [GPUAtlasDrawSequence] defines the draw sequence returned by [GetGPUTextDrawData].
//
// [GPUAtlasDrawSequence]: https://wiki.libsdl.org/SDL3_ttf/TTF_GPUAtlasDrawSequence
type GPUAtlasDrawSequence struct {
	AtlasTexture *sdl.GPUTexture       // Texture atlas that stores the glyphs.
	XY           *sdl.FPoint           // An array of vertex positions.
	UV           *sdl.FPoint           // An array of normalized texture coordinates for each vertex.
	NumVertices  int32                 // Number of vertices.
	Indices      *int32                // An array of indices into the 'vertices' arrays.
	NumIndices   int32                 // Number of indices.
	ImageType    ImageType             // The image type of this draw sequence.
	Next         *GPUAtlasDrawSequence // The next sequence (will be nil in case of the last sequence).
}

// [Version] gets the version of the dynamically linked SDL_ttf library.
//
// [Version]: https://wiki.libsdl.org/SDL3_ttf/TTF_Version
func Version() (major, minor, patch int32) {
	version := ttfVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}

// [GetFreeTypeVersion] queries the version of the FreeType library in use.
//
// [GetFreeTypeVersion]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion
func GetFreeTypeVersion() (major, minor, patch int32) {
	ttfGetFreeTypeVersion(&major, &minor, &patch)
	return
}

// [Init] initializes SDL_ttf.
//
// [Init]: https://wiki.libsdl.org/SDL3_ttf/TTF_Init
func Init() bool {
	return ttfInit()
}

// [GetHarfBuzzVersion] queries the version of the HarfBuzz library in use.
//
// [GetHarfBuzzVersion]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion
func GetHarfBuzzVersion() (major, minor, patch int32) {
	ttfGetHarfBuzzVersion(&major, &minor, &patch)
	return
}

// [AddFallbackFont] adds a fallback font.
//
// [AddFallbackFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont
func AddFallbackFont(font *Font, fallback *Font) bool {
	return ttfAddFallbackFont(font, fallback)
}

// [AppendTextString] appends UTF-8 text to a text object.
//
// [AppendTextString]: https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString
func AppendTextString(text *Text, str string, length uint64) bool {
	return ttfAppendTextString(text, str, length)
}

// [ClearFallbackFonts] removes all fallback fonts.
//
// [ClearFallbackFonts]: https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts
func ClearFallbackFonts(font *Font) {
	ttfClearFallbackFonts(font)
}

// [CloseFont] disposes of a previously-created font.
//
// [CloseFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont
func CloseFont(font *Font) {
	ttfCloseFont(font)
}

// [CopyFont] creates a copy of an existing font.
//
// [CopyFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont
func CopyFont(existingFont *Font) *Font {
	return ttfCopyFont(existingFont)
}

// [CreateGPUTextEngine] creates a text engine for drawing text with the SDL GPU API.
//
// [CreateGPUTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngine
func CreateGPUTextEngine(device *sdl.GPUDevice) *TextEngine {
	return ttfCreateGPUTextEngine(device)
}

// [CreateGPUTextEngineWithProperties] creates a text engine for drawing text with the SDL GPU API, with the specified properties.
//
// [CreateGPUTextEngineWithProperties]: https://wiki.libsdl.org/SDL3_ttf/TTF_CreateGPUTextEngineWithProperties
func CreateGPUTextEngineWithProperties(props sdl.PropertiesID) *TextEngine {
	return ttfCreateGPUTextEngineWithProperties(props)
}

// [CreateRendererTextEngine] creates a text engine for drawing text on an SDL renderer.
//
// [CreateRendererTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine
func CreateRendererTextEngine(renderer *sdl.Renderer) *TextEngine {
	return ttfCreateRendererTextEngine(renderer)
}

// [CreateRendererTextEngineWithProperties] creates a text engine for drawing text on an SDL renderer, with the specified properties.
//
// [CreateRendererTextEngineWithProperties]: https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties
func CreateRendererTextEngineWithProperties(props sdl.PropertiesID) *TextEngine {
	return ttfCreateRendererTextEngineWithProperties(props)
}

// [CreateSurfaceTextEngine] creates a text engine for drawing text on SDL surfaces.
//
// [CreateSurfaceTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine
func CreateSurfaceTextEngine() *TextEngine {
	return ttfCreateSurfaceTextEngine()
}

// [CreateText] creates a text object from UTF-8 text and a text engine.
//
// [CreateText]: https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText
func CreateText(engine *TextEngine, font *Font, text string, length uint64) *Text {
	return ttfCreateText(engine, font, text, length)
}

// [DeleteTextString] deletes UTF-8 text from a text object.
//
// [DeleteTextString]: https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString
func DeleteTextString(text *Text, offset int32, length int32) bool {
	return ttfDeleteTextString(text, offset, length)
}

// [DestroyGPUTextEngine] destroys a text engine created for drawing text with the SDL GPU API.
//
// [DestroyGPUTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine
func DestroyGPUTextEngine(engine *TextEngine) {
	ttfDestroyGPUTextEngine(engine)
}

// [DestroyRendererTextEngine] destroys a text engine created for drawing text on an SDL renderer.
//
// [DestroyRendererTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine
func DestroyRendererTextEngine(engine *TextEngine) {
	ttfDestroyRendererTextEngine(engine)
}

// [DestroySurfaceTextEngine] destroys a text engine created for drawing text on SDL surfaces.
//
// [DestroySurfaceTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine
func DestroySurfaceTextEngine(engine *TextEngine) {
	ttfDestroySurfaceTextEngine(engine)
}

// [DestroyText] destroys a text object created by a text engine.
//
// [DestroyText]: https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText
func DestroyText(text *Text) {
	ttfDestroyText(text)
}

// [DrawRendererText] draws text to an SDL renderer.
//
// [DrawRendererText]: https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText
func DrawRendererText(text *Text, x float32, y float32) bool {
	return ttfDrawRendererText(text, x, y)
}

// [DrawSurfaceText] draws text to an SDL surface.
//
// [DrawSurfaceText]: https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText
func DrawSurfaceText(text *Text, x int32, y int32, surface *sdl.Surface) bool {
	return ttfDrawSurfaceText(text, x, y, surface)
}

// [FontHasGlyph] checks whether a glyph is provided by the font for a UNICODE codepoint.
//
// [FontHasGlyph]: https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph
func FontHasGlyph(font *Font, ch rune) bool {
	return ttfFontHasGlyph(font, ch)
}

// [FontIsFixedWidth] queries whether a font is fixed-width.
//
// [FontIsFixedWidth]: https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth
func FontIsFixedWidth(font *Font) bool {
	return ttfFontIsFixedWidth(font)
}

// [FontIsScalable] queries whether a font is scalable or not.
//
// [FontIsScalable]: https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable
func FontIsScalable(font *Font) bool {
	return ttfFontIsScalable(font)
}

// [GetFontAscent] queries the offset from the baseline to the top of a font.
//
// [GetFontAscent]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent
func GetFontAscent(font *Font) int32 {
	return ttfGetFontAscent(font)
}

// [GetFontDescent] queries the offset from the baseline to the bottom of a font.
//
// [GetFontDescent]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent
func GetFontDescent(font *Font) int32 {
	return ttfGetFontDescent(font)
}

// [GetFontDirection] gets the direction to be used for text shaping by a font.
//
// [GetFontDirection]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection
func GetFontDirection(font *Font) Direction {
	return ttfGetFontDirection(font)
}

// [GetFontDPI] gets font target resolutions, in dots per inch.
//
// [GetFontDPI]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI
func GetFontDPI(font *Font, hdpi *int32, vdpi *int32) bool {
	return ttfGetFontDPI(font, hdpi, vdpi)
}

// [GetFontFamilyName] queries a font's family name.
//
// [GetFontFamilyName]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName
func GetFontFamilyName(font *Font) string {
	return ttfGetFontFamilyName(font)
}

// [GetFontGeneration] gets the font generation.
//
// [GetFontGeneration]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration
func GetFontGeneration(font *Font) uint32 {
	return ttfGetFontGeneration(font)
}

// [GetFontHeight] queries the total height of a font.
//
// [GetFontHeight]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight
func GetFontHeight(font *Font) int32 {
	return ttfGetFontHeight(font)
}

// [GetFontHinting] queries a font's current FreeType hinter setting.
//
// [GetFontHinting]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting
func GetFontHinting(font *Font) HintingFlags {
	return ttfGetFontHinting(font)
}

// [GetFontKerning] queries whether or not kerning is enabled for a font.
//
// [GetFontKerning]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning
func GetFontKerning(font *Font) bool {
	return ttfGetFontKerning(font)
}

// [GetFontLineSkip] queries the spacing between lines of text for a font.
//
// [GetFontLineSkip]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip
func GetFontLineSkip(font *Font) int32 {
	return ttfGetFontLineSkip(font)
}

// [GetFontOutline] queries a font's current outline.
//
// [GetFontOutline]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline
func GetFontOutline(font *Font) int32 {
	return ttfGetFontOutline(font)
}

// [GetFontProperties] gets the properties associated with a font.
//
// [GetFontProperties]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties
func GetFontProperties(font *Font) sdl.PropertiesID {
	return ttfGetFontProperties(font)
}

// [GetFontScript] gets the script used for text shaping a font.
//
// [GetFontScript]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript
func GetFontScript(font *Font) uint32 {
	return ttfGetFontScript(font)
}

// [GetFontSDF] queries whether Signed Distance Field rendering is enabled for a font.
//
// [GetFontSDF]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF
func GetFontSDF(font *Font) bool {
	return ttfGetFontSDF(font)
}

// [GetFontSize] gets the size of a font.
//
// [GetFontSize]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize
func GetFontSize(font *Font) float32 {
	return ttfGetFontSize(font)
}

// [GetFontStyle] queries a font's current style.
//
// [GetFontStyle]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle
func GetFontStyle(font *Font) FontStyleFlags {
	return ttfGetFontStyle(font)
}

// [GetFontStyleName] queries a font's style name.
//
// [GetFontStyleName]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName
func GetFontStyleName(font *Font) string {
	return ttfGetFontStyleName(font)
}

// [GetFontWrapAlignment] queries a font's current wrap alignment option.
//
// [GetFontWrapAlignment]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment
func GetFontWrapAlignment(font *Font) HorizontalAlignment {
	return ttfGetFontWrapAlignment(font)
}

// [GetGlyphImage] gets the pixel image for a UNICODE codepoint.
//
// [GetGlyphImage]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage
func GetGlyphImage(font *Font, ch rune, imageType *ImageType) *sdl.Surface {
	return ttfGetGlyphImage(font, ch, imageType)
}

// [GetGlyphImageForIndex] gets the pixel image for a character index.
//
// [GetGlyphImageForIndex]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex
func GetGlyphImageForIndex(font *Font, glyphIndex uint32, imageType *ImageType) *sdl.Surface {
	return ttfGetGlyphImageForIndex(font, glyphIndex, imageType)
}

// [GetGlyphKerning] queries the kerning size between the glyphs of two UNICODE codepoints.
//
// [GetGlyphKerning]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning
func GetGlyphKerning(font *Font, previousCh, ch rune, kerning *int32) bool {
	return ttfGetGlyphKerning(font, previousCh, ch, kerning)
}

// [GetGlyphMetrics] queries the metrics (dimensions) of a font's glyph for a UNICODE codepoint.
//
// [GetGlyphMetrics]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics
func GetGlyphMetrics(font *Font, ch rune, minx *int32, maxx *int32, miny *int32, maxy *int32, advance *int32) bool {
	return ttfGetGlyphMetrics(font, ch, minx, maxx, miny, maxy, advance)
}

// [GetGlyphScript] gets the script used by a 32-bit codepoint.
//
// [GetGlyphScript]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript
func GetGlyphScript(ch rune) uint32 {
	return ttfGetGlyphScript(ch)
}

// [GetGPUTextDrawData] gets the geometry data needed for drawing the text.
//
// [GetGPUTextDrawData]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData
func GetGPUTextDrawData(text *Text) *GPUAtlasDrawSequence {
	return ttfGetGPUTextDrawData(text)
}

// [GetGPUTextEngineWinding] gets the winding order of the vertices returned by [GetGPUTextDrawData] for a particular GPU text engine.
//
// [GetGPUTextEngineWinding]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding
func GetGPUTextEngineWinding(engine *TextEngine) GPUTextEngineWinding {
	return ttfGetGPUTextEngineWinding(engine)
}

// [GetNextTextSubString] gets the next substring in a text object.
//
// [GetNextTextSubString]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString
func GetNextTextSubString(text *Text, substring *SubString, next *SubString) bool {
	return ttfGetNextTextSubString(text, substring, next)
}

// [GetNumFontFaces] queries the number of faces of a font.
//
// [GetNumFontFaces]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces
func GetNumFontFaces(font *Font) int32 {
	return ttfGetNumFontFaces(font)
}

// [GetPreviousTextSubString] gets the previous substring in a text object.
//
// [GetPreviousTextSubString]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString
func GetPreviousTextSubString(text *Text, substring *SubString, previous *SubString) bool {
	return ttfGetPreviousTextSubString(text, substring, previous)
}

// [GetStringSize] calculates the dimensions of a rendered string of UTF-8 text.
//
// [GetStringSize]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize
func GetStringSize(font *Font, text string, length uint64, w *int32, h *int32) bool {
	return ttfGetStringSize(font, text, length, w, h)
}

// [GetStringSizeWrapped] calculates the dimensions of a rendered string of UTF-8 text.
//
// [GetStringSizeWrapped]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped
func GetStringSizeWrapped(font *Font, text string, length uint64, wrapWidth int32, w *int32, h *int32) bool {
	return ttfGetStringSizeWrapped(font, text, length, wrapWidth, w, h)
}

// [GetTextColor] gets the color of a text object.
//
// [GetTextColor]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor
func GetTextColor(text *Text, r *uint8, g *uint8, b *uint8, a *uint8) bool {
	return ttfGetTextColor(text, r, g, b, a)
}

// [GetTextColorFloat] gets the color of a text object.
//
// [GetTextColorFloat]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat
func GetTextColorFloat(text *Text, r *float32, g *float32, b *float32, a *float32) bool {
	return ttfGetTextColorFloat(text, r, g, b, a)
}

// [GetTextDirection] gets the direction to be used for text shaping a text object.
//
// [GetTextDirection]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection
func GetTextDirection(text *Text) Direction {
	return ttfGetTextDirection(text)
}

// [GetTextEngine] gets the text engine used by a text object.
//
// [GetTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine
func GetTextEngine(text *Text) *TextEngine {
	return ttfGetTextEngine(text)
}

// [GetTextFont] gets the font used by a text object.
//
// [GetTextFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont
func GetTextFont(text *Text) *Font {
	return ttfGetTextFont(text)
}

// [GetTextPosition] gets the position of a text object.
//
// [GetTextPosition]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition
func GetTextPosition(text *Text, x *int32, y *int32) bool {
	return ttfGetTextPosition(text, x, y)
}

// [GetTextProperties] gets the properties associated with a text object.
//
// [GetTextProperties]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties
func GetTextProperties(text *Text) sdl.PropertiesID {
	return ttfGetTextProperties(text)
}

// [GetTextScript] gets the script used for text shaping a text object.
//
// [GetTextScript]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript
func GetTextScript(text *Text) uint32 {
	return ttfGetTextScript(text)
}

// [GetTextSize] gets the size of a text object.
//
// [GetTextSize]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize
func GetTextSize(text *Text, w *int32, h *int32) bool {
	return ttfGetTextSize(text, w, h)
}

// [GetTextSubString] gets the substring of a text object that surrounds a text offset.
//
// [GetTextSubString]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString
func GetTextSubString(text *Text, offset int32, substring *SubString) bool {
	return ttfGetTextSubString(text, offset, substring)
}

// [GetTextSubStringForLine] gets the substring of a text object that contains the given line.
//
// [GetTextSubStringForLine]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine
func GetTextSubStringForLine(text *Text, line int32, substring *SubString) bool {
	return ttfGetTextSubStringForLine(text, line, substring)
}

// [GetTextSubStringForPoint] gets the portion of a text string that is closest to a point.
//
// [GetTextSubStringForPoint]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint
func GetTextSubStringForPoint(text *Text, x int32, y int32, substring *SubString) bool {
	return ttfGetTextSubStringForPoint(text, x, y, substring)
}

// [GetTextSubStringsForRange] gets the substrings of a text object that contain a range of text.
//
// [GetTextSubStringsForRange]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange
func GetTextSubStringsForRange(text *Text, offset int32, length int32, count *int32) {
	ttfGetTextSubStringsForRange(text, offset, length, count)
}

// [GetTextWrapWidth] gets whether wrapping is enabled on a text object.
//
// [GetTextWrapWidth]: https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth
func GetTextWrapWidth(text *Text, wrapWidth *int32) bool {
	return ttfGetTextWrapWidth(text, wrapWidth)
}

// [InsertTextString] inserts UTF-8 text into a text object.
//
// [InsertTextString]: https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString
func InsertTextString(text *Text, offset int32, str string, length uint64) bool {
	return ttfInsertTextString(text, offset, str, length)
}

// [MeasureString] calculates how much of a UTF-8 string will fit in a given width.
//
// [MeasureString]: https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString
func MeasureString(font *Font, text string, length uint64, maxWidth int32, measuredWidth *int32, measuredLength *uint64) bool {
	return ttfMeasureString(font, text, length, maxWidth, measuredWidth, measuredLength)
}

// [OpenFont] creates a font from a file, using a specified point size.
//
// [OpenFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont
func OpenFont(file string, ptsize float32) *Font {
	return ttfOpenFont(file, ptsize)
}

// [OpenFontIO] creates a font from an SDL_IOStream, using a specified point size.
//
// [OpenFontIO]: https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO
func OpenFontIO(src *sdl.IOStream, closeio bool, ptsize float32) *Font {
	return ttfOpenFontIO(src, closeio, ptsize)
}

// [OpenFontWithProperties] creates a font with the specified properties.
//
// [OpenFontWithProperties]: https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties
func OpenFontWithProperties(props sdl.PropertiesID) *Font {
	return ttfOpenFontWithProperties(props)
}

// [Quit] deinitializes SDL_ttf.
//
// [Quit]: https://wiki.libsdl.org/SDL3_ttf/TTF_Quit
func Quit() {
	ttfQuit()
}

// [RemoveFallbackFont] removes a fallback font.
//
// [RemoveFallbackFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont
func RemoveFallbackFont(font *Font, fallback *Font) {
	ttfRemoveFallbackFont(font, fallback)
}

// [RenderGlyphBlended] renders a single UNICODE codepoint at high quality to a new ARGB surface.
//
// [RenderGlyphBlended]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended
func RenderGlyphBlended(font *Font, ch rune, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderGlyphBlended(font, ch, uintptr(fgColor))
}

// [RenderGlyphLCD] renders a single UNICODE codepoint at LCD subpixel quality to a new ARGB surface.
//
// [RenderGlyphLCD]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD
func RenderGlyphLCD(font *Font, ch rune, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderGlyphLCD(font, ch, uintptr(fgColor), uintptr(bgColor))
}

// [RenderGlyphShaded] renders a single UNICODE codepoint at high quality to a new 8-bit surface.
//
// [RenderGlyphShaded]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded
func RenderGlyphShaded(font *Font, ch rune, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderGlyphShaded(font, ch, uintptr(fgColor), uintptr(bgColor))
}

// [RenderGlyphSolid] renders a single 32-bit glyph at fast quality to a new 8-bit surface.
//
// [RenderGlyphSolid]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid
func RenderGlyphSolid(font *Font, ch rune, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderGlyphSolid(font, ch, uintptr(fgColor))
}

// [RenderTextBlended] renders UTF-8 text at high quality to a new ARGB surface.
//
// [RenderTextBlended]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended
func RenderTextBlended(font *Font, text string, length uint64, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextBlended(font, text, length, uintptr(fgColor))
}

// [RenderTextBlendedWrapped] renders word-wrapped UTF-8 text at high quality to a new ARGB surface.
//
// [RenderTextBlendedWrapped]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped
func RenderTextBlendedWrapped(font *Font, text string, length uint64, fg sdl.Color, wrapWidth int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextBlendedWrapped(font, text, length, uintptr(fgColor), wrapWidth)
}

// [RenderTextLCD] renders UTF-8 text at LCD subpixel quality to a new ARGB surface.
//
// [RenderTextLCD]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD
func RenderTextLCD(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextLCD(font, text, length, uintptr(fgColor), uintptr(bgColor))
}

// [RenderTextLCDWrapped] renders word-wrapped UTF-8 text at LCD subpixel quality to a new ARGB surface.
//
// [RenderTextLCDWrapped]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped
func RenderTextLCDWrapped(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color, wrapWidth int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextLCDWrapped(font, text, length, uintptr(fgColor), uintptr(bgColor), wrapWidth)
}

// [RenderTextShaded] renders UTF-8 text at high quality to a new 8-bit surface.
//
// [RenderTextShaded]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded
func RenderTextShaded(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextShaded(font, text, length, uintptr(fgColor), uintptr(bgColor))
}

// [RenderTextShadedWrapped] renders word-wrapped UTF-8 text at high quality to a new 8-bit surface.
//
// [RenderTextShadedWrapped]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped
func RenderTextShadedWrapped(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color, wrapWidth int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextShadedWrapped(font, text, length, uintptr(fgColor), uintptr(bgColor), wrapWidth)
}

// [RenderTextSolid] renders UTF-8 text at fast quality to a new 8-bit surface.
//
// [RenderTextSolid]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid
func RenderTextSolid(font *Font, text string, length uint64, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextSolid(font, text, length, uintptr(fgColor))
}

// [RenderTextSolidWrapped] renders word-wrapped UTF-8 text at fast quality to a new 8-bit surface.
//
// [RenderTextSolidWrapped]: https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped
func RenderTextSolidWrapped(font *Font, text string, length uint64, fg sdl.Color, wrapLength int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextSolidWrapped(font, text, length, uintptr(fgColor), wrapLength)
}

// [SetFontDirection] sets the direction to be used for text shaping by a font.
//
// [SetFontDirection]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection
func SetFontDirection(font *Font, direction Direction) bool {
	return ttfSetFontDirection(font, direction)
}

// [SetFontHinting] sets a font's current hinter setting.
//
// [SetFontHinting]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting
func SetFontHinting(font *Font, hinting HintingFlags) {
	ttfSetFontHinting(font, hinting)
}

// [SetFontKerning] sets if kerning is enabled for a font.
//
// [SetFontKerning]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning
func SetFontKerning(font *Font, enabled bool) {
	ttfSetFontKerning(font, enabled)
}

// [SetFontLanguage] sets language to be used for text shaping by a font.
//
// [SetFontLanguage]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage
func SetFontLanguage(font *Font, languageBCP47 string) bool {
	return ttfSetFontLanguage(font, languageBCP47)
}

// [SetFontLineSkip] sets the spacing between lines of text for a font.
//
// [SetFontLineSkip]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip
func SetFontLineSkip(font *Font, lineskip int32) {
	ttfSetFontLineSkip(font, lineskip)
}

// [SetFontOutline] sets a font's current outline.
//
// [SetFontOutline]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline
func SetFontOutline(font *Font, outline int32) bool {
	return ttfSetFontOutline(font, outline)
}

// [SetFontScript] sets the script to be used for text shaping by a font.
//
// [SetFontScript]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript
func SetFontScript(font *Font, script uint32) bool {
	return ttfSetFontScript(font, script)
}

// [SetFontSDF] enables Signed Distance Field rendering for a font.
//
// [SetFontSDF]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF
func SetFontSDF(font *Font, enabled bool) bool {
	return ttfSetFontSDF(font, enabled)
}

// [SetFontSize] sets a font's size dynamically.
//
// [SetFontSize]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize
func SetFontSize(font *Font, ptsize float32) bool {
	return ttfSetFontSize(font, ptsize)
}

// [SetFontSizeDPI] sets font size dynamically with target resolutions, in dots per inch.
//
// [SetFontSizeDPI]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI
func SetFontSizeDPI(font *Font, ptsize float32, hdpi int32, vdpi int32) bool {
	return ttfSetFontSizeDPI(font, ptsize, hdpi, vdpi)
}

// [SetFontStyle] sets a font's current style.
//
// [SetFontStyle]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle
func SetFontStyle(font *Font, style FontStyleFlags) {
	ttfSetFontStyle(font, style)
}

// [SetFontWrapAlignment] sets a font's current wrap alignment option.
//
// [SetFontWrapAlignment]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment
func SetFontWrapAlignment(font *Font, align HorizontalAlignment) {
	ttfSetFontWrapAlignment(font, align)
}

// [SetGPUTextEngineWinding] sets the winding order of the vertices returned by [GetGPUTextDrawData] for a particular GPU text engine.
//
// [SetGPUTextEngineWinding]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding
func SetGPUTextEngineWinding(engine *TextEngine, winding GPUTextEngineWinding) {
	ttfSetGPUTextEngineWinding(engine, winding)
}

// [SetTextColor] sets the color of a text object.
//
// [SetTextColor]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor
func SetTextColor(text *Text, r uint8, g uint8, b uint8, a uint8) bool {
	return ttfSetTextColor(text, r, g, b, a)
}

// [SetTextColorFloat] sets the color of a text object.
//
// [SetTextColorFloat]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat
func SetTextColorFloat(text *Text, r float32, g float32, b float32, a float32) bool {
	return ttfSetTextColorFloat(text, r, g, b, a)
}

// [SetTextDirection] sets the direction to be used for text shaping a text object.
//
// [SetTextDirection]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection
func SetTextDirection(text *Text, direction Direction) bool {
	return ttfSetTextDirection(text, direction)
}

// [SetTextEngine] sets the text engine used by a text object.
//
// [SetTextEngine]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine
func SetTextEngine(text *Text, engine *TextEngine) bool {
	return ttfSetTextEngine(text, engine)
}

// [SetTextFont] sets the font used by a text object.
//
// [SetTextFont]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont
func SetTextFont(text *Text, font *Font) bool {
	return ttfSetTextFont(text, font)
}

// [SetTextPosition] sets the position of a text object.
//
// [SetTextPosition]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition
func SetTextPosition(text *Text, x int32, y int32) bool {
	return ttfSetTextPosition(text, x, y)
}

// [SetTextScript] sets the script to be used for text shaping a text object.
//
// [SetTextScript]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript
func SetTextScript(text *Text, script uint32) bool {
	return ttfSetTextScript(text, script)
}

// [SetTextString] sets the UTF-8 text used by a text object.
//
// [SetTextString]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString
func SetTextString(text *Text, str string, length uint64) bool {
	return ttfSetTextString(text, str, length)
}

// [SetTextWrapWhitespaceVisible] sets whether whitespace should be visible when wrapping a text object.
//
// [SetTextWrapWhitespaceVisible]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible
func SetTextWrapWhitespaceVisible(text *Text, visible bool) bool {
	return ttfSetTextWrapWhitespaceVisible(text, visible)
}

// [SetTextWrapWidth] sets whether wrapping is enabled on a text object.
//
// [SetTextWrapWidth]: https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth
func SetTextWrapWidth(text *Text, wrapWidth int32) bool {
	return ttfSetTextWrapWidth(text, wrapWidth)
}

// [TextWrapWhitespaceVisible] returns whether whitespace is shown when wrapping a text object.
//
// [TextWrapWhitespaceVisible]: https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible
func TextWrapWhitespaceVisible(text *Text) bool {
	return ttfTextWrapWhitespaceVisible(text)
}

// [UpdateText] updates the layout of a text object.
//
// [UpdateText]: https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText
func UpdateText(text *Text) bool {
	return ttfUpdateText(text)
}

// [WasInit] checks if SDL_ttf is initialized.
//
// [WasInit]: https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit
func WasInit() int32 {
	return ttfWasInit()
}

// [StringToTag] converts from a 4 character string to a 32-bit tag.
//
// [StringToTag]: https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag
func StringToTag(str string) uint32 {
	return ttfStringToTag(str)
}

// [TagToString] converts from a 32-bit tag to a 4 character string.
//
// [TagToString]: https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString
func TagToString(tag uint32) string {
	const size = 5
	var buffer [size]byte
	ttfTagToString(tag, &buffer[0], size)
	return convert.ToString(&buffer[0])
}
