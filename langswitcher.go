package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	q "langswitch.com/quickstart"
)

func connectAPI() (*gmail.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile("quickstart/credentials_lang.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSettingsBasicScope, gmail.GmailSettingsSharingScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := q.GetClient(config)

	return gmail.NewService(ctx, option.WithHTTPClient(client))
}

func switchLanguageGmail(srv *gmail.Service) error {
	lang, err := srv.Users.Settings.GetLanguage("me").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve language settings: %v", err)
	}

	switch lang.DisplayLanguage {
	case "es":
		lang.DisplayLanguage = "fr"
	case "fr":
		lang.DisplayLanguage = "es"
	}

	// switch language for my account for learning purposes and i am to lazy to do it manual :DDD
	_, err = srv.Users.Settings.UpdateLanguage("me", lang).Do()
	return err

}

func main() {

	// connect to gmail service
	srv, err := connectAPI()
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	// change my language preferences
	err = switchLanguageGmail(srv)
	if err != nil {
		log.Fatalf("Couldn't switch language %v", err)
	}

}
