package callhelm

import (
	log "github.com/sirupsen/logrus"
	helmgo "github.com/mittwald/go-helm-client"
)

func DeleteRelease(helmclient helmgo.Client, release string, namespace string) error {
	
	log.WithFields(log.Fields{
		"release": release,
		"namespace": namespace,
	}).Warn("Deleting release ...")

	err := helmclient.UninstallReleaseByName(release)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"release": release,
		"namespace": namespace,
	}).Info("Release deleted ...")	

	return nil
}