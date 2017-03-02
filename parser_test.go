// Copyright © 2017 Mario Carrion <info@carrion.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
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
