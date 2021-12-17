## [DIY vs DBaaS](project/README.md)

## GKE Logs
![Alt text](/gke-logs.png?raw=true)

## Linkerd Canary Release
![Alt text](/linkerd.png?raw=true)

## Rancher vs OpenShift
### Winner: Rancher
* 100% open source
* Free
* No vendor lock
* Less opinionated
* Cluster importing

## Landscape
![Alt text](/landscape.png?raw=true)
### Usages
* Microsoft SQL Server: Outside of the course
* mongoDB: Outside of the course
* PostgreSQL: Used as the chosen database for appliations.
* redis: Outside of the course
* snowflake: Outside of the course
* NATS: Used to create a message que system for sending messages to an URL on todo creation/update.
* HELM: Used to install several resources to cluster
* argo: Used to implement canary release rollout
* flux: Used for GitOps implementation
* circleCI: Outside of the course
* Github Actions: Used implement automatic deployment and environment creation
* Gitlab: Outside of the course
* Travis CI: Outside of the course
* kubernetes: I meaaaaan..
* contour: Used for proxyng when setting up a serverless version for ping-pong app
* traefic: Used for proxyng at least in k3d througout the course.
* Linkerd: Used for implementing a service mesh in our project
* Google Persistent Disk: Used for our postgres data and other persistent storage in google cloud
* Amazon ECR: Outside of the course
* Google Container Registry: Used to store our docker images starting from part 3
* AWS: Outside of the course
* Google Cloud: Used throughout part 3 to deploy our deployment, images etc
* Heroku: Outside of the course
* Prometheus: Used for monitoring and querying the system throughout the course.
* Datadog: Outside of the course
* Grafana: Used for monitoring our system
* Sentry: Outside of the course
* elasitc: Outside of the course
* Grafana loki: Used for getting logs to show up in grafana
* graylog: Outside of the course
* Docker Hub: Used for storing and fetching container images used in the course
* GitHub: this
* git: Used for version control
* knative: Used for implementing a serverless version for ping-pong application
