oklog ingeststore -store.segment-replication-factor 1 &
docker run -d --network=host --name="logspout"   --rm --volume=/var/run/docker.sock:/var/run/docker.sock        gliderlabs/logspout     syslog+tcp://0.0.0.0:7652
