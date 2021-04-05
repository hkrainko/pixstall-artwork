package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	err2 "pixstall-artwork/app/artwork/delivery/http/resp/err"
	"pixstall-artwork/app/artwork/delivery/http/resp/get-artwork"
	"pixstall-artwork/app/artwork/delivery/http/resp/get-artworks"
	"pixstall-artwork/domain/artwork"
	model2 "pixstall-artwork/domain/artwork/model"
	error2 "pixstall-artwork/domain/error"
	"strconv"
)

type ArtworkController struct {
	useCase artwork.UseCase
}

func NewArtworkController(useCase artwork.UseCase) ArtworkController {
	return ArtworkController{
		useCase: useCase,
	}
}

func (a ArtworkController) GetArtworks(ctx *gin.Context) {
	var tokenUserID *string
	if s := ctx.GetString("userId"); s != "" {
		tokenUserID = &s
	}
	filter, err := getFilter(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err2.NewErrorResponse(error2.BadRequestError))
		return
	}
	sorter, err := getSorter(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err2.NewErrorResponse(error2.BadRequestError))
		return
	}
	artworks, err := a.useCase.GetArtworks(ctx, *filter, *sorter)
	if err != nil {
		ctx.AbortWithStatusJSON(err2.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_artworks.NewResponse(*artworks, tokenUserID, filter.Offset, filter.Count))
}

func (a ArtworkController) GetArtwork(ctx *gin.Context) {
	var tokenUserID *string
	if s := ctx.GetString("userId"); s != "" {
		tokenUserID = &s
	}
	artworkID := ctx.Param("id")
	if artworkID == "" {
		ctx.AbortWithStatusJSON(err2.NewErrorResponse(error2.BadRequestError))
		return
	}
	dArtwork, err := a.useCase.GetArtwork(ctx, artworkID)
	if err != nil {
		ctx.AbortWithStatusJSON(err2.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_artwork.NewResponse(*dArtwork, tokenUserID))
}

func (a ArtworkController) UpdateArtwork(ctx *gin.Context) {
	artworkID := ctx.Param("id")
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
		return
	}
	updater := model2.ArtworkUpdater{
		ID:                   artworkID,
	}
	if title, exist := ctx.GetPostForm("title"); exist {
		updater.Title = &title
	}
	if textContent, exist := ctx.GetPostForm("textContent"); exist {
		updater.TextContent = &textContent
	}
	if state, exist := ctx.GetPostForm("state"); exist {
		st := model2.ArtworkState(state)
		updater.State = &st
	}
	err := a.useCase.UpdateArtwork(ctx, tokenUserID, updater)
	if err != nil {

	}

}

func (a ArtworkController) GetArtworkFavors(ctx *gin.Context) {

}

func (a ArtworkController) UpdateArtworkFavors(ctx *gin.Context) {

}

func getFilter(ctx *gin.Context) (*model2.ArtworkFilter, error) {

	filter := model2.ArtworkFilter{
		ArtistID:    nil,
		RequesterID: nil,
		Count:       0,
		Offset:      0,
	}
	if artistId, exist := ctx.GetQuery("artistId"); exist {
		filter.ArtistID = &artistId
	}
	if requesterId, exist := ctx.GetQuery("requesterId"); exist {
		filter.RequesterID = &requesterId
	}
	if count, exist := ctx.GetQuery("count"); exist {
		if countInt, err := strconv.Atoi(count); err == nil {
			filter.Count = countInt
		} else {
			return nil, errors.New("bad request")
		}
	} else {
		return nil, errors.New("bad request")
	}
	if offset, exist := ctx.GetQuery("offset"); exist {
		if countInt, err := strconv.Atoi(offset); err == nil {
			filter.Offset = countInt
		} else {
			return nil, errors.New("bad request")
		}
	} else {
		return nil, errors.New("bad request")
	}

	return &filter, nil
}

func getSorter(ctx *gin.Context) (*model2.ArtworkSorter, error) {
	sorter := model2.ArtworkSorter{}
	if sortBy, exist := ctx.GetQueryArray("sortBy"); exist {
		if len(sortBy) < 2 {
			return nil, errors.New("bad request")
		}
		var asc bool
		switch sortBy[1] {
		case "asc":
			asc = true
		case "dsc":
			asc = false
		default:
			return nil, errors.New("bad request")
		}
		switch sortBy[0] {
		case "price":
			sorter.Price = &asc
		case "state":
			sorter.State = &asc
		case "createTime":
			sorter.CreateTime = &asc
		case "lastUpdateTime":
			sorter.LastUpdateTime = &asc
		default:
			return nil, errors.New("bad request")
		}
		return &sorter, nil
	} else {
		return &sorter, nil
	}
}
