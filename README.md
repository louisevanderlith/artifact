# artifact
Mango API: Artifact

All large objects used by any mango module will be manage via the Artifact API.

## Run with Docker
*$ go build
*$ docker build -t avosa/artifact:latest .
*$ docker rm artifactDEV
*$ docker run -d -e RUNMODE=DEV -p 8082:8082 --network mango_net --name ArtifactDEV avosa/artifact:latest 
*$ docker logs artifactDEV