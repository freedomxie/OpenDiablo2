package d2input

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// HandlerEvent is an event that EventHandlers will process and respond to
type HandlerEvent struct {
	keyMod    d2enum.KeyMod
	buttonMod d2enum.MouseButtonMod
	x         int
	y         int
}

// KeyMod yields the modifier for a key action
func (e *HandlerEvent) KeyMod() d2enum.KeyMod {
	return e.keyMod
}

// ButtonMod yields the modifier for a button action
func (e *HandlerEvent) ButtonMod() d2enum.MouseButtonMod {
	return e.buttonMod
}

// X returns the x screen coordinate for the event
func (e *HandlerEvent) X() int {
	return e.x
}

//Y returns the y screen coordinate for the event
func (e *HandlerEvent) Y() int {
	return e.y
}

type KeyCharsEvent struct {
	HandlerEvent
	chars []rune
}

func (e *KeyCharsEvent) Chars() []rune {
	return e.chars
}

type KeyEvent struct {
	HandlerEvent
	key d2enum.Key
	// Duration represents the number of frames this key has been pressed for
	duration int
}

func (e *KeyEvent) Key() d2enum.Key {
	return e.key
}
func (e *KeyEvent) Duration() int {
	return e.duration
}

type MouseEvent struct {
	HandlerEvent
	mouseButton d2enum.MouseButton
}

func (e *MouseEvent) KeyMod() d2enum.KeyMod {
	return e.HandlerEvent.keyMod
}

func (e *MouseEvent) ButtonMod() d2enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

func (e *MouseEvent) X() int {
	return e.HandlerEvent.x
}

func (e *MouseEvent) Y() int {
	return e.HandlerEvent.y
}

func (e *MouseEvent) Button() d2enum.MouseButton {
	return e.mouseButton
}

type MouseMoveEvent struct {
	HandlerEvent
}

func (e *MouseMoveEvent) KeyMod() d2enum.KeyMod {
	return e.HandlerEvent.keyMod
}

func (e *MouseMoveEvent) ButtonMod() d2enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

func (e *MouseMoveEvent) X() int {
	return e.HandlerEvent.x
}

func (e *MouseMoveEvent) Y() int {
	return e.HandlerEvent.y
}
