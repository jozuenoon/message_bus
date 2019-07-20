# Travel time information system
[![Build Status](https://cloud.drone.io/api/badges/jozuenoon/message_bus/status.svg)](https://cloud.drone.io/jozuenoon/message_bus)
[![codecov](https://codecov.io/gh/jozuenoon/message_bus/branch/master/graph/badge.svg)](https://codecov.io/gh/jozuenoon/message_bus)
[![GolangCI](https://golangci.com/badges/github.com/jozuenoon/message_bus.svg)](https://golangci.com/r/github.com/jozuenoon/message_bus)


This repository is part of bigger conceptual system that would
provide travel time estimations based on mobile devices identification.

Explanation:

Let's say that we have Bluetooth enabled device in car or with us
while commuting. While we pass through our regular paths around
city special detectors grabs MAC address of our device and time
on witch it got catch. There is great concern taken about privacy
and MAC address is hashed before it is passed down through the system.
Next, system tries to figure out what is current travel time 
between points based on many such traces.

## Message bus

Repository provides server for collecting and querying
events received from mobile devices detectors. Additionally 
there is client tool for simple traffic flow simulation.

## Bluetooth / WiFi / IMSI detectors aka. detector of mobile devices

Detectors are devices located around city, highways etc. which
are capable for intercepting Bluetooth, WiFi MAC addresses and
IMSI ids from passing travelers.

## NOTE

This service is purely conceptual however I did work with similar
system which was capable of intercepting Bluetooth MAC addresses.
Adding IMSI catcher is purely theoretical since they are 
basically illegal without special permissions.

# Technical details

This section provides some extra technical details about solution.

### Configuration and deployment

Example configuration could be found in root directory `<service_name>.cofnig.yaml`. For
more details please reference either calling a binary release with `--help` flag or going
directly into `cmd/<service>/main.go` and inspecting `config` structure.

Example for `cqserver`:
```bash
Usage of main:
      --collector_port string      grpc collector port (default ":9000")
      --query_port string          grpc query port (default ":8000")
      --healthcheck_port string    grpc health check port (default ":5000")
      --etcd.prefix string         etcd app prefix (default "tdc")
      --etcd.endpoints string...   etcd endpoints (default "http://127.0.0.1:2379")
      --config_file string         
  -h, --help                       print this help menu
```

Service could be also configured via environment variables, please reference [gonfig](https://github.com/stevenroose/gonfig)
package.

Deployment assumes existence of etcd cluster with endpoints `["http://etcd-0:2379", "http://etcd-1:2379", "http://etcd-2:2379"]`.
This could be achieved easily by deploying [etcd-operator](https://github.com/helm/charts/tree/master/stable/etcd-operator).

Deployment also assumes existence of [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager)
to provide TLS certificates for domains. Expected issuer in this case is `http-issuer`.

### ETCD repository data model discussion

Storing events from detectors would include detector ID as root path
then bucketing based on unix timestamp with bucket size of 100 seconds:
```bash
${namespace}/detectors/<detector_id>/<unix_timestamp/100>/<device_id>
value: {"time": "<rfc3339>"} # storing precise time
```

There is device ID duplicated intentionally inside value for convenience and
precise time of detection.

Consecutive timestamp values would look like following:
```$xslt
1563568100 => 2019-07-19T20:28:20+00:00
1563568200 => 2019-07-19T20:30:00+00:00
1563568300 => 2019-07-19T20:31:40+00:00
```

Thanks to unix timestamp usage we could avoid many problems regarding timezones,
leap years, daylight saving time changes etc.

Since detectors are representing directed graph of traffic flows we can
limit complexity of calculating all graph connections by explicitly declaring
links between detectors by listing it's incoming traffic peers.
```bash
${namespace}/links/<dest_detector_id>/<src_detector_id#0>
${namespace}/links/<dest_detector_id>/<src_detector_id#1>
${namespace}/links/<dest_detector_id>/<src_detector_id#2>
value: {"max_seconds": "600"} # maximum travel time which is taken under consideration
```



Worker which would calculate travel times for given detector should acquire
`<dest_detector_id>` to reserve given batch of work since number of source detectors
should be well beyond 10 resulting in fair amount of work for single run.

### Running ETCD on local host

Create local etcd single node cluster:
```bash
docker-compose up
```


### etcdctl installation

```
go get github.com/coreos/etcd/etcdctl
cd <go-dir>/src/github.com/coreos/etcd/
git checkout -b v3.3.4
CFLAGS="-Wno-error" go install .
```