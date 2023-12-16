package vendorgateway

// Config is required to create new client for communicating with Claude
type Config struct {
	ApiKey       string // API Keys
	DefaultModel string // Choose the default AI model
}
