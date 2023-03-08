package dto

import "hause/entity"

type SetupResponse struct {
	User      entity.User      `json:"user"`
	Portfolio entity.Portfolio `json:"portfolio"`
}
