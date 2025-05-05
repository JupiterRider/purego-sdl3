package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

const FltEpsilon = 0x0.000002p0

func FourCC(a, b, c, d byte) uint32 {
	return uint32(a) | uint32(b)<<8 | uint32(c)<<16 | uint32(d)<<24
}

// func abs(x int32) int32 {
//	return sdlabs(x)
// }

// func acos(x float64) float64 {
//	return sdlacos(x)
// }

// func acosf(x float32) float32 {
//	return sdlacosf(x)
// }

// func aligned_alloc(alignment uint64, size uint64) unsafe.Pointer {
//	return sdlaligned_alloc(alignment, size)
// }

// func aligned_free(mem unsafe.Pointer)  {
//	sdlaligned_free(mem)
// }

// func asin(x float64) float64 {
//	return sdlasin(x)
// }

// func asinf(x float32) float32 {
//	return sdlasinf(x)
// }

// func asprintf(strp **byte, fmt string) int32 {
//	return sdlasprintf(strp, fmt)
// }

// func atan(x float64) float64 {
//	return sdlatan(x)
// }

// func atan2(y float64, x float64) float64 {
//	return sdlatan2(y, x)
// }

// func atan2f(y float32, x float32) float32 {
//	return sdlatan2f(y, x)
// }

// func atanf(x float32) float32 {
//	return sdlatanf(x)
// }

// func atof(str string) float64 {
//	return sdlatof(str)
// }

// func atoi(str string) int32 {
//	return sdlatoi(str)
// }

// func bsearch(key unsafe.Pointer, base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback) unsafe.Pointer {
//	return sdlbsearch(key, base, nmemb, size, compare)
// }

// func bsearch_r(key unsafe.Pointer, base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback_r, userdata unsafe.Pointer) unsafe.Pointer {
//	return sdlbsearch_r(key, base, nmemb, size, compare, userdata)
// }

// func calloc(nmemb uint64, size uint64) unsafe.Pointer {
//	return sdlcalloc(nmemb, size)
// }

// func ceil(x float64) float64 {
//	return sdlceil(x)
// }

// func ceilf(x float32) float32 {
//	return sdlceilf(x)
// }

// func copysign(x float64, y float64) float64 {
//	return sdlcopysign(x, y)
// }

// func copysignf(x float32, y float32) float32 {
//	return sdlcopysignf(x, y)
// }

// func cos(x float64) float64 {
//	return sdlcos(x)
// }

// func cosf(x float32) float32 {
//	return sdlcosf(x)
// }

// func crc16(crc uint16, data unsafe.Pointer, len uint64) uint16 {
//	return sdlcrc16(crc, data, len)
// }

// func crc32(crc uint32, data unsafe.Pointer, len uint64) uint32 {
//	return sdlcrc32(crc, data, len)
// }

// func CreateEnvironment(populated bool) *Environment {
//	return sdlCreateEnvironment(populated)
// }

// func DestroyEnvironment(env *Environment)  {
//	sdlDestroyEnvironment(env)
// }

// func exp(x float64) float64 {
//	return sdlexp(x)
// }

// func expf(x float32) float32 {
//	return sdlexpf(x)
// }

// func fabs(x float64) float64 {
//	return sdlfabs(x)
// }

// func fabsf(x float32) float32 {
//	return sdlfabsf(x)
// }

// func floor(x float64) float64 {
//	return sdlfloor(x)
// }

// func floorf(x float32) float32 {
//	return sdlfloorf(x)
// }

// func fmod(x float64, y float64) float64 {
//	return sdlfmod(x, y)
// }

// func fmodf(x float32, y float32) float32 {
//	return sdlfmodf(x, y)
// }

// Free frees allocated memory.
//
// If mem is nil, this function does nothing.
func Free(mem unsafe.Pointer) {
	purego.SyscallN(sdlfree, uintptr(mem))
}

// func getenv(name string) string {
//	return sdlgetenv(name)
// }

// func getenv_unsafe(name string) string {
//	return sdlgetenv_unsafe(name)
// }

// func GetEnvironment() *Environment {
//	return sdlGetEnvironment()
// }

// func GetEnvironmentVariable(env *Environment, name string) string {
//	return sdlGetEnvironmentVariable(env, name)
// }

// func GetEnvironmentVariables(env *Environment) **byte {
//	return sdlGetEnvironmentVariables(env)
// }

// func GetMemoryFunctions(malloc_func *malloc_func, calloc_func *calloc_func, realloc_func *realloc_func, free_func *free_func)  {
//	sdlGetMemoryFunctions(malloc_func, calloc_func, realloc_func, free_func)
// }

// func GetNumAllocations() int32 {
//	return sdlGetNumAllocations()
// }

