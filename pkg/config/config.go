package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	InProduction  bool
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
}
