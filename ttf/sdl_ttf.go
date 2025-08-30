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

type FontStyleFlags uint32

const (
	StyleNormal        = 0x00
	StyleBold          = 0x01
	StyleItalic        = 0x02
	StyleUnderline     = 0x04
	StyleStrikethrough = 0x08
)

type SubStringFlags uint32

const (
	SubStringDirectionMask = 0x000000FF
	SubStringTextStart     = 0x00000100
	SubStringLineStart     = 0x00000200
	SubStringLineEnd       = 0x00000400
	SubStringTextEnd       = 0x00000800
)

type Direction uint32

const (
	DirectionInvalid Direction = 0
	DirectionLtr     Direction = iota + 3 // Left to Right
	DirectionRtl                          // Right to Left
	DirectionTtb                          // Top to Bottom
	DirectionBtt                          // Bottom to Top
)

type HintingFlags uint32

const (
	HintingNormal HintingFlags = iota
	HintingLight
	HintingMono
	HintingNone
	HintingLightSubpixel
)

type HorizontalAlignment int32

const (
	HorizontalAlignInvalid HorizontalAlignment = iota - 1
	HorizontalAlignLeft
	HorizontalAlignCenter
	HorizontalAlignRight
)

type ImageType uint32

const (
	ImageInvalid ImageType = iota
	ImageAlpha
	ImageColor
	ImageSDF
)

type GPUTextEngineWinding int32

const (
	GPUTextEngineWindingInvalid GPUTextEngineWinding = iota - 1
	GPUTextEngineWindingClockwise
	GPUTextEngineWindingCounterClockwise
)

type Font struct{}

type TextData struct{}

type TextEngine struct{}

type Text struct {
	text     *byte
	NumLines int32
	Refcount int32
	Internal *TextData
}

func (t *Text) Text() string {
	return convert.ToString(t.text)
}

type SubString struct {
	Flags        SubStringFlags
	Offset       int32
	Length       int32
	LineIndex    int32
	ClusterIndex int32
	Rect         sdl.Rect
}

type GPUAtlasDrawSequence struct {
	AtlasTexture *sdl.GPUTexture
	XY           *sdl.FPoint
	UV           *sdl.FPoint
	NumVertices  int32
	Indices      *int32
	NumIndices   int32
	ImageType    ImageType
	Next         *GPUAtlasDrawSequence
}

// Version gets the version of the dynamically linked SDL_ttf library.
func Version() (major, minor, patch int32) {
	version := ttfVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}

// GetFreeTypeVersion queries the version of the FreeType library in use.
func GetFreeTypeVersion() (major, minor, patch int32) {
	ttfGetFreeTypeVersion(&major, &minor, &patch)
	return
}

// Init initializes SDL_ttf.
func Init() bool {
	return ttfInit()
}

// GetHarfBuzzVersion queries the version of the HarfBuzz library in use.
func GetHarfBuzzVersion() (major, minor, patch int32) {
	ttfGetHarfBuzzVersion(&major, &minor, &patch)
	return
}

func AddFallbackFont(font *Font, fallback *Font) bool {
	return ttfAddFallbackFont(font, fallback)
}

func AppendTextString(text *Text, str string, length uint64) bool {
	return ttfAppendTextString(text, str, length)
}

func ClearFallbackFonts(font *Font) {
	ttfClearFallbackFonts(font)
}

func CloseFont(font *Font) {
	ttfCloseFont(font)
}

func CopyFont(existingFont *Font) *Font {
	return ttfCopyFont(existingFont)
}

func CreateGPUTextEngine(device *sdl.GPUDevice) *TextEngine {
	return ttfCreateGPUTextEngine(device)
}

func CreateGPUTextEngineWithProperties(props sdl.PropertiesID) *TextEngine {
	return ttfCreateGPUTextEngineWithProperties(props)
}

func CreateRendererTextEngine(renderer *sdl.Renderer) *TextEngine {
	return ttfCreateRendererTextEngine(renderer)
}

func CreateRendererTextEngineWithProperties(props sdl.PropertiesID) *TextEngine {
	return ttfCreateRendererTextEngineWithProperties(props)
}

func CreateSurfaceTextEngine() *TextEngine {
	return ttfCreateSurfaceTextEngine()
}

func CreateText(engine *TextEngine, font *Font, text string, length uint64) *Text {
	return ttfCreateText(engine, font, text, length)
}

func DeleteTextString(text *Text, offset int32, length int32) bool {
	return ttfDeleteTextString(text, offset, length)
}

