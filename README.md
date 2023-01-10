## 📦 How to Run?

```bash
$ git clone https://github.com/salman0ansari/instagram-without-api
# before running (setup variable in main.go)
$ go run .
```

## 🍪 How to get Instagram Cookie

- Login to Instagram Web
- Go to your https://instagram/yourUsername
- Open your Browser Console (on Chrome just pressing F12)
  1. Select the "Network" tab
  2. Search and click on "timeline/" file; if it is empty just refresh the page
  3. Select "Headers" bar
  4. Be sure the file is Request Method "POST" (if it is "OPTIONS" search the other "timeline/" file in the list)
  5. Scroll down and select "Request Headers" tab
  6. Copy all the code after the word "cookie: " and paste it on `Cookie` variable
  7. Copy all the code after the word "user-agent: " and paste it on `Useragent` variable
  8. Copy all the code after the word "x-ig-app-id: " and paste it on `Appid` variable

```diff
- don't share your cookie code with anyone!!! it is the same of your credentials
```

- That's it, enjoy :)

## 📝 JSON Outputs

- Output example for `igUsername` function

```json
[
  {
    "id": "3012042985755280866",
    "time": 1673283526,
    "imageUrl": "https://scontent-del1-1.cdninstagram.com/v/t51.2885-15/324246910_1506475979847111_5502903366025676939_n.jpg?stp=dst-jpg_e35_p1080x1080&_nc_ht=scontent-del1-1.cdninstagram.com&_nc_cat=1&_nc_ohc=JiQBUhHNRR8AX9ECN3t&edm=AOQ1c0wBAAAA&ccb=7-5&oh=00_AfA-r5M1K9QqvoHQRoNbZBU_i-OhBu36qzLSYrXoxh7JKQ&oe=63C2E2EC&_nc_sid=8fd12b",
    "likes": 300603,
    "comments": 7961,
    "link": "https://www.instagram.com/p/CnM7SMWJoXi/",
    "text": "My Indigenous heritage comes even before my name. It has shaped how I view the world. It is everything to me.” —Scientist and educ....."
  }
]
```

Original Project in Node.js Here [instagram-without-api-node](https://github.com/orsifrancesco/instagram-without-api-node)
