package validator

import (
	"testing"
)

func TestIsDNSName(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"localhost", true},
		{"a.bc", true},
		{"a.b.", true},
		{"a.b..", false},
		{"localhost.local", true},
		{"localhost.localdomain.intern", true},
		{"l.local.intern", true},
		{"ru.link.n.svpncloud.com", true},
		{"-localhost", false},
		{"localhost.-localdomain", false},
		{"localhost.localdomain.-int", false},
		{"_localhost", true},
		{"localhost._localdomain", true},
		{"localhost.localdomain._int", true},
		{"lÖcalhost", false},
		{"localhost.lÖcaldomain", false},
		{"localhost.localdomain.üntern", false},
		{"__", true},
		{"localhost/", false},
		{"127.0.0.1", false},
		{"[::1]", false},
		{"50.50.50.50", false},
		{"localhost.localdomain.intern:65535", false},
		{"漢字汉字", false},
		{"www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de", false},
	}

	for _, test := range tests {
		actual := IsDNSName(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsDNS(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsHost(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		param    string
		expected bool
	}{
		{"localhost", true},
		{"localhost.localdomain", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"::1", true},
		{"play.golang.org", true},
		{"localhost.localdomain.intern:65535", false},
		{"-[::1]", false},
		{"-localhost", false},
		{".localhost", false},
	}
	for _, test := range tests {
		actual := IsHost(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsHost(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsIP(t *testing.T) {
	t.Parallel()

	// Without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsIP(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIP(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", false},
		{"2001:db8:0000:1:1:1:1:1", false},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsIPv4(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIPv4(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv6
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"1.2.3.4", false},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsIPv6(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIPv6(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsPort(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"1", true},
		{"65535", true},
		{"0", false},
		{"65536", false},
		{"65538", false},
	}

	for _, test := range tests {
		actual := IsPort(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsPort(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
