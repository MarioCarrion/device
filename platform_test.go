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

const TEST_PLATFORM_MAC_OSX = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.80 Safari/537.36"
const TEST_PLATFORM_IPHONE = "Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1A543a Safari/419.3"
const TEST_PLATFORM_IPAD = "Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) version/4.0.4 Mobile/7B367 Safari/531.21.10"
const TEST_PLATFORM_IPOD = "Mozilla/5.0 (iPod; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/3A101a Safari/419.3"

func TestIsWindows(t *testing.T) {
	platform1 := NewPlatform("Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)")
	if platform1.IsWindows() != true {
		t.Errorf("TestIsWindows failed #1")
	}

	platform2 := NewPlatform("YOLO")
	if platform2.IsWindows() == true {
		t.Errorf("TestIsWindows failed #2")
	}
}

func TestIsMac(t *testing.T) {
	if NewPlatform(TEST_PLATFORM_MAC_OSX).IsMac() != true {
		t.Errorf("TestIsMac failed #1")
	}

	if NewPlatform(TEST_PLATFORM_IPHONE).IsMac() != true {
		t.Errorf("TestIsMac failed #2")
	}

	if NewPlatform(TEST_PLATFORM_IPAD).IsMac() != true {
		t.Errorf("TestIsMac failed #3")
	}

	if NewPlatform(TEST_PLATFORM_IPOD).IsMac() != true {
		t.Errorf("TestIsMac failed #4")
	}

	if NewPlatform("YOLO").IsMac() == true {
		t.Errorf("TestIsMac failed #5")
	}
}

func TestIsMacOSX(t *testing.T) {
	platform1 := NewPlatform(TEST_PLATFORM_MAC_OSX)
	if platform1.IsMacOSX() != true {
		t.Errorf("TestIsMacOSX failed #1")
	}

	platform2 := NewPlatform("YOLO")
	if platform2.IsMacOSX() == true {
		t.Errorf("TestIsMacOSX failed #2")
	}
}

func TestIsIPhone(t *testing.T) {
	platform1 := NewPlatform(TEST_PLATFORM_IPHONE)
	if platform1.IsIPhone() != true {
		t.Errorf("TestIsIPhone failed #1")
	}

	platform2 := NewPlatform("YOLO")
	if platform2.IsIPhone() == true {
		t.Errorf("TestIsIPhone failed #2")
	}
}

func TestIsIPad(t *testing.T) {
	if NewPlatform(TEST_PLATFORM_IPAD).IsIPad() != true {
		t.Errorf("TestIsIPad failed #1")
	}

	if NewPlatform("YOLO").IsIPad() == true {
		t.Errorf("TestIsIPad failed #2")
	}
}

func TestIsIPod(t *testing.T) {
	if NewPlatform(TEST_PLATFORM_IPOD).IsIPod() != true {
		t.Errorf("TestIsIPod failed #1")
	}

	if NewPlatform("YOLO").IsIPod() == true {
		t.Errorf("TestIsIPod failed #2")
	}
}

func TestIsLinux(t *testing.T) {
	user_agent := "Mozilla/5.0 (X11; Linux i586; rv:31.0) Gecko/20100101 Firefox/31.0"

	platform1 := NewPlatform(user_agent)
	if platform1.IsLinux() != true {
		t.Errorf("TestIsLinux failed #1")
	}

	platform2 := NewPlatform("YOLO")
	if platform2.IsLinux() == true {
		t.Errorf("TestIsLinux failed #2")
	}
}
