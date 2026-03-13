package sdl

// [OpenURL] opens a URL/URI in the browser or other appropriate external application.
//
// [OpenURL]: https://wiki.libsdl.org/SDL3/SDL_OpenURL
func OpenURL(url string) bool {
	return sdlOpenURL(url)
}
