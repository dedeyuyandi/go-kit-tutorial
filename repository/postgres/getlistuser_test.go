package postgres

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/testing/go-kit-tutorial/model"
)

func Test_repo_GetListUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		req model.ParamList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repo{
				db: tt.fields.db,
			}
			got, err := repo.GetListUser(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.GetListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
