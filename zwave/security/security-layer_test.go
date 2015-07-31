package security

import (
	"testing"
	"time"

	"github.com/bjyoungblood/gozw/zwave/command-class"
	"github.com/stretchr/testify/assert"
)

func TestSecurityLayerNonceGeneration(t *testing.T) {
	securityLayer := NewSecurityLayer()
	nonce, err := securityLayer.GenerateInternalNonce()

	assert.NoError(t, err)
	assert.Len(t, nonce, 8)

	savedNonce, err := securityLayer.internalNonceTable.Get(nonce[0])

	assert.NoError(t, err)
	assert.EqualValues(t, savedNonce, nonce)
}

func TestSecurityLayerGetExternalNonce(t *testing.T) {
	securityLayer := NewSecurityLayer()

	nonce, err := securityLayer.GetExternalNonce(1)
	assert.Error(t, err)
	assert.Nil(t, nonce)

	receivedNonce := []byte{0x98, 0xe4, 0x1b, 0x30, 0x84, 0x33, 0xf4, 0x3f}

	securityLayer.ReceiveNonce(1, &commandclass.SecurityNonceReport{
		CommandClass: commandclass.CommandClassSecurity,
		Command:      commandclass.CommandSecurityNonceReport,
		Nonce:        receivedNonce,
	})

	nonce, err = securityLayer.GetExternalNonce(1)
	assert.NoError(t, err)
	assert.EqualValues(t, receivedNonce, nonce)
}

func TestSecurityLayerWaitForExternalNonce(t *testing.T) {
	securityLayer := NewSecurityLayer()

	done := make(chan bool)

	receivedNonce := []byte{0x98, 0xe4, 0x1b, 0x30, 0x84, 0x33, 0xf4, 0x3f}

	go func() {
		nonce, err := securityLayer.WaitForExternalNonce(1)
		assert.NoError(t, err)
		assert.EqualValues(t, receivedNonce, nonce)

		done <- true
	}()

	time.Sleep(time.Millisecond * 50)

	securityLayer.ReceiveNonce(1, &commandclass.SecurityNonceReport{
		CommandClass: commandclass.CommandClassSecurity,
		Command:      commandclass.CommandSecurityNonceReport,
		Nonce:        receivedNonce,
	})

	<-done
}
