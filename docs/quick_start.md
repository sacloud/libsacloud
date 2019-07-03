# Quick Start Guide

ここではlibsacloud DSLを用いてAPIリソースの定義を行う方法について記します。  
libsacloud DSLについては[libsacloud DSL for defination of API](dls_overview.md)を参照してください。

## リソースの定義

リソースの定義は`internal/define`パッケージにて行います。

リソースの定義は以下の手順で行います。

- <リソース名>.goファイルを追加
  - オペレーションで利用するAPIモデルの定義
  - オペレーションの定義
  
### APIリソースの追加

libsacloud DSLでは`internal/define`パッケージで定義されているvariable`Resources`を参照してコード生成を行います。  
これはAPIリソースのリストとなっており、ここに定義を追加することで各種コード生成が行われるようになります。

### APIモデルの定義

APIモデルの定義には`schema.Model`を用いて以下のようにします。

```go
    // Findオペレーション向けのAPIモデル
    findParam := &schema.Model{
    	Fields: []*schema.FieldDesc{
    		conditions.Count(),
    		conditions.From(),
    		conditions.Sort(),
    		conditions.Filter(),
    		conditions.Include(),
    		conditions.Exclude(),
    	},
    }

    // Createオペレーション向けのAPIモデル
    createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Tags(),
			fields.IconID(),
			fields.NoteClass(),
			fields.NoteContent(),
		},
	}
```

`internal/define`パッケージにはモデル定義のショートカットとして`fields`と`conditions`というvariableが定義されています。  
これらを利用せず`schema.FieldDesc`を直接定義することも可能です。

### TODO オペレーションの定義
