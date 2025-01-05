package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type PageData struct {
	IP string
}

func main() {
	// HTML template with a button
	tmpl := template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Find My Static IP</title>
</head>
<body>
    <h1>Your Static IP Address</h1>
    {{if .IP}}
        <p>Your Public IP: <strong>{{.IP}}</strong></p>
    {{else}}
        <form action="/" method="post">
            <button type="submit">Click to Find My Static IP</button>
        </form>
    {{end}}
</body>
</html>
`))

	// Handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Fetch the user's public IP
			ip, err := getPublicIP()
			if err != nil {
				http.Error(w, "Failed to fetch IP", http.StatusInternalServerError)
				return
			}
			data := PageData{IP: ip}
			tmpl.Execute(w, data)
		} else {
			// Render the page with no IP initially
			tmpl.Execute(w, PageData{})
		}
	})

	// Start the web server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Helper function to get the public IP of the user
func getPublicIP() (string, error) {
	// Use an external service to fetch the public IP
	resp, err := http.Get("https://api64.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
