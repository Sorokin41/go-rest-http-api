pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Deploy') {
            steps {
                script {
                    // Развертываем приложение с помощью Docker Compose
                    sh 'docker-compose up -d'
                }
            }
        }
}