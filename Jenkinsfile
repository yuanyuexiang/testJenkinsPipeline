pipeline {
  agent any
  stages {
    stage('buildTheJob') {
      steps {
        build 'go build'
        git(url: 'https://github.com/yuanyuexiang/testJenkinsPipeline.git', branch: 'master')
      }
    }
  }
}