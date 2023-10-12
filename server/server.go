package server

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

var htmlPush = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func GinInitHttp2() *gin.Engine {
	r := gin.Default()
	r.SetHTMLTemplate(htmlPush)
	r.Static("/assets", "./assets")

	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"protocol": c.Request.Proto,
		})
	})

	r.GET("/push", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status":   "success",
			"protocol": c.Request.Proto,
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8000", "./server/key-ssl/server.pem", "./server/key-ssl/server.key")

	return r
}
