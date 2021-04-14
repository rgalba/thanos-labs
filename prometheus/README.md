# Prometheus 

## Getting started

<!-- ```sh
cd prometheus
docker run -d --net=host \
    -v $PWD/prometheus.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-server \
    prom/prometheus
``` -->

Run Prometheus:  

```sh
mkdir prometheus0_us1_data
docker run -d \
    --rm \
    -p 9090:9090 \
    --cpus 2 \
    -v $PWD/prometheus0_us1.yml:/etc/prometheus/prometheus.yml \
    -v $PWD/prometheus0_us1_data:/prometheus \
    --name prometheus-server \
    prom/prometheus
```

Access it [at](http://localhost:9090)  

Run Node Exporter:  
```
docker run -d \
  -v "/proc:/host/proc" \
  -v "/sys:/host/sys" \
  -v "/:/rootfs" \
  -p 9100:9100 \
  --name=prometheus \
  quay.io/prometheus/node-exporter:v0.13.0 \
    -collector.procfs /host/proc \
    -collector.sysfs /host/sys \
    -collector.filesystem.ignored-mount-points "^/(sys|proc|dev|host|etc)($|/)"
```

Test it out: `curl localhost:9100/metrics`  

Run Grafana:  

```

```

Test it out: `curl localhost:3000`  

Deploy Thanos sidecar

```
docker run -d --net=host --rm \ 
    -v $(pwd)/prometheus0_us1.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-0-sidecar-eu1 \
    -u root \
    quay.io/thanos/thanos:v0.18.0 sidecar \
        --http-address 0.0.0.0:19090 \
        --grpc-address 0.0.0.0:19190 \
        --reloader.config-file /etc/prometheus/prometheus.yml \
        --prometheus.url http://127.0.0.1:9090 
        && echo "Started sidecar for Prometheus 0 EU1"
```
