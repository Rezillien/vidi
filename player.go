package main
import (
  "github.com/faiface/beep/mp3";
  "github.com/faiface/beep";
  "github.com/faiface/beep/speaker";
  "time";
  "log";
  "os"
)
func main() {
  f, err := os.Open("file_example_MP3_5MG.mp3")
  if err != nil {
		log.Fatal(err)
	}

  streamer, format, err := mp3.Decode(f)
  if err != nil {
		log.Fatal(err)
	}

  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
  defer streamer.Close()

  done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

}
