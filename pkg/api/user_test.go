package api

import (
	"errors"
	"reflect"
	"testing"
)

type mockUserRepo struct {
}

func (m mockUserRepo) CreateUser(req NewUserRequest) error {
	if req.Name == "test user already created" {
		return errors.New("user already exists in db")
	}
	return nil
}

func TestCreateNewUser(t *testing.T) {
	mockrepo := mockUserRepo{}
	mockUserService := NewUserService(&mockrepo)

	tests := []struct {
		name    string
		request NewUserRequest
		want    error
	}{
		{
			name: "should create new user successfully",
			request: NewUserRequest{
				Name:          "test user",
				Age:           20,
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				WeightGoal:    "maintain",
				Email:         "test_user@gmail.com",
			},
			want: nil,
		},
		{
			name: "should return an error missing email",
			request: NewUserRequest{
				Name:          "test user",
				Age:           20,
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				WeightGoal:    "maintain",
				Email:         "",
			},
			want: errors.New("email required"),
		},
		{
			name: "should an error missing name",
			request: NewUserRequest{
				Name:          "",
				Age:           20,
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				WeightGoal:    "maintain",
				Email:         "test_user@gmail.com",
			},
			want: errors.New("name required"),
		},
		{
			name: "should an error user already exist in db",
			request: NewUserRequest{
				Name:          "test user already created",
				Age:           20,
				Height:        180,
				Sex:           "female",
				ActivityLevel: 5,
				WeightGoal:    "maintain",
				Email:         "test_user@gmail.com",
			},
			want: errors.New("user already exists in db"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockUserService.New(test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}

}