// func GetOriginalMemoryFunctions(malloc_func *malloc_func, calloc_func *calloc_func, realloc_func *realloc_func, free_func *free_func)  {
//	sdlGetOriginalMemoryFunctions(malloc_func, calloc_func, realloc_func, free_func)
// }

// func iconv(cd iconv_t, inbuf **byte, inbytesleft *uint64, outbuf **byte, outbytesleft *uint64) uint64 {
//	return sdliconv(cd, inbuf, inbytesleft, outbuf, outbytesleft)
// }

// func iconv_close(cd iconv_t) int32 {
//	return sdliconv_close(cd)
// }

// func iconv_open(tocode string, fromcode string) iconv_t {
//	return sdliconv_open(tocode, fromcode)
// }

// func iconv_string(tocode string, fromcode string, inbuf string, inbytesleft uint64) string {
//	return sdliconv_string(tocode, fromcode, inbuf, inbytesleft)
// }

// func isalnum(x int32) int32 {
//	return sdlisalnum(x)
// }

// func isalpha(x int32) int32 {
//	return sdlisalpha(x)
// }

// func isblank(x int32) int32 {
//	return sdlisblank(x)
// }

// func iscntrl(x int32) int32 {
//	return sdliscntrl(x)
// }

// func isdigit(x int32) int32 {
//	return sdlisdigit(x)
// }

// func isgraph(x int32) int32 {
//	return sdlisgraph(x)
// }

// func isinf(x float64) int32 {
//	return sdlisinf(x)
// }

// func isinff(x float32) int32 {
//	return sdlisinff(x)
// }

// func islower(x int32) int32 {
//	return sdlislower(x)
// }

// func isnan(x float64) int32 {
//	return sdlisnan(x)
// }

// func isnanf(x float32) int32 {
//	return sdlisnanf(x)
// }

// func isprint(x int32) int32 {
//	return sdlisprint(x)
// }

// func ispunct(x int32) int32 {
//	return sdlispunct(x)
// }

// func isspace(x int32) int32 {
//	return sdlisspace(x)
// }

// func isupper(x int32) int32 {
//	return sdlisupper(x)
// }

// func isxdigit(x int32) int32 {
//	return sdlisxdigit(x)
// }

// func itoa(value int32, str string, radix int32) string {
//	return sdlitoa(value, str, radix)
// }

// func lltoa(value int64, str string, radix int32) string {
//	return sdllltoa(value, str, radix)
// }

// func log(x float64) float64 {
//	return sdllog(x)
// }

// func log10(x float64) float64 {
//	return sdllog10(x)
// }

// func log10f(x float32) float32 {
//	return sdllog10f(x)
// }

// func logf(x float32) float32 {
//	return sdllogf(x)
// }

// func lround(x float64) int64 {
//	return sdllround(x)
// }

// func lroundf(x float32) int64 {
//	return sdllroundf(x)
// }

// func ltoa(value int64, str string, radix int32) string {
//	return sdlltoa(value, str, radix)
// }

// func malloc(size uint64) unsafe.Pointer {
//	return sdlmalloc(size)
// }

// func memcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, len uint64) int32 {
//	return sdlmemcmp(s1, s2, len)
// }

// func memcpy(dst unsafe.Pointer, src unsafe.Pointer, len uint64) unsafe.Pointer {
//	return sdlmemcpy(dst, src, len)
// }

// func memmove(dst unsafe.Pointer, src unsafe.Pointer, len uint64) unsafe.Pointer {
//	return sdlmemmove(dst, src, len)
// }

// func memset(dst unsafe.Pointer, c int32, len uint64) unsafe.Pointer {
//	return sdlmemset(dst, c, len)
// }

// func memset4(dst unsafe.Pointer, val uint32, dwords uint64) unsafe.Pointer {
//	return sdlmemset4(dst, val, dwords)
// }

// func modf(x float64, y *:double) float64 {
//	return sdlmodf(x, y)
// }

// func modff(x float32, y *float32) float32 {
//	return sdlmodff(x, y)
// }

// func murmur3_32(data unsafe.Pointer, len uint64, seed uint32) uint32 {
//	return sdlmurmur3_32(data, len, seed)
// }

// func pow(x float64, y float64) float64 {
//	return sdlpow(x, y)
// }

// func powf(x float32, y float32) float32 {
//	return sdlpowf(x, y)
// }

// func qsort(base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback)  {
//	sdlqsort(base, nmemb, size, compare)
// }

// func qsort_r(base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback_r, userdata unsafe.Pointer)  {
//	sdlqsort_r(base, nmemb, size, compare, userdata)
// }

// func rand(n int32) int32 {
//	return sdlrand(n)
// }

// func rand_bits() uint32 {
//	return sdlrand_bits()
// }

// func rand_bits_r(state *uint64) uint32 {
//	return sdlrand_bits_r(state)
// }

