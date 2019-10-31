package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Levi-ackerman/go-go-go/utils"
	"github.com/gorilla/mux"
)

type FileController struct {
	router *mux.Router
}

func NewFileController(r *mux.Router) *FileController {
	return &FileController{
		router: r,
	}
}

func (service *FileController) Upload() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("File Upload Endpoint Hit")

		req.ParseMultipartForm(10 << 20)

		file, handler, err := req.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Println(file)

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		dir := os.Getenv("images_directory")
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("this directory is not exist, create one")
			err := os.Mkdir(dir, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
		tempFile, err := ioutil.TempFile(dir, "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		//fmt.Fprintf(res, "Successfully Uploaded File\n")

		utils.Respond(res, utils.Message(http.StatusOK, "Successfully Uploaded File"+tempFile.Name()))

	})
}

func (controller *FileController) MakeUserHandler() {
	controller.router.Handle("/api/file/upload", controller.Upload()).Methods("POST")
}
