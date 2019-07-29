package sacloud

var clientFactory = make(map[string]func(APICaller) interface{})

// SetClientFactoryFunc リソースごとのクライアントファクトリーを登録する
func SetClientFactoryFunc(resourceName string, factoryFunc func(caller APICaller) interface{}) {
	clientFactory[resourceName] = factoryFunc
}

var clientFactoryHooks = make(map[string][]func(interface{}) interface{})

// AddClientFacotyHookFunc クライアントファクトリーのフックを登録する
func AddClientFacotyHookFunc(resourceName string, hookFunc func(interface{}) interface{}) {
	clientFactoryHooks[resourceName] = append(clientFactoryHooks[resourceName], hookFunc)
}

// GetClientFactoryFunc リソースごとのクライアントファクトリーを取得する
//
// resourceNameに対するファクトリーが登録されてない場合はpanicする
func GetClientFactoryFunc(resourceName string) func(APICaller) interface{} {
	f, ok := clientFactory[resourceName]
	if !ok {
		panic(resourceName + " is not found in clientFactory")
	}
	if hooks, ok := clientFactoryHooks[resourceName]; ok {
		return func(caller APICaller) interface{} {
			ret := f(caller)
			for _, hook := range hooks {
				ret = hook(ret)
			}
			return ret
		}
	}
	return f
}
