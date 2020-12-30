package providers

import (
	"fmt"
	"github.com/zendesk/goship/color"
	"github.com/zendesk/goship/config"
	"github.com/zendesk/goship/resources"
	"os"
	"os/exec"
	"strings"
)

// Provider defines interface which should be implemented by every provider
type Provider interface {
	Init() error
	Name() string
	GetResources() (resources.ResourceList, error)
}

func RunPreinitHook(hook string) error {
	if config.GlobalConfig.Verbose {
		fmt.Printf("Executing preinit hook: %s\n", hook)
	}
	command := strings.Split(hook, " ")
	cmd := exec.Command(
		command[0], command[1:]...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		color.PrintRed(fmt.Sprintf("Error while running the preinit `%s` hook: %s\n", hook, err.Error()))
	}
	return err
}
// InitProvidersFromConfig returns list of initialized providers for particular types
func InitProvidersFromConfig(providersConfig map[string]interface{}) (p []Provider, err error) {
	if config.GlobalConfig.PreinitHook != "" {
		_ = RunPreinitHook(config.GlobalConfig.PreinitHook)
	}
	for cloudVendor, cloudProvidersConfig := range providersConfig {
		switch cloudVendor {
		case "aws":
			for providerName, providerCfg := range cloudProvidersConfig.(map[string]interface{}) {
				switch providerName {
				case "ec2":
					for _, providerItem := range providerCfg.([]interface{}) {
						providersList, _ := InitAwsEc2ProvidersFromCfg(providerItem.(map[interface{}]interface{}))
						for _, provider := range providersList {
							p = append(p, provider)
						}

					}
				}
			}
		}
	}
	return p, err
}
