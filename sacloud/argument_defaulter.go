package sacloud

type argumentDefaulter interface {
	setDefaults() interface{}
}
