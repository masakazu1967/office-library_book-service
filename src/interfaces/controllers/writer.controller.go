package controllers

import (
	"local.package/domain"
	"local.package/interfaces/database"
	"local.package/usecase"
	"strconv"
)

// WriterController 著者コントローラ
type WriterController struct {
	Interactor usecase.WriterInteractor
}

// NewWriterController 著者コントローラを生成する
func NewWriterController(sqlHandler database.SQLHandler) *WriterController {
	return &WriterController{
		Interactor: usecase.WriterInteractor{
			WriterRepository: &database.WriterRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Create 著者を登録する
func (controller *WriterController) Create(c Context) {
	w := domain.Writer{}
	c.Bind(&w)
	writer, err := controller.Interactor.Add(w)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, writer)
}

// Index 著者一覧を取得する
func (controller *WriterController) Index(c Context) {
	writers, err := controller.Interactor.Writers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, writers)
}

// Show 著者情報を取得する
func (controller *WriterController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	writer, err := controller.Interactor.WriterByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, writer)
}
