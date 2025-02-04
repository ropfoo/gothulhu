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
				"name":   []string{"John"},
				"age":    []string{"25"},
				"gender": []string{"male"},
				"str":    []string{"14"},
				"dex":    []string{"12"},
				"con":    []string{"13"},
				"siz":    []string{"11"},
				"intl":   []string{"15"},
				"wis":    []string{"10"},
				"cha":    []string{"16"},
			},
			want: model.CharacterParams{
				Name:   "John",
				Age:    25,
				Gender: "male",
				Stats: model.Stats{
					STR: []float32{14, 7, 2.8},
					DEX: []float32{12, 6, 2.4},
					CON: []float32{13, 6.5, 2.6},
					SIZ: []float32{11, 5.5, 2.2},
					INT: []float32{15, 7.5, 3},
					WIS: []float32{10, 5, 2},
					CHA: []float32{16, 8, 3.2},
				},
			},
			wantErr: false,
		},
		{
			name: "valid form data",
			formData: url.Values{
				"name":   []string{"Jane"},
				"age":    []string{"30"},
				"gender": []string{"female"},
				"str":    []string{"16"},
				"dex":    []string{"15"},
				"con":    []string{"14"},
				"siz":    []string{"12"},
				"intl":   []string{"13"},
				"wis":    []string{"11"},
				"cha":    []string{"10"},
			},
			want: model.CharacterParams{
				Name:   "Jane",
				Age:    30,
				Gender: "female",
				Stats: model.Stats{
					STR: []float32{16, 8, 3.2},
					DEX: []float32{15, 7.5, 3},
					CON: []float32{14, 7, 2.8},
					SIZ: []float32{12, 6, 2.4},
					INT: []float32{13, 6.5, 2.6},
					WIS: []float32{11, 5.5, 2.2},
					CHA: []float32{10, 5, 2},
				},
			},
			wantErr: false,
		},
		{
			name: "missing optional parameters should use defaults",
			queryParams: url.Values{
				"name": []string{"John"},
				"str":  []string{"14"},
				"dex":  []string{"12"},
			},
			want: model.CharacterParams{
				Name:   "John",
				Age:    0,
				Gender: "",
				Stats: model.Stats{
					STR: []float32{14, 7, 2.8},
					DEX: []float32{12, 6, 2.4},
					CON: []float32{0, 0, 0},
					SIZ: []float32{0, 0, 0},
					INT: []float32{0, 0, 0},
					WIS: []float32{0, 0, 0},
					CHA: []float32{0, 0, 0},
				},
			},
			wantErr: false,
		},
		{
			name: "missing name",
			queryParams: url.Values{
				"age":    []string{"25"},
				"gender": []string{"male"},
				"str":    []string{"14"},
			},
			want: model.CharacterParams{
				Age:    25,
				Gender: "male",
				Stats: model.Stats{
					STR: []float32{14, 7, 2.8},
					DEX: []float32{0, 0, 0},
					CON: []float32{0, 0, 0},
					SIZ: []float32{0, 0, 0},
					INT: []float32{0, 0, 0},
					WIS: []float32{0, 0, 0},
					CHA: []float32{0, 0, 0},
				},
			},
			wantErr: false,
		},
		{
			name: "with invalid age",
			queryParams: url.Values{
				"name":   []string{"Invalid"},
				"age":    []string{"not_a_number"},
				"gender": []string{"male"},
				"str":    []string{"10"},
				"dex":    []string{"10"},
				"con":    []string{"10"},
				"siz":    []string{"10"},
				"intl":   []string{"10"},
				"wis":    []string{"10"},
				"cha":    []string{"10"},
			},
			want: model.CharacterParams{
				Name:   "Invalid",
				Age:    0,
				Gender: "male",
				Stats: model.Stats{
					STR: []float32{10, 5, 2},
					DEX: []float32{10, 5, 2},
					CON: []float32{10, 5, 2},
					SIZ: []float32{10, 5, 2},
					INT: []float32{10, 5, 2},
					WIS: []float32{10, 5, 2},
					CHA: []float32{10, 5, 2},
				},
			},
			wantErr: false,
		},
		{
			name: "with invalid gender",
			queryParams: url.Values{
				"name":   []string{"Invalid"},
				"age":    []string{"25"},
				"gender": []string{"invalid_gender"},
				"str":    []string{"10"},
				"dex":    []string{"10"},
				"con":    []string{"10"},
				"siz":    []string{"10"},
				"intl":   []string{"10"},
				"wis":    []string{"10"},
				"cha":    []string{"10"},
			},
			want: model.CharacterParams{
				Name:   "Invalid",
				Age:    25,
				Gender: "invalid_gender",
				Stats: model.Stats{
					STR: []float32{10, 5, 2},
					DEX: []float32{10, 5, 2},
					CON: []float32{10, 5, 2},
					SIZ: []float32{10, 5, 2},
					INT: []float32{10, 5, 2},
					WIS: []float32{10, 5, 2},
					CHA: []float32{10, 5, 2},
				},
			},
			wantErr: false,
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
				if got.Gender != tt.want.Gender {
					t.Errorf("GetCharacterParams() got gender = %v, want %v", got.Gender, tt.want.Gender)
				}
				if !compareStats(got.Stats, tt.want.Stats) {
					t.Errorf("GetCharacterParams() got stats = %v, want %v", got.Stats, tt.want.Stats)
				}
			}
		})
	}
}

// Add helper function to compare Stats
func compareStats(a, b model.Stats) bool {
	return slicesEqual(a.STR, b.STR) &&
		slicesEqual(a.DEX, b.DEX) &&
		slicesEqual(a.CON, b.CON) &&
		slicesEqual(a.SIZ, b.SIZ) &&
		slicesEqual(a.INT, b.INT) &&
		slicesEqual(a.WIS, b.WIS) &&
		slicesEqual(a.CHA, b.CHA)
}

func slicesEqual(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
