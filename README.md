# helm-release-cleaner
A tool that cleans Helm Releases past a certain age in minutes.

## How it works
This tool is meant to run as a job in a Kubernetes cluster to remove helm-releases in a specific namespace. 

It will list all releases in the namespace and begin to remove each release if it's last deployed time is older than (x) minutes. The allowable age for a release can be specified.

## How to use this
This can be deployed as part of a helm-release as a cronjob to continually run at a specific rate. The cronjob can be deployed as a helm-hook.

### Permissions required
Since helm-releases uses secrets to store release details, this tool will require read/write permissions on secrets. Ideally this tool be granted an admin permission by tying a service account to the admin clusterrole via a rolebinding on the specific namespace.

### Example implementation
An implmentaiton can look as follow (these templates can be added to the required helm-chart to cleanup a release):

```
{{- if eq .Values.cleanUp }}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: {{ .Release.Name }}-cleaner
  annotations:
    # This is what defines this resource as a hook. Without this line, the
    # job is considered part of the release.
    "helm.sh/hook": post-install
    "helm.sh/hook-delete-policy": before-hook-creation
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: admin
subjects:
- kind: ServiceAccount
  name: {{ .Release.Name }}-cleaner
  namespace: {{ .Release.Namespace }}
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: {{ .Release.Name }}-cleaner
  labels:
    app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
    app.kubernetes.io/instance: {{ .Release.Name | quote }}
    app.kubernetes.io/version: {{ .Chart.AppVersion }}
    helm.sh/chart: "{{ .Chart.Name }}-{{ .Chart.Version }}"
  annotations:
    # This is what defines this resource as a hook. Without this line, the
    # job is considered part of the release.
    "helm.sh/hook": post-install
    "helm.sh/hook-delete-policy": before-hook-creation
spec:
  successfulJobsHistoryLimit: 1
  failedJobsHistoryLimit: 3
  schedule: "*/15 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName: {{ .Release.Name }}-cleaner
          containers:
          - name: helm-release-cleaner
            image: abualks/helm-release-cleaner:latest
            imagePullPolicy: IfNotPresent
            command:
            - "./helm-release-cleaner"
            args:
            - "--namespace={{ .Release.Namespace }}"
            - "--cleanup-age={{ .Values.cleanAge | default 120 }}"
          restartPolicy: OnFailure
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .Release.Name }}-cleaner
  annotations:
    # This is what defines this resource as a hook. Without this line, the
    # job is considered part of the release.
    "helm.sh/hook": post-install
    "helm.sh/hook-delete-policy": before-hook-creation
---
{{- end }}
```
