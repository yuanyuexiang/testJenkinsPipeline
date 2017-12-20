pipeline {
  agent any
  stages {
    stage('buildTheJob') {
      steps {
        sh 'GOOS=windows GOARCH=amd64 go build -o jacocoFileCopy.exe'
        sh 'echo "hello"'
        archiveArtifacts 'jacocoFileCopy.exe'
      }
    }
  }
}