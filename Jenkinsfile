pipeline {
  agent any
  stages {
    stage('BUILD-jtest1') {
      steps {
        sh '''echo ":: Sending greetings..."

echo "Hello World"
'''
        sh '''echo ":: Cleaning up old $GOLANG_BUILD_CONTAINER container"
sudo docker rm -f $GOLANG_BUILD_CONTAINER || /bin/true
'''
        sh '''echo ":: Building inside $GOLANG_BUILD_CONTAINER container"

sudo docker run --name $GOLANG_BUILD_CONTAINER --rm -i \\
    -v $WORKSPACE:/go/src/github.com/zoxpx/jtest1 \\
    -w /go/src/github.com/zoxpx/jtest1 golang \\
  go test -v'''
      }
    }
    stage('Print sumting') {
      steps {
        echo 'Me-ssagge'
      }
    }
    post {
        always {
            echo 'One way or another, I have finished'
            deleteDir() /* clean up our workspace */
        }
        success {
            echo 'I succeeeded!'
        }
        unstable {
            echo 'I am unstable :/'
        }
        failure {
            echo 'I failed :('
        }
        changed {
            echo 'Things were different before...'
        }
    }
  }
  environment {
    GOLANG_BUILD_CONTAINER = 'jtest1-builder'
  }
}
