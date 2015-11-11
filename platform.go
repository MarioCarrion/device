package device

// Windows User Agent: https://msdn.microsoft.com/en-us/library/hh920767(v=vs.85).aspx

import (
	"regexp"
)

type PlatformKind int

const (
	APPLE PlatformKind = 1
	MICROSOFT
	LINUX
)

type Platform struct {
	user_agent string
}

func NewPlatform(user_agent string) *Platform {
	return &Platform{user_agent: user_agent}
}

func (p *Platform) IsLinux() bool {
	return regexp.MustCompile(`Linux`).MatchString(p.user_agent)
}

func (p *Platform) IsWindows() bool {
	return regexp.MustCompile(`Windows`).MatchString(p.user_agent)
}

func (p *Platform) IsMac() bool {
	return p.IsMacOSX() || p.IsIPhone() || p.IsIPad() || p.IsIPod()
}

func (p *Platform) IsMacOSX() bool {
	return regexp.MustCompile(`Mac OS X`).MatchString(p.user_agent)
}

func (p *Platform) IsIPhone() bool {
	return regexp.MustCompile(`iPhone`).MatchString(p.user_agent)
}

func (p *Platform) IsIPad() bool {
	return regexp.MustCompile(`iPad`).MatchString(p.user_agent)
}

func (p *Platform) IsIPod() bool {
	return regexp.MustCompile(`iPod`).MatchString(p.user_agent)
}
