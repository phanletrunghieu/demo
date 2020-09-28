## Install Minio for S3 compatible storage

```
kubectl create ns minio
```

```
helm -n minio install minio stable/minio --set accessKey=hieudeptrai,secretKey=hieudeptrai,resources.requests.memory=500Mi,service.type=NodePort
```
## Install velero

```
kubectl create ns velero
```

```
helm -n velero install velero vmware-tanzu/velero --version 2.12.15 -f values.yaml --set-file credentials.secretContents.cloud=./credentials-velero
```

```
brew install velero
```

## Install demo app (mysql)

```
helm -n app install mysql stable/mysql --set mysqlRootPassword=hieudeptrai,persistence.existingClaim=mysql-pvc
```

Try to create some databases.

Or

```
kubectl -n app apply -f pvc.yaml
kubectl -n app apply -f demo-deployment.yaml
```

## Add annotation to pod

```
kubectl -n app annotate pod/mysql-79c8db4766-rvdzp backup.velero.io/backup-volumes=data
```

## Backup

```
velero create backup backup1 --include-namespaces app
```

```
velero create schedule backup-demo-app --schedule="*/30 * * * *" --include-namespaces app --ttl 48h0m0s
```

```
velero get backup
```

## Simulate a disaster

```
helm -n app delete mysql
```

## Restore

```
velero create restore --from-backup backup1
```

