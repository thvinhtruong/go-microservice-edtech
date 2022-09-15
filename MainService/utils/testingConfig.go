package utils

type TestingConfig struct {
}

func (t TestingConfig) getMainServiceServerHost() (string, error) {
	panic("implement me")
}

func (t TestingConfig) getMainServiceServerPort() (int, error) {
	panic("implement me")
}

func (t TestingConfig) getUserServiceServerHost() (string, error) {
	panic("implement me")
}

func (t TestingConfig) getUserServiceServerPort() (int, error) {
	panic("implement me")
}

func (t TestingConfig) getHMACKey() (string, error) {
	return "BaoDepTrai", nil
}
