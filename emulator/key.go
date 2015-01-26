package emulator

import "fmt"


type Key map[rune]rune

func (k Key) Shift() {
  keys := k.getSortedKeys()
  var lastSeen rune;

  for i := 0; i < len(keys); i++ {
    key := keys[i]
    if i == 0 {
      lastSeen = k[key]
      continue;
    }

    store := k[key];
    k[key] = lastSeen;
    lastSeen = store
  }
  k[keys[0]] = lastSeen
}

func (k Key) SetFirstValue(v rune) error {
  firstKey := k.getSortedKeys()[0];
  if k.HasKey(v) != true {
    return fmt.Errorf("%c is not contained in the key.", v);
  } else if firstKey == v {
    return fmt.Errorf("The first key value can't be equal to itself.")
  }

  for k[firstKey] != v {
    k.Shift()
  }
  return nil
}

func (k Key) copy_() Key {
  newKey := Key{}

  for key := range k {
    newKey[key] = k[key];
  }

  return newKey
}

func (k Key) HasKey(key rune) bool {
  for e := range k {
    if e == key {
      return true
    }
  }
  return false
}

func (k Key) getSortedKeys() []rune {
  sortedKeys := make([]rune, 0, len(k));

  for key := range k {
    if len(sortedKeys) == 0 {
      sortedKeys = append(sortedKeys, key);
      continue;
    }

    for i := 0; i < len(sortedKeys); i++ {
      if sortedKeys[i] <= key {
        if i == len(sortedKeys) - 1 {
          sortedKeys = append(sortedKeys, key);
          break;
        }
        continue;
      }

      sortedKeys = append(sortedKeys[:i], append([]rune{key}, sortedKeys[i:]...)...);
      break;
    }
  }

  return sortedKeys
}
