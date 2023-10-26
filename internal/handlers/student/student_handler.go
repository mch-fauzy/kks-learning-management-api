package student

import (
	"net/http"
	"strconv"

	"github.com/kks-learning-management-api/internal/domain/student/model/dto"
	"github.com/kks-learning-management-api/transport/http/response"
)

// ViewStudent View all student
// @Summary View student
// @Description This endpoint for view all student information
// @Tags student
// @Produce json
// @Param page query string false "Number of page"
// @Param pageSize query string false "Total data per Page"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/students [get]
func (h *StudentHandler) ViewStudent(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	request := dto.BuildViewProductRequest(page, pageSize)
	err := request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, metadata, err := h.StudentService.GetStudentList(request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithMetadata(w, http.StatusOK, result, metadata)
}
