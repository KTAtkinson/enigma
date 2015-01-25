package emulator


type Enigma interface {
  EncodeMessage(string) string
  DecodeMessage(string) string
  GetRotorSettings(string) string
}

