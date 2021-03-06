// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package wf_run_sign

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

func GetSignInfoById(id uint) (entity *Entity, err error) {
	entity, err = Model.FindOne(id)
	if err != nil {
		g.Log().Debug(err)
		err = gerror.New("获取会签信息失败")
	}
	return
}

func AddSing(data g.Map, tx *gdb.TX) (int64, error) {
	res, err := Model.TX(tx).Insert(data)
	if err != nil {
		g.Log().Error(err)
		return 0, gerror.New("报错会签信息失败")
	}
	id, err := res.LastInsertId()
	if err != nil {
		g.Log().Error(err)
		return 0, gerror.New("获取插入的会签ID失败")
	}
	return id, err
}

func UpdateSing(id int, data g.Map, tx *gdb.TX) error {
	_, err := Model.TX(tx).Where(Columns.Id, id).Update(data)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("更新会签信息失败")
	}
	return nil
}
