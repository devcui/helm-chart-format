###
# @Author: cuihaonan
# @Email: devcui@outlook.com
# @Date: 2021-04-29 13:47:15
 # @LastEditTime: 2021-04-29 13:54:23
 # @LastEditors: cuihaonan
# @Description: Basic description
 # @FilePath: /helm-char-push/build.sh
# @LICENSE: NONE
###

go build -o helm-chart-format ./main.go
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o helm-chart-format ./main.go
