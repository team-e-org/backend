package authz

type AuthLayerMock struct {
	ExpectedToken          string
	ExpectedTokenError     error
	ExpectedTokenData      *TokenData
	ExpectedTokenDataError error
}

func (m *AuthLayerMock) AuthenticateUser(email string, password string) (string, error) {
	return m.ExpectedToken, m.ExpectedTokenError
}

func (m *AuthLayerMock) GetTokenData(token string) (*TokenData, error) {
	return m.ExpectedTokenData, m.ExpectedTokenDataError
}
