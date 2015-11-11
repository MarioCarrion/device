package device

import "testing"

func TestParse(t *testing.T) {
	user_agent := "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko"

	device := Parse(user_agent)

	if device.BrowserName != "chrome" {
		t.Errorf("TestParse(BrowserName) returned: %q", device.BrowserName)
	}

	if device.BrowserVersion != "34.0.1847.18" {
		t.Errorf("TestParse(BrowserVersion) returned: %q", device.BrowserVersion)
	}

	if device.OSName != "ios" {
		t.Errorf("TestParse(OSName) returned: %q", device.OSName)
	}

	if device.OSVersion != "7.1.1" {
		t.Errorf("TestParse(OSVersion) returned: %q", device.OSVersion)
	}

	if device.Device != "ipad" {
		t.Errorf("TestParse(Device) returned: %q", device.Device)
	}

	if device.DeviceType != DESKTOP {
		t.Errorf("TestParse(DeviceType) returned: %i", device.DeviceType)
	}

	if device.DeviceVersion != "" {
		t.Errorf("TestParse(DeviceVersion) returned: %q", device.DeviceVersion)
	}

	if device.DeviceDetails != "" {
		t.Errorf("TestParse(DeviceDetails) returned: %q", device.DeviceDetails)
	}
}