// func rand_r(state *uint64, n int32) int32 {
//	return sdlrand_r(state, n)
// }

// func randf() float32 {
//	return sdlrandf()
// }

// func randf_r(state *uint64) float32 {
//	return sdlrandf_r(state)
// }

// func realloc(mem unsafe.Pointer, size uint64) unsafe.Pointer {
//	return sdlrealloc(mem, size)
// }

// func round(x float64) float64 {
//	return sdlround(x)
// }

// func roundf(x float32) float32 {
//	return sdlroundf(x)
// }

// func scalbn(x float64, n int32) float64 {
//	return sdlscalbn(x, n)
// }

// func scalbnf(x float32, n int32) float32 {
//	return sdlscalbnf(x, n)
// }

// func setenv_unsafe(name string, value string, overwrite int32) int32 {
//	return sdlsetenv_unsafe(name, value, overwrite)
// }

// func SetEnvironmentVariable(env *Environment, name string, value string, overwrite bool) bool {
//	return sdlSetEnvironmentVariable(env, name, value, overwrite)
// }

// func SetMemoryFunctions(malloc_func malloc_func, calloc_func calloc_func, realloc_func realloc_func, free_func free_func) bool {
//	return sdlSetMemoryFunctions(malloc_func, calloc_func, realloc_func, free_func)
// }

// func sin(x float64) float64 {
//	return sdlsin(x)
// }

// func sinf(x float32) float32 {
//	return sdlsinf(x)
// }

// func size_add_check_overflow(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_add_check_overflow(a, b, ret)
// }

// func size_add_check_overflow_builtin(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_add_check_overflow_builtin(a, b, ret)
// }

// func size_mul_check_overflow(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_mul_check_overflow(a, b, ret)
// }

// func size_mul_check_overflow_builtin(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_mul_check_overflow_builtin(a, b, ret)
// }

// func snprintf(text string, maxlen uint64, fmt string) int32 {
//	return sdlsnprintf(text, maxlen, fmt)
// }

// func sqrt(x float64) float64 {
//	return sdlsqrt(x)
// }

// func sqrtf(x float32) float32 {
//	return sdlsqrtf(x)
// }

// func srand(seed uint64)  {
//	sdlsrand(seed)
// }

// func sscanf(text string, fmt string) int32 {
//	return sdlsscanf(text, fmt)
// }

// func StepBackUTF8(start string, pstr **byte) uint32 {
//	return sdlStepBackUTF8(start, pstr)
// }

// func StepUTF8(pstr **byte, pslen *uint64) uint32 {
//	return sdlStepUTF8(pstr, pslen)
// }

// func strcasecmp(str1 string, str2 string) int32 {
//	return sdlstrcasecmp(str1, str2)
// }

// func strcasestr(haystack string, needle string) string {
//	return sdlstrcasestr(haystack, needle)
// }

// func strchr(str string, c int32) string {
//	return sdlstrchr(str, c)
// }

// func strcmp(str1 string, str2 string) int32 {
//	return sdlstrcmp(str1, str2)
// }

// func strdup(str string) string {
//	return sdlstrdup(str)
// }

// func strlcat(dst string, src string, maxlen uint64) uint64 {
//	return sdlstrlcat(dst, src, maxlen)
// }

// func strlcpy(dst string, src string, maxlen uint64) uint64 {
//	return sdlstrlcpy(dst, src, maxlen)
// }

// func strlen(str string) uint64 {
//	return sdlstrlen(str)
// }

// func strlwr(str string) string {
//	return sdlstrlwr(str)
// }

// func strncasecmp(str1 string, str2 string, maxlen uint64) int32 {
//	return sdlstrncasecmp(str1, str2, maxlen)
// }

// func strncmp(str1 string, str2 string, maxlen uint64) int32 {
//	return sdlstrncmp(str1, str2, maxlen)
// }

// func strndup(str string, maxlen uint64) string {
//	return sdlstrndup(str, maxlen)
// }

// func strnlen(str string, maxlen uint64) uint64 {
//	return sdlstrnlen(str, maxlen)
// }

// func strnstr(haystack string, needle string, maxlen uint64) string {
//	return sdlstrnstr(haystack, needle, maxlen)
// }

// func strpbrk(str string, breakset string) string {
//	return sdlstrpbrk(str, breakset)
// }

// func strrchr(str string, c int32) string {
//	return sdlstrrchr(str, c)
// }

// func strrev(str string) string {
//	return sdlstrrev(str)
// }

func Strstr(haystack string, needle string) string {
	return sdlstrstr(haystack, needle)
}

// func strtod(str string, endp **byte) float64 {
//	return sdlstrtod(str, endp)
// }

// func strtok_r(str string, delim string, saveptr **byte) string {
//	return sdlstrtok_r(str, delim, saveptr)
// }

