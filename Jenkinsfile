node {
    stage('Checkout') {
      checkout scm 
    }
    stage ("build"){
        sh 'pwd'
        sh 'ls'
        sh 'cp main.go $GOPATH/src/'
        sh 'cd $GOPATH/src'
        sh 'pwd'
        sh 'ls'
        
        sh 'go get github.com/gorilla/mux'
        sh 'go build main.go'
    }
}
