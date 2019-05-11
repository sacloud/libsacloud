package schema

import (
	"testing"

	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/stretchr/testify/require"
)

type dummyArgument struct {
	importStatements  []string
	packageName       string
	argName           string
	typeName          string
	zeroInitializer   string
	zeroValueOnSource string
	destination       string
}

// ImportStatements コード生成時に利用するimport文を生成する
func (d *dummyArgument) ImportStatements() []string {
	return d.importStatements
}

// PackageName インポートパスからパッケージ名を取得する
func (d *dummyArgument) PackageName() string {
	return d.packageName
}

// ArgName 引数の変数名、コード生成で利用される
func (d *dummyArgument) ArgName() string {
	return d.argName
}

// TypeName 型名の文字列表現、コード生成で利用される
func (d *dummyArgument) TypeName() string {
	return d.typeName
}

// ZeroInitializer 値を0初期化する文のコードの文字列表現、コード生成で利用される
func (d *dummyArgument) ZeroInitializer() string {
	return d.zeroInitializer
}

// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
func (d *dummyArgument) ZeroValueOnSource() string {
	return d.zeroValueOnSource
}

// DestinationFieldName マッピング先となるフィールド名を取得
func (d *dummyArgument) DestinationFieldName() string {
	return d.destination
}

// DestinationModel マッピング先のModelを取得
func (d *dummyArgument) DestinationModel() *Model {
	return nil
}

func TestSimpleArgument(t *testing.T) {

	expects := []struct {
		expect *dummyArgument
		input  *SimpleArgument
	}{
		{
			expect: &dummyArgument{
				importStatements:  wrapByDoubleQuote("github.com/sacloud/libsacloud-v2/sacloud/types"),
				packageName:       "types",
				argName:           "id",
				typeName:          "types.ID",
				zeroInitializer:   "types.ID(int64(0))",
				zeroValueOnSource: "types.ID(int64(0))",
			},
			input: ArgumentID.(*SimpleArgument),
		},
		{
			expect: &dummyArgument{
				importStatements:  []string{},
				packageName:       "",
				argName:           "zone",
				typeName:          "string",
				zeroInitializer:   `""`,
				zeroValueOnSource: `""`,
			},
			input: ArgumentZone.(*SimpleArgument),
		},
		{
			expect: &dummyArgument{
				importStatements:  wrapByDoubleQuote("github.com/sacloud/libsacloud-v2/sacloud"),
				packageName:       "sacloud",
				argName:           "client",
				typeName:          "*sacloud.Client",
				zeroInitializer:   "&sacloud.Client{}",
				zeroValueOnSource: "nil",
			},
			input: &SimpleArgument{
				Name: "client",
				Type: meta.Static(sacloud.Client{}),
			},
		},
	}

	for _, tc := range expects {
		require.Equal(t, tc.expect.ImportStatements(), tc.input.ImportStatements())
		require.Equal(t, tc.expect.PackageName(), tc.input.PackageName())
		require.Equal(t, tc.expect.ArgName(), tc.input.ArgName())
		require.Equal(t, tc.expect.TypeName(), tc.input.TypeName())
		require.Equal(t, tc.expect.ZeroInitializer(), tc.input.ZeroInitializer())
		require.Equal(t, tc.expect.ZeroValueOnSource(), tc.input.ZeroValueOnSource())
	}
}

func TestMappableArgument(t *testing.T) {
	var emptyList []string
	expects := []struct {
		expect *dummyArgument
		input  *MappableArgument
	}{
		{
			expect: &dummyArgument{
				importStatements:  emptyList,
				packageName:       "",
				argName:           "arg1",
				typeName:          "*Argument",
				zeroInitializer:   "&Argument{}",
				zeroValueOnSource: "nil",
				destination:       "Destination",
			},
			input: &MappableArgument{
				Name:        "arg1",
				Destination: "Destination",
				Model: &Model{
					Name: "Argument",
					Fields: []*FieldDesc{
						{
							Name: "field1",
							Type: meta.Static(""),
						},
						{
							Name: "field2",
							Type: meta.Static(""),
						},
					},
				},
			},
		},
	}

	for _, tc := range expects {
		require.Equal(t, tc.expect.ImportStatements(), tc.input.ImportStatements())
		require.Equal(t, tc.expect.PackageName(), tc.input.PackageName())
		require.Equal(t, tc.expect.ArgName(), tc.input.ArgName())
		require.Equal(t, tc.expect.TypeName(), tc.input.TypeName())
		require.Equal(t, tc.expect.ZeroInitializer(), tc.input.ZeroInitializer())
		require.Equal(t, tc.expect.ZeroValueOnSource(), tc.input.ZeroValueOnSource())
		require.Equal(t, tc.expect.DestinationFieldName(), tc.input.DestinationFieldName())
	}
}
