package samsung

import (
	"time"
)

type keyAction int

const (
	keyActionClick   keyAction = 1
	keyActionPress   keyAction = 2
	keyActionRelease keyAction = 3
)

const (
	defaultKeyWaitDelay = 500 * time.Millisecond
)

type keyCommand struct {
	key    Key
	action keyAction
	wait   time.Duration
}

type KeySequence []keyCommand

func (k *KeySequence) Click(key Key) *KeySequence {
	*k = append(*k, keyCommand{
		key:    key,
		action: keyActionClick,
		wait:   defaultKeyWaitDelay,
	})

	return k
}

func (k *KeySequence) Press(key Key) *KeySequence {
	*k = append(*k, keyCommand{
		key:    key,
		action: keyActionPress,
		wait:   defaultKeyWaitDelay,
	})

	return k
}

func (k *KeySequence) Release(key Key) *KeySequence {
	*k = append(*k, keyCommand{
		key:    key,
		action: keyActionRelease,
		wait:   defaultKeyWaitDelay,
	})

	return k
}

func (k *KeySequence) Wait(wait time.Duration) *KeySequence {
	v := *k
	if len(v) == 0 {
		return k
	}

	index := len(v) - 1
	v[index].wait = wait

	return k
}

func (k *KeySequence) Repeat(n int) *KeySequence {
	v := *k
	if len(v) == 0 {
		return k
	}

	index := len(v) - 1
	element := v[index]
	for i := 0; i < n; i++ {
		*k = append(*k, element)
	}

	return k
}
