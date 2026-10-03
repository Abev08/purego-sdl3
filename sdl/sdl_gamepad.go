package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

// [Gamepad] defines the structure used to identify an SDL gamepad.
//
// [Gamepad]: https://wiki.libsdl.org/SDL3/SDL_Gamepad
type Gamepad struct{}

// [GamepadType] defines the standard gamepad types.
//
// [GamepadType]: https://wiki.libsdl.org/SDL3/SDL_GamepadType
type GamepadType uint32

const (
	GamepadTypeUnknown GamepadType = iota
	GamepadTypeStandard
	GamepadTypeXbox360
	GamepadTypeXboxOne
	GamepadTypePS3
	GamepadTypePS4
	GamepadTypePS5
	GamepadTypeNintendoSwitchPro
	GamepadTypeNintendoSwitchJoyConLeft
	GamepadTypeNintendoSwitchJoyConRight
	GamepadTypeNintendoSwitchJoyConPair
	GamepadTypeGamecube
	GamepadTypeSteam
	GamepadTypeCount
)

// [GamepadButton] defines the list of buttons available on a gamepad.
//
// [GamepadButton]: https://wiki.libsdl.org/SDL3/SDL_GamepadButton
type GamepadButton int32

const (
	GamepadButtonInvalid GamepadButton = iota - 1
	GamepadButtonSouth                 // Bottom face button (e.g. Xbox A button).
	GamepadButtonEast                  // Right face button (e.g. Xbox B button).
	GamepadButtonWest                  // Left face button (e.g. Xbox X button).
	GamepadButtonNorth                 // Top face button (e.g. Xbox Y button).
	GamepadButtonBack
	GamepadButtonGuide
	GamepadButtonStart
	GamepadButtonLeftStick
	GamepadButtonRightStick
	GamepadButtonLeftShoulder
	GamepadButtonRightShoulder
	GamepadButtonDpadUp
	GamepadButtonDpadDown
	GamepadButtonDpadLeft
	GamepadButtonDpadRight
	GamepadButtonMisc1        // Additional button (e.g. Xbox Series X share button, PS5 microphone button, Nintendo Switch Pro capture button, Amazon Luna microphone button, Google Stadia capture button).
	GamepadButtonRightPaddle1 // Upper or primary paddle, under your right hand (e.g. Xbox Elite paddle P1, DualSense Edge RB button, Right Joy-Con SR button).
	GamepadButtonLeftPaddle1  // Upper or primary paddle, under your left hand (e.g. Xbox Elite paddle P3, DualSense Edge LB button, Left Joy-Con SL button).
	GamepadButtonRightPaddle2 // Lower or secondary paddle, under your right hand (e.g. Xbox Elite paddle P2, DualSense Edge right Fn button, Right Joy-Con SL button).
	GamepadButtonLeftPaddle2  // Lower or secondary paddle, under your left hand (e.g. Xbox Elite paddle P4, DualSense Edge left Fn button, Left Joy-Con SR button).
	GamepadButtonTouchpad     // PS4/PS5 touchpad button.
	GamepadButtonMisc2        // Additional button.
	GamepadButtonMisc3        // Additional button (e.g. Nintendo GameCube left trigger click).
	GamepadButtonMisc4        // Additional button (e.g. Nintendo GameCube right trigger click).
	GamepadButtonMisc5        // Additional button.
	GamepadButtonMisc6        // Additional button.
	GamepadButtonCount
)

// [GamepadButtonLabel] defines the set of gamepad button labels.
//
// [GamepadButtonLabel]: https://wiki.libsdl.org/SDL3/SDL_GamepadButtonLabel
type GamepadButtonLabel uint32

const (
	GamepadButtonLabelUnknown GamepadButtonLabel = iota
	GamepadButtonLabelA
	GamepadButtonLabelB
	GamepadButtonLabelX
	GamepadButtonLabelY
	GamepadButtonLabelCross
	GamepadButtonLabelCircle
	GamepadButtonLabelSquare
	GamepadButtonLabelTriangle
)

// [GamepadAxis] defines the list of axes available on a gamepad.
//
// [GamepadAxis]: https://wiki.libsdl.org/SDL3/SDL_GamepadAxis
type GamepadAxis int32

const (
	GamepadAxisInvalid GamepadAxis = iota - 1
	GamepadAxisLeftX
	GamepadAxisLeftY
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
	GamepadAxisCount
)

// [GamepadBindingType] defines the types of gamepad control bindings.
//
// [GamepadBindingType]: https://wiki.libsdl.org/SDL3/SDL_GamepadBindingType
type GamepadBindingType uint32

