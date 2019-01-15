package model

import (
	"errors"
	"regexp"
)

type Maintainer struct {
	Name  string
	Email string
}

type App struct {
	Title       string
	Version     string
	Maintainers []Maintainer
	Company     string
	Website     string
	Source      string
	License     string
	Description string
}

func (c *App) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		Title       string       `yalm:"title"`
		Version     string       `yalm:"version"`
		Maintainers []Maintainer `yaml:"maintainers"`
		Company     string       `yaml:"company"`
		Website     string       `yaml:"website"`
		Source      string       `yaml:"source"`
		License     string       `yaml:"license"`
		Description string       `yaml:"description"`
	}

	if err := unmarshal(&tmp); err != nil {
		return err
	}
	if tmp.Title == "" {
		return errors.New("payload: missing `title`")
	}
	if tmp.Version == "" {
		return errors.New("payload: missing `version`")
	}
	if tmp.Maintainers == nil {
		return errors.New("payload: missing `maintainers`")
	}
	if tmp.Company == "" {
		return errors.New("payload: missing `company`")
	}
	if tmp.Website == "" {
		return errors.New("payload: missing `website`")
	}
	if tmp.Source == "" {
		return errors.New("payload: missing `source`")
	}
	if tmp.License == "" {
		return errors.New("payload: missing `license`")
	}
	if tmp.Description == "" {
		return errors.New("payload: missing `description`")
	}
	for _, m := range tmp.Maintainers {
		if m.Name == "" {
			return errors.New("payload: missing `name`")
		}
		if !ValidateEmail(m.Email) {
			return errors.New("payload: invalid `Email`")
		}
	}

	c.Title = tmp.Title
	c.Version = tmp.Version
	c.Maintainers = tmp.Maintainers
	c.Company = tmp.Company
	c.Website = tmp.Website
	c.Source = tmp.Source
	c.License = tmp.License
	c.Description = tmp.Description

	return nil
}

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
