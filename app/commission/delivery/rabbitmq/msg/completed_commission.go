package msg

import (
	model2 "pixstall-artwork/domain/commission/model"
)

type CompletedCommission struct {
	model2.Commission `json:",inline"`
}