// func strtol(str string, endp **byte, base int32) int64 {
//	return sdlstrtol(str, endp, base)
// }

// func strtoll(str string, endp **byte, base int32) int64 {
//	return sdlstrtoll(str, endp, base)
// }

// func strtoul(str string, endp **byte, base int32) uint64 {
//	return sdlstrtoul(str, endp, base)
// }

// func strtoull(str string, endp **byte, base int32) uint64 {
//	return sdlstrtoull(str, endp, base)
// }

// func strupr(str string) string {
//	return sdlstrupr(str)
// }

// func swprintf(text *wchar_t, maxlen uint64, fmt *wchar_t) int32 {
//	return sdlswprintf(text, maxlen, fmt)
// }

// func tan(x float64) float64 {
//	return sdltan(x)
// }

// func tanf(x float32) float32 {
//	return sdltanf(x)
// }

// func tolower(x int32) int32 {
//	return sdltolower(x)
// }

// func toupper(x int32) int32 {
//	return sdltoupper(x)
// }

// func trunc(x float64) float64 {
//	return sdltrunc(x)
// }

// func truncf(x float32) float32 {
//	return sdltruncf(x)
// }

// func UCS4ToUTF8(codepoint uint32, dst string) string {
//	return sdlUCS4ToUTF8(codepoint, dst)
// }

// func uitoa(value uint32, str string, radix int32) string {
//	return sdluitoa(value, str, radix)
// }

// func ulltoa(value uint64, str string, radix int32) string {
//	return sdlulltoa(value, str, radix)
// }

// func ultoa(value uint64, str string, radix int32) string {
//	return sdlultoa(value, str, radix)
// }

// func unsetenv_unsafe(name string) int32 {
//	return sdlunsetenv_unsafe(name)
// }

// func UnsetEnvironmentVariable(env *Environment, name string) bool {
//	return sdlUnsetEnvironmentVariable(env, name)
// }

// func utf8strlcpy(dst string, src string, dst_bytes uint64) uint64 {
//	return sdlutf8strlcpy(dst, src, dst_bytes)
// }

// func utf8strlen(str string) uint64 {
//	return sdlutf8strlen(str)
// }

// func utf8strnlen(str string, bytes uint64) uint64 {
//	return sdlutf8strnlen(str, bytes)
// }

// func vasprintf(strp **byte, fmt string, ap va_list) int32 {
//	return sdlvasprintf(strp, fmt, ap)
// }

// func vsnprintf(text string, maxlen uint64, fmt string, ap va_list) int32 {
//	return sdlvsnprintf(text, maxlen, fmt, ap)
// }

// func vsscanf(text string, fmt string, ap va_list) int32 {
//	return sdlvsscanf(text, fmt, ap)
// }

// func vswprintf(text *wchar_t, maxlen uint64, fmt *wchar_t, ap va_list) int32 {
//	return sdlvswprintf(text, maxlen, fmt, ap)
// }

// func wcscasecmp(str1 *wchar_t, str2 *wchar_t) int32 {
//	return sdlwcscasecmp(str1, str2)
// }

// func wcscmp(str1 *wchar_t, str2 *wchar_t) int32 {
//	return sdlwcscmp(str1, str2)
// }

// func wcsdup(wstr *wchar_t) *wchar_t {
//	return sdlwcsdup(wstr)
// }

// func wcslcat(dst *wchar_t, src *wchar_t, maxlen uint64) uint64 {
//	return sdlwcslcat(dst, src, maxlen)
// }

// func wcslcpy(dst *wchar_t, src *wchar_t, maxlen uint64) uint64 {
//	return sdlwcslcpy(dst, src, maxlen)
// }

// func wcslen(wstr *wchar_t) uint64 {
//	return sdlwcslen(wstr)
// }

// func wcsncasecmp(str1 *wchar_t, str2 *wchar_t, maxlen uint64) int32 {
//	return sdlwcsncasecmp(str1, str2, maxlen)
// }

// func wcsncmp(str1 *wchar_t, str2 *wchar_t, maxlen uint64) int32 {
//	return sdlwcsncmp(str1, str2, maxlen)
// }

// func wcsnlen(wstr *wchar_t, maxlen uint64) uint64 {
//	return sdlwcsnlen(wstr, maxlen)
// }

// func wcsnstr(haystack *wchar_t, needle *wchar_t, maxlen uint64) *wchar_t {
//	return sdlwcsnstr(haystack, needle, maxlen)
// }

// func wcsstr(haystack *wchar_t, needle *wchar_t) *wchar_t {
//	return sdlwcsstr(haystack, needle)
// }

// func wcstol(str *wchar_t, endp **wchar_t, base int32) int64 {
//	return sdlwcstol(str, endp, base)
// }
