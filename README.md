# Thanos Katakoda

[link](https://katacoda.com/thanos/courses/thanos/1-globalview)

## Setup environment

```sh
multipass launch --name thanos-box -m 3GB -c 2
multipass mount 
multipass shell thanos-box
sudo snap install docker
```

## Intro: Global View and seamless HA for Prometheus

`prometheus0_eu1.yml`
```yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
  external_labels:
    cluster: eu1
    replica: 0

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['127.0.0.1:9090']
```

`prometheus0_us1.yml`
```yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
  external_labels:
    cluster: us1
    replica: 0

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['127.0.0.1:9091','127.0.0.1:9092']
```

`prometheus1_us1.yml`
```yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
  external_labels:
    cluster: us1
    replica: 1

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['127.0.0.1:9091','127.0.0.1:9092']
```

**2 - Prepare persistent volumes**

```shell
mkdir -p prometheus0_eu1_data prometheus0_us1_data prometheus1_us1_data
```

**3 - Deploying EU1**

```yml
docker run -d --net=host --rm \
    -v $(pwd)/prometheus0_eu1.yml:/etc/prometheus/prometheus.yml \
    -v $(pwd)/prometheus0_eu1_data:/prometheus \
    -u root \
    --name prometheus-0-eu1 \
    quay.io/prometheus/prometheus:v2.14.0 \
    --config.file=/etc/prometheus/prometheus.yml \
    --storage.tsdb.path=/prometheus \
    --web.listen-address=:9090 \
    --web.external-url=http://192.168.203.140 \
    --web.enable-lifecycle \
    --web.enable-admin-api && echo "Prometheus EU1 started!"
```

> Access it at http://192.168.203.140:9090/graph

**4 - Deploying US0**

```yml
sudo docker run -d --net=host --rm -v $(pwd)/prometheus0_us1.yml:/etc/prometheus/prometheus.yml -v $(pwd)/prometheus0_us1_data:/prometheus -u root --name prometheus-0-us1 quay.io/prometheus/prometheus:v2.14.0 --config.file=/etc/prometheus/prometheus.yml --storage.tsdb.path=/prometheus --web.listen-address=:9091 --web.external-url=http://192.168.203.140 --web.enable-lifecycle --web.enable-admin-api && echo "Prometheus 0 US1 started!"
```

> Access it at http://192.168.203.140:9091/graph

**4.1 - Deploying US1**

```yml
docker run -d --net=host --rm -v $(pwd)/prometheus1_us1.yml:/etc/prometheus/prometheus.yml -v $(pwd)/prometheus1_us1_data:/prometheus -u root --name prometheus-1-us1 quay.io/prometheus/prometheus:v2.14.0 --config.file=/etc/prometheus/prometheus.yml --storage.tsdb.path=/prometheus --web.listen-address=:9092 --web.external-url=https://192.168.203.137 --web.enable-lifecycle --web.enable-admin-api && echo "Prometheus 1 US1 started!"
```

> Access it at http://192.168.203.140:9092/graph

# Thanos Sidecars

Thanos is a single Go binary capable to run in different modes. Each mode represents a different component and can be invoked in a single command.

```
sudo docker run --rm quay.io/thanos/thanos:v0.18.0 --help
```

**2.1 - Adding sidecar to "EU1" Prometheus**

```sh
sudo docker run -d --net=host --rm -v $(pwd)/prometheus0_eu1.yml:/etc/prometheus/prometheus.yml --name prometheus-0-sidecar-eu1 -u root quay.io/thanos/thanos:v0.18.0 sidecar --http-address 0.0.0.0:19090 --grpc-address 0.0.0.0:19190 --reloader.config-file /etc/prometheus/prometheus.yml --prometheus.url http://127.0.0.1:9090 && echo "Started sidecar for Prometheus 0 EU1"
```

**2.2 - Adding sidecars to each replica of Prometheus in "US1"**

```sh
sudo docker run -d --net=host --rm -v $(pwd)/prometheus0_us1.yml:/etc/prometheus/prometheus.yml --name prometheus-0-sidecar-us1 -u root quay.io/thanos/thanos:v0.18.0 sidecar --http-address 0.0.0.0:19091 --grpc-address 0.0.0.0:19191 --reloader.config-file /etc/prometheus/prometheus.yml --prometheus.url http://127.0.0.1:9091 && echo "Started sidecar for Prometheus 0 US1"
```

```sh
sudo docker run -d --net=host --rm -v $(pwd)/prometheus1_us1.yml:/etc/prometheus/prometheus.yml --name prometheus-1-sidecar-us1 -u root quay.io/thanos/thanos:v0.18.0 sidecar --http-address 0.0.0.0:19092 --grpc-address 0.0.0.0:19192 --reloader.config-file /etc/prometheus/prometheus.yml --prometheus.url http://127.0.0.1:9092 && echo "Started sidecar for Prometheus 1 US1"
```

Now, to check if sidecars are running well, let's modify Prometheus scrape configuration to include our added sidecars.  
