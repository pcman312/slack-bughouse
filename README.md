# Bughouse Team Creator

### Configuring Slack
1. Create new slack app: https://api.slack.com/apps/
1. Create new slash command in the app
  1. Command `/bughouse`
  1. Request url (http://_\<url\_for\_container\>_/teams)
  1. Give a meaningful short description
  1. Usage hint should be `player1 player2 player3 player4`

### Building / Running
1. To pull the latest version
  1. `docker pull jturpin/slack-bughouse`
1. To run 
  1. `docker run -d -p 80:9090 jturpin/slack-bughouse`
  1. note that the container port is `9090` we run ours behind an AWS ELB to easily add HTTPS support

### Testing
This is very much a quick hack to get back to bughouse playing quicker. That said, output can be easily seen via curl

1. Start container 
  1. `docker run -d -p 80:9090 jturpin/slack-bughouse`
1. Use curl to hit endpoint
  1. `curl -v http://localhost/test?text=player1%20player2%20player3%20player4`
  1. Ensure status code 200
  1. Verify Content-Type: `application/json`
  1. Verify valid json returned

#### Questions
Find [me](https://twitter.com/jim_turpin) on Twitter.