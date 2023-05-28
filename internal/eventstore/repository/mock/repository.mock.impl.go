package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/repository"
)

type MockRepository struct {
	*MockPusher
	*MockQuerier
}

func NewRepo(t *testing.T) *MockRepository {
	controller := gomock.NewController(t)
	return &MockRepository{
		MockPusher:  NewMockPusher(controller),
		MockQuerier: NewMockQuerier(controller),
	}
}

func (m *MockRepository) ExpectFilterNoEventsNoError() *MockRepository {
	m.MockQuerier.ctrl.T.Helper()

	m.MockQuerier.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(nil, nil)
	return m
}

func (m *MockRepository) ExpectFilterEvents(events ...*repository.Event) *MockRepository {
	m.MockQuerier.ctrl.T.Helper()

	m.MockQuerier.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(events, nil)
	return m
}

func (m *MockRepository) ExpectFilterEventsError(err error) *MockRepository {
	m.MockQuerier.ctrl.T.Helper()

	m.MockQuerier.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(nil, err)
	return m
}

func (m *MockRepository) ExpectInstanceIDs(hasFilters []*repository.Filter, instanceIDs ...string) *MockRepository {
	m.MockQuerier.ctrl.T.Helper()

	matcher := gomock.Any()
	if len(hasFilters) > 0 {
		matcher = &filterQueryMatcher{Filters: [][]*repository.Filter{hasFilters}}
	}
	m.MockQuerier.EXPECT().InstanceIDs(gomock.Any(), matcher).Return(instanceIDs, nil)
	return m
}

func (m *MockRepository) ExpectInstanceIDsError(err error) *MockRepository {
	m.MockQuerier.ctrl.T.Helper()

	m.MockQuerier.EXPECT().InstanceIDs(gomock.Any(), gomock.Any()).Return(nil, err)
	return m
}

func (m *MockRepository) ExpectPush(expectedCommands []eventstore.Command) *MockRepository {
	m.MockPusher.EXPECT().Push(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, commands ...eventstore.Command) ([]eventstore.Event, error) {
			m.MockPusher.ctrl.T.Helper()

			if len(expectedCommands) != len(commands) {
				return nil, fmt.Errorf("unexpected amount of commands: want %d, got %d", len(expectedCommands), len(commands))
			}
			for i, expectedCommand := range expectedCommands {
				if !assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Aggregate(), commands[i].Aggregate()) {
					m.MockPusher.ctrl.T.Errorf("invalid command.Aggregate [%d]: expected: %#v got: %#v", i, expectedCommand.Aggregate(), commands[i].Aggregate())
				}
				if !assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Creator(), commands[i].Creator()) {
					m.MockPusher.ctrl.T.Errorf("invalid command.Creator [%d]: expected: %#v got: %#v", i, expectedCommand.Creator(), commands[i].Creator())
				}
				if !assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Type(), commands[i].Type()) {
					m.MockPusher.ctrl.T.Errorf("invalid command.Type [%d]: expected: %#v got: %#v", i, expectedCommand.Type(), commands[i].Type())
				}
				if !assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Revision(), commands[i].Revision()) {
					m.MockPusher.ctrl.T.Errorf("invalid command.Revision [%d]: expected: %#v got: %#v", i, expectedCommand.Revision(), commands[i].Revision())
				}
				if !assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Payload(), commands[i].Payload()) {
					m.MockPusher.ctrl.T.Errorf("invalid command.Payload [%d]: expected: %#v got: %#v", i, expectedCommand.Payload(), commands[i].Payload())
				}
				if !assert.ElementsMatch(m.MockPusher.ctrl.T, expectedCommand.UniqueConstraints(), commands[i].UniqueConstraints()) {
					m.MockPusher.ctrl.T.Errorf("invalid command.UniqueConstraints [%d]: expected: %#v got: %#v", i, expectedCommand.UniqueConstraints(), commands[i].UniqueConstraints())
				}
			}
			// assert.ElementsMatch(m.MockPusher.ctrl.T, expectedCommands, commands)
			// if expectedUniqueConstraints == nil {
			// 	expectedUniqueConstraints = []*repository.UniqueConstraint{}
			// }
			// assert.Equal(m.MockPusher.ctrl.T, expectedUniqueConstraints, uniqueConstraints)
			events := make([]eventstore.Event, len(commands))
			for i, command := range commands {
				events[i] = &mockEvent{
					Command: command,
				}
			}
			return events, nil
		},
	)
	return m
}

func (m *MockRepository) ExpectPushFailed(err error, expectedCommands []eventstore.Command) *MockRepository {
	m.MockPusher.ctrl.T.Helper()

	m.MockPusher.EXPECT().Push(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, commands ...eventstore.Command) ([]eventstore.Event, error) {
			if len(expectedCommands) != len(commands) {
				return nil, fmt.Errorf("unexpected amount of commands: want %d, got %d", len(expectedCommands), len(commands))
			}
			for i, expectedCommand := range expectedCommands {
				assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Aggregate(), commands[i].Aggregate())
				assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Creator(), commands[i].Creator())
				assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Type(), commands[i].Type())
				assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Revision(), commands[i].Revision())
				assert.Equal(m.MockPusher.ctrl.T, expectedCommand.Payload(), commands[i].Payload())
				assert.ElementsMatch(m.MockPusher.ctrl.T, expectedCommand.UniqueConstraints(), commands[i].UniqueConstraints())
			}

			return nil, err
		},
	)
	return m
}

type mockEvent struct {
	eventstore.Command
	sequence  uint64
	createdAt time.Time
}

// DataAsBytes implements eventstore.Event
func (e *mockEvent) DataAsBytes() []byte {
	if e.Payload() == nil {
		return nil
	}
	payload, err := json.Marshal(e.Payload())
	if err != nil {
		panic("unable to unmarshal")
	}
	return payload
}

func (e *mockEvent) Unmarshal(ptr any) error {
	reflect.ValueOf(ptr).Set(reflect.ValueOf(e.Command.Payload()))
	return nil
}

func (e *mockEvent) Sequence() uint64 {
	return e.sequence
}

func (e *mockEvent) CreatedAt() time.Time {
	return e.createdAt
}
