package domain

import "testing"

func TestGetFullName(t *testing.T) {
	tests := []struct {
		user     User
		expected string
	}{
		{User{FirstName: "John", MiddleName: "Paul", LastName: "Doe"}, "John Paul Doe"},
		{User{FirstName: "Jane", MiddleName: "Mary", LastName: "Smith"}, "Jane Mary Smith"},
		{User{FirstName: "Mike", MiddleName: "Jordan", LastName: "Johnson"}, "Mike Jordan Johnson"},
	}

	for _, testCases := range tests {
		result := testCases.user.GetFullName()
		if result != testCases.expected {
			t.Errorf("Expected '%s', got '%s'", testCases.expected, result)
		}
	}
}