const (
	GamepadBindTypeNone GamepadBindingType = iota
	GamepadBindTypeButton
	GamepadBindTypeAxis
	GamepadBindTypeHat
)

// [GamepadBinding] defines a mapping between one joystick input to a gamepad control.
//
// [GamepadBinding]: https://wiki.libsdl.org/SDL3/SDL_GamepadBinding
type GamepadBinding struct {
	InputType  GamepadBindingType
	input      [12]byte
	OutputType GamepadBindingType
	output     [12]byte
}

func (g *GamepadBinding) InputButton() int32 {
	return *(*int32)(unsafe.Pointer(&g.input))
}

func (g *GamepadBinding) InputAxis() struct{ Axis, AxisMin, AxisMax int32 } {
	return *(*struct{ Axis, AxisMin, AxisMax int32 })(unsafe.Pointer(&g.input))
}

func (g *GamepadBinding) InputHat() struct{ Hat, HatMask int32 } {
	return *(*struct{ Hat, HatMask int32 })(unsafe.Pointer(&g.input))
}

func (g *GamepadBinding) OutputButton() GamepadButton {
	return *(*GamepadButton)(unsafe.Pointer(&g.output))
}

func (g *GamepadBinding) OutputAxis() struct {
	Axis             GamepadAxis
	AxisMin, AxisMax int32
} {
	return *(*struct {
		Axis             GamepadAxis
		AxisMin, AxisMax int32
	})(unsafe.Pointer(&g.output))
}

// [AddGamepadMapping] adds support for gamepads that SDL is unaware of or change the binding of an existing gamepad.
//
// [AddGamepadMapping]: https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping
// func AddGamepadMapping(mapping string) int32 {
//	return sdlAddGamepadMapping(mapping)
// }

// [AddGamepadMappingsFromIO] loads a set of gamepad mappings from an [IOStream].
//
// [AddGamepadMappingsFromIO]: https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO
// func AddGamepadMappingsFromIO(src *IOStream, closeio bool) int32 {
//	return sdlAddGamepadMappingsFromIO(src, closeio)
// }

// [AddGamepadMappingsFromFile] loads a set of gamepad mappings from a file.
//
// [AddGamepadMappingsFromFile]: https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile
// func AddGamepadMappingsFromFile(file string) int32 {
//	return sdlAddGamepadMappingsFromFile(file)
// }

// [ReloadGamepadMappings] reinitializes the SDL mapping database to its initial state.
//
// [ReloadGamepadMappings]: https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings
// func ReloadGamepadMappings() bool {
//	return sdlReloadGamepadMappings()
// }

// [GetGamepadMappings] gets the current gamepad mappings.
//
// [GetGamepadMappings]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings
// func GetGamepadMappings(count *int32) **byte {
//	return sdlGetGamepadMappings(count)
// }

// [GetGamepadMappingForGUID] gets the gamepad mapping string for a given GUID.
//
// [GetGamepadMappingForGUID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID
// func GetGamepadMappingForGUID(guid GUID) string {
//	return sdlGetGamepadMappingForGUID(guid)
// }

// [GetGamepadMapping] gets the current mapping of a gamepad.
//
// [GetGamepadMapping]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping
// func GetGamepadMapping(gamepad *Gamepad) string {
//	return sdlGetGamepadMapping(gamepad)
// }

// [SetGamepadMapping] sets the current mapping of a joystick or gamepad.
//
// [SetGamepadMapping]: https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping
// func SetGamepadMapping(instance_id JoystickID, mapping string) bool {
//	return sdlSetGamepadMapping(instance_id, mapping)
// }

// [HasGamepad] returns whether a gamepad is currently connected.
//
// [HasGamepad]: https://wiki.libsdl.org/SDL3/SDL_HasGamepad
// func HasGamepad() bool {
//	return sdlHasGamepad()
// }

// [GetGamepads] returns a list of currently connected gamepads or nil on failure.
//
// The returned slice must not be freed with e.g. [Free].
//
// [GetGamepads]: https://wiki.libsdl.org/SDL3/SDL_GetGamepads
func GetGamepads() []JoystickID {
	var count int32
	gamepads := sdlGetGamepads(&count)
	defer Free(unsafe.Pointer(gamepads))
	return mem.Copy(gamepads, count)
}

