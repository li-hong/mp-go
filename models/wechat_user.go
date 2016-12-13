package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type WechatUser struct {
	Id             string    `orm:"column(id);pk"`
	CreateTime     time.Time `orm:"column(create_time);type(datetime);null"`
	LastModifyTime time.Time `orm:"column(last_modify_time);type(datetime);null"`
	City           string    `orm:"column(city);size(255);null";json:"city"`
	Country        string    `orm:"column(country);size(255);null";json:"country"`
	Groupid        int       `orm:"column(groupid);null";json:"groupid"`
	Headimgurl     string    `orm:"column(headimgurl);size(255);null";json:"headimgurl"`
	Language       string    `orm:"column(language);size(255);null";json:"language"`
	Nickname       string    `orm:"column(nickname);size(255);null";json:"nickname"`
	Openid         string    `orm:"column(openid);size(255);null";json:"openid"`
	PrivilegeStr   string    `orm:"column(privilege_str);size(255);null"`
	Province       string    `orm:"column(province);size(255);null";json:"province"`
	Remark         string    `orm:"column(remark);size(255);null"`
	Sex            string    `orm:"column(sex);size(10);null";json:"sex"`
	Subscribe      int       `orm:"column(subscribe);null";json:"subscribe"`
	SubscribeTime  int       `orm:"column(subscribe_time);null";json:"subscribe_time"`
	TagidListStr   string    `orm:"column(tagid_list_str);size(255);null"`
	Type           string    `orm:"column(type);size(1);null"`
	Unionid        string    `orm:"column(unionid);size(255);null";json:"unionid"`
	AccountId      string    `orm:"column(account_id);size(100);null"`
	MemberInfoId   string    `orm:"column(member_info_id);size(64);null"`
	Mobile         string    `orm:"column(mobile);size(100);null"`
	Birthday       time.Time `orm:"column(birthday);type(date);null"`
	Profession     string    `orm:"column(profession);size(100);null"`
	Height         string    `orm:"column(height);size(10);null"`
	Weight         string    `orm:"column(weight);size(10);null"`
	Username       string    `orm:"column(username);size(100);null"`
	Age            int       `orm:"column(age);null"`
	DeleteFlag     int       `orm:"column(delete_flag);null"`
	BindTime       time.Time `orm:"column(bind_time);type(datetime);null"`
	MpPlatform     string    `orm:"column(mp_platform);size(100);null"`
	NewRemark      string    `orm:"column(new_remark);size(200);null"`
}

func (t *WechatUser) TableName() string {
	return "wechat_user"
}

func init() {
	orm.RegisterModel(new(WechatUser))
}

// AddWechatUser insert a new WechatUser into database and returns
// last inserted Id on success.
func AddWechatUser(m *WechatUser) (id string, err error) {
	o := orm.NewOrm()
	_, err = o.Insert(m)
	return
}

// GetWechatUserById retrieves WechatUser by Id. Returns error if
// Id doesn't exist
func GetWechatUserById(id string) (v *WechatUser, err error) {
	o := orm.NewOrm()
	v = &WechatUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetWechatUserByOpenid retrieves WechatUser by Openid. Returns error if
// Openid doesn't exist
func GetWechatUserByOpenid(Openid string) (v *WechatUser, err error) {
	o := orm.NewOrm()

	v = &WechatUser{}
	err = o.QueryTable(new(WechatUser)).Filter("openid", Openid).One(v)

	if err == orm.ErrMultiRows {
		// 多条的时候报错
		beego.Error(err)
		return nil, err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		beego.Info(err)
		return nil, nil
	}
	return v, nil
}

// GetAllWechatUser retrieves all WechatUser matches certain condition. Returns empty list if
// no records exist
func GetAllWechatUser(query map[string]string, fields []string, sortby []string, order []string,
offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(WechatUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []WechatUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateWechatUser updates WechatUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateWechatUserById(m *WechatUser) (err error) {
	o := orm.NewOrm()
	v := WechatUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWechatUser deletes WechatUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWechatUser(id string) (err error) {
	o := orm.NewOrm()
	v := WechatUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&WechatUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
