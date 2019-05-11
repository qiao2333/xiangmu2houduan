package impl

//Myinterface 通用接口
type Myinterface interface {
	Update(id int, obj *interface{}) (bool, string)
	Delete(id int) (bool, string)
	Insert(obj *interface{}) (bool, string)
}