func DestroyGPUTextEngine(engine *TextEngine) {
	ttfDestroyGPUTextEngine(engine)
}

func DestroyRendererTextEngine(engine *TextEngine) {
	ttfDestroyRendererTextEngine(engine)
}

func DestroySurfaceTextEngine(engine *TextEngine) {
	ttfDestroySurfaceTextEngine(engine)
}

func DestroyText(text *Text) {
	ttfDestroyText(text)
}

func DrawRendererText(text *Text, x float32, y float32) bool {
	return ttfDrawRendererText(text, x, y)
}

func DrawSurfaceText(text *Text, x int32, y int32, surface *sdl.Surface) bool {
	return ttfDrawSurfaceText(text, x, y, surface)
}

func FontHasGlyph(font *Font, ch rune) bool {
	return ttfFontHasGlyph(font, ch)
}

func FontIsFixedWidth(font *Font) bool {
	return ttfFontIsFixedWidth(font)
}

func FontIsScalable(font *Font) bool {
	return ttfFontIsScalable(font)
}

func GetFontAscent(font *Font) int32 {
	return ttfGetFontAscent(font)
}

func GetFontDescent(font *Font) int32 {
	return ttfGetFontDescent(font)
}

func GetFontDirection(font *Font) Direction {
	return ttfGetFontDirection(font)
}

func GetFontDPI(font *Font, hdpi *int32, vdpi *int32) bool {
	return ttfGetFontDPI(font, hdpi, vdpi)
}

func GetFontFamilyName(font *Font) string {
	return ttfGetFontFamilyName(font)
}

func GetFontGeneration(font *Font) uint32 {
	return ttfGetFontGeneration(font)
}

func GetFontHeight(font *Font) int32 {
	return ttfGetFontHeight(font)
}

func GetFontHinting(font *Font) HintingFlags {
	return ttfGetFontHinting(font)
}

func GetFontKerning(font *Font) bool {
	return ttfGetFontKerning(font)
}

func GetFontLineSkip(font *Font) int32 {
	return ttfGetFontLineSkip(font)
}

func GetFontOutline(font *Font) int32 {
	return ttfGetFontOutline(font)
}

func GetFontProperties(font *Font) sdl.PropertiesID {
	return ttfGetFontProperties(font)
}

func GetFontScript(font *Font) uint32 {
	return ttfGetFontScript(font)
}

func GetFontSDF(font *Font) bool {
	return ttfGetFontSDF(font)
}

func GetFontSize(font *Font) float32 {
	return ttfGetFontSize(font)
}

func GetFontStyle(font *Font) FontStyleFlags {
	return ttfGetFontStyle(font)
}

func GetFontStyleName(font *Font) string {
	return ttfGetFontStyleName(font)
}

func GetFontWrapAlignment(font *Font) HorizontalAlignment {
	return ttfGetFontWrapAlignment(font)
}

func GetGlyphImage(font *Font, ch rune, imageType *ImageType) *sdl.Surface {
	return ttfGetGlyphImage(font, ch, imageType)
}

func GetGlyphImageForIndex(font *Font, glyphIndex uint32, imageType *ImageType) *sdl.Surface {
	return ttfGetGlyphImageForIndex(font, glyphIndex, imageType)
}

func GetGlyphKerning(font *Font, previousCh, ch rune, kerning *int32) bool {
	return ttfGetGlyphKerning(font, previousCh, ch, kerning)
}

func GetGlyphMetrics(font *Font, ch rune, minx *int32, maxx *int32, miny *int32, maxy *int32, advance *int32) bool {
	return ttfGetGlyphMetrics(font, ch, minx, maxx, miny, maxy, advance)
}

func GetGlyphScript(ch rune) uint32 {
	return ttfGetGlyphScript(ch)
}

func GetGPUTextDrawData(text *Text) *GPUAtlasDrawSequence {
	return ttfGetGPUTextDrawData(text)
}

func GetGPUTextEngineWinding(engine *TextEngine) GPUTextEngineWinding {
	return ttfGetGPUTextEngineWinding(engine)
}

func GetNextTextSubString(text *Text, substring *SubString, next *SubString) bool {
	return ttfGetNextTextSubString(text, substring, next)
}

func GetNumFontFaces(font *Font) int32 {
	return ttfGetNumFontFaces(font)
}

func GetPreviousTextSubString(text *Text, substring *SubString, previous *SubString) bool {
	return ttfGetPreviousTextSubString(text, substring, previous)
}

