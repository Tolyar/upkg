package sysinfo

import "runtime"

// Detect OS.
// Look to go tool dist list, on full OS and platfrom versions.

func Platform() string {
	// linux, darwin(mac os), windows, freebsd, ...
	return runtime.GOOS
}

func Arch() string {
	// one of 386, amd64, arm, s390x, etc.
	return runtime.GOARCH
}
