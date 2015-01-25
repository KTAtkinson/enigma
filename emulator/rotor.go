package emulator


type RotorEncoder interface{
  EncodeMessage(string) string

  GetInitialPosition rune
  GetTotalShifts int

  IsAtStart bool
  Reset
  Shift
}


type Rotor struct {
  totalShifts int
  init rune
  key Key
}

func (rotor Rotor) EncodeMessage(message string) string {
  encoded = new string
  for _, r := range message {
    encoded += Key[r];
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

func (rotor Rotor) IsAtStart() {
  return rotor.init == rotor.key['a']
}


func NewAlphabeticRotor(init rune) Rotor, err {
  if init == 'a' {
    return nil, 'Rotor can\'t be inititialized with "a" as the initial rotor position.'
  key := Key{}
  for _, r := range EN {
    key[r] = r
  }

  rotor := rotor{key:key, init: init};
  rotor.Reset();
  return rotor, nil
}
