package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallback(t *testing.T) {
	connMap := make(map[ConnKey]*ConnectionStats)
	connKey := ConnKey{
		SrcIP:   "10.10.10.10",
		SrcPort: 40040,
		DstIP:   "10.10.10.23",
		DstPort: 80,
	}
	statesResource := createStatesResource()
	connStats := &ConnectionStats{
		Pid:              0,
		Comm:             "test",
		ConnKey:          connKey,
		InitialTimestamp: 0,
		EndTimestamp:     0,
		Code:             0,
	}
	connStats.StateMachine = NewStateMachine(Inprogress, statesResource, connStats)
	connMap[connKey] = connStats

	stats, err := connStats.StateMachine.ReceiveEvent(tcpConnectNoError, connMap)
	assert.NoError(t, err)
	assert.Equal(t, Inprogress, connStats.StateMachine.currentStateType)

	stats, err = connStats.StateMachine.ReceiveEvent(tcpSetStateToEstablished, connMap)
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, Success, connStats.StateMachine.currentStateType)

	connStats.StateMachine = NewStateMachine(Inprogress, statesResource, connStats)
	stats, err = connStats.StateMachine.ReceiveEvent(expiredEvent, connMap)
	assert.Equal(t, Failure, connStats.StateMachine.currentStateType)
	_, ok := connMap[connKey]
	assert.Equal(t, false, ok)
}
