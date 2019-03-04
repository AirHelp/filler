#!groovy
@Library('jenkins-pipeline-library') _

def label = "filler-${UUID.randomUUID().toString()}"

podTemplate(label: label, yaml: readTrusted('JenkinsPods.yaml')) {
  node(label) {
    gitCheckout()

    ci {
      stage('download Go deps') {
        container('golang'){
          sh 'apk add --no-cache git'
          sh 'go mod download'
        }
      }

      stage('go test') {
        container('golang'){
          sh 'go test -cover -v ./...'
        }
      }

      stage('go formatting') {
        container('golang'){
          sh 'gofmt -s -w .'
        }
      }

      stage('go vet') {
        container('golang'){
          sh 'go vet -v ./...'
        }
      }
    }

  }
}
