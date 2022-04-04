package verifyurl

import "testing"

var urltest1 = "https://archives.bassdrivearchive.com/4%20-%20Thursday/Funked%20Up%20Radio%20-%20DFunk/%5b2018.07.19%5d%20Funked%20Up%20Radio%20-%20DFunk.mp3"
var urltest2 = "https://archives.bassdrivearchive.com/1%20-%20Monday/Translation%20Sound%20-%20Rogue%20State/%5b2022.01.24%5d%20Translation%20Sound%20-%20Rogue%20State.mp3"
var urltest3 = "badlyformed"

func TestVerify(t *testing.T) {

	result := Verify(urltest1)

	if result == nil {
		t.Errorf("Test 1 should have been invalid.")
	}

	result = Verify(urltest2)

	if result != nil {
		t.Errorf("Test 2 should have been valid.")
	}

	result = Verify(urltest3)

	if result == nil {
		t.Errorf("This one should have been badly formed.")
	}
}
