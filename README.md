# artifact
Mango API: Artifact

All large objects used by any mango module will be manage via the Artifact API.

## Run with Docker
* $ docker build -t avosa/artifact:dev .
* $ docker rm ArtifactDEV
* $ docker run -d -e RUNMODE=DEV -p 8082:8082 --network mango_net --name ArtifactDEV avosa/artifact:dev 
* $ docker logs ArtifactDEV