/*
Copyright © 2022 Noah Ispas <noahispas.public@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package internal

import "time"

type Project struct {
	Name string `json:name`
}

type ProjectResponse struct {
	Value []Project `json:value`
}

type Pipeline struct {
	Name string `json:name`
	Id   int    `json:id`
}

type PipelineResponse struct {
	Value []Pipeline `json:value`
}

type PipelineRun struct {
	Id       int       `json:id`
	Name     string    `json:name`
	State    string    `json:state`
	Result   string    `json:result`
	Created  time.Time `json:"createdDate"`
	Finished time.Time `json:"finishedDate"`
}

type PipelineRunResponse struct {
	Value []PipelineRun `json:value`
}

type PipelineRunRequestParameter struct {
	Resources PipelineRunRequestParameterResources `json:resources`
}

type PipelineRunRequestParameterResources struct {
	Repositories PipelineRunRequestParameterRepositories `json:repositories`
}

type PipelineRunRequestParameterRepositories struct {
	Self PipelineRunRequestParameterSelf `json:self`
}

type PipelineRunRequestParameterSelf struct {
	RefName string `json:refName`
}

type PipelineRunLogsResponse struct {
	Count int `json:count`
}

type PipelineRunLogs struct {
	Url string `json:url`
}
