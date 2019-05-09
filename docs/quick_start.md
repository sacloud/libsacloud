# Quick Start Guide

ここではlibsacloud DSLを用いてAPIリソースの定義を行う方法について記します。  
libsacloud DSLについては[libsacloud DSL for defination of API](dls_overview.md)を参照してください。

## リソースの定義

リソースの定義は`internal/define`パッケージにて行います。

リソースの定義は以下の手順で行います。

- <リソース名>.goファイルを追加
- func init()の実装
  - オペレーションで利用するAPIモデルの定義
  - `Resources`に対してDefineメソッドの呼び出し
  - オペレーションの定義
  
### func init()の実装

libsacloud DSLでは`internal/define`パッケージで定義されているvariable`Resources`を参照してコード生成を行います。  
これはAPIリソースのリストとなっており、ここに定義を追加することで各種コード生成が行われるようになります。

`Resources`へ定義を追加するために各APIリソースごとにソースファイルを分けた上でfunc init()を実装するようにしています。

### APIモデルの定義

func init()の中でAPIモデルの定義を行います。
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

### Resourcesに対するDefineメソッドの呼び出し

`Resources`はFluent APIを提供しており、メソッドチェーンで定義を投入することができるようになっています。
APIリソースを登録するには`Define`メソッドを呼び出します。

```go
	Resources.Define("Note")
```

定義したAPIリソースに対し設定を行いたい場合はメソッドチェーンで指定します。
> 通常はAPIリソースの各フィールドにはデフォルト値が設定されるため.Nameや.PathNameなどを呼び出す必要はありません

```go
	Resources.Define("Note").
		Name("Note").               // リソース名を指定
		PathName("note").           // URLパスを指定
		PathSuffix("api/cloud/1.1") // パスサフィックスを指定
```

### オペレーションの追加

APIリソースに対しオペレーションを追加します。

```go
	Resources.Define("Note").
	    Operation(op1).
	    Operation(op2).
```

Find操作+CRUD操作の場合ショートカット`OperationCRUD`メソッドが利用できます。

```go
	Resources.Define("Note").
		OperationCRUD(meta.Static(naked.Note{}), findParam, createParam, updateParam, result)
```
