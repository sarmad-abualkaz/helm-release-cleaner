package callhelm

import (
	helmgo "github.com/mittwald/go-helm-client"
	log "github.com/sirupsen/logrus"
)

func CreateClient(namespace string, repoCache string, repoConfig string, debug bool, linting bool) (helmgo.Client, error){

	log.WithFields(log.Fields{
		"namespace": namespace,
		"repo-cache": repoCache,
		"repo-config": repoConfig,
		"debug": debug,
		"linting": linting,
	  }).Info("Initializing client ...")

	opt := &helmgo.Options{
		Namespace:        namespace, // Change this to the namespace you wish the client to operate in.
		RepositoryCache:  repoCache,
		RepositoryConfig: repoConfig,
		Debug:            debug,
		Linting:          linting,
		DebugLog:         func(format string, v ...interface{}) {},
	}

	helmClient, err := helmgo.New(opt)
	
	if err != nil {
		return nil, err
	}

	return helmClient, nil

}
