package unittest_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"practice/unittest"
	"testing"
)

func TestGetPositionSlice(t *testing.T) {
	index := 3
	var expected int64 = 4
	result, err := unittest.GetPositionSlice(index, 1, 2, 3, 4)

	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify that the value is as expected")
	/*
		if err != nil {
			t.Errorf("Unexpect Error %s", err.Error())
			return
		}
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
			return
		}
	*/
	t.Log("successful")
}

func TestGetPositionSlice_out_of_range(t *testing.T) {

	indx := 4
	var expected int64 = 0
	result, err := unittest.GetPositionSlice(indx, 1, 2, 3, 4)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if !errors.Is(err, unittest.ErrIndexOutRange) {
		t.Errorf("Expected %s , got %s", unittest.ErrIndexOutRange.Error(), err.Error())
	}
	if result != expected {
		t.Errorf("Expected %d , got %d", expected, result)

	}

}
