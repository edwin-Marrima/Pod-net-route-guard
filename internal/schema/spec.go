package schema

type Config struct {
	Rules Rules `json:"rules"`
}

type Rules struct {
	NAT []NAT `json:"nat"`
}
type NAT struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Source      *Source      `yaml:"source"`
	Destination *Destination `yaml:"destination"`
	Action      *Action      `yaml:"action"`
}
type Source struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Protocol string `yaml:"protocol"`
}
type Destination struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}
type Action struct {
	RedirectTo *RedirectTo `yaml:"redirect_to"`
}
type RedirectTo struct {
	Port string `yaml:"port"`
}
