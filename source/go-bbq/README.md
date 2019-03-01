# go-bbq
Go BBQ - a golang backend for the most over engineered bbq thermometer ever project.

## To Run Locally
Create text files

| Name  |  Contents of File | Location of file |
|-------|-------------------|------------------|
| BBQ_DB_USER.txt | Name of PGSQL User to connect with | ./secrets/BBQ_DB_USER.txt |
| BBQ_DB_PASSWORD.txt | Password of PGSQL User | ./secrets/BBQ_DB_PASSWORD.txt |
| BBQ_DB_NAME.txt | Name of PGSQL DB to connect to | ./secrets/BBQ_DB_NAME.txt |
| BBQ_DB_HOST.txt | Address of PGSQL Server to connect to | ./secrets/BBQ_DB_HOST.txt |
| BBQ_REDIS_MASTER.txt | Address of Redis master in fqdns:port format| to ./secrets/BBQ_REDIS_MASTER.txt |
| BBQ_REDIS_PASSWORD.txt | Redis password | to ./secrets/BBQ_REDIS_PASSWORD.txt |

This should only used for local development.  K8S is the only currently supported environment for this, in K8S use standard secrets.  /secrets/*.txt files are gitignored.