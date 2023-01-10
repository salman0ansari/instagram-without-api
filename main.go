package main

import (
	"fmt"
)

var (
	Cookie    = ``
	Useragent = ""
	Appid     = ""
)

func igUsername(option requestOptions) []Item {
	var url = fmt.Sprintf(`https://i.instagram.com/api/v1/users/web_profile_info/?username=%v`, option.id)
	list := callApi(url, option.headers, option)
	return list
}

func main() {
	_ = igUsername(requestOptions{
		id: "instagram", // change this to any username
		headers: map[string]string{
			"cookie":      Cookie,
			"user-agent":  Useragent,
			"x-ig-app-id": Appid,
		},
		maxImages: 1, // max images is 12
	})
	// for _, v := range result {
	// 	fmt.Println(v)
	// }
}
