package sdl

const PropHidapiLibusbDeviceHandlePointer = "SDL.hidapi.libusb.device.handle"

// [HidDevice] is an opaque handle representing an open HID device.
//
// [HidDevice]: https://wiki.libsdl.org/SDL3/SDL_hid_device
type HidDevice struct{}

// [HidBusType] defines HID underlying bus types.
//
// [HidBusType]: https://wiki.libsdl.org/SDL3/SDL_hid_bus_type
type HidBusType uint32

const (
	HidApiBusUnknown   HidBusType = iota // Unknown bus type.
	HidApiBusUSB                         // USB bus. Specifications: https://usb.org/hid.
	HidApiBusBluetooth                   // Bluetooth or Bluetooth LE bus. Specifications: https://www.bluetooth.com/specifications/specs/human-interface-device-profile-1-1-1/, https://www.bluetooth.com/specifications/specs/hid-service-1-0/, https://www.bluetooth.com/specifications/specs/hid-over-gatt-profile-1-0/.
	HidApiBusI2C                         // I2C bus. Specifications: https://docs.microsoft.com/previous-versions/windows/hardware/design/dn642101(v=vs.85).
	HidApiBusSPI                         // SPI bus. Specifications: https://www.microsoft.com/download/details.aspx?id=103325.
)

// [HidDeviceInfo] describes information about a connected HID device
//
// [HidDeviceInfo]: https://wiki.libsdl.org/SDL3/SDL_hid_device_info
type HidDeviceInfo struct {
	Path               *byte  // Platform-specific device path.
	VendorId           uint16 // Device Vendor ID.
	ProductId          uint16 // Device Product ID.
	SerialNumber       *byte  // Serial Number.
	ReleaseNumber      uint16 // Device Release Number in binary-coded decimal, also known as Device Version Number.
	ManufacturerString *byte  // Manufacturer.
	ProductString      *byte  // Product.
	UsagePage          uint16 // Usage Page for this Device/Interface (Windows/Mac/hidraw only).
	Usage              uint16 // Usage for this Device/Interface (Windows/Mac/hidraw only).
	InterfaceNumber    int32  // The USB interface which this logical device represents. Valid only if the device is a USB HID device. Set to -1 in all other cases.
	InterfaceClass     int32  // Additional information about the USB interface. Valid on libusb and Android implementations.
	InterfaceSubclass  int32
	InterfaceProtocol  int32
	BusType            HidBusType     // Underlying bus type.
	Next               *HidDeviceInfo // Pointer to the next device.
}

// [HidInit] initializes the HIDAPI library.
//
// [HidInit]: https://wiki.libsdl.org/SDL3/SDL_hid_init
// func HidInit() int32 {
// 	return sdlhid_init()
// }

// [HidExit] finalizes the HIDAPI library.
//
// [HidExit]: https://wiki.libsdl.org/SDL3/SDL_hid_exit
// func HidExit() int32 {
// 	return sdlhid_exit()
// }

// [HidDeviceChangeCount] checks to see if devices may have been added or removed.
//
// [HidDeviceChangeCount]: https://wiki.libsdl.org/SDL3/SDL_hid_device_change_count
// func HidDeviceChangeCount() uint32 {
// 	return sdlhid_device_change_count()
// }

// [HidEnumerate] enumerates the HID Devices.
//
// [HidEnumerate]: https://wiki.libsdl.org/SDL3/SDL_hid_enumerate
// func HidEnumerate(vendorId uint16, productId uint16) *HidDeviceInfo {
// 	return sdlhid_enumerate(vendorId, productId)
// }

// [HidFreeEnumeration] frees an enumeration linked list.
//
// [HidFreeEnumeration]: https://wiki.libsdl.org/SDL3/SDL_hid_free_enumeration
// func HidFreeEnumeration(devs *HidDeviceInfo) {
//	sdlhid_free_enumeration(devs)
// }

// [HidOpen] opens a HID device using a Vendor ID (VID), Product ID (PID) and optionally.
//
// [HidOpen]: https://wiki.libsdl.org/SDL3/SDL_hid_open
// func HidOpen(vendorId uint16, productId uint16, serialNumber *byte) *HidDevice {
// return sdlhid_open(vendor_id, product_id, serial_number)
// }

// [HidOpenPath] opens a HID device by its path name.
//
// [HidOpenPath]: https://wiki.libsdl.org/SDL3/SDL_hid_open_path
// func HidOpenPath(path *byte) *HidDevice {
//	return sdlhid_open_path(path)
// }

// [HidGetProperties] gets the properties associated with an [HidDevice].
//
// Available since SDL 3.4.0.
//
// [HidGetProperties]: https://wiki.libsdl.org/SDL3/SDL_hid_get_properties
// func HidGetProperties(dev *HidDevice) PropertiesID {
// 	return sdlhid_get_properties(dev)
// }

