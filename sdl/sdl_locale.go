package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

// [Locale] is a struct to provide locale data.
//
// [Locale]: https://wiki.libsdl.org/SDL3/SDL_Locale
type Locale struct {
	Language string // A language name, like "en" for English.
	Country  string // A country, like "US" for America. Can be NULL.
}

// [GetPreferredLocales] reports the user's preferred locale.
//
// [GetPreferredLocales]: https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales
func GetPreferredLocales() []*Locale {
	var count int32
	locales := sdlGetPreferredLocales(&count)
	if locales == nil {
		return nil
	}
	defer Free(unsafe.Pointer(locales))

	result := make([]*Locale, count)

	for i, v := range unsafe.Slice(locales, count) {
		locale := new(Locale)
		locale.Language = convert.ToString(v.language)
		locale.Country = convert.ToString(v.country)
		result[i] = locale
	}

	return result
}
