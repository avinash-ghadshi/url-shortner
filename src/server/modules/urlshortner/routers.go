package urlshortner

import (
	"github.com/gin-gonic/gin"
)

func UshortApis(r *gin.RouterGroup) {
	r.GET("/getorigin/:url", VisitURL)
	r.GET("/getshort/:url/:expiry?", GenerateShortURL)
	r.DELETE("/delete/:url", DeleteURL)
}
