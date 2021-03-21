package config

type Extended struct {
	Services map[string]Service
}

type Service struct {
	ComposeService
	Base string `yaml:"base"` // base path to the service
}
type Build struct {
	ComposeBuild
	Cmd string `yaml:"command"` // path to build command that run in context
}
