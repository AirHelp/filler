#!groovy
@Library('jenkins-pipeline-library') _

def label = "filler-${UUID.randomUUID().toString()}"

podTemplate(label: label, yaml: readTrusted('JenkinsPods.yaml')) {
  node(label) {
    gitCheckout()

    ci {
      stage("Build test image") {
        container('docker') {
          sh "docker build -t local/fillertests -f Dockerfile-test ."
        }
      }

      stage("Run fmt and vet") {
        parallel (
          fmt: {
            container('docker') {
              sh "docker run --rm local/fillertests gofmt -s -w ."
            }
          },
          vet: {
            container('docker') {
              sh "docker run --rm local/fillertests go vet -v ./..."
            }
          }
        )
      }

      stage("Run tests") {
        container('docker') {
          sh "docker run --rm local/fillertests"
        }
      }
    }
  }
}
