package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		num1    float64
		num2    float64
		op      string
		want    float64
		wantErr bool
	}{
		{"add", 2, 3, "+", 5, false},
		{"subtract", 5, 3, "-", 2, false},
		{"multiply", 2, 3, "*", 6, false},
		{"divide", 6, 3, "/", 2, false},
		{"divide by zero", 6, 0, "/", 0, true},
		{"unknown operator", 6, 3, "%", 0, true},
		{"empty operator", 6, 3, "", 0, true},
		{"negative numbers", -2, -3, "+", -5, false},
		{"decimals", 1.5, 2.5, "+", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.num1, tt.num2, tt.op)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("calculate(%v, %v, %q) = %v, want an error", tt.num1, tt.num2, tt.op, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("calculate(%v, %v, %q) returned unexpected error: %v", tt.num1, tt.num2, tt.op, err)
			}
			if got != tt.want {
				t.Errorf("calculate(%v, %v, %q) = %v, want %v", tt.num1, tt.num2, tt.op, got, tt.want)
			}
		})
	}
}
