package option

import (
	"errors"
	"reflect"
	"testing"
)

func TestApply(t *testing.T) {
	testcases := []struct {
		name    string
		gotT    *testModel
		gotOpts []Option[*testModel]
		want    *testModel
		wantErr error
	}{
		{
			name: "pass: one option",
			gotT: &testModel{},
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					t.name = "foobar"
					return nil
				},
			},
			want: &testModel{
				name: "foobar",
			},
		},
		{
			name: "pass: multiple options",
			gotT: &testModel{},
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					t.name = "foo"
					return nil
				},
				func(t *testModel) error {
					t.age = 42
					return nil
				},
			},
			want: &testModel{
				name: "foo",
				age:  42,
			},
		},
		{
			name: "fail: first option returns err",
			gotT: &testModel{},
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					return errBadNameTest
				},
			},
			wantErr: errBadNameTest,
		},
		{
			name: "fail: second option returns err",
			gotT: &testModel{},
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					t.name = "foo"
					return nil
				},
				func(t *testModel) error {
					return errErrBadAgeTest
				},
			},
			wantErr: errErrBadAgeTest,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Apply(tc.gotT, tc.gotOpts...)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("got err: %v / want err: %v", err, tc.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("got: %v / want: %v", err, tc.wantErr)
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	testcases := []struct {
		name    string
		gotOpts []Option[*testModel]
		want    *testModel
		wantErr error
	}{
		{
			name:    "pass: zero option",
			gotOpts: []Option[*testModel]{},
			want:    &testModel{},
		},
		{
			name: "pass: one option",
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					t.name = "foobar"
					return nil
				},
			},
			want: &testModel{
				name: "foobar",
			},
		},
		{
			name: "pass: multiple options",
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					t.name = "foo"
					return nil
				},
				func(t *testModel) error {
					t.age = 42
					return nil
				},
			},
			want: &testModel{
				name: "foo",
				age:  42,
			},
		},
		{
			name: "fail: first option returns err",
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					return errBadNameTest
				},
			},
			wantErr: errBadNameTest,
		},
		{
			name: "fail: second option returns err",
			gotOpts: []Option[*testModel]{
				func(t *testModel) error {
					t.name = "foo"
					return nil
				},
				func(t *testModel) error {
					return errErrBadAgeTest
				},
			},
			wantErr: errErrBadAgeTest,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := New(tc.gotOpts...)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("got err: %v / want err: %v", err, tc.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("got: %v / want: %v", err, tc.wantErr)
				}
			}
		})
	}
}

// testModel is just a dummy type used for tests.
type testModel struct {
	name string
	age  int
}

var errBadNameTest = errors.New("fake error for bad name")
var errErrBadAgeTest = errors.New("fake error for bad age")
