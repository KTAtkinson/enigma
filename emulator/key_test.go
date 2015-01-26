package emulator

import "testing"


var testKey = Key{'c': 'c', 'a': 'a', 'b': 'b'}


func TestShift(t *testing.T) {
  expected := Key{'c': 'b', 'a': 'c', 'b': 'a'};
  keyCopy := testKey.copy_()
  keyCopy.Shift();

  for k := range testKey {
    if expected[k] != keyCopy[k] {
      t.Errorf("Expected key '%c' to equal %c, not %c.", k, expected[k], keyCopy[k])
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

func TestHasKey(t *testing.T) {
  tests := []struct {
    key rune
    outcome bool} {
      {key: 'a', outcome: true}, {key: 'd', outcome: false}}

   for _, e := range tests {
     result := testKey.HasKey(e.key);
     if result != e.outcome {
       t.Errorf("Expected %s, got %s.", e.outcome, result)
     }
   }
}

func TestSetFirstValue(t *testing.T) {
  expected := Key{'a': 'c', 'b': 'a', 'c': 'b'};
  keyCopy := testKey.copy_();
  keyCopy.SetFirstValue('c')

  for v := range testKey {
    if expected[v] != keyCopy[v] {
      t.Errorf("Expected %c to be %c, not %c", v, expected[v], keyCopy[v])
    }
  }
}
