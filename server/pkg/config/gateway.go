package config

import "os"

type GatewayInfo struct {
	Port      string
	Version   string
	AdminPass string
}

func parseGatewayConf(gatewayConf map[string]interface{}) *GatewayInfo {
	adminPass := os.Getenv("GATEWAY_ADMIN_PASS")
	if adminPass == "" {
		adminPass = gatewayConf["admin_pass"].(string)
	}
	return &GatewayInfo{
		Port:      gatewayConf["port"].(string),
		Version:   gatewayConf["version"].(string),
		AdminPass: adminPass,
	}
}