// [IsGamepad] checks if the given joystick is supported by the gamepad interface.
//
// [IsGamepad]: https://wiki.libsdl.org/SDL3/SDL_IsGamepad
// func IsGamepad(instance_id JoystickID) bool {
//	return sdlIsGamepad(instance_id)
// }

// [GetGamepadNameForID] gets the implementation dependent name of a gamepad.
//
// [GetGamepadNameForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID
func GetGamepadNameForID(instanceId JoystickID) string {
	return sdlGetGamepadNameForID(instanceId)
}

// [GetGamepadPathForID] gets the implementation dependent path of a gamepad.
//
// [GetGamepadPathForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID
// func GetGamepadPathForID(instance_id JoystickID) string {
//	return sdlGetGamepadPathForID(instance_id)
// }

// [GetGamepadPlayerIndexForID] gets the player index of a gamepad.
//
// [GetGamepadPlayerIndexForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID
// func GetGamepadPlayerIndexForID(instance_id JoystickID) int32 {
//	return sdlGetGamepadPlayerIndexForID(instance_id)
// }

// [GetGamepadGUIDForID] gets the implementation-dependent GUID of a gamepad.
//
// [GetGamepadGUIDForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID
// func GetGamepadGUIDForID(instance_id JoystickID) GUID {
//	return sdlGetGamepadGUIDForID(instance_id)
// }

// [GetGamepadVendorForID] gets the USB vendor ID of a gamepad, if available.
//
// [GetGamepadVendorForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID
// func GetGamepadVendorForID(instance_id JoystickID) uint16 {
//	return sdlGetGamepadVendorForID(instance_id)
// }

// [GetGamepadProductForID] gets the USB product ID of a gamepad, if available.
//
// [GetGamepadProductForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID
// func GetGamepadProductForID(instance_id JoystickID) uint16 {
//	return sdlGetGamepadProductForID(instance_id)
// }

// [GetGamepadProductVersionForID] gets the product version of a gamepad, if available.
//
// [GetGamepadProductVersionForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID
// func GetGamepadProductVersionForID(instance_id JoystickID) uint16 {
//	return sdlGetGamepadProductVersionForID(instance_id)
// }

// [GetGamepadTypeForID] gets the type of a gamepad.
//
// [GetGamepadTypeForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID
// func GetGamepadTypeForID(instance_id JoystickID) GamepadType {
//	return sdlGetGamepadTypeForID(instance_id)
// }

// [GetRealGamepadTypeForID] gets the type of a gamepad, ignoring any mapping override.
//
// [GetRealGamepadTypeForID]: https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID
// func GetRealGamepadTypeForID(instance_id JoystickID) GamepadType {
//	return sdlGetRealGamepadTypeForID(instance_id)
// }

// [GetGamepadMappingForID] gets the mapping of a gamepad.
//
// [GetGamepadMappingForID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID
// func GetGamepadMappingForID(instance_id JoystickID) string {
//	return sdlGetGamepadMappingForID(instance_id)
// }

// [OpenGamepad] opens a gamepad for use.
//
// [OpenGamepad]: https://wiki.libsdl.org/SDL3/SDL_OpenGamepad
func OpenGamepad(instanceId JoystickID) *Gamepad {
	return sdlOpenGamepad(instanceId)
}

// [GetGamepadFromID] gets the [Gamepad] associated with a joystick instance ID, if it has been opened.
//
// [GetGamepadFromID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID
func GetGamepadFromID(instanceId JoystickID) *Gamepad {
	return sdlGetGamepadFromID(instanceId)
}

// [GetGamepadFromPlayerIndex] gets the [Gamepad] associated with a player index.
//
// [GetGamepadFromPlayerIndex]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex
// func GetGamepadFromPlayerIndex(player_index int32) *Gamepad {
//	return sdlGetGamepadFromPlayerIndex(player_index)
// }

// [GetGamepadProperties] gets the properties associated with an opened gamepad.
//
// [GetGamepadProperties]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties
// func GetGamepadProperties(gamepad *Gamepad) PropertiesID {
//	return sdlGetGamepadProperties(gamepad)
// }

// [GetGamepadID] gets the instance ID of an opened gamepad.
//
// [GetGamepadID]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadID
// func GetGamepadID(gamepad *Gamepad) JoystickID {
//	return sdlGetGamepadID(gamepad)
// }

// [GetGamepadName] gets the implementation-dependent name for an opened gamepad.
//
// [GetGamepadName]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadName
func GetGamepadName(gamepad *Gamepad) string {
	return sdlGetGamepadName(gamepad)
}

