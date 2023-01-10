package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func imageToBase64(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(body), nil
}

func callApi(url string, headers map[string]string, option requestOptions) []Item {
	// ig username
	if option.id == "" {
		option.id = "instagram"
	}

	// image need to be fetch (max 12)
	if option.maxImages == 0 {
		option.maxImages = 12
	}

	// output file name
	if option.file == "" {
		option.file = "instagram-cache.json"
	}

	// preety json?
	if !option.pretty {
		option.pretty = false
	}

	//
	if option.time == 0 {
		option.time = 3600
	}

	// convert image to base64
	if !option.base64images {
		option.base64images = false
	}

	// required
	if option.headers == nil {
		option.headers = make(map[string]string)
	}

	result := []Item{}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var jsonBody map[string]interface{}
	json.Unmarshal(body, &jsonBody)

	data := jsonBody["data"].(map[string]interface{})
	user := data["user"].(map[string]interface{})
	edgeOwner := user["edge_owner_to_timeline_media"].(map[string]interface{})
	edges := edgeOwner["edges"].([]interface{})

	maxImages := option.maxImages
	base64images := option.base64images
	pretty := option.pretty
	file := option.file
	for _, el := range edges {
		el := el.(map[string]interface{})
		node := el["node"].(map[string]interface{})

		if maxImages > 0 && len(result) >= maxImages {
			break
		}

		var image string
		if base64images {
			image, _ = imageToBase64(node["display_url"].(string))
		}

		item := Item{
			ID:       node["id"].(string),
			Time:     node["taken_at_timestamp"].(float64),
			ImageURL: node["display_url"].(string),
			Likes:    node["edge_liked_by"].(map[string]interface{})["count"].(float64),
			Comments: node["edge_media_to_comment"].(map[string]interface{})["count"].(float64),
			Link:     "https://www.instagram.com/p/" + node["shortcode"].(string) + "/",
			Text:     node["edge_media_to_caption"].(map[string]interface{})["edges"].([]interface{})[0].(map[string]interface{})["node"].(map[string]interface{})["text"].(string),
			Image:    image,
		}
		result = append(result, item)
	}

	if pretty {
		jsonBytes, _ := json.MarshalIndent(result, "", "  ")
		os.WriteFile(file, jsonBytes, 0644)
	} else {
		jsonBytes, _ := json.Marshal(result)
		os.WriteFile(file, jsonBytes, 0644)
	}

	return result
}
