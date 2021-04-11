package fileconf

// Compose keeps configuration for docker-compose.
type Compose struct {
	Services map[string]ComposeService
}

// ComposeService keeps configuration for the single service.
type ComposeService struct {
	Repo    string       `yaml:"remote"`  // remote repository
	Image   string       `yaml:"image"`   // use docker image instead
	Ports   []string     `yaml:"ports"`   // bind network ports for docker
	Command string       `yaml:"command"` // path to a command with optional args for run a service
	Build   ComposeBuild `yaml:"build"`   // compatible with docker-compose
}

// ComposeBuild keeps configuration for the build recipes.
type ComposeBuild struct {
	Context    string `yaml:"context"`    // path to build context
	Dockerfile string `yaml:"dockerfile"` // path to alternate dockerfile, used when build command empty
}
