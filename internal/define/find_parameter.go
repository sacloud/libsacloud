package define

import "github.com/sacloud/libsacloud-v2/internal/schema"

var findParameter = &schema.Model{
	Fields: []*schema.FieldDesc{
		conditions.Count(),
		conditions.From(),
		conditions.Sort(),
		conditions.Filter(),
		conditions.Include(),
		conditions.Exclude(),
	},
}
