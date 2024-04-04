# c2pa-go
This library implements c2pa-c wrapper for go.

> [!NOTE]
> This library is for personal research at this momoment.
>
> Supported functions are limited.

## Preparation
Please build c2pa-c and set some environment variable to find required libs.

https://github.com/contentauth/c2pa-c

## Usage

### Prepare manifest data
```
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
            Data: ...
        },
    },
}
manifestString, err := json.MarshalIndent(manifest, "", "  ")
```

### Prepare SignerInfo with pair key
```
publicKey, err := ioutil.ReadFile("public_key.pem")
if err != nil {
    ...
}
publicKeyString := string(publicKey)

privateKey, err := ioutil.ReadFile("private_key.pem")
if err != nil {
    ...
}
privateKeyString := string(privateKey)

signerInfo := SignerInfo{
    Alg: "ps256",
    SignCert: publicKeyString,
    PrivateKey: privateKeyString,
    TaURL: "http://timestamp.digicert.com",
}
```

### Invoke `SignFile`.

You can get signed image with your manifest.

```
SignFile("./test.jpg", "signed.jpg", string(manifestString), signerInfo, "./data")
```

### Check the result
You can check the result with this web tool.

https://contentcredentials.org/verify

> [!TIP]
> You need to understand basic requrement for pair key.
> 
>Or you can also find sasmple pair key [here](https://github.com/contentauth/c2patool/tree/main/sample).
> 
> About the detail for `Signing manifeests`, please read [this page](https://opensource.contentauthenticity.org/docs/manifest/signing-manifests).

