# Learn Prometheus

## Run project
```
go build .
./learn-prometheus
```

## Run prometheus instances as docker container
```
docker run -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 9090:9090 --name learn-prometheus -d prom/prometheus:v2.37.6
```