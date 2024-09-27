package handler

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

// cmpTransformJSON はcmp.DiffでJSON文字列([]byte)を比較のためのオプション
//
// この設定を入れることでJSON文字列の改行や空白を無視してくれる。
func cmpTransformJSON(t *testing.T) cmp.Option {
	return cmp.FilterValues(
		func(x, y []byte) bool {
			return json.Valid(x) && json.Valid(y)
		},
		cmp.Transformer("ParseJSON", func(in []byte) (out interface{}) {
			err := json.Unmarshal(in, &out)
			require.NoError(t, err)
			return out
		}),
	)
}

// ignoreMapEntires go-cmpのDiffで特定のmapの特定のkeyを無視するオプションを返す
//
// cmpTransformJSONを使っているときに'CreatedAt'や'UpdatedAt'を無視する場合に便利。
func ignoreMapEntires(want map[string]any) cmp.Option {
	return cmpopts.IgnoreMapEntries(func(k string, _ any) bool {
		_, ok := want[k]
		return ok
	})
}
