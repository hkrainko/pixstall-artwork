package http

import (
	"github.com/gin-gonic/gin"
	get_artworks "pixstall-artwork/app/artwork/delivery/model/get-artworks"
	error2 "pixstall-artwork/domain/error"
)

type ArtworkController struct {

}

func NewArtworkController() ArtworkController {
	return ArtworkController{

	}
}

func (a ArtworkController) GetArtworks(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(get_artworks.NewErrorResponse(error2.UnAuthError))
		return
	}
}