// [GetGamepadPath] gets the implementation-dependent path for an opened gamepad.
//
// [GetGamepadPath]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath
// func GetGamepadPath(gamepad *Gamepad) string {
//	return sdlGetGamepadPath(gamepad)
// }

// [GetGamepadType] gets the type of an opened gamepad.
//
// [GetGamepadType]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadType
func GetGamepadType(gamepad *Gamepad) GamepadType {
	return sdlGetGamepadType(gamepad)
}

// [GetRealGamepadType] gets the type of an opened gamepad, ignoring any mapping override.
//
// [GetRealGamepadType]: https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType
// func GetRealGamepadType(gamepad *Gamepad) GamepadType {
//	return sdlGetRealGamepadType(gamepad)
// }

// [GetGamepadPlayerIndex] gets the player index of an opened gamepad.
//
// [GetGamepadPlayerIndex]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex
// func GetGamepadPlayerIndex(gamepad *Gamepad) int32 {
//	return sdlGetGamepadPlayerIndex(gamepad)
// }

// [SetGamepadPlayerIndex] sets the player index of an opened gamepad.
//
// [SetGamepadPlayerIndex]: https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex
// func SetGamepadPlayerIndex(gamepad *Gamepad, player_index int32) bool {
//	return sdlSetGamepadPlayerIndex(gamepad, player_index)
// }

// [GetGamepadVendor] gets the USB vendor ID of an opened gamepad, if available.
//
// [GetGamepadVendor]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor
// func GetGamepadVendor(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadVendor(gamepad)
// }

// [GetGamepadProduct] gets the USB product ID of an opened gamepad, if available.
//
// [GetGamepadProduct]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct
// func GetGamepadProduct(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadProduct(gamepad)
// }

// [GetGamepadProductVersion] gets the product version of an opened gamepad, if available.
//
// [GetGamepadProductVersion]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion
// func GetGamepadProductVersion(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadProductVersion(gamepad)
// }

// [GetGamepadFirmwareVersion] gets the firmware version of an opened gamepad, if available.
//
// [GetGamepadFirmwareVersion]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion
// func GetGamepadFirmwareVersion(gamepad *Gamepad) uint16 {
//	return sdlGetGamepadFirmwareVersion(gamepad)
// }

// [GetGamepadSerial] returns the serial number of an opened gamepad, or "" if unavailable.
//
// [GetGamepadSerial]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial
func GetGamepadSerial(gamepad *Gamepad) string {
	return sdlGetGamepadSerial(gamepad)
}

// [GetGamepadSteamHandle] gets the Steam Input handle of an opened gamepad, if available.
//
// [GetGamepadSteamHandle]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle
// func GetGamepadSteamHandle(gamepad *Gamepad) uint64 {
//	return sdlGetGamepadSteamHandle(gamepad)
// }

// [GetGamepadConnectionState] gets the connection state of a gamepad.
//
// [GetGamepadConnectionState]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState
// func GetGamepadConnectionState(gamepad *Gamepad) JoystickConnectionState {
//	return sdlGetGamepadConnectionState(gamepad)
// }

// [GetGamepadPowerInfo] gets the battery state of a gamepad.
//
// [GetGamepadPowerInfo]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo
// func GetGamepadPowerInfo(gamepad *Gamepad, percent *int32) PowerState {
//	return sdlGetGamepadPowerInfo(gamepad, percent)
// }

// [GamepadConnected] checks if a gamepad has been opened and is currently connected.
//
// [GamepadConnected]: https://wiki.libsdl.org/SDL3/SDL_GamepadConnected
// func GamepadConnected(gamepad *Gamepad) bool {
//	return sdlGamepadConnected(gamepad)
// }

// [GetGamepadJoystick] gets the underlying joystick from a gamepad.
//
// [GetGamepadJoystick]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick
// func GetGamepadJoystick(gamepad *Gamepad) *Joystick {
//	return sdlGetGamepadJoystick(gamepad)
// }

// [SetGamepadEventsEnabled] sets the state of gamepad event processing.
//
// [SetGamepadEventsEnabled]: https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled
// func SetGamepadEventsEnabled(enabled bool)  {
//	sdlSetGamepadEventsEnabled(enabled)
// }

// [GamepadEventsEnabled] queries the state of gamepad event processing.
//
// [GamepadEventsEnabled]: https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled
// func GamepadEventsEnabled() bool {
//	return sdlGamepadEventsEnabled()
// }

