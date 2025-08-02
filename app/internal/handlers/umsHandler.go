package handlers
import (
	"net/http"
	"app/internal/services"
)
type UmsHandler struct {
	umsService *services.UmsService
}	
func NewUmsHandler(umsService *services.UmsService) *UmsHandler {
	return &UmsHandler{umsService: umsService}
}
func (h *UmsHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
}