package main

import (
	"bufio"
	"fmt"
	"os"

	//LsmDB 的实际导入路径
	"log"

	"github.com/HanochYuuka/myparser/common"
	"github.com/HanochYuuka/myparser/executor"
	"github.com/HanochYuuka/myparser/lexer"
	"github.com/HanochYuuka/myparser/parser"
	"github.com/HanochYuuka/myparser/planner"

	db "github.com/ZLSMDB/stpdb-demo/src"
)

func main() {
	input1 := "SELECT name, age FROM users WHERE id = 1"
	lexer1 := lexer.NewLexer(input1)
	parser1 := parser.NewParser(lexer1)
	ast1 := parser1.ParseQuery()
	//astnode存放了分割好的字段
	fmt.Println("AST:", ast1.String())

	// // 假设我们已经解析了一个 AST 节点
	// astNode := &ast.ASTSelectQuery{
	// 	Columns: []string{"name", "age"},
	// 	From:    "users",
	// 	Where:   "id = 1",
	// }

	// 创建 Planner 和 Executor
	planner := planner.NewPlanner()
	executor := executor.NewDefaultExecutor()

	// 生成执行计划
	plan := planner.GeneratePlan(ast1)

	ldb, err := db.NewLevelDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer ldb.Close()

	// 创建 ExecutionContext

	ctx := common.ExecutionContext{
		DB:       nil, //数据库实例
		Metadata: nil, // 假设 metadata 已经初始化
	}

	// 执行计划
	result, err := executor.Execute(plan, ctx)
	if err != nil {
		fmt.Println("Error executing plan:", err)
		return
	}

	// 处理执行结果
	fmt.Println("Rows affected:", result.RowsAffected)
	for _, record := range result.Records {
		fmt.Println(record)
	}
	repl()
}

func repl() {
	// // 初始化数据库
	// ldb, err := db.NewLevelDB("test.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer ldb.Close()

	// 创建 Planner 和 Executor
	planner := planner.NewPlanner()
	executor := executor.NewDefaultExecutor()

	// 创建 ExecutionContext
	// ctx := common.ExecutionContext{
	// 	DB:       ldb, // 数据库实例
	// 	Metadata: nil, // 假设 metadata 已经初始化
	// }

	ctx := common.ExecutionContext{
		DB:       nil, // 数据库实例
		Metadata: nil, // 假设 metadata 已经初始化
	}

	// 创建一个 bufio.Scanner 用于读取用户输入
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("SQL REPL (Read-Eval-Print Loop)")
	fmt.Println("Enter SQL statements or type 'EXIT' to quit.")

	for {
		// 提示用户输入
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// 检查用户是否输入了退出命令
		if input == "EXIT" || input == "QUIT" {
			fmt.Println("Exiting SQL REPL. Goodbye!")
			break
		}

		// 解析用户输入的 SQL 语句
		lexer1 := lexer.NewLexer(input)
		parser1 := parser.NewParser(lexer1)
		ast1 := parser1.ParseQuery()

		// 打印解析后的 AST
		fmt.Println("AST:", ast1.String())

		// 生成执行计划
		plan := planner.GeneratePlan(ast1)

		// 执行计划
		result, err := executor.Execute(plan, ctx)
		if err != nil {
			fmt.Println("Error executing plan:", err)
			continue
		}

		// 处理执行结果
		fmt.Println("Rows affected:", result.RowsAffected)
		for _, record := range result.Records {
			fmt.Println(record)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
