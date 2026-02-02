package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var firstNames = []string{
	"Chad", "Brent", "Matt", "Kyle", "Zack", "Travis", "Hunter", "Cody",
	"Blake", "Derek", "Trevor", "Dustin", "Tanner", "Bryce", "Connor",
	"Elon", "Sam", "Marc", "Peter", "Dario", "Demis", "Sundar", "Satya",
}

var lastNameParts1 = []string{
	"Cheese", "Block", "Crypto", "Cloud", "Data", "Quantum", "Synergy",
	"Pivot", "Disrupt", "Moon", "Rocket", "Hustle", "Grind", "Hack",
	"Stack", "Token", "Node", "Scale", "Venture", "Alpha", "Sigma",
}

var lastNameParts2 = []string{
	"flo", "man", "berg", "stein", "worth", "ton", "ski", "son",
	"bro", "dude", "chief", "lord", "master", "guru", "ninja",
	"wizard", "whale", "shark", "hawk", "wolf", "fox",
}

var timeframes = []string{
	"in the next 6 months", "within 2 years", "by next Tuesday",
	"before my Series B closes", "in exactly 18 months", "by Q4",
	"within the next 3 sprints", "before the heat death of the universe",
	"in 20 months", "by 2027", "next week probably", "any day now",
	"in like 5 years max", "before my stock options vest",
}

var predictions = []string{
	"AGI is coming %s!",
	"AI will replace all developers %s!",
	"We'll achieve superintelligence %s!",
	"AI will solve climate change %s!",
	"Robots will do all human jobs %s!",
	"We're definitely hitting the singularity %s!",
	"AI consciousness will emerge %s!",
	"Human-level AI is guaranteed %s!",
	"AI will cure all diseases %s!",
	"We won't need to work anymore %s!",
}

var buzzwords = []string{
	"leveraging synergies", "disrupting the paradigm", "scaling horizontally",
	"moving fast and breaking things", "thinking 10x", "building in public",
	"shipping MVPs", "iterating rapidly", "democratizing intelligence",
	"unlocking value", "maximizing engagement", "crushing it",
	"pivoting to AI", "going viral", "optimizing for growth",
	"embracing the grindset", "manifesting success", "hacking the matrix",
}

var openings = []string{
	"In a groundbreaking LinkedIn post,",
	"During a 47-tweet thread,",
	"In an exclusive podcast appearance,",
	"While keynoting at DisruptCon 2026,",
	"In a viral TikTok,",
	"During a 3AM Twitter Spaces session,",
	"In a Medium article with 47 claps,",
	"While microdosing at a WeWork,",
	"In a Substack nobody asked for,",
	"During a fireside chat with himself,",
}

var closings = []string{
	"When asked for evidence, they pointed to vibes.",
	"The prediction was made with zero technical background.",
	"Investors immediately added $500M to their valuation.",
	"Their crypto token pumped 400% on the news.",
	"They then pivoted to selling an AI course for $4,999.",
	"The thread ended with a link to their Calendly.",
	"They are definitely not trying to raise a Series C.",
	"This marks their 47th prediction this quarter.",
	"They assured everyone this time is different.",
	"Sources confirm they have never written a line of code.",
}

var filler = []string{
	"This is not financial advice, but definitely invest in my startup.",
	"The future belongs to those who build. And by build, I mean tweet.",
	"We're not just changing the game, we're inventing a new game entirely.",
	"If you're not thinking about AI 24/7, you're already obsolete.",
	"The only limit is your mindset. And maybe physics. But mostly mindset.",
	"I've been saying this for weeks. Well, days. Hours actually.",
	"My unpopular opinion: [extremely popular opinion].",
	"Let that sink in. 🧵👇",
}

func generateName() string {
	first := firstNames[rand.Intn(len(firstNames))]
	last := lastNameParts1[rand.Intn(len(lastNameParts1))] + lastNameParts2[rand.Intn(len(lastNameParts2))]
	return first + " " + last
}

func generateHeadline() string {
	prediction := predictions[rand.Intn(len(predictions))]
	timeframe := timeframes[rand.Intn(len(timeframes))]
	return fmt.Sprintf(prediction, timeframe)
}

func generateArticle() string {
	var sb strings.Builder

	name := generateName()
	headline := generateHeadline()

	sb.WriteString("╔══════════════════════════════════════════════════════════════╗\n")
	sb.WriteString("║                    🚀 BREAKING AI NEWS 🚀                    ║\n")
	sb.WriteString("╚══════════════════════════════════════════════════════════════╝\n\n")

	sb.WriteString(fmt.Sprintf("📰 HEADLINE: %s\n\n", headline))
	sb.WriteString(fmt.Sprintf("✍️  By: %s, Self-Proclaimed AI Thought Leader\n", name))
	sb.WriteString(fmt.Sprintf("📍 Location: Probably a Tesla or a podcast studio\n"))
	sb.WriteString(fmt.Sprintf("📅 Date: %s\n\n", time.Now().Format("January 2, 2006")))

	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	// Opening
	opening := openings[rand.Intn(len(openings))]
	sb.WriteString(fmt.Sprintf("%s %s declared that \"%s\"\n\n", opening, name, headline))

	// Body with buzzwords
	numBuzzwords := 2 + rand.Intn(3)
	sb.WriteString("\"We're ")
	for i := 0; i < numBuzzwords; i++ {
		buzzword := buzzwords[rand.Intn(len(buzzwords))]
		if i == numBuzzwords-1 {
			sb.WriteString("and " + buzzword)
		} else {
			sb.WriteString(buzzword + ", ")
		}
	}
	sb.WriteString(",\" they added.\n\n")

	// Filler quote
	sb.WriteString(fmt.Sprintf("\"%s\"\n\n", filler[rand.Intn(len(filler))]))

	// Closing
	sb.WriteString(closings[rand.Intn(len(closings))] + "\n\n")

	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	sb.WriteString("🔔 Like and subscribe for more totally real AI news!\n")
	sb.WriteString("💎 Not affiliated with any actual journalism whatsoever.\n")

	return sb.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateArticle())
}
