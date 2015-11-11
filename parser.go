package device

import "fmt" //

type DeviceType int

const (
  UNKNOWN DeviceType = iota
	DESKTOP
	TABLET
	PHONE
	CONSOLE
)

type Device struct {
	BrowserName    string
	BrowserVersion string
	OSName         string
	OSVersion      string
	Device         string
	DeviceType     DeviceType
	DeviceVersion  string
	DeviceDetails  string
}

func Parse(user_agent string) Device {
	device_type := UNKNOWN

	fmt.Println("value ", device_type)

	platform := NewPlatform(user_agent)
	if platform.IsWindows() {
		device_type = NewWindowsDevice(user_agent).DeviceType()
	}
	fmt.Println("value ", device_type)

	return Device{
		BrowserName:    "chrome",
		BrowserVersion: "34.0.1847.18",
		OSName:         "ios",
		OSVersion:      "7.1.1",
		Device:         "ipad",
		DeviceType:     device_type,
		DeviceVersion:  "",
		DeviceDetails:  ""}
}
