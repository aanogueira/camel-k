/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package maven

import (
	"encoding/json"
	"github.com/aanogueira/camel-k/pkg/util/log"
)

type MavenLog struct {
	Level            string `json:"level"`
	Ts               string `json:"ts"`
	Logger           string `json:"logger"`
	Msg              string `json:"msg"`
	Class            string `json:"class"`
	CallerMethodName string `json:"caller_method_name"`
	CallerFileName   string `json:"caller_file_name"`
	CallerLineNumber int    `json:"caller_line_number"`
	Thread           string `json:"thread"`
}

const (
	TRACE = "TRACE"
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

func ParseLog(line string) (mavenLog MavenLog, error error) {
	error = json.Unmarshal([]byte(line), &mavenLog)

	return mavenLog, error
}

func NormalizeLog(mavenLog MavenLog) {
	logger := log.WithName("maven.build")

	switch mavenLog.Level {
	case DEBUG, TRACE:
		logger.Debug(mavenLog.Msg)
	case INFO, WARN:
		logger.Info(mavenLog.Msg)
	case ERROR, FATAL:
		logger.Errorf(nil, mavenLog.Msg)
	}
}
