package sdl

import (
	"unsafe"
)

const PropNameString = "SDL.name" // Available since SDL 3.4.0.

// [PropertiesID] is an ID that represents a properties set.
//
// [PropertiesID]: https://wiki.libsdl.org/SDL3/SDL_PropertiesID
type PropertiesID uint32

// [PropertyType] defines the SDL property type.
//
// [PropertyType]: https://wiki.libsdl.org/SDL3/SDL_PropertyType
type PropertyType uint32

const (
	PropertyTypeInvalid PropertyType = iota
	PropertyTypePointer
	PropertyTypeString
	PropertyTypeNumber
	PropertyTypeFloat
	PropertyTypeBoolean
)

// [GetGlobalProperties] gets the global SDL properties.
//
// [GetGlobalProperties]: https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties
func GetGlobalProperties() PropertiesID {
	return sdlGetGlobalProperties()
}

// [CreateProperties] creates a group of properties.
//
// [CreateProperties]: https://wiki.libsdl.org/SDL3/SDL_CreateProperties
func CreateProperties() PropertiesID {
	return sdlCreateProperties()
}

// [CopyProperties] copies a group of properties.
//
// [CopyProperties]: https://wiki.libsdl.org/SDL3/SDL_CopyProperties
func CopyProperties(src PropertiesID, dst PropertiesID) bool {
	return sdlCopyProperties(src, dst)
}

// [LockProperties] locks a group of properties.
//
// [LockProperties]: https://wiki.libsdl.org/SDL3/SDL_LockProperties
func LockProperties(props PropertiesID) bool {
	return sdlLockProperties(props)
}

// [UnlockProperties] unlocks a group of properties.
//
// [UnlockProperties]: https://wiki.libsdl.org/SDL3/SDL_UnlockProperties
func UnlockProperties(props PropertiesID) {
	sdlUnlockProperties(props)
}

// [CleanupPropertyCallback] is a callback used to free resources when a property is deleted.
//
// [CleanupPropertyCallback]: https://wiki.libsdl.org/SDL3/SDL_CleanupPropertyCallback
type CleanupPropertyCallback uintptr

// [SetPointerPropertyWithCleanup] sets a pointer property in a group of properties with a cleanup function that is called when the property is deleted.
//
// [SetPointerPropertyWithCleanup]: https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup
func SetPointerPropertyWithCleanup(props PropertiesID, name string, value unsafe.Pointer, cleanup CleanupPropertyCallback, userdata unsafe.Pointer) bool {
	return sdlSetPointerPropertyWithCleanup(props, name, value, cleanup, userdata)
}

// [SetPointerProperty] sets a pointer property in a group of properties.
//
// [SetPointerProperty]: https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty
func SetPointerProperty(props PropertiesID, name string, value unsafe.Pointer) bool {
	return sdlSetPointerProperty(props, name, value)
}

// [SetStringProperty] sets a string property in a group of properties.
//
// [SetStringProperty]: https://wiki.libsdl.org/SDL3/SDL_SetStringProperty
func SetStringProperty(props PropertiesID, name string, value string) bool {
	return sdlSetStringProperty(props, name, value)
}

// [SetNumberProperty] sets an integer property in a group of properties.
//
// [SetNumberProperty]: https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty
func SetNumberProperty(props PropertiesID, name string, value int64) bool {
	return sdlSetNumberProperty(props, name, value)
}

// [SetFloatProperty] sets a floating point property in a group of properties.
//
// [SetFloatProperty]: https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty
func SetFloatProperty(props PropertiesID, name string, value float32) bool {
	return sdlSetFloatProperty(props, name, value)
}

// [SetBooleanProperty] sets a boolean property in a group of properties.
//
// [SetBooleanProperty]: https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty
func SetBooleanProperty(props PropertiesID, name string, value bool) bool {
	return sdlSetBooleanProperty(props, name, value)
}

// [HasProperty] returns whether a property exists in a group of properties.
//
// [HasProperty]: https://wiki.libsdl.org/SDL3/SDL_HasProperty
func HasProperty(props PropertiesID, name string) bool {
	return sdlHasProperty(props, name)
}

// [GetPropertyType] gets the type of a property in a group of properties.
//
// [GetPropertyType]: https://wiki.libsdl.org/SDL3/SDL_GetPropertyType
func GetPropertyType(props PropertiesID, name string) PropertyType {
	return sdlGetPropertyType(props, name)
}

// [GetPointerProperty] gets a pointer property from a group of properties.
//
// [GetPointerProperty]: https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty
func GetPointerProperty(props PropertiesID, name string, defaultValue unsafe.Pointer) unsafe.Pointer {
	return sdlGetPointerProperty(props, name, defaultValue)
}

// [GetStringProperty] gets a string property from a group of properties.
//
// [GetStringProperty]: https://wiki.libsdl.org/SDL3/SDL_GetStringProperty
func GetStringProperty(props PropertiesID, name string, defaultValue string) string {
	return sdlGetStringProperty(props, name, defaultValue)
}

// [GetNumberProperty] gets a number property from a group of properties.
//
// [GetNumberProperty]: https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty
func GetNumberProperty(props PropertiesID, name string, defaultValue int64) int64 {
	return sdlGetNumberProperty(props, name, defaultValue)
}

// [GetFloatProperty] gets a floating point property from a group of properties.
//
// [GetFloatProperty]: https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty
func GetFloatProperty(props PropertiesID, name string, defaultValue float32) float32 {
	return sdlGetFloatProperty(props, name, defaultValue)
}

// [GetBooleanProperty] gets a boolean property from a group of properties.
//
// [GetBooleanProperty]: https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty
func GetBooleanProperty(props PropertiesID, name string, defaultValue bool) bool {
	return sdlGetBooleanProperty(props, name, defaultValue)
}

// [ClearProperty] clears a property from a group of properties.
//
// [ClearProperty]: https://wiki.libsdl.org/SDL3/SDL_ClearProperty
func ClearProperty(props PropertiesID, name string) bool {
	return sdlClearProperty(props, name)
}

// [EnumeratePropertiesCallback] is a callback used to enumerate all the properties in a group of properties.
//
// [EnumeratePropertiesCallback]: https://wiki.libsdl.org/SDL3/SDL_EnumeratePropertiesCallback
type EnumeratePropertiesCallback uintptr

// [EnumerateProperties] enumerates the properties contained in a group of properties.
//
// [EnumerateProperties]: https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties
func EnumerateProperties(props PropertiesID, callback EnumeratePropertiesCallback, userdata unsafe.Pointer) bool {
	return sdlEnumerateProperties(props, callback, userdata)
}

// [DestroyProperties] destroys a group of properties.
//
// [DestroyProperties]: https://wiki.libsdl.org/SDL3/SDL_DestroyProperties
func DestroyProperties(props PropertiesID) {
	sdlDestroyProperties(props)
}
