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
    fmt.Println(string(store));
    fmt.Println(string(k[key]));
    lastSeen = store
  }
  k[keys[0]] = lastSeen
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
