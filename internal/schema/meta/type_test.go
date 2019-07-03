package meta

import (
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestStaticType_TypeImplements(t *testing.T) {

	expects := []struct {
		caseName                 string
		instance                 interface{}
		goType                   string
		goPkg                    string
		goImportPath             string
		goTypeSourceCode         string
		zeroInitializeSourceCode string
		zeroValueSourceCode      string
		fatal                    bool
	}{
		{
			caseName: "unsupported primitive type",
			instance: int32(0),
			fatal:    true,
		},
		{
			caseName:                 "bool",
			instance:                 true,
			goType:                   "bool",
			goPkg:                    "",
			goImportPath:             "",
			goTypeSourceCode:         "bool",
			zeroInitializeSourceCode: "false",
			zeroValueSourceCode:      "false",
		},
		{
			caseName:                 "int",
			instance:                 0,
			goType:                   "int",
			goPkg:                    "",
			goImportPath:             "",
			goTypeSourceCode:         "int",
			zeroInitializeSourceCode: "0",
			zeroValueSourceCode:      "0",
		},
		{
			caseName:                 "int64",
			instance:                 int64(0),
			goType:                   "int64",
			goPkg:                    "",
			goImportPath:             "",
			goTypeSourceCode:         "int64",
			zeroInitializeSourceCode: "int64(0)",
			zeroValueSourceCode:      "int64(0)",
		},
		{
			caseName:                 "string",
			instance:                 "",
			goType:                   "string",
			goPkg:                    "",
			goImportPath:             "",
			goTypeSourceCode:         "string",
			zeroInitializeSourceCode: `""`,
			zeroValueSourceCode:      `""`,
		},
		{
			caseName:                 "map",
			instance:                 map[string]interface{}{},
			goType:                   "map[string]interface {}",
			goPkg:                    "",
			goImportPath:             "",
			goTypeSourceCode:         "map[string]interface {}",
			zeroInitializeSourceCode: "map[string]interface {}{}",
			zeroValueSourceCode:      "nil",
		},
		{
			caseName:                 "slice",
			instance:                 []string{},
			goType:                   "[]string",
			goPkg:                    "",
			goImportPath:             "",
			goTypeSourceCode:         "[]string",
			zeroInitializeSourceCode: "[]string{}",
			zeroValueSourceCode:      "nil",
		},
		{
			caseName:                 "enum",
			instance:                 types.EAvailability(""),
			goType:                   "types.EAvailability",
			goPkg:                    "types",
			goImportPath:             "github.com/sacloud/libsacloud/sacloud/types",
			goTypeSourceCode:         "types.EAvailability",
			zeroInitializeSourceCode: `types.EAvailability("")`,
			zeroValueSourceCode:      `types.EAvailability("")`,
		},
		{
			caseName:                 "another package struct",
			instance:                 sacloud.Client{},
			goType:                   "sacloud.Client",
			goPkg:                    "sacloud",
			goImportPath:             "github.com/sacloud/libsacloud/sacloud",
			goTypeSourceCode:         "*sacloud.Client",
			zeroInitializeSourceCode: "&sacloud.Client{}",
			zeroValueSourceCode:      "nil",
		},
	}

	for _, expect := range expects {
		t.Run(expect.caseName, func(t *testing.T) {
			defer func() {
				err := recover()
				require.Equal(t, expect.fatal, err != nil)
			}()
			var tp Type = Static(expect.instance)
			require.Equal(t, expect.goType, tp.GoType())
			require.Equal(t, expect.goPkg, tp.GoPkg())
			require.Equal(t, expect.goImportPath, tp.GoImportPath())
			require.Equal(t, expect.goTypeSourceCode, tp.GoTypeSourceCode())
			require.Equal(t, expect.zeroInitializeSourceCode, tp.ZeroInitializeSourceCode())
			require.Equal(t, expect.zeroValueSourceCode, tp.ZeroValueSourceCode())
		})
	}

}
