package student

import (
	"github.com/go-chi/chi"
	"github.com/kks-learning-management-api/internal/domain/student/service"
)

type StudentHandler struct {
	StudentService service.StudentService
}

func ProvideStudentHandler(studentService service.StudentService) StudentHandler {
	return StudentHandler{
		StudentService: studentService,
	}
}

func (h *StudentHandler) Router(r chi.Router) {

	r.Route("/students", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/{studentId}", h.ViewStudentById)
			r.Get("/", h.ViewStudent)
		})
	})
}