func GetStringSize(font *Font, text string, length uint64, w *int32, h *int32) bool {
	return ttfGetStringSize(font, text, length, w, h)
}

func GetStringSizeWrapped(font *Font, text string, length uint64, wrapWidth int32, w *int32, h *int32) bool {
	return ttfGetStringSizeWrapped(font, text, length, wrapWidth, w, h)
}

func GetTextColor(text *Text, r *uint8, g *uint8, b *uint8, a *uint8) bool {
	return ttfGetTextColor(text, r, g, b, a)
}

func GetTextColorFloat(text *Text, r *float32, g *float32, b *float32, a *float32) bool {
	return ttfGetTextColorFloat(text, r, g, b, a)
}

func GetTextDirection(text *Text) Direction {
	return ttfGetTextDirection(text)
}

func GetTextEngine(text *Text) *TextEngine {
	return ttfGetTextEngine(text)
}

func GetTextFont(text *Text) *Font {
	return ttfGetTextFont(text)
}

func GetTextPosition(text *Text, x *int32, y *int32) bool {
	return ttfGetTextPosition(text, x, y)
}

func GetTextProperties(text *Text) sdl.PropertiesID {
	return ttfGetTextProperties(text)
}

func GetTextScript(text *Text) uint32 {
	return ttfGetTextScript(text)
}

func GetTextSize(text *Text, w *int32, h *int32) bool {
	return ttfGetTextSize(text, w, h)
}

func GetTextSubString(text *Text, offset int32, substring *SubString) bool {
	return ttfGetTextSubString(text, offset, substring)
}

func GetTextSubStringForLine(text *Text, line int32, substring *SubString) bool {
	return ttfGetTextSubStringForLine(text, line, substring)
}

func GetTextSubStringForPoint(text *Text, x int32, y int32, substring *SubString) bool {
	return ttfGetTextSubStringForPoint(text, x, y, substring)
}

func GetTextSubStringsForRange(text *Text, offset int32, length int32, count *int32) {
	ttfGetTextSubStringsForRange(text, offset, length, count)
}

func GetTextWrapWidth(text *Text, wrapWidth *int32) bool {
	return ttfGetTextWrapWidth(text, wrapWidth)
}

func InsertTextString(text *Text, offset int32, str string, length uint64) bool {
	return ttfInsertTextString(text, offset, str, length)
}

func MeasureString(font *Font, text string, length uint64, maxWidth int32, measuredWidth *int32, measuredLength *uint64) bool {
	return ttfMeasureString(font, text, length, maxWidth, measuredWidth, measuredLength)
}

func OpenFont(file string, ptsize float32) *Font {
	return ttfOpenFont(file, ptsize)
}

func OpenFontIO(src *sdl.IOStream, closeio bool, ptsize float32) *Font {
	return ttfOpenFontIO(src, closeio, ptsize)
}

func OpenFontWithProperties(props sdl.PropertiesID) *Font {
	return ttfOpenFontWithProperties(props)
}

func Quit() {
	ttfQuit()
}

func RemoveFallbackFont(font *Font, fallback *Font) {
	ttfRemoveFallbackFont(font, fallback)
}

func RenderGlyphBlended(font *Font, ch rune, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderGlyphBlended(font, ch, uintptr(fgColor))
}

func RenderGlyphLCD(font *Font, ch rune, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderGlyphLCD(font, ch, uintptr(fgColor), uintptr(bgColor))
}

func RenderGlyphShaded(font *Font, ch rune, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderGlyphShaded(font, ch, uintptr(fgColor), uintptr(bgColor))
}

func RenderGlyphSolid(font *Font, ch rune, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderGlyphSolid(font, ch, uintptr(fgColor))
}

func RenderTextBlended(font *Font, text string, length uint64, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextBlended(font, text, length, uintptr(fgColor))
}

func RenderTextBlendedWrapped(font *Font, text string, length uint64, fg sdl.Color, wrapWidth int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextBlendedWrapped(font, text, length, uintptr(fgColor), wrapWidth)
}

func RenderTextLCD(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextLCD(font, text, length, uintptr(fgColor), uintptr(bgColor))
}

func RenderTextLCDWrapped(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color, wrapWidth int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextLCDWrapped(font, text, length, uintptr(fgColor), uintptr(bgColor), wrapWidth)
}

func RenderTextShaded(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextShaded(font, text, length, uintptr(fgColor), uintptr(bgColor))
}

