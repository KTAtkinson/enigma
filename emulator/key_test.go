package emulator

import "testing"


var testKey = Key{'c': 'b', 'a': 'c', 'b': 'a'}


func TestShift(t *testing.T) {
  expected := Key{'c': 'a', 'a': 'b', 'b': 'c'};
  testKey.Shift();

  for k := range testKey {
    if expected[k] != testKey[k] {
      t.Errorf("Expected key '%c' to equal %c, not %c.", k, expected[k], testKey[k])
    }
  }
}

func TestGetSortedKeys(t *testing.T) {
  expected := []rune{'a', 'b', 'c'}
  sorted := testKey.getSortedKeys()

  err := false
  for i, v := range expected {
    if v != sorted[i] {
      err = true;
      break;
    }
  }
  if err {
    t.Errorf("Expected keys in %c order, found them in %c order,", expected, sorted)
  }
}



