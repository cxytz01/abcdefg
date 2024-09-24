# How to deploy?
a. create an postgresql database, and test by below command, make sure it is correct, then you can use it in the producer/consumer configure file.
```
docker run -it --rm bitnami/postgresql:latest psql "postgresql://$user:$password@$ip:5432/$instance?sslmode=disable"
```
b. create a kafka instance, with no special options, the more partitions we have set, the more throughput we can get from it, the bottleneck is in the DB. 

c. prepare the producer/consumer configure files and execute file, please refer to separate README.md

[./producer/README.md](https://github.com/cxytz01/abcdefg/blob/main/producer/README.md)

[./consumer/README.md](https://github.com/cxytz01/abcdefg/blob/main/consumer/README.md)


d. use the swagger api [/api/v1/campaign, which indicated in swagger doc] to upload the csv file [./producer/tools/recipients.csv](https://github.com/cxytz01/abcdefg/blob/main/producer/tools/recipients.csv)

e. and you can call the api [/api/v1/messages] manually or use the cronjob shell [./producer/tools/crontrigger.sh](https://github.com/cxytz01/abcdefg/blob/main/producer/tools/crontrigger.sh) to dispatch the messages to kafka when the scheduled time is reached.

f. the consumer stdout will print the result.

# Deployment diagram
![deployment](./assets/deployment.svg)