package handler

import (
	"fmt"
	"net/http"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/pkg/utils"
)

func (h *Handler) increment(w http.ResponseWriter, r *http.Request) {
	var incrementInput entity.IncrementDTO
	if err := utils.BindJSON(r.Body, &incrementInput); err != nil {
		err = h.handleError(w, http.StatusBadRequest, fmt.Errorf("%s: %w", jsonBindErr, err))
		if err != nil {
			h.Logger.Error(err.Error())
		}
		return
	}

	h.Logger.Info(fmt.Sprintf("increment start: %+v", incrementInput))

	incrementResult, err := h.services.MathOperation.Increment(h.ctx, &incrementInput)
	if err != nil {
		err = h.handleError(w, http.StatusInternalServerError, fmt.Errorf("increment error: %w", err))
		if err != nil {
			h.Logger.Error(err.Error())
		}
		return
	}

	err = utils.SendJSON(w, http.StatusOK, utils.JSON{"value": incrementResult.Value})
	if err != nil {
		h.Logger.Error(fmt.Errorf("%s: %w", jsonSendErr, err).Error())
		return
	}
	h.Logger.Info(fmt.Sprintf("increment success: %+v", incrementResult))
}
