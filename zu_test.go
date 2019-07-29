package zu

import "testing"

func TestZu(t *testing.T) {

}

func app() error {
	for Next() {
		if quit.IsDown() {
			break
		}
	}
	return nil
}
