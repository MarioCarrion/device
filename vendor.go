package device

// Windows User Agent: https://msdn.microsoft.com/en-us/library/hh920767(v=vs.85).aspx

import (
	"regexp"
)

type VendorKind int

const (
	APPLE VendorKind = 1
	MICROSOFT
	LINUX
)

type Vendor struct {
	user_agent string
	kind       VendorKind
}

func NewVendor(user_agent string) *Vendor {
	return &Vendor{user_agent: user_agent}
}

func (v *Vendor) IsWindows() bool {
	return regexp.MustCompile(`Windows`).MatchString(v.user_agent)
}

func (v *Vendor) IsApple() bool {
	return v.IsMacOSX() || v.IsIPhone() || v.IsIPad() || v.IsIPod()
}

func (v *Vendor) IsMacOSX() bool {
	return regexp.MustCompile(`Mac OS X`).MatchString(v.user_agent)
}

func (v *Vendor) IsIPhone() bool {
	return regexp.MustCompile(`iPhone`).MatchString(v.user_agent)
}

func (v *Vendor) IsIPad() bool {
	return regexp.MustCompile(`iPad`).MatchString(v.user_agent)
}

func (v *Vendor) IsIPod() bool {
	return regexp.MustCompile(`iPod`).MatchString(v.user_agent)
}

// Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)
