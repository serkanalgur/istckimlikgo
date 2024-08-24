package istckimlikgo

import "testing"

func TestValidate(t *testing.T) {
	t.Run("Valiadte", func(t *testing.T) {
		toBeValid, err := validate("11111111111")
		if err != nil {
			t.Errorf("Error = %v", err)
			return
		} else if toBeValid != true {
			t.Errorf("Error = %v", err)
		}
	})
}
