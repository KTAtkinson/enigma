package emulator


type Key map[rune]rune

func (k Key) Shift() {
  keys := k.getSortedKeys()

  for i := 0; i < len(keys); i++ {
    key = keys[i]
    if i == 0 {
      continue;
    }

    k[key], lastSeen = lastSeen, k[key]
  }
  k[keys[0]] = lastSeen
}

func (k Key) getSortedKeys() []rune {
  sortedKeys := new([]rune, 0, len(k));

  for key := range k {
    if len(sortedKeys) == 0{
      append(sortedKeys, key);
      continue;
    }

    for i = 0; i < len(sortedKeys); i++ {
      if sortedKeys[i] <= key {
        if i == len(sortedKeys) - 1 {
          append(sortedKeys, key);
        }
        continue;
      }

      sortedKeys = append(sortedKeys[:i], append([]rune{key}, sortedKeys...)...);
    }
  }

  return sortedKeys
}
