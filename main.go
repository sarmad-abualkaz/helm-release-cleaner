	package main

import (
	"flag"
	"strings"
	"github.com/sarmad-abualkaz/helm-release-cleaner/cmd"

	log "github.com/sirupsen/logrus"
)

func main(){
	
	cleanupAge := flag.Int("cleanup-age", 120, "age of release to be cleaned in minutes")
	dryRun := flag.Bool("dry-run", false, "dry run")
	namespace := flag.String("namespace", "", "kubernetes namespace where secret exists")
	releases := flag.String("releases", "", "list of releases to delete")
	repoCache := flag.String("repo-cache", "/tmp/.helmcache", "Repoistory Cache")
	repoConfig := flag.String("repo-config", "/tmp/.helmrepo", "Repoistory Config")
	clientDebug := flag.Bool("client-debug", false, "Enable debug level on helm client")
	clientLinting := flag.Bool("client-linting", true, "Enable linting on helm client")

	releasesMap := make(map[string]bool)
	
	flag.Parse()

	// log program starting
	log.WithFields(log.Fields{
		"cleanup-age": *cleanupAge,
		"dry-run": *dryRun,
		"namespace": *namespace,
		"releases": *releases,
	  }).Info("program started ...")


	// call cmd.updatECRSecret with params
	if *releases == "" {

		log.WithFields(log.Fields{
			"cleanup-age": *cleanupAge,
			"dry-run": *dryRun,
			"namespace": *namespace,
		  }).Info("kicking generic release cleaner - no releases specified...")

	} else {

		log.WithFields(log.Fields{
			"cleanup-age": *cleanupAge,
			"dry-run": *dryRun,
			"namespace": *namespace,
			"releases": *releases,
		  }).Info("kicking specific release cleaner...")

		  releasesList := strings.Split(*releases, " ")

		  for _, release := range(releasesList) {
			releasesMap[release] = true
			
		}

	}

	cmd.CleanReleases(*dryRun, *namespace, releasesMap, *cleanupAge, *repoCache, *repoConfig, *clientDebug, *clientLinting)

}
