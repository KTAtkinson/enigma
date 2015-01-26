package emulator


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

func (rotor Rotor) Reset() error {
  err := rotor.key.SetFirstValue(rotor.init);
  rotor.totalShifts = 0;
  return err
}

func (rotor Rotor) copy_() Rotor {
  newRotor := Rotor{key: rotor.key.copy_(), init: rotor.init};
  newRotor.Reset()
  return newRotor
}

func (rotor Rotor) Shift() {
  rotor.key.Shift();
  rotor.totalShifts += 1
}

func (rotor Rotor) IsAtStart() bool {
  return rotor.init == rotor.key['a']
}


func NewAlphabeticRotor(init rune) (Rotor, error) {
 key := Key{}
  for _, r := range EN {
    key[r] = r
  }

  rotor := Rotor{key: key, init: init};
  err := rotor.key.SetFirstValue(rotor.init);
  if err != nil {
    return rotor, err
  }

  return rotor, nil
}
