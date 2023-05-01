package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contactMe)
	e.GET("/my-project", myProject)
	// e.GET("/detail-project", detailProject)
	e.GET("/testimonial", testimonial)
	e.POST("/add-project", addProject)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

type Project struct {
	Pname       string
	StartDate   string
	EndDate     string
	Description string
	Tech        []string
}

var dataProject = []Project{
	{
		Pname:       "Teknologi AI: Perkembangan Terkini, Manfaat, dan Tantangan di Era Digital",
		StartDate:   "1 Jan 2023",
		EndDate:     "1 Maret 2023",
		Description: "Teknologi kecerdasan buatan (Artificial Intelligence atau AI) adalah cabang ilmu komputer yang bertujuan untuk menciptakan mesin yang dapat berpikir, belajar, dan bertindak seperti manusia. Dalam beberapa dekade terakhir, AI telah mengalami kemajuan pesat dan banyak digunakan di berbagai sektor seperti industri, bisnis, pemerintahan, kesehatan, pendidikan, dan banyak lagi.",
		Tech:        []string{"NodeJs", "NextJs", "ReactJs", "TypeScript"},
	},
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/myProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
	}

	projects := map[string]interface{}{
		"Project": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
}

// func detailProject(c echo.Context) error {
// 	var tmpl, err = template.ParseFiles("views/detailProject.html")

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
// 	}

// 	return tmpl.Execute(c.Response(), nil)
// }

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

// ==============================================================================

func addProject(c echo.Context) error {
	pName := c.FormValue("project-name")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("description")
	nodeBox := c.FormValue("nodeBox")
	nextBox := c.FormValue("nextBox")
	reactBox := c.FormValue("reactBox")
	typeScriptBox := c.FormValue("typeScriptBox")

	var tech []string
	if nodeBox == "NodeJs" {
		tech = append(tech, "NodeJs")
	}
	if nextBox == "NextJs" {
		tech = append(tech, "NextJs")
	}
	if reactBox == "ReactJs" {
		tech = append(tech, "ReactJs")
	}
	if typeScriptBox == "TypeScript" {
		tech = append(tech, "TypeScript")
	}

	var addProject = Project{
		Pname:       pName,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
		Tech:        tech,
	}

	fmt.Println(addProject)
	dataProject = append(dataProject, addProject)

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tmpl, err := template.ParseFiles("views/detailProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	var projectdata = Project{}
	for index, data := range dataProject {
		if id == index {
			projectdata = Project{
				Pname:       data.Pname,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				Tech:        data.Tech,
			}
		}
	}

	data := map[string]interface{}{
		"Project": projectdata,
	}

	return tmpl.Execute(c.Response(), data)
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataProject = append(dataProject[:id], dataProject[id:+1]...)

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

//
