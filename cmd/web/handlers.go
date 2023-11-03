package main

import (
	"errors"
	"fmt"
	"net/http"
	"pilrugen.com/snippetbox/pkg/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	//for _, snippet := range s {
	//	fmt.Fprintf(w, "%v\n", snippet)
	//}

	// Создаем экземпляр структуры templateData,
	// содержащий срез с заметками.
	//data := &templateData{Snippets: s}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	//files := []string{
	//	"./ui/html/home.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	//err = ts.Execute(w, data)
	//if err != nil {
	//	// app.errorLog.Println(err.Error())
	//	// http.Error(w, "Internal Server Error", 500)
	//	app.serverError(w, err)
	//}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

	//for _, snippet := range s {
	//	fmt.Fprintf(w, "%v\n", snippet)
	//}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNotRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	//data := &templateData{Snippet: s}
	//
	//files := []string{
	//	"./ui/html/show.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}

	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.serverError(w, err)
	//}
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...\n", id)
	fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		// http.Error(w, "Метот запрещён!", 405)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "История про улитку"
	content := "Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их."
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// w.Write([]byte("Форма для создания новой заметки..."))
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