func RenderTextShadedWrapped(font *Font, text string, length uint64, fg sdl.Color, bg sdl.Color, wrapWidth int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	bgColor := uint32(bg.A)<<24 + uint32(bg.B)<<16 + uint32(bg.G)<<8 + uint32(bg.R)
	return ttfRenderTextShadedWrapped(font, text, length, uintptr(fgColor), uintptr(bgColor), wrapWidth)
}

func RenderTextSolid(font *Font, text string, length uint64, fg sdl.Color) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextSolid(font, text, length, uintptr(fgColor))
}

func RenderTextSolidWrapped(font *Font, text string, length uint64, fg sdl.Color, wrapLength int32) *sdl.Surface {
	fgColor := uint32(fg.A)<<24 + uint32(fg.B)<<16 + uint32(fg.G)<<8 + uint32(fg.R)
	return ttfRenderTextSolidWrapped(font, text, length, uintptr(fgColor), wrapLength)
}

func SetFontDirection(font *Font, direction Direction) bool {
	return ttfSetFontDirection(font, direction)
}

func SetFontHinting(font *Font, hinting HintingFlags) {
	ttfSetFontHinting(font, hinting)
}

func SetFontKerning(font *Font, enabled bool) {
	ttfSetFontKerning(font, enabled)
}

func SetFontLanguage(font *Font, languageBCP47 string) bool {
	return ttfSetFontLanguage(font, languageBCP47)
}

func SetFontLineSkip(font *Font, lineskip int32) {
	ttfSetFontLineSkip(font, lineskip)
}

func SetFontOutline(font *Font, outline int32) bool {
	return ttfSetFontOutline(font, outline)
}

func SetFontScript(font *Font, script uint32) bool {
	return ttfSetFontScript(font, script)
}

func SetFontSDF(font *Font, enabled bool) bool {
	return ttfSetFontSDF(font, enabled)
}

func SetFontSize(font *Font, ptsize float32) bool {
	return ttfSetFontSize(font, ptsize)
}

func SetFontSizeDPI(font *Font, ptsize float32, hdpi int32, vdpi int32) bool {
	return ttfSetFontSizeDPI(font, ptsize, hdpi, vdpi)
}

func SetFontStyle(font *Font, style FontStyleFlags) {
	ttfSetFontStyle(font, style)
}

func SetFontWrapAlignment(font *Font, align HorizontalAlignment) {
	ttfSetFontWrapAlignment(font, align)
}

func SetGPUTextEngineWinding(engine *TextEngine, winding GPUTextEngineWinding) {
	ttfSetGPUTextEngineWinding(engine, winding)
}

func SetTextColor(text *Text, r uint8, g uint8, b uint8, a uint8) bool {
	return ttfSetTextColor(text, r, g, b, a)
}

func SetTextColorFloat(text *Text, r float32, g float32, b float32, a float32) bool {
	return ttfSetTextColorFloat(text, r, g, b, a)
}

func SetTextDirection(text *Text, direction Direction) bool {
	return ttfSetTextDirection(text, direction)
}

func SetTextEngine(text *Text, engine *TextEngine) bool {
	return ttfSetTextEngine(text, engine)
}

func SetTextFont(text *Text, font *Font) bool {
	return ttfSetTextFont(text, font)
}

func SetTextPosition(text *Text, x int32, y int32) bool {
	return ttfSetTextPosition(text, x, y)
}

func SetTextScript(text *Text, script uint32) bool {
	return ttfSetTextScript(text, script)
}

func SetTextString(text *Text, str string, length uint64) bool {
	return ttfSetTextString(text, str, length)
}

func SetTextWrapWhitespaceVisible(text *Text, visible bool) bool {
	return ttfSetTextWrapWhitespaceVisible(text, visible)
}

func SetTextWrapWidth(text *Text, wrapWidth int32) bool {
	return ttfSetTextWrapWidth(text, wrapWidth)
}

func TextWrapWhitespaceVisible(text *Text) bool {
	return ttfTextWrapWhitespaceVisible(text)
}

func UpdateText(text *Text) bool {
	return ttfUpdateText(text)
}

func WasInit() int32 {
	return ttfWasInit()
}

func StringToTag(str string) uint32 {
	return ttfStringToTag(str)
}

func TagToString(tag uint32) string {
	const size = 5
	var buffer [size]byte
	ttfTagToString(tag, &buffer[0], size)
	return convert.ToString(&buffer[0])
}
