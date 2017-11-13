node {
    stage('Checkout') {
      checkout scm 
    }
    stage ("build"){
        sh 'pwd'
        sh 'ls'
        sh 'export GOPATH="/opt/bitnami/apps/jenkins/jenkins_home/workspace"'
        sh 'go get github.com/gorilla/mux'
        sh 'go build main.go'
    }
}
