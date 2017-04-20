package handlers

import (
	"encoding/json"
	"net/http"

	"path"

	"github.com/alex-ramos/info344-in-class/tasksvr/models/tasks"
)

//HandleTasks will handle requests for the /v1/tasks resource
func (ctx *Context) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		newtask := &tasks.NewTask{}
		if err := decoder.Decode(newtask); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := newtask.Validate(); err != nil {
			http.Error(w, "error validating task: "+err.Error(), http.StatusBadRequest)
			return
		}

		task, err := ctx.TasksStore.Insert(newtask)
		if err != nil {
			http.Error(w, "error inserting task : "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)

	case "GET":
		tasks, err := ctx.TasksStore.GetAll()
		if err != nil {
			http.Error(w, "error getting tasks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(tasks)

	}
}

//HandleSpecificTask will handle requests for the /v1/tasks/some-task-id resource
func (ctx *Context) HandleSpecificTask(w http.ResponseWriter, r *http.Request) {
	_, id := path.Split(r.URL.Path)

	switch r.Method {
	case "GET":
		task, err := ctx.TasksStore.Get(id)
		if err != nil {
			http.Error(w, "Error finding task :"+err.Error(), http.StatusBadRequest)
			return
		}

		//this can be made into one function called "Respond"
		//it is used throughout the api source code
		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)

	case "PATCH":
		decoder := json.NewDecoder(r.Body)
		task := &tasks.Task{}
		if err := decoder.Decode(task); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		task.ID = id

		if err := ctx.TasksStore.Update(task); err != nil {}
			if err != nil {
				http.Error(w, "error updating task : "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)

	}

}
