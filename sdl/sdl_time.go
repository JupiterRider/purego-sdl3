package sdl

// [DateTime] is a structure holding a calendar date and time broken down into its components.
//
// [DateTime]: https://wiki.libsdl.org/SDL3/SDL_DateTime
type DateTime struct {
	Year       int32 // Year.
	Month      int32 // Month [01-12].
	Day        int32 // Day of the month [01-31].
	Hour       int32 // Hour [0-23].
	Minute     int32 // Minute [0-59].
	Second     int32 // Seconds [0-60].
	Nanosecond int32 // Nanoseconds [0-999999999].
	DayOfWeek  int32 // Day of the week [0-6] (0 being Sunday).
	UtcOffset  int32 // Seconds east of UTC.
}

// [DateFormat] is a structure specifying the preferred date format of the current system locales.
//
// [DateFormat]: https://wiki.libsdl.org/SDL3/SDL_DateFormat
type DateFormat uint32

const (
	DateFormatYYYYMMDD DateFormat = iota // Year/Month/Day.
	DateFormatDDMMYYYY                   // Day/Month/Year.
	DateFormatMMDDYYYY                   // Month/Day/Year.
)

// [TimeFormat] is a structure specifying the preferred time format of the current system locales.
//
// [TimeFormat]: https://wiki.libsdl.org/SDL3/SDL_TimeFormat
type TimeFormat uint32

const (
	TimeFormat24HR TimeFormat = iota // 24 hour time.
	TimeFormat12HR                   // 12 hour time.
)

// func DateTimeToTime(dt *DateTime, ticks *Time) bool {
//	return sdlDateTimeToTime(dt, ticks)
// }

// func GetCurrentTime(ticks *Time) bool {
//	return sdlGetCurrentTime(ticks)
// }

// func GetDateTimeLocalePreferences(dateFormat *DateFormat, timeFormat *TimeFormat) bool {
//	return sdlGetDateTimeLocalePreferences(dateFormat, timeFormat)
// }

// func GetDayOfWeek(year int32, month int32, day int32) int32 {
//	return sdlGetDayOfWeek(year, month, day)
// }

// func GetDayOfYear(year int32, month int32, day int32) int32 {
//	return sdlGetDayOfYear(year, month, day)
// }

// func GetDaysInMonth(year int32, month int32) int32 {
//	return sdlGetDaysInMonth(year, month)
// }

// func TimeFromWindows(dwLowDateTime uint32, dwHighDateTime uint32) Time {
//	return sdlTimeFromWindows(dwLowDateTime, dwHighDateTime)
// }

// func TimeToDateTime(ticks Time, dt *DateTime, localTime bool) bool {
//	return sdlTimeToDateTime(ticks, dt, localTime)
// }

// func TimeToWindows(ticks Time, dwLowDateTime *uint32, dwHighDateTime *uint32)  {
//	sdlTimeToWindows(ticks, dwLowDateTime, dwHighDateTime)
// }
