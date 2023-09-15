pipeline {

  environment {
    dockerimagename = ""
    dockerImage = ""
  }

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/kerygits/hello-devops-project'
      }
    }

    stage('Build image') {
      steps{
        script {
          dockerImage = docker.build dockerimagename
        }
      }
    }

    stage('Pushing Image') {
      environment {
               registryCredential = 'dockerhub-credentials'
           }
      steps{
        script {
          docker.withRegistry( 'https://registry.hub.docker.com', registryCredential ) {
            dockerImage.push("")
          }
        }
      }
    }

    stage('Deploying Go app container to Kubernetes') {
      steps {
      
      script {
          kubernetesDeploy(configs: "deployment.yaml, service.yaml", kubeconfigId: "kubeconfig")
  
            }
      	 
               
        }
      }
    }

  }


