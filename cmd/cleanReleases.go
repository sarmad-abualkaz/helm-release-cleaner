package cmd

import (
	
	"github.com/sarmad-abualkaz/helm-release-cleaner/util"

	log "github.com/sirupsen/logrus"
	callhelm "github.com/sarmad-abualkaz/helm-release-cleaner/helm"
)

func CleanReleases(dryRun bool, namespace string, releasesMap map[string]bool, cleanup int, repoCache string, repoConfig string, debug bool, linting bool) {


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

		if util.ShouldDelete(release, cleanup, releasesMap) {
				
			if !dryRun {
				deleteReleaseErr := callhelm.DeleteRelease(helmclient, release.Name, namespace)
				
				if deleteReleaseErr != nil {

					log.WithFields(log.Fields{
						"release": release.Name,
						"namespace": release.Namespace,
						"Error": deleteReleaseErr.Error(),
					  }).Error("Failed to delete release ...")
				}
			}

		}

	}
}
