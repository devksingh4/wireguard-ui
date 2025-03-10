package util

import "strings"

// Runtime config
var (
	DisableLogin   bool
	BindAddress    string
	SmtpHostname   string
	SmtpPort       int
	SmtpTLS		   bool
	SmtpUsername   string
	SmtpPassword   string
	SmtpNoTLSCheck bool
	SmtpAuthType   string
	SendgridApiKey string
	EmailFrom      string
	EmailFromName  string
	EmailSubject   string
	EmailContent   string
	SessionSecret  []byte
	WgConfTemplate string
	BasePath       string
)

const (
	DefaultUsername                        = "admin"
	DefaultPassword                        = "admin"
	DefaultServerAddress                   = "10.252.1.0/24"
	DefaultServerPort                      = 51820
	DefaultDNS                             = "1.1.1.1"
	DefaultMTU                             = 1450
	DefaultPersistentKeepalive             = 15
	DefaultForwardMark                     = "0xca6c"
	DefaultConfigFilePath                  = "/etc/wireguard/wg0.conf"
	UsernameEnvVar                         = "WGUI_USERNAME"
	PasswordEnvVar                         = "WGUI_PASSWORD"
	EndpointAddressEnvVar                  = "WGUI_ENDPOINT_ADDRESS"
	DNSEnvVar                              = "WGUI_DNS"
	MTUEnvVar                              = "WGUI_MTU"
	PersistentKeepaliveEnvVar              = "WGUI_PERSISTENT_KEEPALIVE"
	ForwardMarkEnvVar                      = "WGUI_FORWARD_MARK"
	ConfigFilePathEnvVar                   = "WGUI_CONFIG_FILE_PATH"
	ServerAddressesEnvVar                  = "WGUI_SERVER_INTERFACE_ADDRESSES"
	ServerListenPortEnvVar                 = "WGUI_SERVER_LISTEN_PORT"
	ServerPostUpScriptEnvVar               = "WGUI_SERVER_POST_UP_SCRIPT"
	ServerPostDownScriptEnvVar             = "WGUI_SERVER_POST_DOWN_SCRIPT"
	DefaultClientAllowedIpsEnvVar          = "WGUI_DEFAULT_CLIENT_ALLOWED_IPS"
	DefaultClientExtraAllowedIpsEnvVar     = "WGUI_DEFAULT_CLIENT_EXTRA_ALLOWED_IPS"
	DefaultClientUseServerDNSEnvVar        = "WGUI_DEFAULT_CLIENT_USE_SERVER_DNS"
	DefaultClientEnableAfterCreationEnvVar = "WGUI_DEFAULT_CLIENT_ENABLE_AFTER_CREATION"
)

func ParseBasePath(basePath string) string {
	if !strings.HasPrefix(basePath, "/") {
		basePath = "/" + basePath
	}
	if strings.HasSuffix(basePath, "/") {
		basePath = strings.TrimSuffix(basePath, "/")
	}
	return basePath
}
