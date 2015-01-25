package emulator

import "fmt"


type RotorEncoder interface {
  EncodeMessage(string) string

  GetInitialPosition() rune
  GetTotalShifts() int

  IsAtStart() bool
  Reset()
  Shift()
}


type Rotor struct {
  totalShifts int
  init rune
  key Key
}

func (rotor Rotor) EncodeMessage(message string) string {
  encoded := "";
  for _, r := range message {
    encoded += string(rotor.key[r]);
  }

  return encoded
}

func (rotor Rotor) GetInitialPosition() rune {
  return rotor.init
}

func (rotor Rotor) GetShiftsPerRotation() int {
  return len(rotor.key)
}

func (rotor Rotor) GetTotalShifts() int {
  return rotor.totalShifts
}

func (rotor Rotor) Reset() {
  for rotor.key['a'] != rotor.init {
    rotor.key.Shift()
  }
}

func (rotor Rotor) Shift() {
  rotor.key.Shift();
  rotor.totalShifts += 1
}

func (rotor Rotor) IsAtStart() bool {
  return rotor.init == rotor.key['a']
}


func NewAlphabeticRotor(init rune) (Rotor, error) {
  if init == 'a' {
    return Rotor{}, fmt.Errorf("Rotor can't be inititialized with 'a' as the initial rotor position.")
  }
  key := Key{}
  for _, r := range EN {
    key[r] = r
  }

  rotor := Rotor{key: key, init: init};
  rotor.Reset();
  return rotor, nil
}
