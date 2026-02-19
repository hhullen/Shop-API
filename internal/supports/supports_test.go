package supports

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type Test struct {
	Field        []byte `fieldKey:"fieldValue"`
	AnotherField int64  `anotherFieldKey:"anotherFieldValue"`
}

func TestIsInContainer(t *testing.T) {
	t.Parallel()

	require.Equal(t, IsInContainer(), false)
}

func TestReadSecret(t *testing.T) {
	t.Parallel()

	secret, err := ReadSecret("./supports_test.go")
	require.Nil(t, err)
	require.NotEqual(t, secret, "")
}

func TestMakeKVMessagesJSON(t *testing.T) {
	t.Parallel()

	b, _ := MakeKVMessagesJSON("key1", "val1", "key2", "val2", "KeyNoVal")
	data := map[string]string{}
	err := json.Unmarshal(b, &data)
	require.Nil(t, err)

	require.Equal(t, data["key1"], "val1")
	require.Equal(t, data["key2"], "val2")
	_, exists := data["KeyNoVal"]
	require.False(t, exists)
}

func TestGetUUIDIfEmpty(t *testing.T) {
	t.Parallel()

	var uid uuid.UUID
	newUid := GetUUIDIfEmpty(uid)
	require.NotEqual(t, uid, newUid)
	require.NotEqual(t, newUid, uuid.Nil)
	require.Equal(t, uid, uuid.Nil)
}

func TestGetNowIfZero(t *testing.T) {
	t.Parallel()

	var tt time.Time
	newTt := GetNowIfZero(tt)
	require.NotEqual(t, tt, newTt)
	require.True(t, tt.IsZero())
	require.False(t, newTt.IsZero())
}

func TestValidatePhoneNumber(t *testing.T) {
	t.Parallel()

	err := ValidatePhoneNumber("+79634567733")
	require.Nil(t, err)

	err = ValidatePhoneNumber("+79634567733 RU")
	require.Nil(t, err)

	err = ValidatePhoneNumber("+17182222222 US")
	require.Nil(t, err)

	err = ValidatePhoneNumber("+17182222222")
	require.NotNil(t, err)

	err = ValidatePhoneNumber("0934 RU")
	require.NotNil(t, err)

	err = ValidatePhoneNumber("someShit EN")
	require.NotNil(t, err)
}

func TestGetStructFieldByTagKey(t *testing.T) {
	t.Parallel()

	tt := &Test{Field: []byte("Hey MF!")}

	val, field, err := GetStructFieldByTagKey(tt, "fieldKey")

	require.Nil(t, err)
	require.Equal(t, val, "fieldValue")
	newVal := "A new ones"
	field.SetBytes([]byte(newVal))

	require.Equal(t, []byte(newVal), tt.Field)

	_, _, err = GetStructFieldByTagKey(tt, "noSuchKey")

	require.NotNil(t, err)

}

func TestIsFieldByteSlice(t *testing.T) {
	t.Parallel()

	tt := &Test{Field: []byte("Hey MF!")}

	_, field, err := GetStructFieldByTagKey(tt, "fieldKey")
	require.Nil(t, err)
	require.True(t, IsFieldByteSlice(field))

	_, field, err = GetStructFieldByTagKey(tt, "anotherFieldKey")
	require.Nil(t, err)
	require.False(t, IsFieldByteSlice(field))
}

func TestGetDateAsFileName(t *testing.T) {
	t.Parallel()

	dd := GetDateAsFileName(time.Now())
	require.NotEmpty(t, dd)
	require.False(t, strings.Contains(dd, ":"))
	require.False(t, strings.Contains(dd, " "))
}
