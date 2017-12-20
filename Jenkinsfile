pipeline {
  agent any
  stages {
    stage('buildTheJob') {
      steps {
        sh 'go build'
        sh 'echo "hello"'
      }
    }
  }
}