// [GetGamepadBindings] returns the SDL joystick layer bindings for a gamepad or nil on failure.
//
// The returned slice must not be freed with e.g. [Free].
//
// [GetGamepadBindings]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings
func GetGamepadBindings(gamepad *Gamepad) []*GamepadBinding {
	var count int32
	bindings := sdlGetGamepadBindings(gamepad, &count)
	if bindings == nil {
		return nil
	}
	defer Free(unsafe.Pointer(bindings))
	return mem.DeepCopy(bindings, count)
}

// [UpdateGamepads] manually pump gamepad updates if not using the loop.
//
// [UpdateGamepads]: https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads
// func UpdateGamepads()  {
//	sdlUpdateGamepads()
// }

// [GetGamepadTypeFromString] converts a string into [GamepadType] enum.
//
// [GetGamepadTypeFromString]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString
// func GetGamepadTypeFromString(str string) GamepadType {
//	return sdlGetGamepadTypeFromString(str)
// }

// [GetGamepadStringForType] converts from an [GamepadType] enum to a string.
//
// [GetGamepadStringForType]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType
func GetGamepadStringForType(gamepadType GamepadType) string {
	return sdlGetGamepadStringForType(gamepadType)
}

// [GetGamepadAxisFromString] converts a string into [GamepadAxis] enum.
//
// [GetGamepadAxisFromString]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString
// func GetGamepadAxisFromString(str string) GamepadAxis {
//	return sdlGetGamepadAxisFromString(str)
// }

// [GetGamepadStringForAxis] converts from an [GamepadAxis] enum to a string.
//
// [GetGamepadStringForAxis]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis
// func GetGamepadStringForAxis(axis GamepadAxis) string {
//	return sdlGetGamepadStringForAxis(axis)
// }

// [GamepadHasAxis] queries whether a gamepad has a given axis.
//
// [GamepadHasAxis]: https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis
// func GamepadHasAxis(gamepad *Gamepad, axis GamepadAxis) bool {
//	return sdlGamepadHasAxis(gamepad, axis)
// }

// [GetGamepadAxis] gets the current state of an axis control on a gamepad.
//
// [GetGamepadAxis]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis
// func GetGamepadAxis(gamepad *Gamepad, axis GamepadAxis) int16 {
//	return sdlGetGamepadAxis(gamepad, axis)
// }

// [GetGamepadButtonFromString] converts a string into an [GamepadButton] enum.
//
// [GetGamepadButtonFromString]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString
// func GetGamepadButtonFromString(str string) GamepadButton {
//	return sdlGetGamepadButtonFromString(str)
// }

// [GetGamepadStringForButton] returns the name for the given button.
//
// [GetGamepadStringForButton]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton
func GetGamepadStringForButton(button GamepadButton) string {
	return sdlGetGamepadStringForButton(button)
}

// [GamepadHasButton] queries whether a gamepad has a given button.
//
// [GamepadHasButton]: https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton
// func GamepadHasButton(gamepad *Gamepad, button GamepadButton) bool {
//	return sdlGamepadHasButton(gamepad, button)
// }

// [GetGamepadButton] gets the current state of a button on a gamepad.
//
// [GetGamepadButton]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton
// func GetGamepadButton(gamepad *Gamepad, button GamepadButton) bool {
//	return sdlGetGamepadButton(gamepad, button)
// }

// [GetGamepadButtonLabelForType] gets the label of a button on a gamepad.
//
// [GetGamepadButtonLabelForType]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType
// func GetGamepadButtonLabelForType(type GamepadType, button GamepadButton) GamepadButtonLabel {
//	return sdlGetGamepadButtonLabelForType(type, button)
// }

// [GetGamepadButtonLabel] gets the label of a button on a gamepad.
//
// [GetGamepadButtonLabel]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel
// func GetGamepadButtonLabel(gamepad *Gamepad, button GamepadButton) GamepadButtonLabel {
//	return sdlGetGamepadButtonLabel(gamepad, button)
// }

// [GetNumGamepadTouchpads] gets the number of touchpads on a gamepad.
//
// [GetNumGamepadTouchpads]: https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads
// func GetNumGamepadTouchpads(gamepad *Gamepad) int32 {
//	return sdlGetNumGamepadTouchpads(gamepad)
// }

// [GetNumGamepadTouchpadFingers] gets the number of supported simultaneous fingers on a touchpad on a game.
//
// [GetNumGamepadTouchpadFingers]: https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers
// func GetNumGamepadTouchpadFingers(gamepad *Gamepad, touchpad int32) int32 {
//	return sdlGetNumGamepadTouchpadFingers(gamepad, touchpad)
// }

