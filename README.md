# artifact
Mango API: Artifact

All large objects used by any mango module will be manage via the Artifact API.

## Run with Docker
*$ go build
*$ docker build -t avosa/artifact:dev .
*$ docker rm artifactDEV
*$ docker run -d -p 8082:8082 --network mango_net --name artifactDEV avosa/artifact:dev 
*$ docker logs artifactDEV