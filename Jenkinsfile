pipeline {
    agent any

    environment {
        DOCKERHUB_USERNAME = 'mailtester'
        APP_NAME = 'go-web-app'
        IMAGE_TAG = '${BUILD_NUMBER}'
    }

    stages {
        stage('Code Checkout') {
            steps {
                echo 'Checking out code...'
                checkout scm
            }
        }

        stage('Build') {
            steps {
                echo 'Building...'
                sh '''
                go version
                go build -o app .
                '''
            }
        }

        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test ./...'
            }

        stage('Code Quality') {
            steps {
                echo 'Checking code quality...'
                sh '''
                curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.8.0
                $(go env GOPATH)/bin/golangci-lint run
                '''
            }
        }

        stage('Docker Login') {
            steps {
                echo 'Logging into Docker registry...'
                withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials', usernameVariable: 'DOCKERHUB_USERNAME', passwordVariable: 'DOCKERHUB_PASSWORD')]) {
                    sh 'echo $DOCKERHUB_PASSWORD | docker login -u $DOCKERHUB_USERNAME --password-stdin'
                }
            }
        }

        stage('Docker Build & Push') {
            steps {
                echo 'Building and pushing Docker image...'
                dockerImage = docker.build("$DOCKERHUB_USERNAME/$APP_NAME:$IMAGE_TAG")
                dockerImage.push()
            }
        }

        stage('Update Helm Charts') {
            steps {
                echo 'Updating Helm Charts Image tag and push to github...'
                withCredentials([usernamePassword(credentialsId: 'github-credentials', usernameVariable: 'GITHUB_USERNAME', passwordVariable: 'GITHUB_PASSWORD')]) {
                    sh '''
                    git config --global user.name "Dan George"
                    git config --global user.email "5060367+dangeorge@users.noreply.github.com"
                    git remote set-url origin https://$GITHUB_USERNAME:$GITHUB_PASSWORD@github.com/$GITHUB_USERNAME/$APP_NAME.git
                    git add .
                    git commit -m "Update Helm Charts Image tag"
                    git push origin HEAD:main
                    '''
            }
        }
    }
}