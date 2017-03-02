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
	userAgent string
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
