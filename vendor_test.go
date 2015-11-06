package device

import "testing"

func TestIsWindows(t *testing.T) {
	vendor1 := NewVendor("Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)")
	if vendor1.IsWindows() != true {
		t.Errorf("TestIsWindows failed #1")
	}

	vendor2 := NewVendor("YOLO")
	if vendor2.IsWindows() == true {
		t.Errorf("TestIsWindows failed #2")
	}
}

func TestIsMacOSX(t *testing.T) {
	user_agent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.80 Safari/537.36"

	vendor1 := NewVendor(user_agent)
	if vendor1.IsMacOSX() != true {
		t.Errorf("TestIsMacOSX failed #1")
	}

	vendor2 := NewVendor("YOLO")
	if vendor2.IsMacOSX() == true {
		t.Errorf("TestIsMacOSX failed #2")
	}
}

func TestIsIPhone(t *testing.T) {
	user_agent := "Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1A543a Safari/419.3"

	vendor1 := NewVendor(user_agent)
	if vendor1.IsIPhone() != true {
		t.Errorf("TestIsIPhone failed #1")
	}

	vendor2 := NewVendor("YOLO")
	if vendor2.IsIPhone() == true {
		t.Errorf("TestIsIPhone failed #2")
	}
}

func TestIsIPad(t *testing.T) {
	user_agent := "Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) version/4.0.4 Mobile/7B367 Safari/531.21.10"

	vendor1 := NewVendor(user_agent)
	if vendor1.IsIPad() != true {
		t.Errorf("TestIsIPad failed #1")
	}

	vendor2 := NewVendor("YOLO")
	if vendor2.IsIPad() == true {
		t.Errorf("TestIsIPad failed #2")
	}
}

func TestIsIPod(t *testing.T) {
	user_agent := "Mozilla/5.0 (iPod; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/3A101a Safari/419.3"

	vendor1 := NewVendor(user_agent)
	if vendor1.IsIPod() != true {
		t.Errorf("TestIsIPod failed #1")
	}

	vendor2 := NewVendor("YOLO")
	if vendor2.IsIPod() == true {
		t.Errorf("TestIsIPod failed #2")
	}
}
