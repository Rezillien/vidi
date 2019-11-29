package main
import (
  "github.com/gvalkov/golang-evdev";
  "github.com/faiface/beep/mp3";
  "github.com/faiface/beep/speaker";
  "fmt";
  "time";
  "log";
  "os"
)
func main() {
  device, _ := evdev.Open("/dev/input/event2")

  fmt.Println(device)
  fmt.Println(device.Capabilities)

  fmt.Println(evdev.REP)

  keymap := device.Capabilities[struct{Type int; Name string}{Type: 1, Name: "EV_KEY"}]

  fmt.Printf("%+v\n", device)

  f, err := os.Open("file_example_MP3_5MG.mp3")
  if err != nil {
		log.Fatal(err)
	}

  streamer, format, err := mp3.Decode(f)
  if err != nil {
		log.Fatal(err)
	}

  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
  speaker.Play(streamer)
  defer streamer.Close()

  for ;; {
    event, _ := device.ReadOne()
    if(event.Type != 1){
      continue
    }
    fmt.Printf("%+v\n",event)
    fmt.Println(keymap[event.Code-1])
  }
}

