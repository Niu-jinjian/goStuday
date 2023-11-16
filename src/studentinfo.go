package main

import "fmt"

var StudentInfo map[string]map[string]string

func AddStudentInfo(studentinfo map[string]map[string]string, name, id, age, add string) map[string]map[string]string {
	if studentinfo == nil {
		studentinfo = make(map[string]map[string]string)
	}
	studentinfo[name] = make(map[string]string)
	studentinfo[name]["id"] = id
	studentinfo[name]["age"] = age
	studentinfo[name]["add"] = add

	return studentinfo
}

func PrintStudentInfo(studentinfo map[string]map[string]string, name string) {
	if studentname, ok := studentinfo[name]; ok {
		fmt.Printf("学生姓名: %s, 学生ID: %s, 学生年龄: %s, 学生地址: %s\n", name,
			studentname["id"], studentname["age"], studentname["add"])
	} else {
		fmt.Printf("学生 %s 不存在\n", name)
	}
}

func main() {
	Tominfo := AddStudentInfo(StudentInfo, "Tom", "1001", "23", "shanghai")
	PrintStudentInfo(Tominfo, "Tom")

}
