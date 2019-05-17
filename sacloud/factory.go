package sacloud

var clientFactory = make(map[string]func(APICaller) interface{})

// SetClientFactoryFunc リソースごとのクライアントファクトリーを登録する
func SetClientFactoryFunc(resourceName string, factoryFunc func(caller APICaller) interface{}) {
	clientFactory[resourceName] = factoryFunc
}

// GetClientFactoryFunc リソースごとのクライアントファクトリーを取得する
//
// resourceNameに対するファクトリーが登録されてない場合はpanicする
func GetClientFactoryFunc(resourceName string) func(APICaller) interface{} {
	f, ok := clientFactory[resourceName]
	if !ok {
		panic(resourceName + " is not found in clientFactory")
	}
	return f
}
