package list

import "testing"

func TestGetLatestVersion(t *testing.T) {
	t.Parallel()

	t.Run("not exist image", func(t *testing.T) {
		t.Parallel()

		res, err := GetLatestVersion("abcde")
		if err == nil {
			t.Errorf("want error")
		}

		if res != "" {
			t.Errorf("want empty")
		}
	})
	t.Run("exist image", func(t *testing.T) {
		t.Parallel()

		res, err := GetLatestVersion("gcr.io/cloud-spanner-emulator/emulator")
		if err != nil {
			t.Fatal(err)
		}

		if res == "" {
			t.Errorf("want not empty")
		}
	})
}
