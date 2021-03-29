package msg

import (
	"pixstall-artwork/domain/artwork/model"
	model2 "pixstall-artwork/domain/commission/model"
)

type CompletedCommission struct {
	model2.Commission `json:",inline"`
}

func (c CompletedCommission) ToArtworkCreator() model.ArtworkCreator {

	dayUsed := c.CompleteTime.Sub(c.CreateTime)

	return model.ArtworkCreator{
		CommissionID:         c.ID,
		OpenCommissionID:     c.OpenCommissionID,
		ArtistID:             c.ArtistID,
		ArtistName:           c.ArtistName,
		ArtistProfilePath:    c.ArtistProfilePath,
		RequesterID:          c.RequesterID,
		RequesterName:        c.RequesterName,
		RequesterProfilePath: c.RequesterProfilePath,
		DayUsed:              dayUsed,
		Size:                 c.Size,
		Volume:               0,
		Resolution:           c.Resolution,
		Format:               c.ExportFormat,
		IsR18:                c.IsR18,
		Anonymous:            c.Anonymous,
		DisplayImagePath:     *c.DisplayImagePath,
		CompletionFilePath:   *c.CompletionFilePath,
		Rating:               *c.Rating,
		Comment:              c.Comment,
		CompleteTime:         c.CreateTime,
	}
}
