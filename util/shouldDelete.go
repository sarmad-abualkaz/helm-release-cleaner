package util

import 	(
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/time"
	
	log "github.com/sirupsen/logrus"
)

func ShouldDelete(release *release.Release, cleanup int, releasesMap map[string]bool) bool {

	if len(releasesMap) > 0 {

		if _, found := releasesMap[release.Name] ; found {

			log.WithFields(log.Fields{
				"namespace": release.Namespace,
				"release": release.Name,
				"lastDeplay": release.Info.LastDeployed.UTC(),
			}).Info("Release is in provided list. Proceed with check ...")

			return ShouldDeleteHelper(release, cleanup)

		} else {

			log.WithFields(log.Fields{
				"namespace": release.Namespace,
				"release": release.Name,
				"lastDeplay": release.Info.LastDeployed.UTC(),
			}).Info("Release is NOT in provided list. Skipping check ...")
		}
	
	} else {

		// always return cleanup if map is blank
		return ShouldDeleteHelper(release, cleanup)
	}

	return false

}


func ShouldDeleteHelper(release *release.Release, cleanup int) bool {

	currentTime := time.Now().UTC()
	lastDeployedTime := release.Info.LastDeployed.UTC()
	diff := currentTime.Sub(lastDeployedTime).Minutes()

	if diff > float64(cleanup) {

		log.WithFields(log.Fields{
			"namespace": release.Namespace,
			"release": release.Name,
			"lastDeplay": lastDeployedTime,
			"age (in min)": diff,
		}).Warn("Release is going to be deleted ...")

		return true

	}

	log.WithFields(log.Fields{
		"namespace": release.Namespace,
		"release": release.Name,
		"lastDeplay": release.Info.LastDeployed.UTC(),
		"age (in min)": diff,
	}).Info("Release is still within allowed age ...")

	return false
}
