package external

import (
	"testing"
)

func TestGetLocals_OK(t *testing.T) {
	res, err := GetExternalPokemon(1)

	if err != nil {
		t.Errorf("fails to get pokemon %s", err)
		t.FailNow()
	}

	t.Logf("Result: %v", res)

}
