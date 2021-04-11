package config

// Extended keeps full configuration for MRT that includes
// fields for the docker-compose.
type Extended struct {
	Services map[string]Service
}

// Service keeps configuration for the single service. It based on
// docker-compose Service.
type Service struct {
	ComposeService
	Base string `yaml:"base"` // base path to the service
}

// Build keeps configuration for build recipes. It based on
// docker-compose Build.
type Build struct {
	ComposeBuild
	Cmd string `yaml:"command"` // path to build command that run in context
}
