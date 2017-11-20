pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }
    
  }
  stages {
    stage('Build') {
      agent {
        dockerfile {
          filename 'Dockerfile'
        }
        
      }
      steps {
        echo 'Starting Build'
        sh 'docker build -t jturpin/slack-bughouse .'
      }
    }
    stage('test') {
      agent {
        docker {
          image 'jturpin/slack'
        }
        
      }
      steps {
        sh '''docker run -d -p 8080:8080 jturpin/slack-bughouse
curl -v http://localhost:8080/teams?text=player1%20player2%20player3%20player4'''
      }
    }
  }
}