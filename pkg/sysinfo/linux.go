package sysinfo

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Detect linux distribution.

const (
	MainPath     = "/etc/os-release"
	FallbackPath = "/usr/lib/os-release"
)

var re = regexp.MustCompile(`^([\w_]+)=(.+)$`)

func LinuxRelease() (map[string]string, error) {
	var f *os.File
	var err error

	f, err = os.Open(MainPath)
	if err != nil {
		// Try to use fallback.
		f, err = os.Open(FallbackPath)
		if err != nil {
			return nil, err
		}
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := make(map[string]string)
	for s.Scan() {
		if r := re.FindStringSubmatch(s.Text()); r != nil {
			if len(r) == 3 {
				key := r[1]
				val := strings.Trim(r[2], `"`)
				m[key] = val
			}
		}
	}

	return m, nil
}
