package sdl

// [DateTime] specifies a structure holding a calendar date and time broken down into its components.
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

// [DateFormat] specifies the preferred date format of the current system locale.
//
// [DateFormat]: https://wiki.libsdl.org/SDL3/SDL_DateFormat
type DateFormat uint32

const (
	DateFormatYYYYMMDD DateFormat = iota // Year/Month/Day.
	DateFormatDDMMYYYY                   // Day/Month/Year.
	DateFormatMMDDYYYY                   // Month/Day/Year.
)

// [TimeFormat] specifies the preferred time format of the current system locale.
//
// [TimeFormat]: https://wiki.libsdl.org/SDL3/SDL_TimeFormat
type TimeFormat uint32

const (
	TimeFormat24HR TimeFormat = iota // 24 hour time.
	TimeFormat12HR                   // 12 hour time.
)

// [GetDateTimeLocalePreferences] gets the current preferred date and time format for the system locale.
//
// [GetDateTimeLocalePreferences]: https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences
// func GetDateTimeLocalePreferences(dateFormat *DateFormat, timeFormat *TimeFormat) bool {
//	return sdlGetDateTimeLocalePreferences(dateFormat, timeFormat)
// }

// [GetCurrentTime] gets the current value of the system realtime clock in nanoseconds since.
//
// [GetCurrentTime]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentTime
// func GetCurrentTime(ticks *Time) bool {
//	return sdlGetCurrentTime(ticks)
// }

// [TimeToDateTime] converts an [Time] in nanoseconds since the epoch to a calendar time in.
//
// [TimeToDateTime]: https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime
// func TimeToDateTime(ticks Time, dt *DateTime, localTime bool) bool {
//	return sdlTimeToDateTime(ticks, dt, localTime)
// }

// [DateTimeToTime] converts a calendar time to an [Time] in nanoseconds since the epoch.
//
// [DateTimeToTime]: https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime
// func DateTimeToTime(dt *DateTime, ticks *Time) bool {
//	return sdlDateTimeToTime(dt, ticks)
// }

// [TimeToWindows] converts an SDL time into a Windows FILETIME (100-nanosecond intervals.
//
// [TimeToWindows]: https://wiki.libsdl.org/SDL3/SDL_TimeToWindows
// func TimeToWindows(ticks Time, dwLowDateTime *uint32, dwHighDateTime *uint32) {
//	sdlTimeToWindows(ticks, dwLowDateTime, dwHighDateTime)
// }

// [TimeFromWindows] converts a Windows FILETIME (100-nanosecond intervals since January 1,.
//
// [TimeFromWindows]: https://wiki.libsdl.org/SDL3/SDL_TimeFromWindows
// func TimeFromWindows(dwLowDateTime uint32, dwHighDateTime uint32) Time {
//	return sdlTimeFromWindows(dwLowDateTime, dwHighDateTime)
// }

// [GetDaysInMonth] gets the number of days in a month for a given year.
//
// [GetDaysInMonth]: https://wiki.libsdl.org/SDL3/SDL_GetDaysInMonth
// func GetDaysInMonth(year int32, month int32) int32 {
//	return sdlGetDaysInMonth(year, month)
// }

// [GetDayOfYear] gets the day of year for a calendar date.
//
// [GetDayOfYear]: https://wiki.libsdl.org/SDL3/SDL_GetDayOfYear
// func GetDayOfYear(year int32, month int32, day int32) int32 {
//	return sdlGetDayOfYear(year, month, day)
// }

// [GetDayOfWeek] gets the day of week for a calendar date.
//
// [GetDayOfWeek]: https://wiki.libsdl.org/SDL3/SDL_GetDayOfWeek
// func GetDayOfWeek(year int32, month int32, day int32) int32 {
//	return sdlGetDayOfWeek(year, month, day)
// }
