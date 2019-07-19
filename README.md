# Travel time information

This repository is part of bigger conceptual system that would
provide travel time estimations based on mobile devices identification.

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

## Bluetooth/WiFi/IMSI detectors aka. detector of mobile devices

Detectors are devices located around city, runways etc. which
are capable for intercepting Bluetooth, WiFi MAC addresses and
IMSI ids from passing travelers.

## NOTE

This service is purely conceptual however I did work with similar
system which was capable of intercepting Bluetooth MAC addresses.
Adding IMSI catcher is purely theoretical since they are 
basically illegal without special permissions.

# Technical details

[TBA]

### ETCD repository data model discussion

Storing events from detectors would include detector ID as root path
then bucketing based on unix timestamp with bucket size of 100 seconds:
```bash
${namespace}/detectors/<detector_id>/<unix_timestamp/100>/<device_id>
value: {"time": <rfc3339>}
```

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
links between detectors by declaring for each detector it's incoming
traffic peers.
```bash
${namespace}/<dest_detector_id>/<src_detector_id>
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