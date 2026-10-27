package imagescan

import (
	"testing"
)

func TestScanImage(t *testing.T) {
	err := ScanImage("debian-jdk:trixie")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
