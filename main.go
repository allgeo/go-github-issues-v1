package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	token := os.Getenv("GITHUB_ACCESS_TOKEN")

	var owner string
	fmt.Print("Enter the owner of the repository: ")
	fmt.Scanln(&owner)

	var repo string
	fmt.Print("Enter the repository name: ")
	fmt.Scanln(&repo)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return
	}

	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request")
		return
	}

	defer resp.Body.Close()

	var issues []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		fmt.Printf("Error decoding response. Please make sure %v is the correct repo owner and %v is the correct repository\n", owner, repo)
		return
	}

	fmt.Println()
	color.Green("####################################################")
	color.Green("Retrieved %d issues from %s/%s", len(issues), owner, repo)
	color.Green("####################################################")
	fmt.Println()	

	for i, issue := range issues {
		color.Green("Issue %d\n", i+1)
		printColor("Issue Title: ", color.FgCyan)
		fmt.Println(issue["title"].(string))
		printColor("Issue URL: ", color.FgBlue)
		fmt.Println(issue["html_url"].(string))
		fmt.Println("----------------------------------------------------------------------------")
	}
}

func printColor(text string, colorCode color.Attribute) {
	color.New(colorCode).Print(text)
}
