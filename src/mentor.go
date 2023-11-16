package main

import "fmt"

var mentor map[string][]string

func MentorSystem(mentorstudent map[string][]string, mentor_name, student_name string) map[string][]string {
	if mentorstudent == nil {
		mentorstudent = make(map[string][]string)
	}
	// 由于append函数返回一个新的切片，因此我们可以将返回的新切片赋值给mentorstudent[mentor_name]，以实现切片的追加操作
	mentorstudent[mentor_name] = append(mentorstudent[mentor_name], student_name)
	return mentorstudent
}

func initmentor(mentor map[string][]string, mentor_name string, student_names ...string) map[string][]string {
	for _, student_name := range student_names {
		mentor = MentorSystem(mentor, mentor_name, student_name)
	}
	return mentor
}

func main() {
	laozhang := initmentor(mentor, "laozhang", "xiaowang", "xiaoli", "xiaochen")
	laocheng := initmentor(mentor, "laocheng", "xiaoniu", "xiaohu", "xiaozhou")
	fmt.Println(laozhang)
	fmt.Println("laozhang导师的学生有:", laozhang["laozhang"])
	fmt.Println(laocheng)
	fmt.Println("laocheng导师的学生有:", laocheng["laocheng"])

}
