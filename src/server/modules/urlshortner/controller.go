package urlshortner

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
	"urlshortener/server/libs/common"
	"urlshortener/server/libs/response"
	"urlshortener/server/libs/utils"
	"urlshortener/server/middleware/urlstore"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Adhar string `json:"aadharNumber" binding:"required"`
}

var Store = urlstore.NewURLStore()

func VisitURL(ctx *gin.Context) {
	url := ctx.Param("url")
	bslongURL, found := Store.Get(url)
	if !found {
		response.Send404(ctx)
		return
	}
	fmt.Println(bslongURL)
	longURL, err := base64.StdEncoding.DecodeString(bslongURL)
	if err != nil {
		response.Send422(ctx)
		return
	}
	fmt.Println(string(longURL))
	response.Send200(ctx, string(longURL))
}

func GenerateShortURL(ctx *gin.Context) {
	encurl := ctx.Param("url")
	burl, err := base64.StdEncoding.DecodeString(encurl)
	if err != nil {
		response.Send400(ctx, "Invalid url")
		return
	}
	url := string(burl)
	fmt.Println(url)
	if !utils.IsValidURL(&url) {
		response.Send400(ctx, "Invalid url")
		return
	}

	expiry := ctx.Param("expiry")
	exp := common.Config.App.Expiry

	if expiry != "" {
		if utils.IsValidExpiry(&expiry) {
			response.Send400(ctx, "Invalid expiry")
			return
		}
		x, err := strconv.Atoi(expiry)
		if err != nil {
			response.Send400(ctx, "Invalid expiry")
			return
		}
		exp = int64(x)
	}
	exp = exp * 60
	encodedURL := base64.StdEncoding.EncodeToString([]byte(url))
	us := &urlstore.URLSchema{
		LongURL: encodedURL,
		Timeout: time.Now().Unix() + exp,
	}
	short := utils.GenerateShortURL()
	Store.Save(us, short)
	req := ctx.Request
	scheme := "http"
	if req.URL.Scheme == "" {
		if req.TLS != nil {
			scheme = "https"
		}
	} else {
		scheme = req.URL.Scheme
	}
	completeURL := fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI)
	trimmedurl := utils.TrimURL(completeURL)
	response.Send200(ctx, fmt.Sprintf("%s%s", trimmedurl, short))
}

func DeleteURL(ctx *gin.Context) {
	url := ctx.Param("url")
	Store.Delete(url)
	response.Send200(ctx, "URL deleted sucessfully")
}
