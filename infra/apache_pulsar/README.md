This is the docker container for Apache Pulsar. This container is expected to be run only during the development locally.

## Build the apache pulsar docker container
`sudo docker build -t apache_pulsar .`

## Run apache pulsar
`sudo docker run -p 6650:6650 -p 8080:8080 apache_pulsar`

