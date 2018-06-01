pipeline {
  agent any
  stages {
    stage('BUILD-jtest1') {
      steps {
        sh '''echo ":: Sending greetings..."

echo "Hello World"
'''
        sh '''GOLANG_BUILD_CONTAINER=${GOLANG_BUILD_CONTAINER:-"jtest1-builder"}

echo ":: Cleaning up old $GOLANG_BUILD_CONTAINER container"
sudo docker rm -f $GOLANG_BUILD_CONTAINER || /bin/true
'''
        sh '''GOLANG_BUILD_CONTAINER=${GOLANG_BUILD_CONTAINER:-"jtest1-builder"}

echo ":: Building inside $GOLANG_BUILD_CONTAINER container"

sudo docker run --name $GOLANG_BUILD_CONTAINER --rm -i \\
    -v $WORKSPACE:/go/src/github.com/zoxpx/jtest1 \\
    -w /go/src/github.com/zoxpx/jtest1 golang \\
  go test -v'''
      }
    }
    stage('Post results') {
      steps {
        echo 'Me-ssagge'
      }
    }
  }
  environment {
    Fua = 'Bar'
  }
}