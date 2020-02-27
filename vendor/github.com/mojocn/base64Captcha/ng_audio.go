// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

import (
	"bytes"
	"encoding/binary"

	"io"
	"math"
)

//ConfigAudio captcha config for captcha-engine-audio.
type ConfigAudio struct {
	// CaptchaLen Default number of digits in captcha solution.
	CaptchaLen int
	// Language possible values for lang are "en", "ja", "ru", "zh".
	Language string
}

const sampleRate = 8000 // Hz

var endingBeepSound []byte

func init() {
	endingBeepSound = changeSpeed(beepSound, 1.4)
}

//Audio captcha-audio-engine return type.
type Audio struct {
	CaptchaItem
	body        *bytes.Buffer
	digitSounds [][]byte
	rng         siprng
}

//EngineAudioCreate create captcha with configAudio.
func EngineAudioCreate(id string, config ConfigAudio) *Audio {
	digits := randomDigits(config.CaptchaLen)
	audio := newAudio(id, digits, config.Language)
	audio.VerifyValue = parseDigitsToString(digits)
	return audio
}

// newAudio returns a new audio captcha with the given digits, where each digit
// must be in range 0-9. Digits are pronounced in the given language. If there
// are no sounds for the given language, English is used.
//
// Possible values for lang are "en", "ja", "ru", "zh".
func newAudio(id string, digits []byte, lang string) *Audio {
	a := new(Audio)

	// Initialize PRNG.
	a.rng.Seed(deriveSeed(audioSeedPurpose, id, digits))

	if sounds, ok := digitSounds[lang]; ok {
		a.digitSounds = sounds
	} else {
		a.digitSounds = digitSounds["en"]
	}
	numsnd := make([][]byte, len(digits))
	for i, n := range digits {
		snd := a.randomizedDigitSound(n)
		setSoundLevel(snd, 1.5)
		numsnd[i] = snd
	}
	// Random intervals between digits (including beginning).
	intervals := make([]int, len(digits)+1)
	intdur := 0
	for i := range intervals {
		dur := a.rng.Int(sampleRate, sampleRate*2) // 1 to 2 seconds
		intdur += dur
		intervals[i] = dur
	}
	// Generate background sound.
	bg := a.makeBackgroundSound(a.longestDigitSndLen()*len(digits) + intdur)
	// Create buffer and write audio to it.
	sil := makeSilence(sampleRate / 5)
	bufcap := 3*len(beepSound) + 2*len(sil) + len(bg) + len(endingBeepSound)
	a.body = bytes.NewBuffer(make([]byte, 0, bufcap))
	// Write prelude, three beeps.
	a.body.Write(beepSound)
	a.body.Write(sil)
	a.body.Write(beepSound)
	a.body.Write(sil)
	a.body.Write(beepSound)
	// Write digits.
	pos := intervals[0]
	for i, v := range numsnd {
		mixSound(bg[pos:], v)
		pos += len(v) + intervals[i+1]
	}
	a.body.Write(bg)
	// Write ending (one beep).
	a.body.Write(endingBeepSound)
	return a
}

// encodedLen returns the length of WAV-encoded audio captcha.
func (a *Audio) encodedLen() int {
	return len(waveHeader) + 4 + a.body.Len()
}

func (a *Audio) makeBackgroundSound(length int) []byte {
	b := a.makeWhiteNoise(length, 4)
	for i := 0; i < length/(sampleRate/10); i++ {
		snd := reversedSound(a.digitSounds[a.rng.Intn(10)])
		//snd = changeSpeed(snd, a.rng.Float(0.8, 1.2))
		place := a.rng.Intn(len(b) - len(snd))
		setSoundLevel(snd, a.rng.Float(0.04, 0.08))
		mixSound(b[place:], snd)
	}
	return b
}

func (a *Audio) randomizedDigitSound(n byte) []byte {
	s := a.randomSpeed(a.digitSounds[n])
	setSoundLevel(s, a.rng.Float(0.85, 1.2))
	return s
}

func (a *Audio) longestDigitSndLen() int {
	n := 0
	for _, v := range a.digitSounds {
		if n < len(v) {
			n = len(v)
		}
	}
	return n
}

func (a *Audio) randomSpeed(b []byte) []byte {
	pitch := a.rng.Float(0.95, 1.1)
	return changeSpeed(b, pitch)
}

func (a *Audio) makeWhiteNoise(length int, level uint8) []byte {
	noise := a.rng.Bytes(length)
	adj := 128 - level/2
	for i, v := range noise {
		v %= level
		v += adj
		noise[i] = v
	}
	return noise
}

// mixSound mixes src into dst. Dst must have length equal to or greater than
// src length.
func mixSound(dst, src []byte) {
	for i, v := range src {
		av := int(v)
		bv := int(dst[i])
		if av < 128 && bv < 128 {
			dst[i] = byte(av * bv / 128)
		} else {
			dst[i] = byte(2*(av+bv) - av*bv/128 - 256)
		}
	}
}

func setSoundLevel(a []byte, level float64) {
	for i, v := range a {
		av := float64(v)
		switch {
		case av > 128:
			if av = (av-128)*level + 128; av < 128 {
				av = 128
			}
		case av < 128:
			if av = 128 - (128-av)*level; av > 128 {
				av = 128
			}
		default:
			continue
		}
		a[i] = byte(av)
	}
}

// changeSpeed returns new PCM bytes from the bytes with the speed and pitch
// changed to the given value that must be in range [0, x].
func changeSpeed(a []byte, speed float64) []byte {
	b := make([]byte, int(math.Floor(float64(len(a))*speed)))
	var p float64
	for _, v := range a {
		for i := int(p); i < int(p+speed); i++ {
			b[i] = v
		}
		p += speed
	}
	return b
}

func makeSilence(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = 128
	}
	return b
}

func reversedSound(a []byte) []byte {
	n := len(a)
	b := make([]byte, n)
	for i, v := range a {
		b[n-1-i] = v
	}
	return b
}

// WriteTo writes captcha audio in WAVE format into the given io.Writer, and
// returns the number of bytes written and an error if any.
func (a *Audio) WriteTo(w io.Writer) (n int64, err error) {
	// Calculate padded length of PCM chunk data.
	bodyLen := uint32(a.body.Len())
	paddedBodyLen := bodyLen
	if bodyLen%2 != 0 {
		paddedBodyLen++
	}
	totalLen := uint32(len(waveHeader)) - 4 + paddedBodyLen
	// Header.
	header := make([]byte, len(waveHeader)+4) // includes 4 bytes for chunk size
	copy(header, waveHeader)
	// Put the length of whole RIFF chunk.
	binary.LittleEndian.PutUint32(header[4:], totalLen)
	// Put the length of WAVE chunk.
	binary.LittleEndian.PutUint32(header[len(waveHeader):], bodyLen)
	// Write header.
	nn, err := w.Write(header)
	n = int64(nn)
	if err != nil {
		return
	}
	// Write data.
	n, err = a.body.WriteTo(w)
	n += int64(nn)
	if err != nil {
		return
	}
	// Pad byte if chunk length is odd.
	// (As header has even length, we can check if n is odd, not chunk).
	if bodyLen != paddedBodyLen {
		w.Write([]byte{0})
		n++
	}
	return
}

// BinaryEncoding encodes an sound to wave and returns
// the result as a byte slice.
func (a *Audio) BinaryEncoding() []byte {
	var buf bytes.Buffer
	if _, err := a.WriteTo(&buf); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}
