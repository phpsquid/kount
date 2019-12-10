package settings

type Settings struct {
	merchantID string
	risURL     string
	apiKey     string
	configKey  string
}

func (s *Settings) GetMerchantID() string {
	return s.merchantID
}

func (s *Settings) GetAPIKey() string {
	return s.apiKey
}

func (s *Settings) GetRISURL() string {
	return s.risURL
}

func (s *Settings) GetConfigKey() string {
	return s.configKey
}

func New(merchantID, risURL, apiKey, configKey string) *Settings {
	return &Settings{
		merchantID: merchantID,
		risURL:     risURL,
		apiKey:     apiKey,
		configKey:  configKey,
	}
}
