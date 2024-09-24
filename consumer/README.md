# Consumer

The consumer agent watches the Kafka mq and fetches the message to third pary immediately.

## Run
- create a configure file, please refer to ./config/service.conf_example
- >_$ make binary
- >_$ bin/consumer --conf ./config/service.conf

## Note
The consumer agent is a single process/thread/routine, if you want more throughput, please start multiple consumer agents, the amounts of agents should be less and equal to kafka partitions.

I prefer to use single process/thread/routine with multiple replicas rather than multiple processes/threads/routines.

## Tools
you could start/stop a bunch of consumer agents by using the [tool/* scripts](https://github.com/cxytz01/abcdefg/tree/main/consumer/tools).