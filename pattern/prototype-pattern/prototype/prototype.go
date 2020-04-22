package prototype

// 【原型模式】

type School struct {
	Name    string
	Address string
	Level   string
}

// 创建School实例，返回对应指针
func NewSchool(name, address, level string) *School {
	return &School{
		Name:    name,
		Address: address,
		Level:   level,
	}
}

// School也要实现Clone()方法
func (s *School) Clone() *School {
	return NewSchool(s.Name, s.Address, s.Level)
}

// 重写School的String()方法，便于输出信息
func (s *School) String() string {
	return "{" + s.Name + " " + s.Address + " " + s.Level + "}"
}

type Resume struct {
	Name   string
	Gender string
	// 由于Go中结构体是值类型而不是指针类型，创建新的复合结构体时，会把值复制一份过去，所以这里使用School指针
	School        *School
	Apply4Company string
	Apply4Job     string
}

func NewResume(name, gender string, school *School, company, job string) *Resume {
	return &Resume{
		Name:          name,
		Gender:        gender,
		School:        school,
		Apply4Company: company,
		Apply4Job:     job,
	}
}

// 原型的Clone()方法
func (r *Resume) Clone() *Resume {
	// 注意这里的r.School.Clone()
	return NewResume(r.Name, r.Gender, r.School.Clone(), r.Apply4Company, r.Apply4Job)
}
