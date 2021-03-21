package config

type Configuration struct {
	Services map[string]Service
}
type Service struct {
	Base    string   `yaml:"base"`    // base path to the service
	Repo    string   `yaml:"remote"`  // remote repository
	Image   string   `yaml:"image"`   // use docker image instead (compatible with docker-compose)
	Ports   []string `yaml:"ports"`   // bind network ports for docker (compatible with docker-compose)
	Command string   `yaml:"command"` // path to a command with optional args for run a service (compatible with docker-compose)
	Build   Build    `yaml:"build"`   // compatible with docker-compose
}
type Build struct {
	Context    string `yaml:"context"`    // path to build context (compatible with docker-compose)
	Cmd        string `yaml:"command"`    // path to build command that run in context
	Dockerfile string `yaml:"dockerfile"` // path to alternate dockerfile, used when build command empty (compatible with docker-compose)
}
