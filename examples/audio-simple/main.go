package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	var spec sdl.AudioSpec
	var audioData []uint8
	if !sdl.LoadWAV("tone.wav", &spec, &audioData, nil) {
		panic(sdl.GetError())
	}
	// if !sdl.LoadWAVIO(sdl.IOFromFile("tone.wav", "r"), true, &spec, &audioData, nil) { // Alternative way to load audio data
	// 	panic(sdl.GetError())
	// }

	defer sdl.Quit()
	sdl.Init(sdl.InitAudio)

	dataIndex, dataLen, loopCounter := int32(0), int32(len(audioData)), 0
	audioStream := sdl.OpenAudioDeviceStream(
		sdl.AudioDeviceDefaultPlayback,
		&spec,
		sdl.NewAudioStreamCallback(func(userdata unsafe.Pointer, stream *sdl.AudioStream, additionalAmount int32, totalAmount int32) {
			length := dataLen - dataIndex
			if additionalAmount < length {
				length = additionalAmount
			}

			sdl.PutAudioStreamData(stream, &audioData[dataIndex], length)

			dataIndex += length
			if dataIndex >= dataLen { // Loop the sound
				dataIndex = 0
				loopCounter++

				// In each loop increase gain and speed of the sound sample
				gain := sdl.GetAudioStreamGain(stream)
				sdl.SetAudioStreamGain(stream, gain+0.1)
				freq := sdl.GetAudioStreamFrequencyRatio(stream)
				sdl.SetAudioStreamFrequencyRatio(stream, freq+0.10) // 10% faster
			}
		}),
		nil)
	if audioStream == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyAudioStream(audioStream)

	sdl.SetAudioStreamGain(audioStream, 0.1) // Initial gain

	var inSpec, outSpec sdl.AudioSpec
	sdl.GetAudioStreamFormat(audioStream, &inSpec, &outSpec)
	fmt.Printf("Created audio stream with callback function that is using:\n - audio device ID: %d,\n - input (playback) spec: %#v,\n - output (recording) spec: %#v\n",
		sdl.GetAudioStreamDevice(audioStream), inSpec, outSpec)

	fmt.Println("Playing audio stream with callback function")
	if !sdl.ResumeAudioStreamDevice(audioStream) {
		panic(sdl.GetError())
	}

	// Keep the audio sample playing for 5 loops
	for loopCounter < 5 {
		time.Sleep(time.Millisecond * 10)
	}
	sdl.PauseAudioStreamDevice(audioStream)
	time.Sleep(time.Second)

	//
	//
	//
	// The same as above but without using callback function
	loopCounter, dataIndex = 0, 0

	audioStream2 := sdl.OpenAudioDeviceStream(sdl.AudioDeviceDefaultPlayback, &spec, 0, nil)
	if audioStream2 == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyAudioStream(audioStream2)

	sdl.SetAudioStreamGain(audioStream2, 0.1) // Initial gain

	sdl.GetAudioStreamFormat(audioStream2, &inSpec, &outSpec)
	fmt.Printf("Created audio stream without callback function that is using:\n - audio device ID: %d,\n - input (playback) spec: %#v,\n - output (recording) spec: %#v\n",
		sdl.GetAudioStreamDevice(audioStream), inSpec, outSpec)

	sdl.PutAudioStreamData(audioStream2, &audioData[0], dataLen) // Initial data to be played

	fmt.Println("Playing audio stream without callback function")
	if !sdl.ResumeAudioStreamDevice(audioStream2) {
		panic(sdl.GetError())
	}

	for loopCounter < 5 {
		// If the amount of data in the buffer is small put entire sound sample into the buffer, probably not the good approach for large sounds?
		if sdl.GetAudioStreamQueued(audioStream2) <= 256 {
			sdl.PutAudioStreamData(audioStream2, &audioData[0], dataLen)

			// In each loop increase gain and speed of the sound sample
			gain := sdl.GetAudioStreamGain(audioStream2)
			sdl.SetAudioStreamGain(audioStream2, gain+0.1)
			freq := sdl.GetAudioStreamFrequencyRatio(audioStream2)
			sdl.SetAudioStreamFrequencyRatio(audioStream2, freq+0.10) // 10% faster

			loopCounter++
		}

		time.Sleep(time.Millisecond * 10)
	}
}
