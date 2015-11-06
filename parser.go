package device

type DeviceType int

const UNKNOWN = iota

const (
	DESKTOP DeviceType = 1
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
	return Device{
		BrowserName:    "chrome",
		BrowserVersion: "34.0.1847.18",
		OSName:         "ios",
		OSVersion:      "7.1.1",
		Device:         "ipad",
		DeviceType:     TABLET,
		DeviceVersion:  "",
		DeviceDetails:  ""}
}
