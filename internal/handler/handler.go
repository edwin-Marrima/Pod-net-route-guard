package handler

import (
	"github.com/edwin-Marrima/Pod-net-route-guard/internal/schema"
	"gopkg.in/yaml.v3"
	"os"
)

func readConfiguration(configFilePath string) (*schema.Config, error) {
	cfg, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var config schema.Config
	err = yaml.Unmarshal(cfg, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func Apply() error {
	// read configuration
	return nil
}

func natRuleEngine(config *schema.Config) ([][]string, error) {
	rules := make([][]string, len(config.Rules.NAT))
	for i, rule := range config.Rules.NAT {
		r := []string{}
		if rule.Source != nil {
			if rule.Source.IP != "" {
				r = append(r, "-s", rule.Source.IP)
			}
			if rule.Source.Port != "" {
				r = append(r, "--sport", rule.Source.Port)
			}
			if rule.Source.Protocol != "" {
				r = append(r, "-p", rule.Source.Protocol)
			}
		}
		if rule.Destination != nil {
			if rule.Destination.IP != "" {
				r = append(r, "-d", rule.Destination.IP)
			}
			if rule.Destination.Port != "" {
				r = append(r, "--dport", rule.Destination.Port)
			}
		}
		if rule.Action != nil {
			r = append(r, "-j", "REDIRECT")
			if rule.Action.RedirectTo != nil {
				r = append(r, "--to-ports", rule.Action.RedirectTo.Port)
			}
		}
		rules[i] = r
	}
	return rules, nil
}
