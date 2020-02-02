
# Steps required for k8s
# 1) Set Secrets
kubectl create secret generic bbq-db-creds \
    --from-file=/home/scott/source/github/ssargent/bbq/src/bbq-apiserver/secrets/BBQ_DB_HOST.txt \
    --from-file=/home/scott/source/github/ssargent/bbq/src/bbq-apiserver/secrets/BBQ_DB_NAME.txt \
    --from-file=/home/scott/source/github/ssargent/bbq/src/bbq-apiserver/secrets/BBQ_DB_USER.txt \
    --from-file=/home/scott/source/github/ssargent/bbq/src/bbq-apiserver/secrets/BBQ_DB_PASSWORD.txt


# 2) Setup Redis via Helm
# 3) Load TLS Cert/Key 
kubectl create secret tls mythicalcodelabs-tls --key ~/infrastructure/keys/mythicalcodelabs.com.key.pem --cert ~/infrastructure/keys/star.mythicalcodelabs.com.cert.pem
# 4) 