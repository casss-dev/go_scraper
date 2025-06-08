package playwright

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// creates the playwright instance and optionally installs dependencies
func CreatePlaywright(opts ...OptsFunc) (*playwright.Playwright, error) {
	o := defaultBrowserOpts()
	for _, opt := range opts {
		opt(&o)
	}
	if o.Install {
		if err := playwright.Install(); err != nil {
			return nil, fmt.Errorf("error installing playwright dependencies: %w", err)
		}
	}
	var (
		pw  *playwright.Playwright
		err error
	)
	if o.RunOptions != nil {
		pw, err = playwright.Run(o.RunOptions)
	} else {
		pw, err = playwright.Run()
	}
	if err != nil {
		return nil, fmt.Errorf("error running playwright: %w", err)
	}
	return pw, err
}

// creates the playwright browser instance
func CreateBrowser(pw *playwright.Playwright, opts ...OptsFunc) (playwright.Browser, error) {
	o := defaultBrowserOpts()
	for _, opt := range opts {
		opt(&o)
	}

	if o.ExecutablePath == nil {
		return pw.Firefox.Connect(o.WebsocketEndpoint, playwright.BrowserTypeConnectOptions{
			SlowMo: o.SlowMo,
		})
	}
	return pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		ExecutablePath: o.ExecutablePath,
		Headless:       &o.Headless,
		Args:           []string{},
		SlowMo:         o.SlowMo,
	})
}
