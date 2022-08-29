package callhelm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	helmTime "helm.sh/helm/v3/pkg/time"
	"helm.sh/helm/v3/pkg/release"
	mockhelmclient "github.com/mittwald/go-helm-client/mock"
)

func TestReturnReleases(t *testing.T) {

	type args struct {
		releases []*release.Release
	}
	tests := []struct {
		name string
		args args
		want []*release.Release
	}{
		//1st test:
		{
			name: "Should correctly find two releases",
			args: args{
				releases: []*release.Release {
					{
						Name: "test",
					 	Namespace: "abualks",
						Info: &release.Info{
							LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
						},
					},
					{
						Name: "test2",
						Namespace: "abualks",
						Info: &release.Info{
							LastDeployed: helmTime.Date(2020, 2, 1, 12, 30, 0, 0, time.UTC),
						},

					},
				},
			},
			want: []*release.Release {
				{
					Name: "test",
					 Namespace: "abualks",
					Info: &release.Info{
						LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					},
				},
				{
					Name: "test2",
					Namespace: "abualks",
					Info: &release.Info{
						LastDeployed: helmTime.Date(2020, 2, 1, 12, 30, 0, 0, time.UTC),
					},

				},
			},
		},
		//2nd test:
		{
			name: "Should correctly find a single release",
			args: args{
				releases: []*release.Release {
					{
						Name: "test",
					 	Namespace: "abualks",
						Info: &release.Info{
							LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
						},
					},
				},
			},
			want: []*release.Release {
				{
					Name: "test",
					 Namespace: "abualks",
					Info: &release.Info{
						LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					},
				},
			},
		},
		// 3rd test:
		{
			name: "Should correctly find no releases",
			args: args{
				releases: []*release.Release{},
			},
			want: []*release.Release{},
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

			client.EXPECT().ListDeployedReleases().Return(tt.args.releases, nil)
			
			got, _ := ReturnReleases(client)
			
			for i, release := range(got){
			
				if !reflect.DeepEqual(release.Name, tt.want[i].Name){
					t.Errorf("ReturnReleases() = release name is %v, want %v",release.Name, tt.want[i].Name)
				}

				if !reflect.DeepEqual(release.Namespace, tt.want[i].Namespace){
					t.Errorf("ReturnReleases() = release namespace is %v, want %v",release.Namespace, tt.want[i].Namespace)
				}

				if !reflect.DeepEqual(release.Info.LastDeployed, tt.want[i].Info.LastDeployed){
					t.Errorf("ReturnReleases() = release last deployed time is %v, want %v",release.Info.LastDeployed, tt.want[i].Info.LastDeployed)
				}

			}
		})
	}
}


func TestReturnReleasesFailOnClient(t *testing.T) {

	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		//1st test:
		{
			name: "Should fail on client error",
			args: args{
				err: fmt.Errorf("Kubernetes cluster unreachable: the server is currently unable to handle the request"),
			},
			want: fmt.Errorf("Kubernetes cluster unreachable: the server is currently unable to handle the request"),
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

			client.EXPECT().ListDeployedReleases().Return(nil, tt.args.err)
			
			_, got := ReturnReleases(client)

			if !reflect.DeepEqual(got, tt.want){
				t.Errorf("GetConfigMap() = %v, want %v", got, tt.want)
			}
			
		})
	}
}
