# k8s-PostgreSQL-Deployment

## Secrets

```bash
echo "password" | base64

kubectl apply -f deployment-postgres-secret.yml

kubectl get secret postgres-secret-config -o yaml
```

## PersistentVolume

```bash
kubectl apply -f deployment-postgres-pv.yml

kubectl get pv postgres-pv-volume
```

## PersistentVolumeClaim

```bash
kubectl apply -f deployment-postgres-pv-claim.yml

kubectl get pvc postgres-pv-claim
```

## App

```bash
kubectl apply -f deployment-postgres.yml

kubectl get deployments
```
