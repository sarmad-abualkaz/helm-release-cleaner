package util

import (
	"reflect"
	"testing"
	"time"

	helmTime "helm.sh/helm/v3/pkg/time"
	"helm.sh/helm/v3/pkg/release"
)

func TestShouldDelete(t *testing.T) {

	type args struct {
		release *release.Release
		cleanup int
		releaseMap map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		// 1st test - nothing passed (delete release):
		{
			name: "Should delete_release - generically",
			args: args{
						release: &release.Release{
							Name: "test",
							Namespace: "abualks",
							Info: &release.Info{
								LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
							},
						},	
				cleanup: 120,
				releaseMap: make(map[string]bool),
			},
			want: true,
		},
		// 2nd test - map passed (delete release):
		{
			name: "Should delete release - specifically",
			args: args{
						release: &release.Release{
							Name: "specific-release",
							Namespace: "abualks",
							Info: &release.Info{
								LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
							},
						},	
				cleanup: 120,
				releaseMap: map[string]bool{
					"specific-release":  true,
					"specific-release-2": true,
				},
			},
			want: true,
		},

		// 3rd test - map passed (not specified):
		{
			name: "Should NOT delete release - missing in map",
			args: args{
						release: &release.Release{
							Name: "dont-delete-release",
							Namespace: "abualks",
							Info: &release.Info{
								LastDeployed: helmTime.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
							},
						},	
				cleanup: 120,
				releaseMap: map[string]bool{
					"specific-release":  true,
					"specific-release-2": true,
				},
			},
			want: false,
		},

		// 4th test - map passed (specified) not old:
		{
			name: "Should NOT delete release - missing in map",
			args: args{
						release: &release.Release{
							Name: "specific-release-2",
							Namespace: "abualks",
							Info: &release.Info{
								LastDeployed: helmTime.Now().UTC(),
							},
						},	
				cleanup: 120,
				releaseMap: map[string]bool{
					"specific-release":  true,
					"specific-release-2": true,
				},
			},
			want: false,
		},

		// 5th test - map not passed not old:
		{
			name: "Should NOT delete release - missing in map",
			args: args{
						release: &release.Release{
							Name: "specific-release-3",
							Namespace: "abualks",
							Info: &release.Info{
								LastDeployed: helmTime.Now().UTC(),
							},
						},	
				cleanup: 120,
				releaseMap: make(map[string]bool),
			},
			want: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := ShouldDelete(tt.args.release, tt.args.cleanup, tt.args.releaseMap)

			if !reflect.DeepEqual(got, tt.want){
				t.Errorf("GetConfigMap() = %v, want %v", got, tt.want)
			}
		})
	}
}	