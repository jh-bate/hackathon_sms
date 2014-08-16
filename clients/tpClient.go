package clients

type PlatformClient interface {
	LoadInto(data interface{}) error
	LoadFrom(userid string) (interface{}, error)
}

type TidepoolClient struct {
}

func NewPlatformClient() *TidepoolClient {
	return &TidepoolClient{}
}

func (tc *TidepoolClient) LoadInto(data interface{}) error {
	return nil
}

func (to *TidepoolClient) LoadFrom(userid string) (data interface{}, err error) {
	return data, nil
}
