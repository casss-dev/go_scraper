package playwright

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/casss-dev/go_scraper/pkg/util"
	"github.com/playwright-community/playwright-go"
)

const (
	localWsEndpoint = "ws://localhost:9222/ws"
)

type Options struct {
	RunOptions        *playwright.RunOptions
	WebsocketEndpoint string
	// the executable path of the browser
	ExecutablePath *string
	// if a headless browser should be used (default: `true`)
	Headless bool
	SlowMo   *float64
	Install  bool
}

func defaultBrowserOpts() Options {
	return Options{
		WebsocketEndpoint: localWsEndpoint,
		Headless:          true,
	}
}

// installs dependencies automatically
func WithInstall() OptsFunc {
	return func(o *Options) {
		o.Install = true
	}
}

func WithHeadless(headless bool) OptsFunc {
	return func(cbo *Options) {
		cbo.Headless = headless
	}
}

func WithExecPath(path string) OptsFunc {
	return func(cbo *Options) {
		cbo.ExecutablePath = &path
	}
}

func WithCamoufoxPath() OptsFunc {
	return func(cbo *Options) {
		var path string
		opsys := runtime.GOOS
		switch opsys {
		case "darwin":
			path = macOSCamoufoxPath()
			cbo.ExecutablePath = &path
		default:
			const run = "python -m camoufox path"
			panic(fmt.Sprintf("exec path for '%s' is not implemented. Run '%s' to get the path", opsys, run))
		}
	}
}

func WithSlowMo(slowDown float64) OptsFunc {
	return func(cbo *Options) {
		cbo.SlowMo = &slowDown
	}
}

type OptsFunc func(*Options)

func macOSCamoufoxPath() string {
	home := util.Must(os.UserHomeDir())
	return filepath.Join(
		home,
		"Library",
		"Caches",
		"camoufox",
		"Camoufox.app",
		"Contents",
		"MacOS",
		"camoufox",
	)
}
