package models

type ConfigExprt struct {
	Exporter struct {
		Port        string `yaml:"port"`
		Log_Path    string `yaml:"log_path"`
		Prefix      string `yaml:"prefix"`
		Grep_period int    `yaml:"grep_period"`
	} `yaml:"exporter"`
	TemplateErr []string `yaml:"template_err"`
}