// [HidWrite] writes an Output report to a HID device.
//
// [HidWrite]: https://wiki.libsdl.org/SDL3/SDL_hid_write
// func HidWrite(dev *HidDevice, data *uint8, length uint64) int32 {
//	return sdlhid_write(dev, data, length)
// }

// [HidReadTimeout] reads an Input report from a HID device with timeout.
//
// [HidReadTimeout]: https://wiki.libsdl.org/SDL3/SDL_hid_read_timeout
// func HidReadTimeout(dev *HidDevice, data *uint8, length uint64, milliseconds int32) int32 {
//	return sdlhid_read_timeout(dev, data, length, milliseconds)
// }

// [HidRead] reads an Input report from a HID device.
//
// [HidRead]: https://wiki.libsdl.org/SDL3/SDL_hid_read
// func HidRead(dev *HidDevice, data *uint8, length uint64) int32 {
//	return sdlhid_read(dev, data, length)
// }

// [HidSetNonblocking] sets the device handle to be non-blocking.
//
// [HidSetNonblocking]: https://wiki.libsdl.org/SDL3/SDL_hid_set_nonblocking
// func HidSetNonblocking(dev *HidDevice, nonblock int32) int32 {
//	return sdlhid_set_nonblocking(dev, nonblock)
// }

// [HidSendFeatureReport] sends a Feature report to the device.
//
// [HidSendFeatureReport]: https://wiki.libsdl.org/SDL3/SDL_hid_send_feature_report
// func HidSendFeatureReport(dev *HidDevice, data *uint8, length uint64) int32 {
//	return sdlhid_send_feature_report(dev, data, length)
// }

// [HidGetFeatureReport] gets a feature report from a HID device.
//
// [HidGetFeatureReport]: https://wiki.libsdl.org/SDL3/SDL_hid_get_feature_report
// func HidGetFeatureReport(dev *HidDevice, data *uint8, length uint64) int32 {
//	return sdlhid_get_feature_report(dev, data, length)
// }

// [HidGetInputReport] gets an input report from a HID device.
//
// [HidGetInputReport]: https://wiki.libsdl.org/SDL3/SDL_hid_get_input_report
// func HidGetInputReport(dev *HidDevice, data *uint8, length uint64) int32 {
//	return sdlhid_get_input_report(dev, data, length)
// }

// [HidClose] closes a HID device.
//
// [HidClose]: https://wiki.libsdl.org/SDL3/SDL_hid_close
// func HidClose(dev *HidDevice) int32 {
//	return sdlhid_close(dev)
// }

// [HidGetManufacturerString] gets The Manufacturer String from a HID device.
//
// [HidGetManufacturerString]: https://wiki.libsdl.org/SDL3/SDL_hid_get_manufacturer_string
// func HidGetManufacturerString(dev *HidDevice, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_manufacturer_string(dev, string, maxlen)
// }

// [HidGetProductString] gets The Product String from a HID device.
//
// [HidGetProductString]: https://wiki.libsdl.org/SDL3/SDL_hid_get_product_string
// func HidGetProductString(dev *HidDevice, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_product_string(dev, string, maxlen)
// }

// [HidGetSerialNumberString] gets The Serial Number String from a HID device.
//
// [HidGetSerialNumberString]: https://wiki.libsdl.org/SDL3/SDL_hid_get_serial_number_string
// func HidGetSerialNumberString(dev *HidDevice, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_serial_number_string(dev, string, maxlen)
// }

// [HidGetIndexedString] gets a string from a HID device, based on its string index.
//
// [HidGetIndexedString]: https://wiki.libsdl.org/SDL3/SDL_hid_get_indexed_string
// func HidGetIndexedString(dev *HidDevice, string_index int32, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_indexed_string(dev, string_index, string, maxlen)
// }

// [HidGetDeviceInfo] gets the device info from a HID device.
//
// [HidGetDeviceInfo]: https://wiki.libsdl.org/SDL3/SDL_hid_get_device_info
// func HidGetDeviceInfo(dev *HidDevice) *HidDeviceInfo {
//	return sdlhid_get_device_info(dev)
// }

// [HidGetReportDescriptor] gets a report descriptor from a HID device.
//
// [HidGetReportDescriptor]: https://wiki.libsdl.org/SDL3/SDL_hid_get_report_descriptor
// func HidGetReportDescriptor(dev *HidDevice, buf *uint8, buf_size uint64) int32 {
//	return sdlhid_get_report_descriptor(dev, buf, buf_size)
// }

// [HidBleScan] starts or stop a BLE scan on iOS and tvOS to pair Steam Controllers.
//
// [HidBleScan]: https://wiki.libsdl.org/SDL3/SDL_hid_ble_scan
// func HidBleScan(active bool)  {
//	sdlhid_ble_scan(active)
// }
