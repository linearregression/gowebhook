# gowebhook
A bare bones example where I demonstrate how to use the [Go Facebook Graph API SDK] (https://github.com/huandu/facebook) to set up a webhook to a Facebook App.

## Website Setup
+ Install Golang
+ Create an [App] (https://developers.facebook.com/docs/apps/register)
+ Install [Heroku CLI] (https://devcenter.heroku.com/articles/heroku-command-line)
+ Clone this repo to your $GOPATH correctly
+ `cd` into this repo and run `go get github.com/tools/godep`
+ Run `godep save`. This tells Heroku which packages are imported in the application
+ Add Go buildpack `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
+ Add and commit all generated files, then run `git push heroku master`
+ Open your application on the specified URL

## Facebook App Setup
+ Add the URL of the form foo.com to your App's App Domains found under Settings/Basic on your App dashboard.
+ On the same page click `Add Platform` and select website. Add the URL of the form https://foo.com/ as the site URL.
+ To enable authentication with Facebook, under Facebook Login/Settings add the URL of the form https://foo.com/ as your `Valid OAuth redirect URI`
+ Your setup is now complete and you will now be able to see the subscriptions under Webhooks on the dashboard.

In order to see activity live, run `heroku logs --tail` in the cloned repo, log yourself into the app, then create a post (or generate some other kind of activity which you are monitoring)!

Voila!

Inspired by
+ [Building Webapps With Go] (https://codegangsta.gitbooks.io/building-web-apps-with-go/content/index.html)
+ [Socketloop tutorial] (https://socketloop.com/tutorials/golang-login-authenticate-with-facebook-example)


