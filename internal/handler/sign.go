package handler

import (
	"fmt"
	"net/http"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/pkg/utils"
)

func (h *Handler) sha512Sign(w http.ResponseWriter, r *http.Request) {
	var inputSHA512 entity.HMACSHA512DTO
	if err := utils.BindJSON(r.Body, &inputSHA512); err != nil {
		err = h.handleError(w, http.StatusBadRequest, fmt.Errorf("%s: %w", jsonBindErr, err))
		if err != nil {
			h.Logger.Error(err.Error())
		}
		return
	}

	h.Logger.Info(fmt.Sprintf("signing start: %+v", inputSHA512))

	signature := h.services.Signature.SHA512Sign(&inputSHA512)

	err := utils.SendJSON(w, http.StatusOK, utils.JSON{"signature": signature.Value})
	if err != nil {
		h.Logger.Error(fmt.Errorf("%s: %w", jsonSendErr, err).Error())
		return
	}
	h.Logger.Info(fmt.Sprintf("signing success: %+v", signature))
}
