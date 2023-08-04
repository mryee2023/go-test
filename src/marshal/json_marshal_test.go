package marshal

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson_MarshalJSON(t *testing.T) {
	type fields struct {
		Name  string
		Age   int
		Phone string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Test_MarshalJSON",
			fields: fields{
				Name:  "Github",
				Age:   20,
				Phone: "13800138000",
			},
			want:    []byte(`{"name":"Github","age":20,"phone":"***"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Person{
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Phone: tt.fields.Phone,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_MarshalJSONWithValue(t *testing.T) {
	type fields struct {
		Name  string
		Age   int
		Phone string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Test_MarshalJSON",
			fields: fields{
				Name:  "Github",
				Age:   20,
				Phone: "13800138000",
			},
			want: `{"name":"Github","age":20,"phone":"13800138000"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Phone: tt.fields.Phone,
			}
			got, err := json.Marshal(p)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, string(got), "MarshalJSON() = %v, want %v", string(got), tt.want)
		})
	}
}
