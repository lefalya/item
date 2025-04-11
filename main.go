package item

import (
	"github.com/google/uuid"
	"math/rand"
	"reflect"
	"time"
)

const (
	RANDID_LENGTH = 16
)

type Blueprint interface {
	SetUUID()
	GetUUID() string
	SecureUUID()
	SetRandId()
	GetRandId() string
	SetCreatedAt(time time.Time)
	GetCreatedAt() time.Time
	SetUpdatedAt(time time.Time)
	GetUpdatedAt() time.Time
}

type Foundation struct {
	UUID      string    `json:"uuid,omitempty" bson:"uuid"`
	RandId    string    `json:"randid,omitempty" bson:"randid"`
	CreatedAt time.Time `json:"-" bson:"-"`
	UpdatedAt time.Time `json:"-" bson:"-"`
}

func (i *Foundation) SetUUID() {
	i.UUID = uuid.New().String()
}

func (i *Foundation) GetUUID() string {
	return i.UUID
}

func (i *Foundation) SecureUUID() {
	i.UUID = ""
}

func (i *Foundation) SetRandId() {
	i.RandId = RandId()
}

func (i *Foundation) GetRandId() string {
	return i.RandId
}

func (i *Foundation) SetCreatedAt(time time.Time) {
	i.CreatedAt = time
}

func (i *Foundation) SetUpdatedAt(time time.Time) {
	i.UpdatedAt = time
}

func (i *Foundation) GetCreatedAt() time.Time {
	return i.CreatedAt
}

func (i *Foundation) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}

func InitItem[T Blueprint](item T) {
	currentTime := time.Now().In(time.UTC)
	value := reflect.ValueOf(item).Elem()

	// Iterate through the fields of the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		// Check if the field is a pointer and is nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			// Allocate a new value for the pointer and set it
			field.Set(reflect.New(field.Type().Elem()))
		}
	}

	item.SetUUID()
	item.SetRandId()
	item.SetCreatedAt(currentTime)
	item.SetUpdatedAt(currentTime)
}

func RandId() string {
	// Define the characters that can be used in the random string
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Initialize an empty string to store the result
	result := make([]byte, RANDID_LENGTH)

	// Generate random characters for the string
	for i := 0; i < RANDID_LENGTH; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}

	return string(result)
}
