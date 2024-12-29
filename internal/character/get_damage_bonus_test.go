package character

import (
	"testing"

	"github.com/ropfoo/gothulhu/internal/model"
)

func TestGetDamageBonus(t *testing.T) {
	tests := []struct {
		name     string
		str      int
		siz      int
		expected model.DamageBonus
	}{
		{
			name:     "Very small str+siz (<=64)",
			str:      30,
			siz:      30,
			expected: model.DamageBonus{Damage: "-2", Stature: "-2"},
		},
		{
			name:     "Small str+siz (65-84)",
			str:      40,
			siz:      40,
			expected: model.DamageBonus{Damage: "-1", Stature: "-1"},
		},
		{
			name:     "Average str+siz (85-124)",
			str:      50,
			siz:      50,
			expected: model.DamageBonus{Damage: "0", Stature: "0"},
		},
		{
			name:     "Above average str+siz (125-164)",
			str:      70,
			siz:      70,
			expected: model.DamageBonus{Damage: "1d4", Stature: "+1"},
		},
		{
			name:     "High str+siz (165-204)",
			str:      90,
			siz:      90,
			expected: model.DamageBonus{Damage: "1d6", Stature: "+2"},
		},
		{
			name:     "Very high str+siz (205-284)",
			str:      120,
			siz:      120,
			expected: model.DamageBonus{Damage: "2d6", Stature: "+3"},
		},
		{
			name:     "Extremely high str+siz (285-364)",
			str:      150,
			siz:      150,
			expected: model.DamageBonus{Damage: "3d6", Stature: "+4"},
		},
		{
			name:     "Massive str+siz (365-444)",
			str:      200,
			siz:      200,
			expected: model.DamageBonus{Damage: "4d6", Stature: "+5"},
		},
		{
			name:     "Maximum str+siz (>444)",
			str:      250,
			siz:      250,
			expected: model.DamageBonus{Damage: "5d6", Stature: "+6"},
		},
		{
			name:     "Edge case - boundary value 124/125",
			str:      62,
			siz:      62,
			expected: model.DamageBonus{Damage: "0", Stature: "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDamageBonus(tt.str, tt.siz)
			if got != tt.expected {
				t.Errorf("getDamageBonus(%d, %d) = %v, want %v",
					tt.str, tt.siz, got, tt.expected)
			}
		})
	}
}
