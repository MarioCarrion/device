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

import "fmt"

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
