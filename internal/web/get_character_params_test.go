package web

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ropfoo/gothulhu/internal/model"
)

func TestGetCharacterParams(t *testing.T) {
	tests := []struct {
		name        string
		queryParams url.Values
		formData    url.Values
		want        model.CharacterParams
		wantErr     bool
	}{
		{
			name: "valid query parameters",
			queryParams: url.Values{
				"name": []string{"John"},
				"age":  []string{"25"},
				"str":  []string{"14"},
				"dex":  []string{"12"},
				"con":  []string{"13"},
				"siz":  []string{"11"},
				"intl": []string{"15"},
				"wis":  []string{"10"},
				"cha":  []string{"16"},
			},
			want: model.CharacterParams{
				Name: "John",
				Age:  25,
				Stats: model.Stats{
					STR: 14,
					DEX: 12,
					CON: 13,
					SIZ: 11,
					INT: 15,
					WIS: 10,
					CHA: 16,
				},
			},
			wantErr: false,
		},
		{
			name: "valid form data",
			formData: url.Values{
				"name": []string{"Jane"},
				"age":  []string{"30"},
				"str":  []string{"16"},
				"dex":  []string{"15"},
				"con":  []string{"14"},
				"siz":  []string{"12"},
				"intl": []string{"13"},
				"wis":  []string{"11"},
				"cha":  []string{"10"},
			},
			want: model.CharacterParams{
				Name: "Jane",
				Age:  30,
				Stats: model.Stats{
					STR: 16,
					DEX: 15,
					CON: 14,
					SIZ: 12,
					INT: 13,
					WIS: 11,
					CHA: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid age",
			queryParams: url.Values{
				"name": []string{"Invalid"},
				"age":  []string{"not_a_number"},
				"str":  []string{"10"},
				"dex":  []string{"10"},
				"con":  []string{"10"},
				"siz":  []string{"10"},
				"intl": []string{"10"},
				"wis":  []string{"10"},
				"cha":  []string{"10"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request with query parameters
			reqURL := "http://example.com"
			if len(tt.queryParams) > 0 {
				reqURL += "?" + tt.queryParams.Encode()
			}

			req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(tt.formData.Encode()))
			if err != nil {
				t.Fatal(err)
			}

			if len(tt.formData) > 0 {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

			got, err := GetCharacterParams(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCharacterParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.Name != tt.want.Name {
					t.Errorf("GetCharacterParams() got name = %v, want %v", got.Name, tt.want.Name)
				}
				if got.Age != tt.want.Age {
					t.Errorf("GetCharacterParams() got age = %v, want %v", got.Age, tt.want.Age)
				}
				if got.Stats != tt.want.Stats {
					t.Errorf("GetCharacterParams() got stats = %v, want %v", got.Stats, tt.want.Stats)
				}
			}
		})
	}
}
