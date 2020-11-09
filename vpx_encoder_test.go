package lenss

import "testing"

func TestEncoderInitialize(t *testing.T) {
	_, err := NewEncoder(CodecVP8, 480, 320, 30, 2, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
}
