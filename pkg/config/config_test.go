package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	text := `
aggregates:
  Car:
    id:
      name: id
      type: string
    properties:
      BaseProperty:
        type: PersonalName
        defaultValue:
          FirstName: JOHN
          LastName: DOE
      Description:
        type: string
      Wheels:
        itemType: Wheel
      Tires:
        itemType: Tire
`

	config, err := NewConfigWithByte([]byte(text))
	if err != nil {
		t.Error(err)
	} else {
		println(config)
	}
}

func TestNewConfigWithDir(t *testing.T) {
	config, err := NewConfigWithDir("./DDDML")
	if err != nil {
		t.Error(err)
	} else {
		println(config)
	}
}
