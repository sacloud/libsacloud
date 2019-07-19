# libsacloud DSL for defination of API

(このドキュメントは古くなっています。v2リリース時に最新の内容を反映します)

ここではlibsacloudでAPIを定義するために用いられているDSL`libsacloud DSL`について記します。

## libsacloudで利用されるDSLの主要概念/用語

libsacloud DSLではコード生成用の定義を行うために以下のような概念を用います。

- APIリソース
- オペレーション
- APIモデル
- nakedモデルとmapconvでのデータ変換

以下では擬似コードを用いて各概念の説明を行います。  
実際のDSLの利用方法は[Quick Start Guide](quick_start.md)を参照してください。

### APIリソース

`APIリソース`とは、APIで操作する対象となるさくらのクラウド上のリソースの種別を指します。
例: サーバ、ディスクなど

ソース上は`internal/schema`の`Resource`structがAPIリソースを表現しています。

APIリソースは自身に定義されているフィールドの値からAPIコール時のURLプレフィックスを決定します。

例えばAPIリソースとしてサーバを定義する場合、以下のようなフィールド/値を持ちます。

```go
// 以下は例示のためのイメージで実際の実装とは異なります
var serverResource = &schema.Resource {
	Name       : "Server",
	PathName   : "server",
	PathSuffix : "api/cloud/1.1",
	Operations : ...,
}
```

この場合、サーバ関連のAPIを利用する際のURLプレフィックスは

    https://secure.sakura.ad.jp/cloud/zone/{{.zone}}/api/cloud/1.1/server

となります。実際のURLはAPIリソースが保持するオペレーションによって決定されます。
(なおURLプレフィックスはDSLの定義や実行時の設定により変更可能です)

### オペレーション

`オペレーション`とはAPIリソースに対する操作を指します。  
例: 検索(Find)やCRUD(Create/Read/Update/Delete)など

ソース上は`innternal/schema`の`Operation`structがオペレーションを表現しています。

オペレーションは自身に定義されているフィールドの値から実際のAPIエンドポイントの決定やパラメータの処理、戻り値の処理などの詳細を決定します。  

例えばサーバリソースに対しての読み取り(Read)操作を定義する場合、以下のようなフィールド/値を持ちます。

```go
// 以下は例示のためのイメージで実際の実装とは異なります
var serverReadOperation = &schema.Operation {
	Resource         : serverResource, 
	Name             : "Read",
	Method           : http.MethodGet,
	PathFormat       : "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/{{.id}}",
	Arguments        : ...,
	Results          : ...,
	RequestEnvelope  : ...,
	ResponseEnvelope : ...,
}
```

この場合、サーバリソースへの読み取りを行うAPIのエンドポイントは

    https://secure.sakura.ad.jp/cloud/zone/{{.zone}}/api/cloud/1.1/server/<サーバのID>

となり、このエンドポイントに対しGETでリクエストを行うようなコードが生成されます。  

#### オペレーションへの引数/戻り値

オペレーションは引数(Arguments)と戻り値(Results)を持ちます。  
引数にはプリミティブなデータ型(IDやゾーンなど)や後述するAPIモデルを指定します。

例えば、引数として以下のように指定されていた場合、

```go
// 以下は例示のためのイメージで実際の実装とは異なります
var serverReadOperation = &schema.Operation {
    // ...
	Name             : "Update",
	Arguments        : []schema.Argument {
	    &schema.SimpleArgument {
	        Name: "zone",
	        Type: TypeString,
	    },
	    &schema.SimpleArgument {
	        Name: "id",
	        Type: TypeInt64,
	    },
	    &schema.MappableArgument{
	        Name: "param",
	        Type: &schema.Model{ // 引数の型情報としてAPIモデルを指定
	            Name: "ServerUpdateRequest",
	            Fields: []*schema.Field{
	                // ...
	            },
	        }
	    },
	},
	// ...
}
```

生成されるコードは以下のようになります。

    func (s *ServerOp) Update(ctx context.Context, zone string, id int64, param *ServerUpdateRequest) (*Server, error)

戻り値についても同様にAPIモデルを指定することで生成されるコードの戻り値が決定されるようになっています。

### APIモデル

`APIモデル`とはlibsacloud利用者が利用するコードにおけるAPIコール時のリクエスト/レスポンスなどのデータ型を指します。  

ソース上は`internal/schema`の`Model`structがAPIモデルを表現しています。

例えばAPIモデルとしてスタートアップスクリプトの作成パラメータを定義する場合、以下のようなフィールド/値を持ちます。

```go
var noteCreateRequest = &schema.Model {
	Name      : "NoteCreateRequest",
	Fields    : []*schema.Field {
        {
            Name: "Name",
       		Type: meta.TypeString,
        },
        {
            Name: "Tags",
       		Type: meta.TypeStringSlice,
        },
        // ...
	}
	NakedType : ...
}
``` 

この場合、生成されるコードは以下のようになります。

```go
type NoteCreateRequest struct {
	Name    string 
	Tags    []string
	// ...
}
```

APIモデルは後述するnakedモデルのフラットな表現となっており、実際のAPI呼び出し時に必要になる複雑な階層を持つパラメータをフラットでシンプルな形で提供します。

### nakedモデルとmapconvでのデータ変換

さくらのクラウドではAPI呼び出し時に指定するパラメータの階層が深い場合があります。  
これをシンプルに扱うために前述のAPIモデルが提供されます。  

しかしAPIモデルはそのままではAPI呼び出しに利用できる形になっていません。
そこでnakedモデルという形で実際のAPI呼び出し時のパラメータの型指定をしておき、mapconvという仕組みでAPIモデルとnakedモデルの変換を行います。

#### nakedモデル

nakedモデルとはさくらのクラウドAPI呼び出し時の実際のパラメータ/戻り値を定義したstructです。
ソースコード上は`sacloud/naked`パッケージ配下に格納されています。

例: スタートアップスクリプト用のnakedモデル

```go
// Note スタートアップスクリプト
type Note struct {
	ID           int64               `json:"ID,omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:"Name,omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:"Description,omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"Tags" yaml:"tags"`
	Availability enums.EAvailability `json:"Availability,omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        enums.EScope        `json:"Scope,omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Class        string              `json:"Class,omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Content      string              `json:"Content,omitempty" yaml:"content,omitempty" structs:",omitempty"`
	Icon         *Icon               `json:"Icon,omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:"CreatedAt,omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:"ModifiedAt,omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
}
```

nakedモデルは手作業で実装する必要があります。さくらのクラウド APIドキュメントやコントロールパネルからAPI呼び出しログなどを参照してstructを実装しています。

#### mapconvでのデータ変換

APIモデルで定義される各フィールドに`mapconv`タグを付与しておくことでAPIモデルとnakedモデル間でデータ変換が行えるようになっています。

例: スタートアップスクリプト用のAPIモデル

```go
type NoteCreateRequest struct {
	Name    string 
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}
```

IconIDに`mapconv:"Icon.ID"`というタグが付与されています。
これにより、IconIDは対応するnakedモデルの`Icon`フィールド配下の`ID`フィールドへと値がコピーされます。

