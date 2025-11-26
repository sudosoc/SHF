package core

type Module struct {
    ID         string   `yaml:"id"`
    Name       string   `yaml:"name"`
    Type       string   `yaml:"type"`
    Path       string   `yaml:"path"`
    Aliases    []string `yaml:"aliases"`
    Tags       []string `yaml:"tags"`
    PrimaryArg string   `yaml:"primary_arg"`
    HelpFile   string   `yaml:"help_file"`
    BaseDir    string   `yaml:"-"`
}
