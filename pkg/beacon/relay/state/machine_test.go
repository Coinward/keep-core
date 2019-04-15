package state

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/keep-network/keep-core/pkg/beacon/relay/group"
	"github.com/keep-network/keep-core/pkg/chain"
	chainLocal "github.com/keep-network/keep-core/pkg/chain/local"
	"github.com/keep-network/keep-core/pkg/net"
	netLocal "github.com/keep-network/keep-core/pkg/net/local"
)

var testLog map[uint64][]string
var blockCounter chain.BlockCounter

func TestExecute(t *testing.T) {
	testLog = make(map[uint64][]string)

	localChain := chainLocal.Connect(10, 5, big.NewInt(200))
	blockCounter, _ = localChain.BlockCounter()
	provider := netLocal.Connect()
	channel, err := provider.ChannelFor("transitions_test")
	if err != nil {
		t.Fatal(err)
	}

	go func(blockCounter chain.BlockCounter) {
		blockCounter.WaitForBlockHeight(1)
		channel.Send(&TestMessage{"message_1"})

		blockCounter.WaitForBlockHeight(5)
		channel.Send(&TestMessage{"message_2"})

		blockCounter.WaitForBlockHeight(9)
		channel.Send(&TestMessage{"message_3"})
	}(blockCounter)

	channel.RegisterUnmarshaler(func() net.TaggedUnmarshaler {
		return &TestMessage{}
	})

	initialState := testState1{
		memberIndex: group.MemberIndex(1),
		channel:     channel,
	}

	stateMachine := NewMachine(channel, blockCounter, initialState)

	finalState, err := stateMachine.Execute()
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

	if _, ok := finalState.(*testState4); !ok {
		t.Errorf("state is not final [%v]", finalState)
	}

	expectedTestLog := map[uint64][]string{
		1: []string{
			"1-state.testState1-initiate",
			"1-state.testState1-receive-message_1",
		},
		4: []string{"1-state.testState2-initiate"},
		5: []string{"1-state.testState2-receive-message_2"},
		7: []string{"1-state.testState3-initiate"},
		10: []string{
			"1-state.testState4-initiate",
			"1-state.testState4-receive-message_3",
		},
	}

	if !reflect.DeepEqual(expectedTestLog, testLog) {
		t.Errorf("\nexpected: %v\nactual:   %v\n", expectedTestLog, testLog)
	}
}

func addToTestLog(testState State, functionName string) {
	currentBlock, _ := blockCounter.CurrentBlock()
	testLog[currentBlock] = append(
		testLog[currentBlock],
		fmt.Sprintf(
			"%v-%v-%v",
			testState.MemberIndex(),
			reflect.TypeOf(testState),
			functionName,
		),
	)
}

type testState1 struct {
	memberIndex group.MemberIndex
	channel     net.BroadcastChannel
}

func (ts testState1) ActiveBlocks() uint64 { return 2 }
func (ts testState1) Initiate() error {
	addToTestLog(ts, "initiate")
	return nil
}
func (ts testState1) Receive(msg net.Message) error {
	addToTestLog(
		ts,
		fmt.Sprintf("receive-%v", msg.Payload().(*TestMessage).content),
	)
	return nil
}
func (ts testState1) Next() State                    { return &testState2{ts} }
func (ts testState1) MemberIndex() group.MemberIndex { return ts.memberIndex }

type testState2 struct {
	testState1
}

func (ts testState2) ActiveBlocks() uint64 { return 2 }
func (ts testState2) Initiate() error {
	addToTestLog(ts, "initiate")
	return nil
}
func (ts testState2) Receive(msg net.Message) error {
	addToTestLog(
		ts,
		fmt.Sprintf("receive-%v", msg.Payload().(*TestMessage).content),
	)
	return nil
}
func (ts testState2) Next() State                    { return &testState3{ts} }
func (ts testState2) MemberIndex() group.MemberIndex { return ts.memberIndex }

type testState3 struct {
	testState2
}

func (ts testState3) ActiveBlocks() uint64 { return 2 }

func (ts testState3) Initiate() error {
	addToTestLog(ts, "initiate")
	return nil
}

func (ts testState3) Receive(msg net.Message) error {
	addToTestLog(
		ts,
		fmt.Sprintf("receive-%v", msg.Payload().(*TestMessage).content),
	)
	return nil
}

func (ts testState3) Next() State                    { return &testState4{ts} }
func (ts testState3) MemberIndex() group.MemberIndex { return ts.memberIndex }

type testState4 struct {
	testState3
}

func (ts testState4) ActiveBlocks() uint64 { return 2 }
func (ts testState4) Initiate() error {
	addToTestLog(ts, "initiate")
	return nil
}
func (ts testState4) Receive(msg net.Message) error {
	addToTestLog(
		ts,
		fmt.Sprintf("receive-%v", msg.Payload().(*TestMessage).content),
	)
	return nil
}
func (ts testState4) Next() State                    { return nil }
func (ts testState4) MemberIndex() group.MemberIndex { return ts.memberIndex }

type TestMessage struct {
	content string
}

func (tm *TestMessage) Marshal() ([]byte, error) {
	return []byte(tm.content), nil
}

func (tm *TestMessage) Unmarshal(bytes []byte) error {
	tm.content = string(bytes)
	return nil
}

func (tm *TestMessage) Type() string {
	return "test_message"
}
