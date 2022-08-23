package cmd

import (
	
	"helm.sh/helm/v3/pkg/time"

	log "github.com/sirupsen/logrus"
	callhelm "github.com/sarmad-abualkaz/helm-release-cleaner/helm"
)

func CleanSepecificReleases(dryRun bool, namespace string, releasesMap map[string]bool, cleanup int, repoCache string, repoConfig string, debug bool, linting bool) {


	helmclient, helmclienterr := callhelm.CreateClient(namespace, repoCache, repoConfig, debug, linting)

	if helmclienterr != nil {
		log.WithFields(log.Fields{
			"namespace": namespace,
			"repo-cache": repoCache,
			"repo-config": repoConfig,
			"debug": debug,
			"linting": linting,
			"Error": helmclienterr.Error(),
		  }).Fatal("Failed to initialize client ...")
		  
		panic(helmclienterr)
	}

	releases, releaseserr := callhelm.ReturnReleases(helmclient)
	
	if releaseserr != nil {

		if releaseserr != nil {
			log.WithFields(log.Fields{
				"namespace": namespace,
				"repo-cache": repoCache,
				"repo-config": repoConfig,
				"debug": debug,
				"linting": linting,
				"Error": releaseserr.Error(),
			  }).Fatal("Failed to list releases ...")

		panic(releaseserr)
		}
	}
	
	for _, release := range(releases){

		if _, found := releasesMap[release.Name] ; found {

			log.WithFields(log.Fields{
				"namespace": namespace,
				"release": release.Name,
				"lastDeplay": release.Info.LastDeployed.UTC(),
			}).Info("Release is in provided list. Proceed with check ...")

			currentTime := time.Now().UTC()
			lastDeployedTime := release.Info.LastDeployed.UTC()
			diff := currentTime.Sub(lastDeployedTime).Minutes()
	
			if diff > float64(cleanup) {
				log.WithFields(log.Fields{
					"namespace": namespace,
					"release": release.Name,
					"lastDeplay": lastDeployedTime,
					"age (in min)": diff,
				}).Warn("Release is going to be deleted ...")
					
				if !dryRun {
					deleteReleaseErr := callhelm.DeleteRelease(helmclient, release.Name, namespace)
					
					if deleteReleaseErr != nil {
	
						log.WithFields(log.Fields{
							"release": release,
							"namespace": namespace,
							"Error": deleteReleaseErr.Error(),
						  }).Error("Failed to delete release ...")
					}
				}
	
			} else {
				log.WithFields(log.Fields{
					"namespace": namespace,
					"release": release.Name,
					"lastDeplay": release.Info.LastDeployed.UTC(),
					"age (in min)": diff,
				}).Info("Release is still within allowed age ...")
			}
		} else {
			
			log.WithFields(log.Fields{
				"namespace": namespace,
				"release": release.Name,
				"lastDeplay": release.Info.LastDeployed.UTC(),
			}).Info("Release is NOT in provided list. Skipping check ...")

		}
	}
}
