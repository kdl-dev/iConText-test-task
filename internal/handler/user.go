package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/pkg/utils"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var userInput entity.UserDTO
	if err := utils.BindJSON(r.Body, &userInput); err != nil {
		err = h.handleError(w, http.StatusBadRequest, fmt.Errorf("%s: %w", jsonBindErr, err))
		if err != nil {
			h.Logger.Error(err.Error())
		}
		return
	}

	h.Logger.Info(fmt.Sprintf("create user start: %+v", userInput))

	userOutput, err := h.services.User.CreateUser(context.Background(), &userInput)
	if err != nil {
		err = h.handleError(w, http.StatusInternalServerError, fmt.Errorf("create user error: %w", err))
		if err != nil {
			h.Logger.Error(err.Error())
		}
		return
	}

	err = utils.SendJSON(w, http.StatusOK, utils.JSON{"id": userOutput.ID})
	if err != nil {
		h.Logger.Error(fmt.Errorf("%s: %w", jsonSendErr, err).Error())
		return
	}
	h.Logger.Info(fmt.Sprintf("create user success: %+v", userOutput))
}
