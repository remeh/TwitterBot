/*
 *   Copyright 2015 Benoit LETONDOR
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/jsgoecke/go-wit"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	INTENT_HI           = "hi"
	INTENT_NICE_ARTICLE = "nice_article"
	INTENT_THANK_FOLLOW = "thank_follow"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func buildReply(tweet anaconda.Tweet) (string, error) {
	message := cleanTweetMessage(tweet.Text)
	if message == "" {
		return "", nil
	}

	// Process a text message
	request := &wit.MessageRequest{}
	request.Query = message

	result, err := witclient.Message(request)
	if err != nil {
		return "", err
	}

	// TODO(remy): quick fix
	if len(result.Outcomes) == 0 {
		return "", nil
	}

	outcome := result.Outcomes[0]
	intent := outcome.Intent
	if outcome.Confidence < 0.5 {
		log.Println("Not enough confidence for intent : " + intent)
		return "", nil
	}

	if intent == INTENT_HI {
		return buildHiIntentResponse(tweet), nil
	} else if intent == INTENT_NICE_ARTICLE {
		return buildNiceArticleIntentResponse(tweet), nil
	} else if intent == INTENT_THANK_FOLLOW {
		return buildThanksFollowIntentResponse(tweet), nil
	}

	return "", nil
}

func buildHiIntentResponse(tweet anaconda.Tweet) string {
	greetings := []string{"hello!", "hey", "yo"}

	return buildMention(tweet.User, greetings[rand.Intn(len(greetings))])
}

func buildNiceArticleIntentResponse(tweet anaconda.Tweet) string {
	greetings := []string{"hello!", "hey", "hi", "well,", ""}
	thanks := []string{"thanks", "thank you", "many thanks", "thx"}
	messages := []string{"reading", "your tweet", "your message"}

	greet := greetings[rand.Intn(len(greetings))]
	thank := thanks[rand.Intn(len(thanks))]
	message := messages[rand.Intn(len(messages))]

	return buildMention(tweet.User, greet+" "+thank+" for "+message)
}

func buildThanksFollowIntentResponse(tweet anaconda.Tweet) string {
	greetings := []string{"hello!", "hey", "hi", "well,", ""}
	thanks := []string{"thanks", "thank you", "many thanks", "thx"}
	follows := []string{"following me", "the follow"}
	messages := []string{"your message", "your tweet", "your mention"}
	reciprocals := []string{"too", "as well", ""}

	following, err := isUserFollowing(tweet.User.ScreenName)
	if following && err == nil {
		greet := greetings[rand.Intn(len(greetings))]
		thank := thanks[rand.Intn(len(thanks))]
		follow := follows[rand.Intn(len(follows))]
		reciprocal := reciprocals[rand.Intn(len(reciprocals))]

		return buildMention(tweet.User, greet+" "+thank+" for "+follow+" "+reciprocal)
	} else {
		greet := greetings[rand.Intn(len(greetings))]
		thank := thanks[rand.Intn(len(thanks))]
		message := messages[rand.Intn(len(messages))]

		return buildMention(tweet.User, greet+" "+thank+" for "+message)
	}
}

func buildIntro() string {
	surprise := []string{"wow", "hey", "huh", "hm", "hum", "huhu", "wao", "awesome", "nice"}
	personal := []string{"that is", "this is", "it's", "i think that it's", "is it", "it's me or it is", "i'm the only one who thinks that it is", "i'm the only one who thinks this is"}
	adj := []string{"awesome", "really awesome", "impressive", "great", "nice", "very nice", "neat", "very well done", "very great", "so great", "really great", "so cool", "really cool", "very cool", "so nice"}
	separators := []string{".", "!", "!!", "..."}

	rv := ""

	if yesorno() {
		rv += randomUpper(randomCapitalize(randomStr(surprise)))

		if yesorno() {
			rv += randomStr(separators)
		}

		rv += " "
	}

	if yesorno() {
		rv += randomCapitalize(randomStr(personal)) + " "
	}

	if len(strings.Trim(rv, " ")) == 0 {
		rv += randomCapitalize(randomStr(adj))
	} else {
		rv += randomStr(adj)
	}

	if yesorno() {
		rv += randomStr(separators)
	}

	return rv
}

func randomUpper(str string) string {
	if yesorno() {
		return strings.ToUpper(str)
	}
	return str
}

func randomCapitalize(str string) string {
	if yesorno() {
		if len(str) > 1 {
			return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
		}
	}
	return str
}

func yesorno() bool {
	return rand.Int()%2 == 0
}

func randomStr(values []string) string {
	r := rand.Int() % len(values)
	return values[r]
}

func buildMention(user anaconda.User, text string) string {
	return "@" + user.ScreenName + " " + text
}

func cleanTweetMessage(message string) string {
	cleaned := ""

	words := strings.Split(message, " ")
	for _, word := range words {
		if strings.HasPrefix(word, "@") || strings.HasPrefix(word, "http") {
			continue
		} else if strings.HasPrefix(word, "#") {
			cleaned += strings.TrimPrefix(word, "#") + " "
		}

		cleaned += word + " "
	}

	return cleaned
}
