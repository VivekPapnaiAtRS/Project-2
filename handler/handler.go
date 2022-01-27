package handler

import (
	"encoding/json"
	"github.com/Project2/models"
	"github.com/Project2/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func CallService(resp http.ResponseWriter, req *http.Request) {
	//make a HTTP Post Request
	res, err := http.Post("http://127.0.0.1:8080/api/text", "application/json; content=UTF-8", req.Body)
	if err != nil {
		logrus.Errorf("CallService: error in calling service: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputData, _ := ioutil.ReadAll(res.Body)

	err = res.Body.Close()
	if err != nil {
		logrus.Errorf("CallService: error in closing response body: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	//converting byte data into Struct
	var outputDataINJSON []models.WordFrequencyStruct
	err = json.Unmarshal(outputData, &outputDataINJSON)
	if err != nil {
		logrus.Errorf("CallService: error in converting to JSON: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.EncodeJSONBody(resp, http.StatusOK, outputDataINJSON)
}
