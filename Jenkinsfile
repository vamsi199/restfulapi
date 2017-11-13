node {
    stage('Checkout') {
      checkout scm 
    }
    stage ("build"){
        sh 'bash script.sh'
        sh 'go build main.go'
    }
}
