node {
    stage('Checkout') {
      checkout scm 
    }
    stage ("build"){
        sh 'pwd'
        sh 'ls'
        sh 'go build main.go'
    }
}
