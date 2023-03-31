package packages

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Tolyar/upkg/pkg/sysinfo"
)

// Detect package manager for current OS.

func PMName() (string, error) {
	os := sysinfo.Platform()

	switch os {
	case "darwin":
		return "brew", nil
	case "linux":
		return LinuxPMName()
	default:
		return "unknown", fmt.Errorf("Unknown package manager fo %s", os)
	}
}

func LinuxPMName() (string, error) {
	lr, err := sysinfo.LinuxRelease()
	if err != nil {
		return "unknown", err
	}
	id, ok := lr["ID"]
	if !ok {
		return "unknown", fmt.Errorf("Unknown ID in /etc/os-release")
	}
	switch id {
	case "arch", "artix", "manjaro":
		return "pacman", nil
	case "centos", "redhat", "rhel":
		_, err := exec.LookPath("dnf")
		if err != nil {
			return "yum", nil
		} else {
			return "dnf", nil
		}

	case "fedora":
		_, err := exec.LookPath("dnf")
		if err != nil {
			return "yum", nil
		} else {
			return "dnf", nil
		}
	case "alpine":
		return "apk", nil
	case "debian", "ubuntu", "astra", "linuxmint":
		return "apt", nil
	}
	if idLike, ok := lr["ID_LIKE"]; ok {
		if strings.Contains(idLike, "debian") {
			return "apt", nil
		} else if strings.Contains(idLike, "rhel") {
			_, err := exec.LookPath("dnf")
			if err != nil {
				return "yum", nil
			} else {
				return "dnf", nil
			}
		}
	}
	// And last chance. Try to find apt, yum and dnf by path.
	if _, err := exec.LookPath("dnf"); err == nil {
		return "dnf", nil
	}
	if _, err := exec.LookPath("yum"); err == nil {
		return "yum", nil
	}
	if _, err := exec.LookPath("apt"); err == nil {
		return "apt", nil
	}
	if _, err := exec.LookPath("pacman"); err == nil {
		return "pacman", nil
	}
	if _, err := exec.LookPath("apk"); err == nil {
		return "apk", nil
	}
	return "unknown", fmt.Errorf("Unknown package for %s", id)
}
