package supports

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nyaruka/phonenumbers"
)

const (
	defaultPhoneRegion = "RU"
)

var validatorInstance validator.Validate = *validator.New()

func StructValidator() *validator.Validate {
	return &validatorInstance
}

func MakeKVMessagesJSON(kvs ...any) (bytes []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("failed MakeKVMessagesJSON: %v", p)
		}
	}()

	msgs := map[string]any{}
	for i := 0; i < len(kvs)-1; i += 2 {
		key := fmt.Sprint(kvs[i])
		value := kvs[i+1]
		msgs[key] = value
	}

	bytes, err = json.Marshal(msgs)
	return
}

func ReadSecret(path string) (string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}

	secret := ""
	_, err = fmt.Fscan(f, &secret)
	if err != nil {
		return "", err
	}

	return string(secret), nil
}

func IsInContainer() bool {
	return os.Getenv("RUNNING_IN_CONTAINER") == "true"
}

func GetUUIDIfEmpty(uid uuid.UUID) uuid.UUID {
	if uid != uuid.Nil {
		return uid
	}
	return uuid.New()

}

func GetNowIfZero(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}
	return t
}

func ValidatePhoneNumber(pn string) error {
	phoneRegion := defaultPhoneRegion
	if part := strings.Split(pn, " "); len(part) == 2 {
		phoneRegion = part[1]
		pn = part[0]
	}

	p, err := phonenumbers.Parse(pn, phoneRegion)
	if err != nil {
		return err
	}

	if !phonenumbers.IsValidNumberForRegion(p, phoneRegion) {
		return fmt.Errorf("Invalid number for region: %s", phoneRegion)
	}

	return nil
}

func GetStructFieldByTagKey(v any, tagKey string) (string, reflect.Value, error) {
	value := reflect.ValueOf(v)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return runBFSFieldSearch(value, tagKey)

}

func runBFSFieldSearch(value reflect.Value, tagKey string) (string, reflect.Value, error) {
	if value.Kind() != reflect.Struct {
		return "", reflect.Value{}, fmt.Errorf("expected struct but got '%s'", value.Kind().String())
	}

	vType := value.Type()

	for i := range value.NumField() {
		fieldType := vType.Field(i)

		tag := fieldType.Tag.Get(tagKey)
		if tag != "" {
			return tag, value.Field(i), nil
		}
	}

	for i := range value.NumField() {
		fieldValue := value.Field(i)
		fieldType := vType.Field(i)

		if fieldType.Anonymous && fieldValue.Kind() == reflect.Struct {
			tag, foundField, err := runBFSFieldSearch(fieldValue, tagKey)
			if err == nil {
				return tag, foundField, nil
			}
		}

	}

	return "", reflect.Value{}, fmt.Errorf(
		"field with tag key'%s' not found in %s", tagKey, vType.Name())
}

func IsFieldByteSlice(field reflect.Value) bool {
	return field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Uint8
}

func GetDateAsFileName(t time.Time) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			t.Format(time.DateTime), " ", "_"),
		":", "-")
}

func GetHash(data []byte) string {
	h := fnv.New64a()
	h.Write(data)
	return strconv.FormatUint(h.Sum64(), 10)
}

func Concat(ss ...string) string {
	length := 0
	for i := range ss {
		length += len(ss[i])
	}

	var b strings.Builder
	b.Grow(length)

	for i := range ss {
		b.WriteString(ss[i])
	}

	return b.String()
}
