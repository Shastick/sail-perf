package sail_perf

import "testing"

func TestInit(t *testing.T) {
	scores := New(10,10)
        if len(scores.samples) != 10 {
		t.Errorf("Length of samples is not 10.")
	}
	if cap(scores.samples) != 10 {
		t.Errorf("Capacity of samples is not 10.")
	}
 	for _, item := range scores.samples {
		if item != nil {
			t.Errorf("Found an initialized sub-slice.")
		}
	}
}
