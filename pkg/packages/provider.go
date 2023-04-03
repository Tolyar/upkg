package packages

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// Describe package managers.

// / Set of arguments for package manager.
type Provider struct {
	bin           string   // Binary name.
	info          []string // Show info about package.
	install       []string // Install package.
	yes           []string // Answer YES to manager.
	remove        []string // Remove package.
	upgrade       []string // Upgrade package.
	upgradeAll    []string // Upgrade all packages.
	search        []string // Search package.
	updateIndex   []string // Update index DB.
	listUpdates   []string // List all upgradable packages.
	listInstalled []string // List all installed packages.
	provides      []string // Find which package provides resource.
}

// Describe one package for lists.
type Package struct {
	Name     string
	Arch     string
	Version  string
	Repo     string
	Provider string
}

var Providers map[string]Provider

func init() {
	p := make(map[string]Provider)

	// yum.
	p["yum"] = Provider{
		bin:           "yum",
		info:          []string{"info"},
		install:       []string{"install"},
		yes:           []string{"-y"},
		remove:        []string{"remove"},
		upgrade:       []string{"update"},
		upgradeAll:    []string{"update"},
		search:        []string{"search"},
		updateIndex:   []string{"makecache"},
		listUpdates:   []string{"list", "updates"},
		listInstalled: []string{"list", "installed"},
		provides:      []string{"provides"},
	}

	// dnf.
	p["dnf"] = Provider{
		bin:           "dnf",
		info:          []string{"info"},
		install:       []string{"install"},
		yes:           []string{"-y"},
		remove:        []string{"remove"},
		upgrade:       []string{"update"},
		upgradeAll:    []string{"update"},
		search:        []string{"search"},
		updateIndex:   []string{"makecache"},
		listUpdates:   []string{"list", "updates"},
		listInstalled: []string{"list", "installed"},
		provides:      []string{"provides"},
	}

	// For cut&paste.
	p["DEMO"] = Provider{
		bin:           "",
		info:          []string{""},
		install:       []string{""},
		yes:           []string{""},
		remove:        []string{""},
		upgrade:       []string{""},
		upgradeAll:    []string{""},
		search:        []string{""},
		updateIndex:   []string{""},
		listUpdates:   []string{""},
		listInstalled: []string{""},
		provides:      []string{""},
	}

	// brew.
	p["brew"] = Provider{
		bin:           "brew",
		info:          []string{"info"},
		install:       []string{"install"},
		yes:           []string{""}, // Not provided.
		remove:        []string{"uninstall"},
		upgrade:       []string{"upgrade"},
		upgradeAll:    []string{"upgrade"},
		search:        []string{"search"},
		updateIndex:   []string{"update"},
		listUpdates:   []string{"outdated"},
		listInstalled: []string{"list"},
		provides:      []string{""}, // Not provided.
	}

	Providers = p
}

func GetProvider() (*Provider, error) {
	pname, err := PMName()
	cobra.CheckErr(err)
	p, ok := Providers[pname]
	if !ok {
		return nil, fmt.Errorf("Provider %s does not supported", pname)
	}

	return &p, nil
}

// Provider methods.

func (p *Provider) run(command []string, args []string) {
	if len(command) == 0 || command[0] == "" {
		log.Fatalf("Command %s does not implemented for provider %s", command, p.bin)
	}
	args = append(command, args...)
	//nolint:gosec
	cmd := exec.Command(p.bin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func (p *Provider) runWithResult(command []string, args []string) []string {
	stdout := make([]string, 0)
	if len(command) == 0 || command[0] == "" {
		log.Fatalf("Command %s does not implemented for provider %s", command, p.bin)
	}
	args = append(command, args...)
	//nolint:gosec
	cmd := exec.Command(p.bin, args...)
	cmd.Stderr = os.Stderr
	s, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(s)
	for sc.Scan() {
		stdout = append(stdout, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return stdout
}

// SHow info.
func (p *Provider) Info(args ...string) {
	p.run(p.info, args)
}

// Install package.
func (p *Provider) Install(args ...string) {
	p.run(p.install, args)
}

// List installed packages.
func (p *Provider) ListInstalled(args ...string) {
	p.run(p.listInstalled, args)
}

func (p *Provider) GetInstalled(args ...string) ([]Package, error) {
	return p.ParseList(p.runWithResult(p.listInstalled, args))
}

// List upgradable packages.
func (p *Provider) ListUpdates(args ...string) {
	p.run(p.listUpdates, args)
}

func (p *Provider) GetUpdates(args ...string) ([]Package, error) {
	return p.ParseList(p.runWithResult(p.listUpdates, args))
}

// Remove package.
func (p *Provider) Remove(args ...string) {
	p.run(p.remove, args)
}

// Search package.
func (p *Provider) Search(args ...string) {
	p.run(p.search, args)
}

// Which packages resource.
func (p *Provider) Provides(args ...string) {
	p.run(p.provides, args)
}

// Update packages.
func (p *Provider) UpdateIndex(args ...string) {
	p.run(p.updateIndex, args)
}

// Upgrade packages.
func (p *Provider) Upgrade(args ...string) {
	p.run(p.upgrade, args)
}

// Upgrade all packages.
func (p *Provider) UpgradeAll(args ...string) {
	p.run(p.upgradeAll, args)
}

// Parse package list.
func (p *Provider) ParseList(list []string) ([]Package, error) {
	switch p.bin {
	case "yum", "dnf":
		return ParseListYUM(list, p.bin), nil
	}

	return nil, fmt.Errorf("can't parse packages list for %s", p.bin)
}

// Return name of detected providers.
func (p *Provider) Name() string {
	return p.bin
}

// ----- provider specific parsers ----

// "^<package>.<arch>\s+<version>\s+<repo>$".
func ParseListYUM(list []string, provider string) []Package {
	pkgs := make([]Package, 0)
	for _, s := range list {
		parsed := strings.Split(s, " ")
		// Skip crap.
		if len(parsed) != 3 {
			continue
		}
		na := strings.Split(parsed[0], ".")
		if len(na) != 2 {
			continue
		}

		p := Package{
			Name:     na[0],
			Arch:     na[1],
			Provider: provider, // yum or dnf.
			Version:  parsed[1],
			Repo:     parsed[2],
		}
		pkgs = append(pkgs, p)
	}

	return pkgs
}
