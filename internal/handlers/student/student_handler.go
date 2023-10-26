package student

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kks-learning-management-api/internal/domain/student/model/dto"
	"github.com/kks-learning-management-api/transport/http/response"
)

// ViewStudentById View student by ID
// @Summary View student by ID
// @Description This endpoint for view student information by ID
// @Tags student
// @Produce json
// @Param studentId path string true "Unique id of the student"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/students/{studentId} [get]
func (h *StudentHandler) ViewStudentById(w http.ResponseWriter, r *http.Request) {
	studentId := chi.URLParam(r, "studentId")

	request := dto.BuildViewStudentByIdRequest(studentId)
	result, err := h.StudentService.GetStudentById(request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, result)
}

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
// func (h *StudentHandler) ViewStudent(w http.ResponseWriter, r *http.Request) {
// 	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
// 	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

// 	request := dto.BuildViewStudentRequest(page, pageSize)
// 	err := request.Validate()
// 	if err != nil {
// 		response.WithError(w, err)
// 		return
// 	}

// 	result, metadata, err := h.StudentService.GetStudentList(request)
// 	if err != nil {
// 		response.WithError(w, err)
// 		return
// 	}
// 	response.WithMetadata(w, http.StatusOK, result, metadata)
// }
