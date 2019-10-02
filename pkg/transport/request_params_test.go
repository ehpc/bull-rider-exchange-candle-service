package transport

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRequestParamsHash(t *testing.T) {
	t.Run("SimpleCases", func(t *testing.T) {
		tests := []struct{
			in RequestParams
			out string
		}{
			{
				RequestParams{
					"k1": "v1",
					"k2": "v2",
				},
				"k1:v1#k2:v2",
			},
			{
				RequestParams{
					"k1": "v1",
				},
				"k1:v1",
			},
			{
				RequestParams{},
				"",
			},
		}
		for _, test := range tests {
			assert.Equal(t, test.out, test.in.Hash())
		}
	})
	t.Run("PreserveOrder", func(t *testing.T) {
		rp := RequestParams{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
			"k4": "v4",
		}
		for i := 0; i<100; i++ {
			assert.Equal(t, "k1:v1#k2:v2#k3:v3#k4:v4", rp.Hash())
		}
	})
}
