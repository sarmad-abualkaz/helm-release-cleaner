	package main

import (
	"flag"
	"github.com/sarmad-abualkaz/helm-release-cleaner/cmd"

	log "github.com/sirupsen/logrus"
)

func main(){
	
	cleanupAge := flag.Int("cleanup-age", 120, "age of release to be cleaned in minutes")
	dryRun := flag.Bool("dry-run", false, "dry run")
	namespace := flag.String("namespace", "", "kubernetes namespace where secret exists")
	repoCache := flag.String("repo-cache", "/tmp/.helmcache", "Repoistory Cache")
	repoConfig := flag.String("repo-config", "/tmp/.helmrepo", "Repoistory Config")
	clientDebug := flag.Bool("client-debug", false, "Enable debug level on helm client")
	clientLinting := flag.Bool("client-linting", true, "Enable linting on helm client")

	
	flag.Parse()

	// log program starting
	log.WithFields(log.Fields{
		"cleanup-age": *cleanupAge,
		"dry-run": *dryRun,
		"namespace": *namespace,
	  }).Info("program started ...")


	// call cmd.updatECRSecret with params
	cmd.CleanReleases(*dryRun, *namespace, *cleanupAge, *repoCache, *repoConfig, *clientDebug, *clientLinting)

}