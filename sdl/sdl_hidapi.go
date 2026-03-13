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

// func hid_ble_scan(active bool)  {
//	sdlhid_ble_scan(active)
// }

// func hid_close(dev *hid_device) int32 {
//	return sdlhid_close(dev)
// }

// func hid_device_change_count() uint32 {
//	return sdlhid_device_change_count()
// }

// func hid_enumerate(vendor_id uint16, product_id uint16) *hid_device_info {
//	return sdlhid_enumerate(vendor_id, product_id)
// }

// func hid_exit() int32 {
//	return sdlhid_exit()
// }

// func hid_free_enumeration(devs *hid_device_info)  {
//	sdlhid_free_enumeration(devs)
// }

// func hid_get_device_info(dev *hid_device) *hid_device_info {
//	return sdlhid_get_device_info(dev)
// }

// func hid_get_feature_report(dev *hid_device, data *uint8, length uint64) int32 {
//	return sdlhid_get_feature_report(dev, data, length)
// }

// func hid_get_indexed_string(dev *hid_device, string_index int32, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_indexed_string(dev, string_index, string, maxlen)
// }

// func hid_get_input_report(dev *hid_device, data *uint8, length uint64) int32 {
//	return sdlhid_get_input_report(dev, data, length)
// }

// func hid_get_manufacturer_string(dev *hid_device, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_manufacturer_string(dev, string, maxlen)
// }

// func hid_get_product_string(dev *hid_device, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_product_string(dev, string, maxlen)
// }

// func hid_get_report_descriptor(dev *hid_device, buf *uint8, buf_size uint64) int32 {
//	return sdlhid_get_report_descriptor(dev, buf, buf_size)
// }

// func hid_get_serial_number_string(dev *hid_device, string *wchar_t, maxlen uint64) int32 {
//	return sdlhid_get_serial_number_string(dev, string, maxlen)
// }

// func hid_init() int32 {
//	return sdlhid_init()
// }

// func hid_open(vendor_id uint16, product_id uint16, serial_number *wchar_t) *hid_device {
//	return sdlhid_open(vendor_id, product_id, serial_number)
// }

// func hid_open_path(path string) *hid_device {
//	return sdlhid_open_path(path)
// }

// Available since SDL 3.4.0.
// func hid_get_properties(dev *HidDevice) PropertiesID {
// 	return sdlhid_get_properties(dev)
// }

// func hid_read(dev *hid_device, data *uint8, length uint64) int32 {
//	return sdlhid_read(dev, data, length)
// }

// func hid_read_timeout(dev *hid_device, data *uint8, length uint64, milliseconds int32) int32 {
//	return sdlhid_read_timeout(dev, data, length, milliseconds)
// }

// func hid_send_feature_report(dev *hid_device, data *uint8, length uint64) int32 {
//	return sdlhid_send_feature_report(dev, data, length)
// }

// func hid_set_nonblocking(dev *hid_device, nonblock int32) int32 {
//	return sdlhid_set_nonblocking(dev, nonblock)
// }

// func hid_write(dev *hid_device, data *uint8, length uint64) int32 {
//	return sdlhid_write(dev, data, length)
// }
