package vault

import "strings"

// FakeVault is a fake implementation of the Vault
type FakeVault struct {
	Requests []string
}

var fakeResponseVaultAwsAzure = map[string]interface{}{"test": "test1"}
var vaultCoNfigData = map[string]interface{}{"test": "test1"}
var fakeResponseVaultConfig = map[string]interface{}{"data": vaultCoNfigData}

func (f *FakeVault) Read(path string, data map[string][]string) (map[string]interface{}, error) {
	f.Requests = append(f.Requests, path)

	if strings.Contains(path, "aws/") || strings.Contains(path, "azure/") || strings.Contains(path, "azuread/") {
		return fakeResponseVaultAwsAzure, nil
	}

	return fakeResponseVaultConfig, nil
}

func (f *FakeVault) Write(path string, data map[string]interface{}) (map[string]interface{}, error) {
	f.Requests = append(f.Requests, path)

	return fakeResponseVaultAwsAzure, nil
}

func (f *FakeVault) GetToken() string {
	return "fake"
}

func (f *FakeVault) GetAddress() string {
	return "fake"
}
