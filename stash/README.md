## Install Minio for S3 compatible storage

```
kubectl create ns minio
```

```
helm install minio stable/minio -n minio --set accessKey=hieudeptrai,secretKey=hieudeptrai,resources.requests.memory=500Mi,service.type=NodePort
```
## Install stash

```
kubectl create ns stash
```

```
helm repo add appscode https://charts.appscode.com/stable/
```

```
helm install stash-operator appscode/stash -n stash --version v0.9.0-rc.6
```

## Setup repository

kubectl create ns app

echo -n 'changeit' > RESTIC_PASSWORD && \
echo -n 'hieudeptrai' > AWS_ACCESS_KEY_ID && \
echo -n 'hieudeptrai' > AWS_SECRET_ACCESS_KEY && \
kubectl create secret generic -n app s3-secret \
    --from-file=./RESTIC_PASSWORD \
    --from-file=./AWS_ACCESS_KEY_ID \
    --from-file=./AWS_SECRET_ACCESS_KEY && \
rm RESTIC_PASSWORD AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY

kubectl -n app apply -f s3-repository.yaml

## Install demo app (mysql)

```
helm -n app install mysql stable/mysql --set mysqlRootPassword=hieudeptrai
```

Try to create some databases.

## Backup

kubectl -n app apply -f backup_configuration.yaml

## Pause backup

kubectl patch backupconfiguration -n app deployment-backup --type="merge" --patch='{"spec": {"paused": true}}'

## Simulate a disaster

```
helm -n app delete mysql
```

## Restore

Reinstall deployment
```
helm -n app install mysql stable/mysql --set mysqlRootPassword=hieudeptrai
```

Restore persistent volume

```
kubectl -n app apply -f restore_session.yaml
```