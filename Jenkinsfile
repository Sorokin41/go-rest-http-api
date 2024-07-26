pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                script {
                    git url: 'https://github.com/Sorokin41/go-rest-http-api.git', branch: 'main'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    sh 'docker-compose up -d'
                }
            }
        }
    }
}