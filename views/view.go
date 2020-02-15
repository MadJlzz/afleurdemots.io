package views

import (
	"bytes"
	"github.com/madjlzz/madlens/context"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

const (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		// do nothing
		vd = d
	default:
		vd = Data{
			Yield: data,
		}
	}
	vd.User = context.User(r.Context())
	var buf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&buf, v.Layout, vd); err != nil {
		http.Error(w, "Something went wrong. If the problem persists, please email support@madlens.io", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles returns a slice of strings representing
// the layout files in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// addTemplatePath takes in a slice of strings
// representing files paths for template, and it prepends
// the TemplateDir directory to each string in the slice.
//
// Eg the input {"home"} would result the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes in a slice of strings
// representing files paths for template, and it appends
// the TemplateExt extension to each string in the slice.
//
// Eg the input {"home"} would result the output
// {"home.gohtml"} if TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
