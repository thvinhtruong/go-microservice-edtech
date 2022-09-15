package utils

type ConfigInterface interface {
	getMainServiceServerHost() (string, error)
	getMainServiceServerPort() (int, error)
	getUserServiceServerHost() (string, error)
	getUserServiceServerPort() (int, error)
	getHMACKey() (string, error)
}
