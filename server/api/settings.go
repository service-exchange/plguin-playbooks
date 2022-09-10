package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/service-exchange/service-exchange-plugin-playbooks/client"
	"github.com/service-exchange/service-exchange-plugin-playbooks/server/config"

	pluginapi "github.com/service-exchange/service-exchange-plugin-api"
)

// SettingsHandler is the API handler.
type SettingsHandler struct {
	*ErrorHandler
	pluginAPI *pluginapi.Client
	config    config.Service
}

// NewSettingsHandler returns a new settings api handler
func NewSettingsHandler(router *mux.Router, api *pluginapi.Client, configService config.Service) *SettingsHandler {
	handler := &SettingsHandler{
		ErrorHandler: &ErrorHandler{},
		pluginAPI:    api,
		config:       configService,
	}

	settingsRouter := router.PathPrefix("/settings").Subrouter()
	settingsRouter.HandleFunc("", handler.getSettings).Methods(http.MethodGet)

	return handler
}

func (h *SettingsHandler) getSettings(w http.ResponseWriter, r *http.Request) {
	cfg := h.config.GetConfiguration()
	settings := client.GlobalSettings{
		EnableExperimentalFeatures: cfg.EnableExperimentalFeatures,
	}

	ReturnJSON(w, &settings, http.StatusOK)
}
