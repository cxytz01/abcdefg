# Producer
The server receives clients' campaigns and recipients, triggered by a scheduler such as cronjob/scheduler framework, etc, you can even trigger the producer manually by http api.

### Run
- create a configure file, please refer to ./config/service.conf_example
- >_$ make binary
- >_$ bin/producer --conf ./config/service.conf

### Test
Please refer to Swagger doc, you could test the api in Swagger at port 7005, the port in Swagger is hardcode. 
`http://127.0.0.1:7005/producer/swagger/index.html`

###### instance

message template:
```
Hello {{.name}}, thank you for shopping with us! Please confirm your phone number "{{.phone}}", use the phone number at checkout to get 20% off your next purchase.
```
schedule time format:
```
2024-09-24T12:10:28+08:00
```

### Note
- DB schemas will be auto created, refer to /producer/pkg/models/models.go
- Trigger api is ```/api/v1/messages```, you could trigger it manually, or use the conjob shell script at /producer/tools/crontrigger.sh
- CSV generate code was put at /producer/tools/generatecsv.go

### TBD
trigger api ```/api/v1/messages``` currently is not singleton, concurrent call will cause the messages repeatedly dispatch to kafka.