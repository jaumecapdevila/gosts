package config

// Loader defines the contract to read and load the app configuration
type Loader interface {
	Load(string) Parameters
}
