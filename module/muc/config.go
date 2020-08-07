/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package muc

import (
	"github.com/pkg/errors"
)

type Config struct {
	MucHost string
}

type configProxy struct {
	MucHost string `yaml:"service"`
}

// UnmarshalYAML satisfies Unmarshaler interface.
func (cfg *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	p := configProxy{}
	if err := unmarshal(&p); err != nil {
		return err
	}
	// NOTE here should be a check if MucHost is a localhost.Right now this is
	// done in New() because the only way to access hosts here is to make it
	// a global variable, and doing it in New() seems cleaner for now
	cfg.MucHost = p.MucHost
	if len(cfg.MucHost) == 0 {
		return errors.New("muc: must specify a service hostname")
	}
	return nil
}