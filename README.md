# goshop


##  结构体代码生成工具

[gen_model](./tools/gen_model/main.go) 是一个简单的model结构体生成工具，自动生成对应model代码文件
- 使用 gen_model 之前，请修改 [gen_model/main.go](./tools/gen_model/main.go), 将`cPath`, `mPath`, `rPath ` 修改为生成文件的目标路径
- 在 [tools/gen_model](./tools/gen_model)目录 执行 `go run main.go`