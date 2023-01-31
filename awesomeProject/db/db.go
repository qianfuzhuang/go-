package db

func Dbconn() {
	println("数据库文件")
}

type Demo struct {
	Name string
	Num  int
}
type demo struct {
	Name string
	Num  int
}

func NewDemo(n string, m int) *demo {
	return &demo{n, m}
}

// 相当于构造器
func NewPer(name string) *Demo {
	return &Demo{
		Name: name,
	}
}
func (p *Demo) SetAge(age int) {
	if age > 0 && age < 150 {
		p.Num = age
	} else {
		println("年龄超出限制")
	}
}
func (p *Demo) GetAge() int {
	return p.Num
}
