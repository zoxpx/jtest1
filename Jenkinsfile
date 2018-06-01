pipeline {
  agent any
  stages {
    stage('') {
      steps {
        sh '''echo "Hello World"

GOLANG_BUILD_CONTAINER="jtest1-builder"
sudo docker rm -f $GOLANG_BUILD_CONTAINER || /bin/true

sudo docker run --name $GOLANG_BUILD_CONTAINER --rm -i -v $WORKSPACE:/go/src/github.com/zoxpx/jtest1 -w /go/src/github.com/zoxpx/jtest1 golang \
    go test -v
'''
      }
    }
  }
  environment {
    Fua = 'Bar'
  }
}
