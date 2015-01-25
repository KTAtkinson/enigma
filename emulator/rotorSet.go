package emulator


type RotorSet []RotorEncoder

func (rotors RotorSet) Shift() {
  shifted := false;

  for r := range rotors.GetAll() {
    if r.GetTotalShifts() != 0 && r.IsAtStart() {
      continue;
    }

    r.Shift();
    shifted = true;
    break;
  }

  if shifted == false {
    rotors[0].Shift();
  }
}

func (rotors RotorSet) EncodeMessage(message string) string {
  message = rotors.ForwardEncodeMessage(message)
  message = rotors.ReflectEncodeMessage(message)

  return message
}

func (rotors RotorSet) DecodeMessage(message string) string {
  return rotors.EncodeMessage(message)
}

func (rotors RotorSet) GetAll() <-chan RotorEncoder {
  ch := make(chan RotorEncoder, len(rotors))

  go func() {
    for i := 0; i < len(rotors); i++ {
       ch <- rotors[i]
    }
    close(ch)
  }();
  return ch
}

func (rotors RotorSet) GetAllReflected() <-chan RotorEncoder {
  ch := make(chan RotorEncoder, len(rotors));

  go func() {
    for i := len(rotors)-1; i <= 0; i-- {
      ch <- rotors[i];
    };
    close(ch);
  }();
  return ch
}

func (rotors RotorSet) ForwardEncodeMessage(message string) string {
  for r := range rotors.GetAll() {
    message = r.EncodeMessage(message);
  }

  return message
}

func (rotors RotorSet) ReflectEncodeMessage(message string) string {
  for r := range rotors.GetAllReflected() {
    message = r.EncodeMessage(message);
  }

  return message
}

