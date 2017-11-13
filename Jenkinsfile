node {
    stage('Checkout') {
        sh 'cd /home/bitnami/go/src'
      checkout scm 
    }
    stage ("build"){
        sh 'pwd'
        sh 'ls'
        sh 'go get github.com/gorilla/mux'
        sh 'go build main.go'
    }
}
