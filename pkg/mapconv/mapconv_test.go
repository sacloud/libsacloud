package mapconv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type dummyTagged struct {
	A          string `mapconv:"ValueA.A"`
	B          string `mapconv:"ValueA.ValueB.B"`
	C          string `mapconv:"ValueA.ValueB.ValueC.C"`
	NoTag      string
	unexported string
}

type dummyNaked struct {
	ValueA *struct {
		A      string
		ValueB *struct {
			B      string
			ValueC *struct {
				C string
			}
		}
	}
	NoTag      string
	unexported string
}

func TestToNaked(t *testing.T) {

	expects := []struct {
		tagged *dummyTagged
		naked  *dummyNaked
		err    error
	}{
		{
			tagged: &dummyTagged{
				A:          "A",
				B:          "B",
				C:          "C",
				NoTag:      "NoTag",
				unexported: "unexported",
			},
			naked: &dummyNaked{
				ValueA: &struct {
					A      string
					ValueB *struct {
						B      string
						ValueC *struct {
							C string
						}
					}
				}{
					A: "A",
					ValueB: &struct {
						B      string
						ValueC *struct {
							C string
						}
					}{
						B: "B",
						ValueC: &struct {
							C string
						}{
							C: "C",
						},
					},
				},
				NoTag: "NoTag",
			},
		},
	}

	for _, expect := range expects {
		naked := &dummyNaked{}
		err := ToNaked(expect.tagged, naked)
		require.Equal(t, expect.err, err)
		if err == nil {
			require.Equal(t, expect.naked, naked)
		}
	}

}

func TestFromNaked(t *testing.T) {

	expects := []struct {
		tagged *dummyTagged
		naked  *dummyNaked
		err    error
	}{
		{
			tagged: &dummyTagged{
				A:     "A",
				B:     "B",
				C:     "C",
				NoTag: "NoTag",
			},
			naked: &dummyNaked{
				ValueA: &struct {
					A      string
					ValueB *struct {
						B      string
						ValueC *struct {
							C string
						}
					}
				}{
					A: "A",
					ValueB: &struct {
						B      string
						ValueC *struct {
							C string
						}
					}{
						B: "B",
						ValueC: &struct {
							C string
						}{
							C: "C",
						},
					},
				},
				NoTag: "NoTag",
			},
		},
	}

	for _, expect := range expects {
		tagged := &dummyTagged{}
		err := FromNaked(expect.naked, tagged)
		require.Equal(t, expect.err, err)
		if err == nil {
			require.Equal(t, expect.tagged, tagged)
		}
	}

}
