package enums

// EPlanGeneration サーバプラン世代
type EPlanGeneration int

var (
	// PlanGenerationValues 有効なサーバプラン世代の値
	PlanGenerationValues = []int{
		int(PlanGenerations.Default),
		int(PlanGenerations.G100),
		int(PlanGenerations.G200),
	}

	// PlanGenerations サーバプラン世代
	PlanGenerations = struct {
		// Default デフォルト(自動選択)
		Default EPlanGeneration
		// G100 第1世代
		G100 EPlanGeneration
		// G200 第2世代
		G200 EPlanGeneration
	}{
		Default: EPlanGeneration(0),
		G100:    EPlanGeneration(100),
		G200:    EPlanGeneration(200),
	}
)
