package model

import (
	"pixstall-artwork/domain/artwork/model"
	"time"
)

type Commission struct {
	ID               string `json:"id" bson:"id"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	Price                          model.Price `json:"price" bson:"price"`
	DayNeed                        int         `json:"dayNeed" bson:"dayNeed"`
	Size                           *model.Size `json:"size" bson:"size,omitempty"`
	Resolution                     *float64    `json:"resolution" bson:"resolution,omitempty"`
	ExportFormat                   *string     `json:"exportFormat" bson:"exportFormat,omitempty"`
	Desc                           string      `json:"desc" bson:"desc"`
	PaymentMethod                  string      `json:"paymentMethod" bson:"paymentMethod"`
	IsR18                          bool        `json:"isR18" bson:"isR18"`
	BePrivate                      bool        `json:"bePrivate" bson:"bePrivate"`
	Anonymous                      bool        `json:"anonymous" bson:"anonymous"`
	RefImagePaths                  []string    `json:"refImagePaths" bson:"refImagePaths"`
	TimesAllowedDraftToChange      *int        `json:"timesAllowedDraftToChange" bson:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int        `json:"timesAllowedCompletionToChange" bson:"timesAllowedCompletionToChange"`
	DraftChangingRequestTime       int         `json:"draftChangingRequestTime" bson:"draftChangingRequestTime"`
	ProofCopyRevisionRequestTime   int         `json:"proofCopyRevisionRequestTime" bson:"proofCopyRevisionRequestTime"`

	ProofCopyImagePaths []string `json:"proofCopyImagePaths" bson:"proofCopyImagePaths"`
	DisplayImagePath    *string  `json:"displayImagePath,omitempty" bson:"displayImagePath,omitempty"`
	CompletionFilePath  *string  `json:"completionFilePath,omitempty" bson:"completionFilePath,omitempty"`
	Rating              *int     `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment             *string  `json:"comment,omitempty" bson:"comment,omitempty"`

	CreateTime     time.Time       `json:"createTime" bson:"createTime"`
	CompleteTime   *time.Time      `json:"completeTime" bson:"completeTime,omitempty"`
	LastUpdateTime time.Time       `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          CommissionState `json:"state" bson:"state"`
}

type CommissionState string

const (
	CommissionStatePendingValidation                       CommissionState = "PENDING_VALIDATION"
	CommissionStateInvalidatedDueToOpenCommission          CommissionState = "INVALIDATED_DUE_TO_OPEN_COMMISSION"
	CommissionStateInvalidatedDueToUsers                   CommissionState = "INVALIDATED_DUE_TO_USERS"
	CommissionStatePendingArtistApproval                   CommissionState = "PENDING_ARTIST_APPROVAL"
	CommissionStatePendingRequesterModificationValidation  CommissionState = "PENDING_REQUESTER_MODIFICATION_VALIDATION"
	CommissionStateInProgress                              CommissionState = "IN_PROGRESS"
	CommissionStatePendingRequesterAcceptance              CommissionState = "PENDING_REQUESTER_ACCEPTANCE"
	CommissionStateDeclinedByArtist                        CommissionState = "DECLINED_BY_ARTIST"
	CommissionStateCancelledByRequester                    CommissionState = "CANCELED_BY_REQUESTER"
	CommissionStatePendingUploadProduct                    CommissionState = "PENDING_UPLOAD_PRODUCT"
	CommissionStatePendingUploadProductDueToRevisionExceed CommissionState = "PENDING_UPLOAD_PRODUCT_DUE_TO_REVISION_EXCEED"
	CommissionStatePendingRequesterAcceptProduct           CommissionState = "PENDING_REQUESTER_ACCEPT_PRODUCT"
	CommissionStateCompleted                               CommissionState = "COMPLETED"
)

func (c *Commission) ToArtworkCreator() model.ArtworkCreator {

	var dayUsed = c.CompleteTime.Sub(c.CreateTime)

	return model.ArtworkCreator{
		CommissionID:         c.ID,
		OpenCommissionID:     c.OpenCommissionID,
		ArtistID:             c.ArtistID,
		ArtistName:           c.ArtistName,
		ArtistProfilePath:    c.ArtistProfilePath,
		RequesterID:          c.RequesterID,
		RequesterName:        c.RequesterName,
		RequesterProfilePath: c.RequesterProfilePath,
		Price:                c.Price,

		DayUsed:              int(dayUsed),
		Size:                 c.Size,
		Volume:               0,
		Resolution:           0,
		Format:               "",

		IsR18:                c.IsR18,
		Anonymous:            c.Anonymous,
		DisplayImagePath:     *c.DisplayImagePath,
		CompletionFilePath:   *c.CompletionFilePath,
		Rating:               *c.Rating,
		Comment:              c.Comment,
		CompleteTime:         *c.CompleteTime,
	}
}
