# Sample App For End to End Golden Path

## Boostrapping your project

1. Enable services

```shell
gcloud services enable sourcerepo.googleapis.com \
                       cloudbuild.googleapis.com \
                       clouddeploy.googleapis.com \
                       container.googleapis.com
```

1. Create a source repository

```shell
gcloud source repos create sample-app
```

1. Create a Cloud Build trigger for the main branch

```shell
gcloud beta builds triggers create cloud-source-repositories --name="sample-app-master" \
                                                             --repo="sample-app" \
                                                             --branch-pattern="master" \
                                                             --build-config="cloudbuild.yaml"
```

1. Create a `staging` GKE Cluster:

```shell
gcloud container clusters create staging \
    --release-channel regular \
    --addons ConfigConnector \
    --workload-pool=$(gcloud config get-value project).svc.id.goog \
    --enable-stackdriver-kubernetes --region us-central1 \
    --enable-ip-alias
```

1. Create a `prod` GKE cluster:
```shell
gcloud container clusters create prod \
    --release-channel regular \
    --addons ConfigConnector \
    --workload-pool=$(gcloud config get-value project).svc.id.goog \
    --enable-stackdriver-kubernetes --region us-central1 \
    --enable-ip-alias
```

1. Push your source code to the repo:
```shell
git config --global credential.https://source.developers.google.com.helper gcloud.sh
git remote add google https://source.developers.google.com/p/$(gcloud config get-value project)/r/sample-app
git push google $(git branch --show-current):master
```
