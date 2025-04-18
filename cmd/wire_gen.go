//  Copyright (c) 2025 dingodb.com, Inc. All Rights Reserved
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http:www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"dingo-hfmirror/internal/dao"
	"dingo-hfmirror/internal/handler"
	"dingo-hfmirror/internal/router"
	"dingo-hfmirror/internal/service"
	"dingo-hfmirror/pkg/config"
	"dingo-hfmirror/pkg/server"
)

// Injectors from wire.go:

func wireApp(configConfig *config.Config) (*App, func(), error) {
	echo := server.NewEngine()
	fileDao := dao.NewFileDao()
	fileService := service.NewFileService(fileDao)
	sysService := service.NewSysService()
	fileHandler := handler.NewFileHandler(fileService, sysService)
	metaDao := dao.NewMetaDao(fileDao)
	metaService := service.NewMetaService(fileDao, metaDao)
	metaHandler := handler.NewMetaHandler(metaService)
	httpRouter := router.NewHttpRouter(echo, fileHandler, metaHandler)
	serverServer := server.NewServer(configConfig, echo, httpRouter)
	app := newApp(serverServer)
	return app, func() {
	}, nil
}