// [GetGamepadTouchpadFinger] gets the current state of a finger on a touchpad on a gamepad.
//
// [GetGamepadTouchpadFinger]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger
// func GetGamepadTouchpadFinger(gamepad *Gamepad, touchpad int32, finger int32, down *bool, x *float32, y *float32, pressure *float32) bool {
//	return sdlGetGamepadTouchpadFinger(gamepad, touchpad, finger, down, x, y, pressure)
// }

// [GamepadHasSensor] returns whether a gamepad has a particular sensor.
//
// [GamepadHasSensor]: https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor
// func GamepadHasSensor(gamepad *Gamepad, type SensorType) bool {
//	return sdlGamepadHasSensor(gamepad, type)
// }

// [SetGamepadSensorEnabled] sets whether data reporting for a gamepad sensor is enabled.
//
// [SetGamepadSensorEnabled]: https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled
// func SetGamepadSensorEnabled(gamepad *Gamepad, type SensorType, enabled bool) bool {
//	return sdlSetGamepadSensorEnabled(gamepad, type, enabled)
// }

// [GamepadSensorEnabled] queries whether sensor data reporting is enabled for a gamepad.
//
// [GamepadSensorEnabled]: https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled
// func GamepadSensorEnabled(gamepad *Gamepad, type SensorType) bool {
//	return sdlGamepadSensorEnabled(gamepad, type)
// }

// [GetGamepadSensorDataRate] gets the data rate (number of events per second) of a gamepad sensor.
//
// [GetGamepadSensorDataRate]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate
// func GetGamepadSensorDataRate(gamepad *Gamepad, type SensorType) float32 {
//	return sdlGetGamepadSensorDataRate(gamepad, type)
// }

// [GetGamepadSensorData] gets the current state of a gamepad sensor.
//
// [GetGamepadSensorData]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData
// func GetGamepadSensorData(gamepad *Gamepad, type SensorType, data *float32, num_values int32) bool {
//	return sdlGetGamepadSensorData(gamepad, type, data, num_values)
// }

// [RumbleGamepad] starts a rumble effect on a gamepad.
//
// [RumbleGamepad]: https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad
// func RumbleGamepad(gamepad *Gamepad, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
//	return sdlRumbleGamepad(gamepad, low_frequency_rumble, high_frequency_rumble, duration_ms)
// }

// [RumbleGamepadTriggers] starts a rumble effect in the gamepad's triggers.
//
// [RumbleGamepadTriggers]: https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers
// func RumbleGamepadTriggers(gamepad *Gamepad, left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
//	return sdlRumbleGamepadTriggers(gamepad, left_rumble, right_rumble, duration_ms)
// }

// [SetGamepadLED] updates a gamepad's LED color.
//
// [SetGamepadLED]: https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED
// func SetGamepadLED(gamepad *Gamepad, red uint8, green uint8, blue uint8) bool {
//	return sdlSetGamepadLED(gamepad, red, green, blue)
// }

// [SendGamepadEffect] sends a gamepad specific effect packet.
//
// [SendGamepadEffect]: https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect
// func SendGamepadEffect(gamepad *Gamepad, data unsafe.Pointer, size int32) bool {
//	return sdlSendGamepadEffect(gamepad, data, size)
// }

// [CloseGamepad] closes a gamepad previously opened with [OpenGamepad].
//
// [CloseGamepad]: https://wiki.libsdl.org/SDL3/SDL_CloseGamepad
func CloseGamepad(gamepad *Gamepad) {
	sdlCloseGamepad(gamepad)
}

// [GetGamepadAppleSFSymbolsNameForButton] returns the sfSymbolsName for a given button on a gamepad on Apple.
//
// [GetGamepadAppleSFSymbolsNameForButton]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton
// func GetGamepadAppleSFSymbolsNameForButton(gamepad *Gamepad, button GamepadButton) string {
//	return sdlGetGamepadAppleSFSymbolsNameForButton(gamepad, button)
// }

// [GetGamepadAppleSFSymbolsNameForAxis] returns the sfSymbolsName for a given axis on a gamepad on Apple platforms.
//
// [GetGamepadAppleSFSymbolsNameForAxis]: https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis
// func GetGamepadAppleSFSymbolsNameForAxis(gamepad *Gamepad, axis GamepadAxis) string {
//	return sdlGetGamepadAppleSFSymbolsNameForAxis(gamepad, axis)
// }
