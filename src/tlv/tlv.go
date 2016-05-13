package tlv

import (
    "reflect"
)

type Event struct {
  Type int    `json:"type"`
  Len  int    `json:"len"`
  Val  []byte `json:"val"` // Sorry, I can't remember the specifics from the whiteboard challenge :-|
}

func Handle(bytes []byte) []Event {
  const typeLength   = 2 // How many bytes are used for the type (T)
  const lengthLength = 4 // How many bytes determine the length of the value (L)

  var currentlyProcessing = "TYPE"
  var bytesEvaluated = 0

  var tmpType   []byte
  var tmpLength []byte
  var tmpValue  []byte

  var events []Event
  eventNo := 0

  // In lieu of setting up a stream :)
  for _, b := range bytes {
    switch currentlyProcessing {
      case "TYPE":
        tmpType = append(tmpType, b)
        bytesEvaluated++
        if bytesEvaluated == typeLength {
          events = append(events, Event{})
          events[eventNo].Type = fakeConvertToInt(tmpType)
          tmpType = []byte{}
          currentlyProcessing = "LENGTH"
          bytesEvaluated = 0
        }
      case "LENGTH":
        tmpLength = append(tmpLength, b)
        bytesEvaluated++
        if bytesEvaluated == lengthLength {
          events[eventNo].Len = fakeConvertToInt(tmpLength)
          tmpLength = []byte{}
          currentlyProcessing = "VALUE"
          bytesEvaluated = 0
        }
      case "VALUE":
        tmpValue = append(tmpValue, b)
        bytesEvaluated++
        if bytesEvaluated == events[eventNo].Len {
          events[eventNo].Val = tmpValue
          tmpValue = []byte{}
          // Do the things with events :)
          currentlyProcessing = "TYPE"
          eventNo++
          bytesEvaluated = 0
        }
     }
  }
  return events
}

// fakeConvertToInt does a fake conversion from []byte to int
// because Go has no native support for binary literals :-|
// []byte is an alias for uint8
// https://golang.org/pkg/builtin/#byte
func fakeConvertToInt(b []byte) int {
  if reflect.DeepEqual(b,[]byte{0,0}) { return 0 }
  if reflect.DeepEqual(b,[]byte{0,1}) { return 1 }
  if reflect.DeepEqual(b,[]byte{1,0}) { return 2 }
  if reflect.DeepEqual(b,[]byte{1,1}) { return 3 }

  if reflect.DeepEqual(b,[]byte{0,0,0,0}) { return 0 }
  if reflect.DeepEqual(b,[]byte{0,0,0,1}) { return 1 }
  if reflect.DeepEqual(b,[]byte{0,0,1,0}) { return 2 }
  if reflect.DeepEqual(b,[]byte{0,0,1,1}) { return 3 }
  if reflect.DeepEqual(b,[]byte{0,1,0,0}) { return 4 }
  if reflect.DeepEqual(b,[]byte{0,1,0,1}) { return 5 }
  if reflect.DeepEqual(b,[]byte{0,1,1,0}) { return 6 }
  if reflect.DeepEqual(b,[]byte{0,1,1,1}) { return 7 }
  if reflect.DeepEqual(b,[]byte{1,0,0,0}) { return 8 }
  if reflect.DeepEqual(b,[]byte{1,0,0,1}) { return 9 }
  if reflect.DeepEqual(b,[]byte{1,0,1,0}) { return 10 }
  if reflect.DeepEqual(b,[]byte{1,0,1,1}) { return 11 }
  if reflect.DeepEqual(b,[]byte{1,1,0,0}) { return 12 }
  if reflect.DeepEqual(b,[]byte{1,1,0,1}) { return 13 }
  if reflect.DeepEqual(b,[]byte{1,1,1,0}) { return 14 }
  if reflect.DeepEqual(b,[]byte{1,1,1,1}) { return 15 }
  return 0
}
