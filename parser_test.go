package device

import "testing"

func TestParse(t *testing.T) {
	user_agent := "Mozilla/5.0 (iPad; CPU OS 7_1_1 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) CriOS/34.0.1847.18 Mobile/11D201 Safari/9537.53 (000655)"

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

	if device.DeviceType != TABLET {
		t.Errorf("TestParse(DeviceType) returned: %q", device.DeviceType)
	}

	if device.DeviceVersion != "" {
		t.Errorf("TestParse(DeviceVersion) returned: %q", device.DeviceVersion)
	}

	if device.DeviceDetails != "" {
		t.Errorf("TestParse(DeviceDetails) returned: %q", device.DeviceDetails)
	}
}
