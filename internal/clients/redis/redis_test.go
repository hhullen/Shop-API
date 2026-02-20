package redis

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	go_redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

type TestValue struct {
	Value string `json:"value"`
}

type TestClient struct {
	ctx       context.Context
	redisMock *MockUniversalClient
	client    *Client
	testVal   TestValue
}

func NewTestClient(t *testing.T) *TestClient {
	mc := gomock.NewController(t)
	tc := &TestClient{
		ctx:       context.Background(),
		redisMock: NewMockUniversalClient(mc),
	}

	c := buildClient(tc.redisMock)
	tc.client = c

	return tc
}

func TestRead(t *testing.T) {
	t.Parallel()

	t.Run("Read ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		scmd := go_redis.NewStringCmd(tc.ctx)
		v := &TestValue{
			Value: "test_value",
		}
		data, _ := json.Marshal(v)
		scmd.SetVal(string(data))

		tc.redisMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(scmd)

		newVal := &TestValue{}

		ok, err := tc.client.Read("key", &newVal)
		require.Nil(t, err)
		require.True(t, ok)
		require.Equal(t, newVal.Value, v.Value)
	})

	t.Run("Read error on Get", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)
		scmd := go_redis.NewStringCmd(tc.ctx)
		scmd.SetErr(errors.New("error"))

		tc.redisMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(scmd)

		newVal := &TestValue{}

		ok, err := tc.client.Read("key", newVal)
		require.NotNil(t, err)
		require.False(t, ok)
	})
}

func TestWrite(t *testing.T) {
	t.Parallel()

	t.Run("Write ok", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		scmd := go_redis.NewStatusCmd(tc.ctx)

		tc.redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(scmd)

		v := &TestValue{}

		err := tc.client.Write("key", v)
		require.Nil(t, err)
	})

	t.Run("Write error on Set", func(t *testing.T) {
		t.Parallel()

		tc := NewTestClient(t)

		scmd := go_redis.NewStatusCmd(tc.ctx)
		scmd.SetErr(errors.New("error"))

		tc.redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(scmd)

		v := &TestValue{}

		err := tc.client.Write("key", v)
		require.NotNil(t, err)
	})
}
