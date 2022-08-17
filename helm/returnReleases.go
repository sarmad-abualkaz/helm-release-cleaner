package callhelm

import (
	"helm.sh/helm/v3/pkg/release"
	log "github.com/sirupsen/logrus"
	helmgo "github.com/mittwald/go-helm-client"
)

func ReturnReleases(helmclient helmgo.Client) ([]*release.Release, error){
	
	log.WithFields(log.Fields{
		}).Info("Listing Releases ...")

	releases, err := helmclient.ListDeployedReleases()

	if err != nil {
		return nil, err
	}

	return releases, nil
}