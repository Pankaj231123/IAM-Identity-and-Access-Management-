#!/usr/bin/env bash
set -e

# Create folders
mkdir -p cmd/api
mkdir -p internal/config
mkdir -p internal/handler
mkdir -p internal/middleware
mkdir -p internal/model
mkdir -p internal/repository
mkdir -p internal/service
mkdir -p pkg/hash
mkdir -p pkg/token
mkdir -p migrations

# Add placeholder doc.go files so Git tracks empty Go package dirs
echo "package config"     > internal/config/doc.go
echo "package handler"    > internal/handler/doc.go
echo "package middleware" > internal/middleware/doc.go
echo "package model"      > internal/model/doc.go
echo "package repository" > internal/repository/doc.go
echo "package service"    > internal/service/doc.go
echo "package hash"       > pkg/hash/doc.go
echo "package token"      > pkg/token/doc.go

# Keep migrations/ tracked even though it's empty for now
touch migrations/.gitkeep

echo "Scaffold created."