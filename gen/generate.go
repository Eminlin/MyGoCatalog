package main

import (
	"newframework/pkg/config"
	"newframework/pkg/db"
	"newframework/pkg/log"

	"gorm.io/gen"
)

func main() {
	config.Load("conf/config.yml")
	log.InitLogger()
	db.InitDB()

	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dao", //curd代码的输出路径
		ModelPkgPath: "model",        //model代码的输出路径

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db.Database)

	// softDeleteField := gen.FieldType("deleted_at", "gorm.DeletedAt")
	// fieldOpts := []gen.ModelOpt{softDeleteField}
	allModel := g.GenerateAllTable()
	g.ApplyBasic(allModel...)

	// Generate the code
	g.Execute()
}
