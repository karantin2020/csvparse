package csvparse

import (
	"encoding/csv"
	"reflect"
	"testing"
)

func TestProcessCSV(t *testing.T) {
	type args struct {
		rc *csv.Reader
	}
	tests := []struct {
		name  string
		args  args
		want  <-chan []string
		want1 <-chan error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ProcessCSV(tt.args.rc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessCSV() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProcessCSV() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	type args struct {
		rc     *csv.Reader
		allocr Alloc
	}
	tests := []struct {
		name    string
		args    args
		want    []Unmarshaller
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAll(tt.args.rc, tt.args.allocr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
