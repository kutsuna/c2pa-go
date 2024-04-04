package c2pa

import (
	"testing"
	"encoding/json"
	"io/ioutil"
)

func errOrLog(t *testing.T, v interface{}, e error) {
	if e != nil {
		t.Errorf("error: %v", e)
	} else {
		t.Logf("%v", v)
	}
}

func Test(t *testing.T) {
	errOrLog(t, "test started.", nil)

	manifest := Manifest{
		ClaimGenerator: "c2pa-c_test/0.1",
		ClaimGeneratorInfo: []GeneratorInfo{
			{
				Name:    "c2pa-c test",
				Version: "0.1",
			},
		},
		Assertions: []Assertion{
			{
				Label: "c2pa.training-mining",
				Data: AssertionData{
					Entries: map[string]Entry{
						"c2pa.ai_generative_training": {Use: "notAllowed"},
						"c2pa.ai_inference":           {Use: "notAllowed"},
						"c2pa.ai_training":            {Use: "notAllowed"},
						"c2pa.data_mining":            {Use: "notAllowed"},
					},
				},
			},
		},
	}
	manifestString, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		errOrLog(t, nil, err)
	}

	publicKey, err := ioutil.ReadFile("public_key.pem")
    if err != nil {
		errOrLog(t, nil, err)
    }
    publicKeyString := string(publicKey)
	
	privateKey, err := ioutil.ReadFile("private_key.pem")
    if err != nil {
		errOrLog(t, nil, err)
    }
    privateKeyString := string(privateKey)
	
	signerInfo := SignerInfo{
		Alg: "ps256",
		SignCert: publicKeyString,
		PrivateKey: privateKeyString,
		TaURL: "http://timestamp.digicert.com",
	}

	returnValue := SignFile("./test.jpg", "signed.jpg", string(manifestString), signerInfo, "./data")
	errOrLog(t, returnValue, nil)
	
	jsonStore := ReadFile("./signed.jpg", "./data/")
	errOrLog(t, jsonStore, nil)
}
