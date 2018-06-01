pipeline {
  agent any
  stages {
    stage('') {
      steps {
        sh '''echo "Hello World"

env GOOS=linux GOARCH=amd64 GO15VENDOREXPERIMENT=1 go test -v'''
      }
    }
  }
  environment {
    Fua = 'Bar'
  }
}
