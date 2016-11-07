# gowebhook
A bare bones example where I demonstrate how to use the [Go Facebook Graph API SDK] (https://github.com/huandu/facebook) to set up a webhook to a Facebook App.

## Setup
+ Install Golang
+ Create an [App] (https://developers.facebook.com/docs/apps/register)
+ Install [Heroku CLI] (https://devcenter.heroku.com/articles/heroku-command-line)
+ Clone this repo to your $GOPATH correctly
+ `cd` into this repo and run `go get github.com/tools/godep`
+ run `godep save`. This tells Heroku which packages are imported in the application
+ add Go buildpack `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
+ `git push heroku master`
+ open your application on the specified URL

Voila!

Inspired by
+ [Building Webapps With Go] (https://codegangsta.gitbooks.io/building-web-apps-with-go/content/deployment/index.html)
+ [Socketloop tutorial] (https://socketloop.com/tutorials/golang-login-authenticate-with-facebook-example)


