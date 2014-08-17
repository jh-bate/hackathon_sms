package clients

type MockClient struct {
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (mc *MockClient) LoadInto(userid string, data interface{}) error {
	return nil
}

func (mc *MockClient) LoadFrom(userid string) (data interface{}, err error) {
	return data, nil
}
