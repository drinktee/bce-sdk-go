# bce-sdk-go

A golang sdk for baidu cloud.

Base code copy from [https://github.com/guoyao/baidubce-sdk-go](https://github.com/guoyao/baidubce-sdk-go)

## Example

```
	var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

	//var bceConfig = bce.NewConfig(credentials)
	var bceConfig = &bce.Config{
		Credentials: credentials,
		Checksum:    true,
		Timeout:     5 * time.Second,
		Region:      os.Getenv("BOS_REGION"),
	}
	var bccConfig = NewConfig(bceConfig)
	eipClient = NewEIPClient(bccConfig)
	eipClient.SetDebug(true)
	eips, _ := eipClient.GetEips(nil)
	for _, eip := range eips {
		fmt.Printf("%+v \n", eip)
	}    
```
## Contributors

- Guoyao Wu (https://www.guoyao.me)
- me (zjsxzong89@gmail.com)
- Yuxiao Song (billsong6919@gmail.com)

## Reference

- [https://github.com/denverdino/aliyungo](https://github.com/denverdino/aliyungo)