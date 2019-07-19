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

