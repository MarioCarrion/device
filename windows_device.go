package device

import (
	"regexp"
)

type WindowsDevice struct {
	user_agent string
}

func NewWindowsDevice(user_agent string) *WindowsDevice {
	return &WindowsDevice{user_agent: user_agent}
}

func (p *WindowsDevice) DeviceType() DeviceType {
	if regexp.MustCompile(`Xbox`).MatchString(p.user_agent) {
		return CONSOLE
	} else if regexp.MustCompile(`Windows\sPhone`).MatchString(p.user_agent) {
		return PHONE
	} else if regexp.MustCompile(`(ARM|Touch)`).MatchString(p.user_agent) {
		return TABLET
	} else {
		return DESKTOP
	}
}
