pipeline {
  agent any
  stages {
    stage('buildTheJob') {
      parallel {
        stage('buildTheJob') {
          steps {
            sh 'GOOS=windows GOARCH=amd64 go build -o jacocoFileCopy.exe'
            sh 'echo "hello"'
            archiveArtifacts 'jacocoFileCopy.exe'
          }
        }
        stage('hello') {
          steps {
            sh 'echo "hello"'
          }
        }
      }
    }
  }
}