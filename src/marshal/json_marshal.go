package marshal

import (
	"encoding/json"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func (p *Person) MarshalJSON() ([]byte, error) {
	type Alias Person
	return json.Marshal(&struct {
		*Alias
		Phone string `json:"phone,omitempty"`
	}{
		Alias: (*Alias)(p),
		Phone: "***",
	})
}
