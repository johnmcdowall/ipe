// Copyright 2014 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"strings"
)

// The config file
type ConfigFile struct {
	Host     string // The host, eg: :8080 will start on 0.0.0.0:8080
	Expvar   bool
	User     string
	Password string
	Apps     []*App
}

// Error for App not found
var AppNotFoundError = errors.New("App not found")

// Initialize Apps
func (c *ConfigFile) Init() {
	for _, app := range c.Apps {
		app.Init()
	}
}

func (c *ConfigFile) WasProvidedUserAndPassword() bool {
	return len(strings.TrimSpace(c.User)) > 0 && len(strings.TrimSpace(c.Password)) > 0
}

// Returns an App with by appID
func (c *ConfigFile) GetAppByAppID(appID string) (*App, error) {
	for _, app := range c.Apps {
		if app.AppID == appID {
			return app, nil
		}
	}
	return &App{}, AppNotFoundError
}

// Returns an App with by key
func (c *ConfigFile) GetAppByKey(key string) (*App, error) {
	for _, app := range c.Apps {
		if app.Key == key {
			return app, nil
		}
	}
	return &App{}, AppNotFoundError
}
