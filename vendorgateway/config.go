package vendorgateway

// Config is required to create new client for communicating with Claude
type Config struct {
	ApiKey       string // API Keys
	SessionId    string
	DefaultModel string // Choose the default AI model
}
