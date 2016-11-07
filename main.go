package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"

	"github.com/julienschmidt/httprouter"
)

// Request object when an event we are subscribed to occurs
type Request struct {
	Entry  EntryStruct
	Object string
}

// EntryStruct object contained by Entry field of above struct
type EntryStruct struct {
	time          int
	id            string
	changedFields []string
	uid           string
}

// Init initialises the Facebook authentication for our application
func Init(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fbConfig := &oauth2.Config{
		ClientID:     "", // Your APP ID
		ClientSecret: "", // Your APP Secret
		RedirectURL:  "https://call_back_url.com/FBLogin",
		Scopes:       []string{"public_profile", "user_friends", "user_posts"}, // Your APP permissions
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}

	url := fbConfig.AuthCodeURL("")

	w.Write([]byte("<a href='" + url + "'><button>Login with Facebook</button></a>"))
}

// FBSubscribe subscribes the registered user to our application
func FBSubscribe(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := fb.Post("/YOUR_APP_ID/subscriptions", fb.Params{
		"object":       "user", // Can be user, page, permission or payments
		"access_token": "YOUR_APP_ACCESS_TOKEN", // copy the token which is generated from the Graph API explorer after selecting "Get App token" in the top right
		"callback_url": "https://call_back_url.com/facebook",
		"fields":       "feed, friends, likes, link, status, statuses, username", // What you subscribe to
		"verify_token": "token", //accessToken, make it more secure
	})

	if err != nil {
		fmt.Println(err)
	}
}

// GetFacebook endpoint is called by Facebook to verify the validity of
// the webhook callback URL
func GetFacebook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.FormValue("hub.mode") == "subscribe" &&
		r.FormValue("hub.verify_token") == "token" {
		w.Write([]byte(r.FormValue("hub.challenge")))
	} else {
		w.Write([]byte("400"))
	}
}

// PostFacebook endpoint is called whenever a change occurs on anything that
// the webhook has been subscribed to
func PostFacebook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// What our json request object typically looks likes
	/*2016/11/06 23:08:17 {"entry": [{"time": 1478473584,
	            "id": "1101298983238047",
	            "changed_fields": ["feed"
	                              ],
	            "uid": "1547842232012"}
	          ],
	  "object": "user"}
	*/
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("WWW")
	}
	log.Println(string(body))

	res := regexp.MustCompile("([^ ]+[ ][^ ]+)[ ](.*)") // Split after second space
	re := res.FindAllString(string(body), -1)
	s := ""
	for _, v := range re { // size of one
		fmt.Println(v)
		s = v
	}

	strs := strings.Split(s, ",")
	for _, i := range strs {
		fmt.Println(i)
	}

	// TODO construct JSON structs
	fmt.Println(strings.Split(strs[0], " ")[2])
	fmt.Println(strings.Split(strs[1], " ")[2])
	fmt.Println(strings.Split(strings.Split(strs[2], " ")[2], ",")[0]) // Needs to be cleaned
	fmt.Println(strings.Split(strs[3], " ")[2])
	fmt.Println(strings.Split(strs[4], " ")[2])

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := httprouter.New()

	r.GET("/", Init)
	r.GET("/FBSubscribe", FBSubscribe)

	// API endpoints for webhook
	r.GET("/facebook", GetFacebook)
	r.POST("/facebook", PostFacebook)

	http.ListenAndServe(":"+port, r)
}
