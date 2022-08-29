package callhelm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	mockhelmclient "github.com/mittwald/go-helm-client/mock"
)

func TestDeleteRelease(t *testing.T) {
	type args struct {
		name string
		namespace string
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		//1st test:
		{
			name: "Should fail on finding release",
			args: args{
				name: "test-release",
				namespace: "abualks",
				err: fmt.Errorf("uninstall: Release not loaded: speckel-resources-pvcs: release: not found"),
			},
			want: fmt.Errorf("uninstall: Release not loaded: speckel-resources-pvcs: release: not found"),
		},
		// 2st test:
		{
			name: "Should succeed - No Error",
			args: args{
				name: "test-release",
				namespace: "abualks",
				err: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
		
			client := mockhelmclient.NewMockClient(ctrl)
			if client == nil {
				t.Fail()
			}

			client.EXPECT().UninstallReleaseByName(tt.args.name).Return(tt.args.err)
			
			gotErr := DeleteRelease(client, tt.args.name, tt.args.namespace)

			if !reflect.DeepEqual(gotErr, tt.want){
				t.Errorf("GetConfigMap() = %v, want %v", gotErr, tt.want)
			}
			
		})
	}
}