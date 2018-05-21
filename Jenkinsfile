pipeline {
  agent {
    label 'go_build'
  }
  stages {
    stage('Pre-check') {
      steps {
        sh 'go version'
      }
      post {
        failure {
          echo 'Build envorinment error: golang is missing...'
        }
        success {
          echo 'Build envorinment: OK '
        }

      }
    }
    stage('Build') {
      steps {
        sh 'go build app.go'
      }
      post {
        failure {
          echo 'Build error...'
        }
        success {
          echo 'Build: OK '
        }
      }
    }
    stage('Test') {
      steps {
        sh 'go test -v'
      }
      post {
        failure {
          echo 'Test error...'

        }
        success {
          echo 'Test: OK '
        }
      }
    }
  }
}